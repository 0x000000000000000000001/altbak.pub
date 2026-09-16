package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Foldable_identity gopurs_runtime.Value
var once_Data_Foldable_identity sync.Once
func Get_Data_Foldable_identity() gopurs_runtime.Value {
	once_Data_Foldable_identity.Do(func() {
		cache_Data_Foldable_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_Foldable_identity
}

var cache_Data_Foldable_unwrap gopurs_runtime.Value
var once_Data_Foldable_unwrap sync.Once
func Get_Data_Foldable_unwrap() gopurs_runtime.Value {
	once_Data_Foldable_unwrap.Do(func() {
		cache_Data_Foldable_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_Foldable_unwrap
}

var cache_Data_Foldable_monoidEndo gopurs_runtime.Value
var once_Data_Foldable_monoidEndo sync.Once
func Get_Data_Foldable_monoidEndo() gopurs_runtime.Value {
	once_Data_Foldable_monoidEndo.Do(func() {
		cache_Data_Foldable_monoidEndo = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}
	})
	return cache_Data_Foldable_monoidEndo
}

var cache_Data_Foldable_identity1 gopurs_runtime.Value
var once_Data_Foldable_identity1 sync.Once
func Get_Data_Foldable_identity1() gopurs_runtime.Value {
	once_Data_Foldable_identity1.Do(func() {
		cache_Data_Foldable_identity1 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_Foldable_identity1
}

var cache_Data_Foldable_identity2 gopurs_runtime.Value
var once_Data_Foldable_identity2 sync.Once
func Get_Data_Foldable_identity2() gopurs_runtime.Value {
	once_Data_Foldable_identity2.Do(func() {
		cache_Data_Foldable_identity2 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_Foldable_identity2
}

var cache_Data_Foldable_not gopurs_runtime.Value
var once_Data_Foldable_not sync.Once
func Get_Data_Foldable_not() gopurs_runtime.Value {
	once_Data_Foldable_not.Do(func() {
		cache_Data_Foldable_not = Call_Data_HeytingAlgebra_not(Rebox_Data_Foldable_3591112874_2663347022(Rebox_Data_Foldable_2663347022_3591112874(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Get_Data_HeytingAlgebra_heytingAlgebraBoolean()))))
	})
	return cache_Data_Foldable_not
}

var cache_Data_Foldable_identity3 gopurs_runtime.Value
var once_Data_Foldable_identity3 sync.Once
func Get_Data_Foldable_identity3() gopurs_runtime.Value {
	once_Data_Foldable_identity3.Do(func() {
		cache_Data_Foldable_identity3 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_Foldable_identity3
}

var cache_Data_Foldable_identity4 gopurs_runtime.Value
var once_Data_Foldable_identity4 sync.Once
func Get_Data_Foldable_identity4() gopurs_runtime.Value {
	once_Data_Foldable_identity4.Do(func() {
		cache_Data_Foldable_identity4 = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_Foldable_identity4
}

var cache_Data_Foldable_Empty gopurs_runtime.Value
var once_Data_Foldable_Empty sync.Once
func Get_Data_Foldable_Empty() gopurs_runtime.Value {
	once_Data_Foldable_Empty.Do(func() {
		cache_Data_Foldable_Empty = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Data_Foldable_Empty
}

var cache_Data_Foldable_Node gopurs_runtime.Value
var once_Data_Foldable_Node sync.Once
func Get_Data_Foldable_Node() gopurs_runtime.Value {
	once_Data_Foldable_Node.Do(func() {
		cache_Data_Foldable_Node = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2421944209, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Node[gopurs_runtime.Value]{1, value0}))}
})
	})
	return cache_Data_Foldable_Node
}

var cache_Data_Foldable_Append gopurs_runtime.Value
var once_Data_Foldable_Append sync.Once
func Get_Data_Foldable_Append() gopurs_runtime.Value {
	once_Data_Foldable_Append.Do(func() {
		cache_Data_Foldable_Append = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Append[gopurs_runtime.Value]{1, value0, value1}))}
})
})
	})
	return cache_Data_Foldable_Append
}

var cache_Data_Foldable_Foldable_dollar_Dict gopurs_runtime.Value
var once_Data_Foldable_Foldable_dollar_Dict sync.Once
func Get_Data_Foldable_Foldable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Foldable_Foldable_dollar_Dict.Do(func() {
		cache_Data_Foldable_Foldable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Call_Data_Foldable_Foldable_dollar_Dict(func() struct{
	foldMap gopurs_runtime.Value
	foldl gopurs_runtime.Value
	foldr gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	foldMap gopurs_runtime.Value
	foldl gopurs_runtime.Value
	foldr gopurs_runtime.Value
}{}
					clone.foldMap = gopurs_runtime.RecordGet(orig, "foldMap")
					clone.foldl = gopurs_runtime.RecordGet(orig, "foldl")
					clone.foldr = gopurs_runtime.RecordGet(orig, "foldr")
					return clone
				}()))}
})
	})
	return cache_Data_Foldable_Foldable_dollar_Dict
}

var cache_Data_Foldable_semigroupFreeMonoidTree gopurs_runtime.Value
var once_Data_Foldable_semigroupFreeMonoidTree sync.Once
func Get_Data_Foldable_semigroupFreeMonoidTree() gopurs_runtime.Value {
	once_Data_Foldable_semigroupFreeMonoidTree.Do(func() {
		cache_Data_Foldable_semigroupFreeMonoidTree = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, Get_Data_Foldable_Append()}))}
	})
	return cache_Data_Foldable_semigroupFreeMonoidTree
}

var cache_Data_Foldable_monoidFreeMonoidTree gopurs_runtime.Value
var once_Data_Foldable_monoidFreeMonoidTree sync.Once
func Get_Data_Foldable_monoidFreeMonoidTree() gopurs_runtime.Value {
	once_Data_Foldable_monoidFreeMonoidTree.Do(func() {
		cache_Data_Foldable_monoidFreeMonoidTree = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Foldable_semigroupFreeMonoidTree()))}
}), gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}}))}
	})
	return cache_Data_Foldable_monoidFreeMonoidTree
}

var cache_Data_Foldable_foldr gopurs_runtime.Value
var once_Data_Foldable_foldr sync.Once
func Get_Data_Foldable_foldr() gopurs_runtime.Value {
	once_Data_Foldable_foldr.Do(func() {
		cache_Data_Foldable_foldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr
}

var cache_Data_Foldable_indexr gopurs_runtime.Value
var once_Data_Foldable_indexr sync.Once
func Get_Data_Foldable_indexr() gopurs_runtime.Value {
	once_Data_Foldable_indexr.Do(func() {
		cache_Data_Foldable_indexr = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, idx_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_indexr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), idx_1_box.IntVal)
})
	})
	return cache_Data_Foldable_indexr
}

var cache_Data_Foldable_null gopurs_runtime.Value
var once_Data_Foldable_null sync.Once
func Get_Data_Foldable_null() gopurs_runtime.Value {
	once_Data_Foldable_null.Do(func() {
		cache_Data_Foldable_null = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_null(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_Data_Foldable_null
}

var cache_Data_Foldable_oneOf gopurs_runtime.Value
var once_Data_Foldable_oneOf sync.Once
func Get_Data_Foldable_oneOf() gopurs_runtime.Value {
	once_Data_Foldable_oneOf.Do(func() {
		cache_Data_Foldable_oneOf = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](dictPlus_1_box))
})
	})
	return cache_Data_Foldable_oneOf
}

var cache_Data_Foldable_oneOfMap gopurs_runtime.Value
var once_Data_Foldable_oneOfMap sync.Once
func Get_Data_Foldable_oneOfMap() gopurs_runtime.Value {
	once_Data_Foldable_oneOfMap.Do(func() {
		cache_Data_Foldable_oneOfMap = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_oneOfMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus[gopurs_runtime.Value]](dictPlus_1_box))
})
	})
	return cache_Data_Foldable_oneOfMap
}

var cache_Data_Foldable_traverse_ gopurs_runtime.Value
var once_Data_Foldable_traverse_ sync.Once
func Get_Data_Foldable_traverse_() gopurs_runtime.Value {
	once_Data_Foldable_traverse_.Do(func() {
		cache_Data_Foldable_traverse_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_traverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Data_Foldable_traverse_
}

var cache_Data_Foldable_for_ gopurs_runtime.Value
var once_Data_Foldable_for_ sync.Once
func Get_Data_Foldable_for_() gopurs_runtime.Value {
	once_Data_Foldable_for_.Do(func() {
		cache_Data_Foldable_for_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_for_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Data_Foldable_for_
}

var cache_Data_Foldable_sequence_ gopurs_runtime.Value
var once_Data_Foldable_sequence_ sync.Once
func Get_Data_Foldable_sequence_() gopurs_runtime.Value {
	once_Data_Foldable_sequence_.Do(func() {
		cache_Data_Foldable_sequence_ = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_sequence_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_1_box))
})
	})
	return cache_Data_Foldable_sequence_
}

var cache_Data_Foldable_foldl gopurs_runtime.Value
var once_Data_Foldable_foldl sync.Once
func Get_Data_Foldable_foldl() gopurs_runtime.Value {
	once_Data_Foldable_foldl.Do(func() {
		cache_Data_Foldable_foldl = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl
}

var cache_Data_Foldable_indexl gopurs_runtime.Value
var once_Data_Foldable_indexl sync.Once
func Get_Data_Foldable_indexl() gopurs_runtime.Value {
	once_Data_Foldable_indexl.Do(func() {
		cache_Data_Foldable_indexl = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, idx_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_indexl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), idx_1_box.IntVal)
})
	})
	return cache_Data_Foldable_indexl
}

var cache_Data_Foldable_intercalate gopurs_runtime.Value
var once_Data_Foldable_intercalate sync.Once
func Get_Data_Foldable_intercalate() gopurs_runtime.Value {
	once_Data_Foldable_intercalate.Do(func() {
		cache_Data_Foldable_intercalate = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_Data_Foldable_intercalate
}

var cache_Data_Foldable_intercalate__4228433826 gopurs_runtime.Value
var once_Data_Foldable_intercalate__4228433826 sync.Once
func Get_Data_Foldable_intercalate__4228433826() gopurs_runtime.Value {
	once_Data_Foldable_intercalate__4228433826.Do(func() {
		cache_Data_Foldable_intercalate__4228433826 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Foldable_intercalate__4228433826(__eta_norm_1_0_box.StrVal(), Rebox_Data_Foldable_849153993_128126966(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__eta_norm_0_1_box))))
})
	})
	return cache_Data_Foldable_intercalate__4228433826
}

var cache_Data_Foldable_length gopurs_runtime.Value
var once_Data_Foldable_length sync.Once
func Get_Data_Foldable_length() gopurs_runtime.Value {
	once_Data_Foldable_length.Do(func() {
		cache_Data_Foldable_length = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_length(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_length
}

var cache_Data_Foldable_maximumBy gopurs_runtime.Value
var once_Data_Foldable_maximumBy sync.Once
func Get_Data_Foldable_maximumBy() gopurs_runtime.Value {
	once_Data_Foldable_maximumBy.Do(func() {
		cache_Data_Foldable_maximumBy = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_maximumBy(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), cmp_1_box)
})
	})
	return cache_Data_Foldable_maximumBy
}

var cache_Data_Foldable_maximum gopurs_runtime.Value
var once_Data_Foldable_maximum sync.Once
func Get_Data_Foldable_maximum() gopurs_runtime.Value {
	once_Data_Foldable_maximum.Do(func() {
		cache_Data_Foldable_maximum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_maximum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Foldable_maximum
}

var cache_Data_Foldable_minimumBy gopurs_runtime.Value
var once_Data_Foldable_minimumBy sync.Once
func Get_Data_Foldable_minimumBy() gopurs_runtime.Value {
	once_Data_Foldable_minimumBy.Do(func() {
		cache_Data_Foldable_minimumBy = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_minimumBy(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), cmp_1_box)
})
	})
	return cache_Data_Foldable_minimumBy
}

var cache_Data_Foldable_minimum gopurs_runtime.Value
var once_Data_Foldable_minimum sync.Once
func Get_Data_Foldable_minimum() gopurs_runtime.Value {
	once_Data_Foldable_minimum.Do(func() {
		cache_Data_Foldable_minimum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_minimum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Foldable_minimum
}

var cache_Data_Foldable_product gopurs_runtime.Value
var once_Data_Foldable_product sync.Once
func Get_Data_Foldable_product() gopurs_runtime.Value {
	once_Data_Foldable_product.Do(func() {
		cache_Data_Foldable_product = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_product(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_product
}

var cache_Data_Foldable_sum gopurs_runtime.Value
var once_Data_Foldable_sum sync.Once
func Get_Data_Foldable_sum() gopurs_runtime.Value {
	once_Data_Foldable_sum.Do(func() {
		cache_Data_Foldable_sum = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_sum(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_sum
}

var cache_Data_Foldable_foldableTuple gopurs_runtime.Value
var once_Data_Foldable_foldableTuple sync.Once
func Get_Data_Foldable_foldableTuple() gopurs_runtime.Value {
	once_Data_Foldable_foldableTuple.Do(func() {
		cache_Data_Foldable_foldableTuple = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_4173511203_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
})})))}
	})
	return cache_Data_Foldable_foldableTuple
}

var cache_Data_Foldable_foldableMultiplicative gopurs_runtime.Value
var once_Data_Foldable_foldableMultiplicative sync.Once
func Get_Data_Foldable_foldableMultiplicative() gopurs_runtime.Value {
	once_Data_Foldable_foldableMultiplicative.Do(func() {
		cache_Data_Foldable_foldableMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})}))}
	})
	return cache_Data_Foldable_foldableMultiplicative
}

var cache_Data_Foldable_foldableMaybe gopurs_runtime.Value
var once_Data_Foldable_foldableMaybe sync.Once
func Get_Data_Foldable_foldableMaybe() gopurs_runtime.Value {
	once_Data_Foldable_foldableMaybe.Do(func() {
		cache_Data_Foldable_foldableMaybe = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_1146820559_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope151)
mempty_1_0 := Call_Data_Monoid_mempty(dictMonoid_0)
_ = mempty_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t3 = mempty_1_0
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
_ = __t_tag_2
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2)
_ = __t_tag_4
if (__t_tag_4 == nil) {
__t6 = v1_1
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2)
_ = __t_tag_5
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2)
_ = __t_tag_7
if (__t_tag_7 == nil) {
__t9 = v1_1
goto end_branch_9
} else {

}
}
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_2)
_ = __t_tag_8
if (__t_tag_8 != nil) {
__t9 = gopurs_runtime.Apply2(v_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})})))}
	})
	return cache_Data_Foldable_foldableMaybe
}

var cache_Data_Foldable_foldableIdentity gopurs_runtime.Value
var once_Data_Foldable_foldableIdentity sync.Once
func Get_Data_Foldable_foldableIdentity() gopurs_runtime.Value {
	once_Data_Foldable_foldableIdentity.Do(func() {
		cache_Data_Foldable_foldableIdentity = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})}))}
	})
	return cache_Data_Foldable_foldableIdentity
}

var cache_Data_Foldable_foldableEither gopurs_runtime.Value
var once_Data_Foldable_foldableEither sync.Once
func Get_Data_Foldable_foldableEither() gopurs_runtime.Value {
	once_Data_Foldable_foldableEither.Do(func() {
		cache_Data_Foldable_foldableEither = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope179)
mempty_1_0 := Call_Data_Monoid_mempty(dictMonoid_0)
_ = mempty_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t1 = mempty_1_0
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply(v_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t2 = v1_1
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t3 = v1_1
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(v_0, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})}))}
	})
	return cache_Data_Foldable_foldableEither
}

var cache_Data_Foldable_foldableDual gopurs_runtime.Value
var once_Data_Foldable_foldableDual sync.Once
func Get_Data_Foldable_foldableDual() gopurs_runtime.Value {
	once_Data_Foldable_foldableDual.Do(func() {
		cache_Data_Foldable_foldableDual = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})}))}
	})
	return cache_Data_Foldable_foldableDual
}

var cache_Data_Foldable_foldableDisj gopurs_runtime.Value
var once_Data_Foldable_foldableDisj sync.Once
func Get_Data_Foldable_foldableDisj() gopurs_runtime.Value {
	once_Data_Foldable_foldableDisj.Do(func() {
		cache_Data_Foldable_foldableDisj = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})}))}
	})
	return cache_Data_Foldable_foldableDisj
}

var cache_Data_Foldable_foldableConst gopurs_runtime.Value
var once_Data_Foldable_foldableConst sync.Once
func Get_Data_Foldable_foldableConst() gopurs_runtime.Value {
	once_Data_Foldable_foldableConst.Do(func() {
		cache_Data_Foldable_foldableConst = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope220)
mempty_1_0 := Call_Data_Monoid_mempty(dictMonoid_0)
_ = mempty_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})}))}
	})
	return cache_Data_Foldable_foldableConst
}

var cache_Data_Foldable_foldableConj gopurs_runtime.Value
var once_Data_Foldable_foldableConj sync.Once
func Get_Data_Foldable_foldableConj() gopurs_runtime.Value {
	once_Data_Foldable_foldableConj.Do(func() {
		cache_Data_Foldable_foldableConj = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})}))}
	})
	return cache_Data_Foldable_foldableConj
}

var cache_Data_Foldable_foldableAdditive gopurs_runtime.Value
var once_Data_Foldable_foldableAdditive sync.Once
func Get_Data_Foldable_foldableAdditive() gopurs_runtime.Value {
	once_Data_Foldable_foldableAdditive.Do(func() {
		cache_Data_Foldable_foldableAdditive = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})}))}
	})
	return cache_Data_Foldable_foldableAdditive
}

var cache_Data_Foldable_foldMapDefaultR gopurs_runtime.Value
var once_Data_Foldable_foldMapDefaultR sync.Once
func Get_Data_Foldable_foldMapDefaultR() gopurs_runtime.Value {
	once_Data_Foldable_foldMapDefaultR.Do(func() {
		cache_Data_Foldable_foldMapDefaultR = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_Data_Foldable_foldMapDefaultR
}

var cache_Data_Foldable_foldableArray gopurs_runtime.Value
var once_Data_Foldable_foldableArray sync.Once
func Get_Data_Foldable_foldableArray() gopurs_runtime.Value {
	once_Data_Foldable_foldableArray.Do(func() {
		cache_Data_Foldable_foldableArray = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))
}), Get_Data_Foldable_foldlArray(), Get_Data_Foldable_foldrArray()}))}
	})
	return cache_Data_Foldable_foldableArray
}

var cache_Data_Foldable_foldableFreeMonoidTree gopurs_runtime.Value
var once_Data_Foldable_foldableFreeMonoidTree sync.Once
func Get_Data_Foldable_foldableFreeMonoidTree() gopurs_runtime.Value {
	once_Data_Foldable_foldableFreeMonoidTree.Do(func() {
		cache_Data_Foldable_foldableFreeMonoidTree = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFreeMonoidTree()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Foldable_go__go_1_0_0 func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_Foldable_go__go_1_0_0
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
Call_local_Data_Foldable_go__go_1_0_0 = func(acc_2_loop gopurs_runtime.Value, lhs_3_loop gopurs_runtime.Value, rhs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t5 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, acc_2, (*Constructor_Data_Foldable_Node[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0)
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_0_0
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2812549951) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
_ = __t_tag_2
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = rhs_4
continue go__go_1_0_0
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
var __t1 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
continue go__go_1_0_0
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Append[gopurs_runtime.Value]{1, (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1, rhs_4}))}
continue go__go_1_0_0
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_1:
__t3 = __t1
}
end_branch_3:
__t5 = __t3
goto end_branch_5
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
var __t4 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
__t4 = acc_2
goto end_branch_4
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_0_0
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_4:
__t5 = __t4
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
go__go_1_0_0 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Foldable_go__go_1_0_0(acc_2_loop_val, lhs_3_loop_val, rhs_4_loop_val)
})
})
})
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Foldable_go__go_1_0_0(a_2, b_3, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Foldable_go__go_1_6_1 func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_Foldable_go__go_1_6_1
var go__go_1_6_1 gopurs_runtime.Value
_ = go__go_1_6_1
Call_local_Data_Foldable_go__go_1_6_1 = func(acc_2_loop gopurs_runtime.Value, lhs_3_loop gopurs_runtime.Value, rhs_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_1_6_1:
for {
if false { continue go__go_1_6_1 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t11 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, (*Constructor_Data_Foldable_Node[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0, acc_2)
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_6_1
__t11 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_11
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2812549951) {
var __t9 gopurs_runtime.Value
{
var __t_tag_8 gopurs_runtime.Value = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
_ = __t_tag_8
if (__t_tag_8.Type == 9 && __t_tag_8.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_6_1
__t9 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_9
} else {

}
}
{
var __t7 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
rhs_4_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_6_1
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_7
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Append[gopurs_runtime.Value]{1, lhs_3, (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0}))}
rhs_4_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_6_1
__t7 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_7:
__t9 = __t7
}
end_branch_9:
__t11 = __t9
goto end_branch_11
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
var __t10 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
__t10 = acc_2
goto end_branch_10
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_6_1
__t10 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_1_6_1 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Foldable_go__go_1_6_1(acc_2_loop_val, lhs_3_loop_val, rhs_4_loop_val)
})
})
})
return gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Foldable_go__go_1_6_1(a_2, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}, b_3)
})
})}))}
	})
	return cache_Data_Foldable_foldableFreeMonoidTree
}

var cache_Data_Foldable_foldMapDefaultL gopurs_runtime.Value
var once_Data_Foldable_foldMapDefaultL sync.Once
func Get_Data_Foldable_foldMapDefaultL() gopurs_runtime.Value {
	once_Data_Foldable_foldMapDefaultL.Do(func() {
		cache_Data_Foldable_foldMapDefaultL = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMapDefaultL(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_Data_Foldable_foldMapDefaultL
}

var cache_Data_Foldable_foldMap gopurs_runtime.Value
var once_Data_Foldable_foldMap sync.Once
func Get_Data_Foldable_foldMap() gopurs_runtime.Value {
	once_Data_Foldable_foldMap.Do(func() {
		cache_Data_Foldable_foldMap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap
}

var cache_Data_Foldable_foldMap__4139410016 gopurs_runtime.Value
var once_Data_Foldable_foldMap__4139410016 sync.Once
func Get_Data_Foldable_foldMap__4139410016() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__4139410016.Do(func() {
		cache_Data_Foldable_foldMap__4139410016 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Foldable_foldMap__4139410016(__eta_norm_1_0_box, Rebox_Data_Foldable_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__eta_norm_0_1_box))))
})
	})
	return cache_Data_Foldable_foldMap__4139410016
}

var cache_Data_Foldable_foldableApp gopurs_runtime.Value
var once_Data_Foldable_foldableApp sync.Once
func Get_Data_Foldable_foldableApp() gopurs_runtime.Value {
	once_Data_Foldable_foldableApp.Do(func() {
		cache_Data_Foldable_foldableApp = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldableApp(dictFoldable_0_box)
})
	})
	return cache_Data_Foldable_foldableApp
}

var cache_Data_Foldable_foldableCompose gopurs_runtime.Value
var once_Data_Foldable_foldableCompose sync.Once
func Get_Data_Foldable_foldableCompose() gopurs_runtime.Value {
	once_Data_Foldable_foldableCompose.Do(func() {
		cache_Data_Foldable_foldableCompose = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldableCompose(dictFoldable_0_box, dictFoldable1_1_box)
})
	})
	return cache_Data_Foldable_foldableCompose
}

var cache_Data_Foldable_foldableCoproduct gopurs_runtime.Value
var once_Data_Foldable_foldableCoproduct sync.Once
func Get_Data_Foldable_foldableCoproduct() gopurs_runtime.Value {
	once_Data_Foldable_foldableCoproduct.Do(func() {
		cache_Data_Foldable_foldableCoproduct = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldableCoproduct(dictFoldable_0_box, dictFoldable1_1_box)
})
	})
	return cache_Data_Foldable_foldableCoproduct
}

var cache_Data_Foldable_foldableFirst gopurs_runtime.Value
var once_Data_Foldable_foldableFirst sync.Once
func Get_Data_Foldable_foldableFirst() gopurs_runtime.Value {
	once_Data_Foldable_foldableFirst.Do(func() {
		cache_Data_Foldable_foldableFirst = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_1146820559_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_3_0 shape=App(Var) bindingType=(TypeVar m$scope151)
mempty_3_0 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))})
_ = mempty_3_0
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t3 = mempty_3_0
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_2
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_4
if (__t_tag_4 == nil) {
__t6 = z_1
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_5
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_7
if (__t_tag_7 == nil) {
__t9 = z_1
goto end_branch_9
} else {

}
}
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_8
if (__t_tag_8 != nil) {
__t9 = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, z_1)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})})))}
	})
	return cache_Data_Foldable_foldableFirst
}

var cache_Data_Foldable_foldableLast gopurs_runtime.Value
var once_Data_Foldable_foldableLast sync.Once
func Get_Data_Foldable_foldableLast() gopurs_runtime.Value {
	once_Data_Foldable_foldableLast.Do(func() {
		cache_Data_Foldable_foldableLast = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_1146820559_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_3_0 shape=App(Var) bindingType=(TypeVar m$scope151)
mempty_3_0 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))})
_ = mempty_3_0
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t3 = mempty_3_0
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_2
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_4
if (__t_tag_4 == nil) {
__t6 = z_1
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_5
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_7
if (__t_tag_7 == nil) {
__t9 = z_1
goto end_branch_9
} else {

}
}
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_8
if (__t_tag_8 != nil) {
__t9 = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, z_1)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})})))}
	})
	return cache_Data_Foldable_foldableLast
}

var cache_Data_Foldable_foldableProduct gopurs_runtime.Value
var once_Data_Foldable_foldableProduct sync.Once
func Get_Data_Foldable_foldableProduct() gopurs_runtime.Value {
	once_Data_Foldable_foldableProduct.Do(func() {
		cache_Data_Foldable_foldableProduct = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldableProduct(dictFoldable_0_box, dictFoldable1_1_box)
})
	})
	return cache_Data_Foldable_foldableProduct
}

var cache_Data_Foldable_foldlDefault gopurs_runtime.Value
var once_Data_Foldable_foldlDefault sync.Once
func Get_Data_Foldable_foldlDefault() gopurs_runtime.Value {
	once_Data_Foldable_foldlDefault.Do(func() {
		cache_Data_Foldable_foldlDefault = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldlDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_Data_Foldable_foldlDefault
}

var cache_Data_Foldable_foldrDefault gopurs_runtime.Value
var once_Data_Foldable_foldrDefault sync.Once
func Get_Data_Foldable_foldrDefault() gopurs_runtime.Value {
	once_Data_Foldable_foldrDefault.Do(func() {
		cache_Data_Foldable_foldrDefault = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_Data_Foldable_foldrDefault
}

var cache_Data_Foldable_lookup gopurs_runtime.Value
var once_Data_Foldable_lookup sync.Once
func Get_Data_Foldable_lookup() gopurs_runtime.Value {
	once_Data_Foldable_lookup.Do(func() {
		cache_Data_Foldable_lookup = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_lookup(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1_box), a_2_box)
})
	})
	return cache_Data_Foldable_lookup
}

var cache_Data_Foldable_surroundMap gopurs_runtime.Value
var once_Data_Foldable_surroundMap sync.Once
func Get_Data_Foldable_surroundMap() gopurs_runtime.Value {
	once_Data_Foldable_surroundMap.Do(func() {
		cache_Data_Foldable_surroundMap = gopurs_runtime.Func5(func(dictFoldable_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value, t_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_surroundMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box), d_2_box, t_3_box, f_4_box)
})
	})
	return cache_Data_Foldable_surroundMap
}

var cache_Data_Foldable_surround gopurs_runtime.Value
var once_Data_Foldable_surround sync.Once
func Get_Data_Foldable_surround() gopurs_runtime.Value {
	once_Data_Foldable_surround.Do(func() {
		cache_Data_Foldable_surround = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_surround(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box), d_2_box)
})
	})
	return cache_Data_Foldable_surround
}

var cache_Data_Foldable_foldM gopurs_runtime.Value
var once_Data_Foldable_foldM sync.Once
func Get_Data_Foldable_foldM() gopurs_runtime.Value {
	once_Data_Foldable_foldM.Do(func() {
		cache_Data_Foldable_foldM = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldM(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_1_box))
})
	})
	return cache_Data_Foldable_foldM
}

var cache_Data_Foldable_fold gopurs_runtime.Value
var once_Data_Foldable_fold sync.Once
func Get_Data_Foldable_fold() gopurs_runtime.Value {
	once_Data_Foldable_fold.Do(func() {
		cache_Data_Foldable_fold = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_fold(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_Data_Foldable_fold
}

var cache_Data_Foldable_fold__1068017326 gopurs_runtime.Value
var once_Data_Foldable_fold__1068017326 sync.Once
func Get_Data_Foldable_fold__1068017326() gopurs_runtime.Value {
	once_Data_Foldable_fold__1068017326.Do(func() {
		cache_Data_Foldable_fold__1068017326 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_fold__1068017326(__eta_norm_0_0_box)
})
	})
	return cache_Data_Foldable_fold__1068017326
}

var cache_Data_Foldable_findMap gopurs_runtime.Value
var once_Data_Foldable_findMap sync.Once
func Get_Data_Foldable_findMap() gopurs_runtime.Value {
	once_Data_Foldable_findMap.Do(func() {
		cache_Data_Foldable_findMap = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_findMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), p_1_box)
})
	})
	return cache_Data_Foldable_findMap
}

var cache_Data_Foldable_find gopurs_runtime.Value
var once_Data_Foldable_find sync.Once
func Get_Data_Foldable_find() gopurs_runtime.Value {
	once_Data_Foldable_find.Do(func() {
		cache_Data_Foldable_find = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_find(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), p_1_box)
})
	})
	return cache_Data_Foldable_find
}

var cache_Data_Foldable_any gopurs_runtime.Value
var once_Data_Foldable_any sync.Once
func Get_Data_Foldable_any() gopurs_runtime.Value {
	once_Data_Foldable_any.Do(func() {
		cache_Data_Foldable_any = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_any(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_Data_Foldable_any
}

var cache_Data_Foldable_elem gopurs_runtime.Value
var once_Data_Foldable_elem sync.Once
func Get_Data_Foldable_elem() gopurs_runtime.Value {
	once_Data_Foldable_elem.Do(func() {
		cache_Data_Foldable_elem = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_elem(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_Data_Foldable_elem
}

var cache_Data_Foldable_notElem gopurs_runtime.Value
var once_Data_Foldable_notElem sync.Once
func Get_Data_Foldable_notElem() gopurs_runtime.Value {
	once_Data_Foldable_notElem.Do(func() {
		cache_Data_Foldable_notElem = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_notElem(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1_box), x_2_box)
})
	})
	return cache_Data_Foldable_notElem
}

var cache_Data_Foldable_or gopurs_runtime.Value
var once_Data_Foldable_or sync.Once
func Get_Data_Foldable_or() gopurs_runtime.Value {
	once_Data_Foldable_or.Do(func() {
		cache_Data_Foldable_or = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_or(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_Data_Foldable_or
}

var cache_Data_Foldable_all gopurs_runtime.Value
var once_Data_Foldable_all sync.Once
func Get_Data_Foldable_all() gopurs_runtime.Value {
	once_Data_Foldable_all.Do(func() {
		cache_Data_Foldable_all = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_all(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_Data_Foldable_all
}

var cache_Data_Foldable_and gopurs_runtime.Value
var once_Data_Foldable_and sync.Once
func Get_Data_Foldable_and() gopurs_runtime.Value {
	once_Data_Foldable_and.Do(func() {
		cache_Data_Foldable_and = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_and(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_Data_Foldable_and
}

type Constructor_Data_Foldable_Empty[T_a any] struct {
	Rc uint32
}


type Constructor_Data_Foldable_Node[T_a any] struct {
	Rc uint32
	V0 T_a
}


type Constructor_Data_Foldable_Append[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Data_Foldable_Foldable[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4280266298] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Foldable_Foldable[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "foldMap": return gopurs_runtime.Box(c.V0)
		case "foldl": return gopurs_runtime.Box(c.V1)
		case "foldr": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Foldable_Foldable: " + key)
		}
	}
}


func Call_Data_Foldable_Foldable_dollar_Dict(x_0_loop struct{
	foldMap gopurs_runtime.Value
	foldl gopurs_runtime.Value
	foldr gopurs_runtime.Value
}) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
var x_0 struct{
	foldMap gopurs_runtime.Value
	foldl gopurs_runtime.Value
	foldr gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", orig.foldMap, orig.foldl, orig.foldr)
				}())
}

func Call_Data_Foldable_foldr(dict_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_Data_Foldable_indexr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], idx_1_loop int64) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var idx_1 int64 = idx_1_loop
_ = idx_1
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_2, "elem")))}
}), gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, cursor_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(cursor_3, "elem"))
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = func() struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
} {
					orig := cursor_3
					_ = orig
					clone := struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}{}
					clone.elem = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "elem"))
					clone.pos = gopurs_runtime.RecordGet(orig, "pos").IntVal
					return clone
				}()
goto end_branch_2
} else {

}
}
{
var __t0 struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}
{
if (gopurs_runtime.RecordGet(cursor_3, "pos").IntVal) == (idx_1) {
__t0 = struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}{gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{a_2, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()), gopurs_runtime.RecordGet(cursor_3, "pos").IntVal}
goto end_branch_0
} else {

}
}
{
__t0 = struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}{gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(cursor_3, "elem")), (gopurs_runtime.RecordGet(cursor_3, "pos").IntVal) + (int64(1))}
}
end_branch_0:
__t2 = __t0
}
end_branch_2:
return func() gopurs_runtime.Value {
				orig := __t2
				_ = orig
				return gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(orig.elem)}, gopurs_runtime.Int(orig.pos))
				}()
}), func() gopurs_runtime.Value {
				orig := struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}{gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()), int64(0)}
				_ = orig
				return gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(orig.elem)}, gopurs_runtime.Int(orig.pos))
				}()))
}

func Call_Data_Foldable_null(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
}), gopurs_runtime.Bool(true))
}

func Call_Data_Foldable_oneOf(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictPlus_1_loop *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *Constructor_Control_Plus_Plus[gopurs_runtime.Value] = dictPlus_1_loop
_ = dictPlus_1
return gopurs_runtime.Apply2(dictFoldable_0.V2, Call_Control_Alt_alt(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(dictPlus_1.V0, gopurs_runtime.Value{}))), Call_Control_Plus_empty(gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(dictPlus_1)}))
}

func Call_Data_Foldable_oneOfMap(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictPlus_1_loop *Constructor_Control_Plus_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *Constructor_Control_Plus_Plus[gopurs_runtime.Value] = dictPlus_1_loop
_ = dictPlus_1
// TAST (Let): alt_2_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar g$scope33) [(TypeVar b$scope35)]), (TypeApp (TypeVar g$scope33) [(TypeVar b$scope35)])] (TypeApp (TypeVar g$scope33) [(TypeVar b$scope35)]))
alt_2_0 := Call_Control_Alt_alt(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(dictPlus_1.V0, gopurs_runtime.Value{})))
_ = alt_2_0
// TAST (Let): empty_3_1 shape=App(Var) bindingType=(TypeApp (TypeVar g$scope33) [(TypeVar b$scope35)])
empty_3_1 := Call_Control_Plus_empty(gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(dictPlus_1)})
_ = empty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), alt_2_0, f_4), empty_3_1)
})
}

func Call_Data_Foldable_traverse_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): applySecond_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope44) [(TypeVar b$scope42)]), (TypeApp (TypeVar m$scope44) [Unit])] (TypeApp (TypeVar m$scope44) [Unit]))
applySecond_1_0 := Call_Control_Apply_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{})))
_ = applySecond_1_0
return gopurs_runtime.Func2(func(dictFoldable_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), applySecond_1_0, f_3), gopurs_runtime.Apply(dictApplicative_0.V1, Get_Data_Unit_unit()))
})
}

func Call_Data_Foldable_for_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): traverse_1_1_0 shape=App(Var) bindingType=Any
traverse_1_1_0 := Call_Data_Foldable_traverse_(dictApplicative_0)
_ = traverse_1_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
__local_var_3_1 := gopurs_runtime.Apply(traverse_1_1_0, dictFoldable_2)
_ = __local_var_3_1
return gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_3_1, a_5, b_4)
})
})
}

func Call_Data_Foldable_sequence_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictFoldable_1_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var dictFoldable_1 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_1_loop
_ = dictFoldable_1
return gopurs_runtime.Apply2(Call_Data_Foldable_traverse_(dictApplicative_0), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_1)}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Data_Foldable_foldl(dict_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Data_Foldable_indexl(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], idx_1_loop int64) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var idx_1 int64 = idx_1_loop
_ = idx_1
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v_2, "elem")))}
}), gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(cursor_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(cursor_2, "elem"))
_ = __t_tag_1
if (__t_tag_1 != nil) {
__t2 = func() struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
} {
					orig := cursor_2
					_ = orig
					clone := struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}{}
					clone.elem = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(orig, "elem"))
					clone.pos = gopurs_runtime.RecordGet(orig, "pos").IntVal
					return clone
				}()
goto end_branch_2
} else {

}
}
{
var __t0 struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}
{
if (gopurs_runtime.RecordGet(cursor_2, "pos").IntVal) == (idx_1) {
__t0 = struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}{gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{a_3, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()), gopurs_runtime.RecordGet(cursor_2, "pos").IntVal}
goto end_branch_0
} else {

}
}
{
__t0 = struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}{gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(cursor_2, "elem")), (gopurs_runtime.RecordGet(cursor_2, "pos").IntVal) + (int64(1))}
}
end_branch_0:
__t2 = __t0
}
end_branch_2:
return func() gopurs_runtime.Value {
				orig := __t2
				_ = orig
				return gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(orig.elem)}, gopurs_runtime.Int(orig.pos))
				}()
}), func() gopurs_runtime.Value {
				orig := struct{
	elem *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
	pos int64
}{gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()), int64(0)}
				_ = orig
				return gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(orig.elem)}, gopurs_runtime.Int(orig.pos))
				}()))
}

func Call_Data_Foldable_intercalate(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope74)])
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 shape=App(Var) bindingType=(TypeVar m$scope74)
mempty_3_1 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)})
_ = mempty_3_1
return gopurs_runtime.Func2(func(sep_4 gopurs_runtime.Value, xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(dictFoldable_0.V1, gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 struct{
	acc gopurs_runtime.Value
	go__init bool
}
{
if (gopurs_runtime.RecordGet(v_6, "init").IntVal) != (0) {
__t2 = struct{
	acc gopurs_runtime.Value
	go__init bool
}{v1_7, false}
goto end_branch_2
} else {

}
}
{
__t2 = struct{
	acc gopurs_runtime.Value
	go__init bool
}{gopurs_runtime.Apply2(Semigroup0_2_0.V0, gopurs_runtime.RecordGet(v_6, "acc"), gopurs_runtime.Apply2(Semigroup0_2_0.V0, sep_4, v1_7)), false}
}
end_branch_2:
return func() gopurs_runtime.Value {
				orig := __t2
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "init", orig.acc, gopurs_runtime.Bool(orig.go__init))
				}()
}), func() gopurs_runtime.Value {
				orig := struct{
	acc gopurs_runtime.Value
	go__init bool
}{mempty_3_1, true}
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "init", orig.acc, gopurs_runtime.Bool(orig.go__init))
				}(), xs_5), "acc")
})
}

func Call_Data_Foldable_intercalate__4228433826(__eta_norm_1_0_loop string, __eta_norm_0_1_loop *Constructor_Data_List_Types_Cons[string]) string {
intercalate__4228433826:
for {
if false { continue intercalate__4228433826 }
var __eta_norm_1_0 string = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_List_Types_Cons[string] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(Rebox_Data_Foldable_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 struct{
	acc gopurs_runtime.Value
	go__init bool
}
{
if (gopurs_runtime.RecordGet(v_2, "init").IntVal) != (0) {
__t0 = struct{
	acc gopurs_runtime.Value
	go__init bool
}{v1_3, false}
goto end_branch_0
} else {

}
}
{
__t0 = struct{
	acc gopurs_runtime.Value
	go__init bool
}{gopurs_runtime.Str(((gopurs_runtime.RecordGet(v_2, "acc").StrVal()) + (__eta_norm_1_0)) + (v1_3.StrVal())), false}
}
end_branch_0:
return func() gopurs_runtime.Value {
				orig := __t0
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "init", orig.acc, gopurs_runtime.Bool(orig.go__init))
				}()
}), func() gopurs_runtime.Value {
				orig := struct{
	acc gopurs_runtime.Value
	go__init bool
}{Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_1950344881_1201789390(Rebox_Data_Foldable_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString()))))}), true}
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "init", orig.acc, gopurs_runtime.Bool(orig.go__init))
				}(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_128126966_849153993(__eta_norm_0_1))}), "acc").StrVal()
}
}

func Call_Data_Foldable_length(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictSemiring_1_loop *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(c_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemiring_1.V0, dictSemiring_1.V2, c_2)
}), dictSemiring_1.V3)
}

func Call_Data_Foldable_maximumBy(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{v1_3, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_4
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_1
if (__t_tag_1 != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 uint32 = uint32(gopurs_runtime.Apply2(cmp_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3).IntVal)
_ = __t_tag_2
if (uint32(__t_tag_2) == 380165415) {
__t3 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
__t3 = v1_3
}
end_branch_3:
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{__t3, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))})
}

func Call_Data_Foldable_maximum(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope90), (TypeVar a$scope90)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := Call_Data_Ord_compare(dictOrd_0)
_ = compare_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{v1_4, true}
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
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
_ = __t_tag_2
if (__t_tag_2 != nil) {
var __t4 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(compare_1_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0, v1_4)
_ = __t_tag_3
if (uint32(__t_tag_3.IntVal) == 380165415) {
__t4 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
__t4 = v1_4
}
end_branch_4:
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{__t4, true}
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
__t5 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))})
})
}

func Call_Data_Foldable_minimumBy(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{v1_3, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_4
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_1
if (__t_tag_1 != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 uint32 = uint32(gopurs_runtime.Apply2(cmp_1, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3).IntVal)
_ = __t_tag_2
if (uint32(__t_tag_2) == 1527465420) {
__t3 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
__t3 = v1_3
}
end_branch_3:
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{__t3, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))})
}

func Call_Data_Foldable_minimum(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope101), (TypeVar a$scope101)] (ADT ["Data","Ordering","Ordering"] []))
compare_1_0 := Call_Data_Ord_compare(dictOrd_0)
_ = compare_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldl"), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
_ = __t_tag_1
if (__t_tag_1 == nil) {
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{v1_4, true}
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
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
_ = __t_tag_2
if (__t_tag_2 != nil) {
var __t4 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(compare_1_0, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0, v1_4)
_ = __t_tag_3
if (uint32(__t_tag_3.IntVal) == 1527465420) {
__t4 = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
__t4 = v1_4
}
end_branch_4:
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{__t4, true}
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
__t5 = func() *Constructor_Data_Maybe_Just[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))})
})
}

func Call_Data_Foldable_product(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictSemiring_1_loop *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, Call_Data_Semiring_mul(dictSemiring_1), dictSemiring_1.V2)
}

func Call_Data_Foldable_sum(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictSemiring_1_loop *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, Call_Data_Semiring_add(dictSemiring_1), dictSemiring_1.V3)
}

func Call_Data_Foldable_foldMapDefaultR(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope252)])
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 shape=App(Var) bindingType=(TypeVar m$scope252)
mempty_3_1 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)})
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_2_0.V0, gopurs_runtime.Apply(f_4, x_5), acc_6)
}), mempty_3_1)
})
}

func Call_Data_Foldable_foldMapDefaultL(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope296)])
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 shape=App(Var) bindingType=(TypeVar m$scope296)
mempty_3_1 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)})
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(acc_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_2_0.V0, acc_5, gopurs_runtime.Apply(f_4, x_6))
}), mempty_3_1)
})
}

func Call_Data_Foldable_foldMap(dict_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_Data_Foldable_foldMap__4139410016(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) float64 {
foldMap__4139410016:
for {
if false { continue foldMap__4139410016 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply3(Rebox_Data_Foldable_1680800814_1022383170(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_List_Types_foldableList())).V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_321927638_1201789390(Rebox_Data_Foldable_1201789390_321927638(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Interval_Duration_Iso_monoidAdditive()))))}, __eta_norm_1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_2442833393_849153993(__eta_norm_0_1))}).FloatVal()
}
}

func Call_Data_Foldable_foldableApp(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_2, v_3)
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, i_2, v_3)
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, i_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, i_2, v_3)
})}))}
}

func Call_Data_Foldable_foldableCompose(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, f_3), v_4)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2), i_3, v_4)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_0 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeVar b))
__local_var_5_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2)
_ = __local_var_5_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_5_0, a_7, b_6)
}), i_3, v_4)
})}))}
}

func Call_Data_Foldable_foldableCoproduct(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(dictMonoid_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, f_3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, f_3))
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, z_3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2, z_3))
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_2, z_3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2, z_3))
})}))}
}

func Call_Data_Foldable_foldableProduct(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_4173511203_1680800814((&Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope394)])
Semigroup0_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_0
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_3_0.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, f_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, f_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, z_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2, z_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
})})))}
}

func Call_Data_Foldable_foldlDefault(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
var Call_local_Data_Foldable_go__go_4_0_2 func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_Foldable_go__go_4_0_2
var go__go_4_0_2 gopurs_runtime.Value
_ = go__go_4_0_2
Call_local_Data_Foldable_go__go_4_0_2 = func(acc_5_loop gopurs_runtime.Value, lhs_6_loop gopurs_runtime.Value, rhs_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_0_2:
for {
if false { continue go__go_4_0_2 }
var acc_5 gopurs_runtime.Value = acc_5_loop
_ = acc_5
var lhs_6 gopurs_runtime.Value = lhs_6_loop
_ = lhs_6
var rhs_7 gopurs_runtime.Value = rhs_7_loop
_ = rhs_7
var __t5 gopurs_runtime.Value
{
if (lhs_6.Type == 9 && lhs_6.IntVal == 2421944209) {
acc_5_loop = gopurs_runtime.Apply2(c_1, acc_5, (*Constructor_Data_Foldable_Node[gopurs_runtime.Value])(lhs_6.UnsafePtr).V0)
lhs_6_loop = rhs_7
rhs_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_4_0_2
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
if (lhs_6.Type == 9 && lhs_6.IntVal == 2812549951) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_6.UnsafePtr).V1
_ = __t_tag_2
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2065045956) {
acc_5_loop = acc_5
lhs_6_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_6.UnsafePtr).V0
rhs_7_loop = rhs_7
continue go__go_4_0_2
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
var __t1 gopurs_runtime.Value
{
if (rhs_7.Type == 9 && rhs_7.IntVal == 2065045956) {
acc_5_loop = acc_5
lhs_6_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_6.UnsafePtr).V0
rhs_7_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_6.UnsafePtr).V1
continue go__go_4_0_2
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
acc_5_loop = acc_5
lhs_6_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_6.UnsafePtr).V0
rhs_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Append[gopurs_runtime.Value]{1, (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(lhs_6.UnsafePtr).V1, rhs_7}))}
continue go__go_4_0_2
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_1:
__t3 = __t1
}
end_branch_3:
__t5 = __t3
goto end_branch_5
} else {

}
}
{
if (lhs_6.Type == 9 && lhs_6.IntVal == 2065045956) {
var __t4 gopurs_runtime.Value
{
if (rhs_7.Type == 9 && rhs_7.IntVal == 2065045956) {
__t4 = acc_5
goto end_branch_4
} else {

}
}
{
acc_5_loop = acc_5
lhs_6_loop = rhs_7
rhs_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_4_0_2
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_4:
__t5 = __t4
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
go__go_4_0_2 = gopurs_runtime.Func(func(acc_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Foldable_go__go_4_0_2(acc_5_loop_val, lhs_6_loop_val, rhs_7_loop_val)
})
})
})
return Call_local_Data_Foldable_go__go_4_0_2(u_2, gopurs_runtime.Apply3(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Foldable_monoidFreeMonoidTree()))}, Get_Data_Foldable_Node(), xs_3), gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
}

func Call_Data_Foldable_foldrDefault(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
var Call_local_Data_Foldable_go__go_4_0_3 func(gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value) gopurs_runtime.Value
_ = Call_local_Data_Foldable_go__go_4_0_3
var go__go_4_0_3 gopurs_runtime.Value
_ = go__go_4_0_3
Call_local_Data_Foldable_go__go_4_0_3 = func(acc_5_loop gopurs_runtime.Value, lhs_6_loop gopurs_runtime.Value, rhs_7_loop gopurs_runtime.Value) gopurs_runtime.Value {
go__go_4_0_3:
for {
if false { continue go__go_4_0_3 }
var acc_5 gopurs_runtime.Value = acc_5_loop
_ = acc_5
var lhs_6 gopurs_runtime.Value = lhs_6_loop
_ = lhs_6
var rhs_7 gopurs_runtime.Value = rhs_7_loop
_ = rhs_7
var __t5 gopurs_runtime.Value
{
if (rhs_7.Type == 9 && rhs_7.IntVal == 2421944209) {
acc_5_loop = gopurs_runtime.Apply2(c_1, (*Constructor_Data_Foldable_Node[gopurs_runtime.Value])(rhs_7.UnsafePtr).V0, acc_5)
lhs_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_7_loop = lhs_6
continue go__go_4_0_3
__t5 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_5
} else {

}
}
{
if (rhs_7.Type == 9 && rhs_7.IntVal == 2812549951) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_7.UnsafePtr).V0
_ = __t_tag_2
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2065045956) {
acc_5_loop = acc_5
lhs_6_loop = lhs_6
rhs_7_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_7.UnsafePtr).V1
continue go__go_4_0_3
__t3 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_3
} else {

}
}
{
var __t1 gopurs_runtime.Value
{
if (lhs_6.Type == 9 && lhs_6.IntVal == 2065045956) {
acc_5_loop = acc_5
lhs_6_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_7.UnsafePtr).V0
rhs_7_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_7.UnsafePtr).V1
continue go__go_4_0_3
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_1
} else {

}
}
{
acc_5_loop = acc_5
lhs_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Append[gopurs_runtime.Value]{1, lhs_6, (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_7.UnsafePtr).V0}))}
rhs_7_loop = (*Constructor_Data_Foldable_Append[gopurs_runtime.Value])(rhs_7.UnsafePtr).V1
continue go__go_4_0_3
__t1 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_1:
__t3 = __t1
}
end_branch_3:
__t5 = __t3
goto end_branch_5
} else {

}
}
{
if (rhs_7.Type == 9 && rhs_7.IntVal == 2065045956) {
var __t4 gopurs_runtime.Value
{
if (lhs_6.Type == 9 && lhs_6.IntVal == 2065045956) {
__t4 = acc_5
goto end_branch_4
} else {

}
}
{
acc_5_loop = acc_5
lhs_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_7_loop = lhs_6
continue go__go_4_0_3
__t4 = func() gopurs_runtime.Value { panic("unreachable") }()
}
end_branch_4:
__t5 = __t4
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
go__go_4_0_3 = gopurs_runtime.Func(func(acc_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Foldable_go__go_4_0_3(acc_5_loop_val, lhs_6_loop_val, rhs_7_loop_val)
})
})
})
return Call_local_Data_Foldable_go__go_4_0_3(u_2, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Apply3(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Foldable_monoidFreeMonoidTree()))}, Get_Data_Foldable_Node(), xs_3))
}

func Call_Data_Foldable_lookup(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictEq_1_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictEq_1 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_1_loop
_ = dictEq_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply2(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_1089607855_1201789390(Rebox_Data_Foldable_1201789390_1089607855(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Maybe_First_monoidFirst()))))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.Apply2(dictEq_1.V0, a_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}())
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})))
}

func Call_Data_Foldable_surroundMap(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value], d_2_loop gopurs_runtime.Value, t_3_loop gopurs_runtime.Value, f_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
var t_3 gopurs_runtime.Value = t_3_loop
_ = t_3
var f_4 gopurs_runtime.Value = f_4_loop
_ = f_4
return gopurs_runtime.Apply2(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply3(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemigroup_1.V0, d_2, gopurs_runtime.Apply2(dictSemigroup_1.V0, gopurs_runtime.Apply(t_3, a_5), m_6))
}), f_4), d_2)
}

func Call_Data_Foldable_surround(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value], d_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
return gopurs_runtime.Apply4(Get_Data_Foldable_surroundMap(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_1)}, d_2, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Data_Foldable_foldM(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictMonad_1_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonad_1 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_1_loop
_ = dictMonad_1
// TAST (Let): Bind1_2_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope438)])
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_1.V1, gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): Applicative0_3_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope438)])
Applicative0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_1.V0, gopurs_runtime.Value{}))
_ = Applicative0_3_1
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, b0_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_0.V1, b_6, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_4, a_8, a_7)
}))
}), gopurs_runtime.Apply(Applicative0_3_1.V1, b0_5))
})
}

func Call_Data_Foldable_fold(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply2(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Data_Foldable_fold__1068017326(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
fold__1068017326:
for {
if false { continue fold__1068017326 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Apply2(Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Interval_Duration_Iso_monoidFn())), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), __eta_norm_0_0)
}
}

func Call_Data_Foldable_findMap(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(p_1, v1_3))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))})
}

func Call_Data_Foldable_find(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
_ = __t_tag_0
if ((__t_tag_0 == nil)) && ((gopurs_runtime.Apply(p_1, v1_3).IntVal) != (0)) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{v1_3, true}
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
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))})
}

func Call_Data_Foldable_any(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): foldMap1_1_0 shape=App(Var) bindingType=Any
foldMap1_1_0 := Call_Data_Foldable_foldMap(dictFoldable_0)
_ = foldMap1_1_0
return gopurs_runtime.Func(func(dictHeytingAlgebra_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply(foldMap1_1_0, Call_Data_Monoid_Disj_monoidDisj(dictHeytingAlgebra_2)))
})
}

func Call_Data_Foldable_elem(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): any1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope468)] Boolean), (TypeApp (TypeVar f$scope469) [(TypeVar a$scope468)])] Boolean)
any1_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_any(dictFoldable_0), gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_3591112874_2663347022(Rebox_Data_Foldable_2663347022_3591112874(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Get_Data_HeytingAlgebra_heytingAlgebraBoolean()))))})
_ = any1_1_0
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), any1_1_0, Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_2)))
})
}

func Call_Data_Foldable_notElem(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictEq_1_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictEq_1 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_1_loop
_ = dictEq_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_HeytingAlgebra_not(Rebox_Data_Foldable_3591112874_2663347022(Rebox_Data_Foldable_2663347022_3591112874(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Get_Data_HeytingAlgebra_heytingAlgebraBoolean())))), gopurs_runtime.Apply2(Call_Data_Foldable_elem(dictFoldable_0), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_1)}, x_2))
}

func Call_Data_Foldable_or(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
return gopurs_runtime.Apply2(Call_Data_Foldable_any(dictFoldable_0), gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(dictHeytingAlgebra_1)}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Call_Data_Foldable_all(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): foldMap1_1_0 shape=App(Var) bindingType=Any
foldMap1_1_0 := Call_Data_Foldable_foldMap(dictFoldable_0)
_ = foldMap1_1_0
return gopurs_runtime.Func(func(dictHeytingAlgebra_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply(foldMap1_1_0, Call_Data_Monoid_Conj_monoidConj(dictHeytingAlgebra_2)))
})
}

func Call_Data_Foldable_and(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
return gopurs_runtime.Apply2(Call_Data_Foldable_all(dictFoldable_0), gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(dictHeytingAlgebra_1)}, Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))
}

func Rebox_Data_Foldable_1089607855_1201789390(in *Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(in.V1)}
	return out
}

func Rebox_Data_Foldable_1146820559_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Foldable_1201789390_1089607855(in *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) *Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](in.V1)
	return out
}

func Rebox_Data_Foldable_1201789390_1950344881(in *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) *Constructor_Data_Monoid_Monoid[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[string]{}
		out.V0 = in.V0
		out.V1 = in.V1.StrVal()
	return out
}

func Rebox_Data_Foldable_1201789390_321927638(in *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) *Constructor_Data_Monoid_Monoid[float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[float64]{}
		out.V0 = in.V0
		out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Foldable_128126966_849153993(in *Constructor_Data_List_Types_Cons[string]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Str(in.V0)
		out.V1 = Rebox_Data_Foldable_128126966_849153993(in.V1)
	return out
}

func Rebox_Data_Foldable_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
		out.V0 = uint32(in.V0.IntVal)
		out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Foldable_1680800814_1022383170(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Foldable_1950344881_1201789390(in *Constructor_Data_Monoid_Monoid[string]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Data_Foldable_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Foldable_3132786365_138441832(in.V0))}
		out.V1 = Rebox_Data_Foldable_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_Foldable_2663347022_3591112874(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool] {
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

func Rebox_Data_Foldable_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
		out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Foldable_321927638_1201789390(in *Constructor_Data_Monoid_Monoid[float64]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Foldable_3591112874_2663347022(in *Constructor_Data_HeytingAlgebra_HeytingAlgebra[bool]) *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] {
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

func Rebox_Data_Foldable_4173511203_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Foldable_849153993_128126966(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[string] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[string]{}
		out.V0 = in.V0.StrVal()
		out.V1 = Rebox_Data_Foldable_849153993_128126966(in.V1)
	return out
}

func Rebox_Data_Foldable_849153993_2442833393(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
		out.V0 = Rebox_Data_Foldable_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
		out.V1 = Rebox_Data_Foldable_849153993_2442833393(in.V1)
	return out
}

func Get_Data_Foldable_foldlArray() gopurs_runtime.Value {
	return _Gopurs_Data_Foldable_FoldlArray
}

func Get_Data_Foldable_foldrArray() gopurs_runtime.Value {
	return _Gopurs_Data_Foldable_FoldrArray
}

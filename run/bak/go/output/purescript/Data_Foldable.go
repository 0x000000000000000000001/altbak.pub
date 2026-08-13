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
		cache_Data_Foldable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_identity(x_0_box)
})
	})
	return cache_Data_Foldable_identity
}

var cache_Data_Foldable_unwrap gopurs_runtime.Value
var once_Data_Foldable_unwrap sync.Once
func Get_Data_Foldable_unwrap() gopurs_runtime.Value {
	once_Data_Foldable_unwrap.Do(func() {
		cache_Data_Foldable_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Foldable_unwrap
}

var cache_Data_Foldable_monoidEndo gopurs_runtime.Value
var once_Data_Foldable_monoidEndo sync.Once
func Get_Data_Foldable_monoidEndo() gopurs_runtime.Value {
	once_Data_Foldable_monoidEndo.Do(func() {
		cache_Data_Foldable_monoidEndo = func() gopurs_runtime.Value {
// TAST (Let): semigroupEndo1_0_0 -> gopurs_runtime.Value
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(v1_1, x_2))
})
})
}))
_ = semigroupEndo1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))))}
}()
	})
	return cache_Data_Foldable_monoidEndo
}

var cache_Data_Foldable_identity1 gopurs_runtime.Value
var once_Data_Foldable_identity1 sync.Once
func Get_Data_Foldable_identity1() gopurs_runtime.Value {
	once_Data_Foldable_identity1.Do(func() {
		cache_Data_Foldable_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_identity1(x_0_box)
})
	})
	return cache_Data_Foldable_identity1
}

var cache_Data_Foldable_not gopurs_runtime.Value
var once_Data_Foldable_not sync.Once
func Get_Data_Foldable_not() gopurs_runtime.Value {
	once_Data_Foldable_not.Do(func() {
		cache_Data_Foldable_not = Get_Data_HeytingAlgebra_boolNot()
	})
	return cache_Data_Foldable_not
}

var cache_Data_Foldable_identity2 gopurs_runtime.Value
var once_Data_Foldable_identity2 sync.Once
func Get_Data_Foldable_identity2() gopurs_runtime.Value {
	once_Data_Foldable_identity2.Do(func() {
		cache_Data_Foldable_identity2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_identity2(x_0_box)
})
	})
	return cache_Data_Foldable_identity2
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
return gopurs_runtime.Value{Type: 9, IntVal: 2421944209, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Node{1, value0})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Append{1, value0, value1})}
})
})
	})
	return cache_Data_Foldable_Append
}

var cache_Data_Foldable_Foldable_dollarDict gopurs_runtime.Value
var once_Data_Foldable_Foldable_dollarDict sync.Once
func Get_Data_Foldable_Foldable_dollarDict() gopurs_runtime.Value {
	once_Data_Foldable_Foldable_dollarDict.Do(func() {
		cache_Data_Foldable_Foldable_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_Foldable_dollarDict(x_0_box)
})
	})
	return cache_Data_Foldable_Foldable_dollarDict
}

var cache_Data_Foldable_semigroupFreeMonoidTree gopurs_runtime.Value
var once_Data_Foldable_semigroupFreeMonoidTree sync.Once
func Get_Data_Foldable_semigroupFreeMonoidTree() gopurs_runtime.Value {
	once_Data_Foldable_semigroupFreeMonoidTree.Do(func() {
		cache_Data_Foldable_semigroupFreeMonoidTree = gopurs_runtime.RecordDict1("append", Get_Data_Foldable_Append())
	})
	return cache_Data_Foldable_semigroupFreeMonoidTree
}

var cache_Data_Foldable_monoidFreeMonoidTree gopurs_runtime.Value
var once_Data_Foldable_monoidFreeMonoidTree sync.Once
func Get_Data_Foldable_monoidFreeMonoidTree() gopurs_runtime.Value {
	once_Data_Foldable_monoidFreeMonoidTree.Do(func() {
		cache_Data_Foldable_monoidFreeMonoidTree = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_semigroupFreeMonoidTree()
}), gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
	})
	return cache_Data_Foldable_monoidFreeMonoidTree
}

var cache_Data_Foldable_foldr gopurs_runtime.Value
var once_Data_Foldable_foldr sync.Once
func Get_Data_Foldable_foldr() gopurs_runtime.Value {
	once_Data_Foldable_foldr.Do(func() {
		cache_Data_Foldable_foldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr
}

var cache_Data_Foldable_indexr gopurs_runtime.Value
var once_Data_Foldable_indexr sync.Once
func Get_Data_Foldable_indexr() gopurs_runtime.Value {
	once_Data_Foldable_indexr.Do(func() {
		cache_Data_Foldable_indexr = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, idx_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_indexr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), idx_1_box.IntVal)
})
	})
	return cache_Data_Foldable_indexr
}

var cache_Data_Foldable_null gopurs_runtime.Value
var once_Data_Foldable_null sync.Once
func Get_Data_Foldable_null() gopurs_runtime.Value {
	once_Data_Foldable_null.Do(func() {
		cache_Data_Foldable_null = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_null(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_Foldable_null
}

var cache_Data_Foldable_oneOf gopurs_runtime.Value
var once_Data_Foldable_oneOf sync.Once
func Get_Data_Foldable_oneOf() gopurs_runtime.Value {
	once_Data_Foldable_oneOf.Do(func() {
		cache_Data_Foldable_oneOf = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](dictPlus_1_box))
})
	})
	return cache_Data_Foldable_oneOf
}

var cache_Data_Foldable_oneOfMap gopurs_runtime.Value
var once_Data_Foldable_oneOfMap sync.Once
func Get_Data_Foldable_oneOfMap() gopurs_runtime.Value {
	once_Data_Foldable_oneOfMap.Do(func() {
		cache_Data_Foldable_oneOfMap = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_oneOfMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](dictPlus_1_box))
})
	})
	return cache_Data_Foldable_oneOfMap
}

var cache_Data_Foldable_traverse_ gopurs_runtime.Value
var once_Data_Foldable_traverse_ sync.Once
func Get_Data_Foldable_traverse_() gopurs_runtime.Value {
	once_Data_Foldable_traverse_.Do(func() {
		cache_Data_Foldable_traverse_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_traverse_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Data_Foldable_traverse_
}

var cache_Data_Foldable_for_ gopurs_runtime.Value
var once_Data_Foldable_for_ sync.Once
func Get_Data_Foldable_for_() gopurs_runtime.Value {
	once_Data_Foldable_for_.Do(func() {
		cache_Data_Foldable_for_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_for_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Data_Foldable_for_
}

var cache_Data_Foldable_sequence_ gopurs_runtime.Value
var once_Data_Foldable_sequence_ sync.Once
func Get_Data_Foldable_sequence_() gopurs_runtime.Value {
	once_Data_Foldable_sequence_.Do(func() {
		cache_Data_Foldable_sequence_ = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_sequence_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_1_box))
})
	})
	return cache_Data_Foldable_sequence_
}

var cache_Data_Foldable_foldl gopurs_runtime.Value
var once_Data_Foldable_foldl sync.Once
func Get_Data_Foldable_foldl() gopurs_runtime.Value {
	once_Data_Foldable_foldl.Do(func() {
		cache_Data_Foldable_foldl = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl
}

var cache_Data_Foldable_indexl gopurs_runtime.Value
var once_Data_Foldable_indexl sync.Once
func Get_Data_Foldable_indexl() gopurs_runtime.Value {
	once_Data_Foldable_indexl.Do(func() {
		cache_Data_Foldable_indexl = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, idx_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_indexl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), idx_1_box.IntVal)
})
	})
	return cache_Data_Foldable_indexl
}

var cache_Data_Foldable_intercalate gopurs_runtime.Value
var once_Data_Foldable_intercalate sync.Once
func Get_Data_Foldable_intercalate() gopurs_runtime.Value {
	once_Data_Foldable_intercalate.Do(func() {
		cache_Data_Foldable_intercalate = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1_box))
})
	})
	return cache_Data_Foldable_intercalate
}

var cache_Data_Foldable_length gopurs_runtime.Value
var once_Data_Foldable_length sync.Once
func Get_Data_Foldable_length() gopurs_runtime.Value {
	once_Data_Foldable_length.Do(func() {
		cache_Data_Foldable_length = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_length(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_length
}

var cache_Data_Foldable_maximumBy gopurs_runtime.Value
var once_Data_Foldable_maximumBy sync.Once
func Get_Data_Foldable_maximumBy() gopurs_runtime.Value {
	once_Data_Foldable_maximumBy.Do(func() {
		cache_Data_Foldable_maximumBy = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_maximumBy(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), cmp_1_box)
})
	})
	return cache_Data_Foldable_maximumBy
}

var cache_Data_Foldable_maximum gopurs_runtime.Value
var once_Data_Foldable_maximum sync.Once
func Get_Data_Foldable_maximum() gopurs_runtime.Value {
	once_Data_Foldable_maximum.Do(func() {
		cache_Data_Foldable_maximum = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_maximum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_1_box))
})
	})
	return cache_Data_Foldable_maximum
}

var cache_Data_Foldable_minimumBy gopurs_runtime.Value
var once_Data_Foldable_minimumBy sync.Once
func Get_Data_Foldable_minimumBy() gopurs_runtime.Value {
	once_Data_Foldable_minimumBy.Do(func() {
		cache_Data_Foldable_minimumBy = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_minimumBy(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), cmp_1_box)
})
	})
	return cache_Data_Foldable_minimumBy
}

var cache_Data_Foldable_minimum gopurs_runtime.Value
var once_Data_Foldable_minimum sync.Once
func Get_Data_Foldable_minimum() gopurs_runtime.Value {
	once_Data_Foldable_minimum.Do(func() {
		cache_Data_Foldable_minimum = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_minimum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_1_box))
})
	})
	return cache_Data_Foldable_minimum
}

var cache_Data_Foldable_product gopurs_runtime.Value
var once_Data_Foldable_product sync.Once
func Get_Data_Foldable_product() gopurs_runtime.Value {
	once_Data_Foldable_product.Do(func() {
		cache_Data_Foldable_product = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_product(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_product
}

var cache_Data_Foldable_sum gopurs_runtime.Value
var once_Data_Foldable_sum sync.Once
func Get_Data_Foldable_sum() gopurs_runtime.Value {
	once_Data_Foldable_sum.Do(func() {
		cache_Data_Foldable_sum = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_sum(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_sum
}

var cache_Data_Foldable_foldableTuple gopurs_runtime.Value
var once_Data_Foldable_foldableTuple sync.Once
func Get_Data_Foldable_foldableTuple() gopurs_runtime.Value {
	once_Data_Foldable_foldableTuple.Do(func() {
		cache_Data_Foldable_foldableTuple = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableTuple
}

var cache_Data_Foldable_foldableMultiplicative gopurs_runtime.Value
var once_Data_Foldable_foldableMultiplicative sync.Once
func Get_Data_Foldable_foldableMultiplicative() gopurs_runtime.Value {
	once_Data_Foldable_foldableMultiplicative.Do(func() {
		cache_Data_Foldable_foldableMultiplicative = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableMultiplicative
}

var cache_Data_Foldable_foldableMaybe gopurs_runtime.Value
var once_Data_Foldable_foldableMaybe sync.Once
func Get_Data_Foldable_foldableMaybe() gopurs_runtime.Value {
	once_Data_Foldable_foldableMaybe.Do(func() {
		cache_Data_Foldable_foldableMaybe = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t1 = mempty_1_0
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just)(v1_3.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t2 = v1_1
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Maybe_Just)(v2_2.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t3 = v1_1
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(v_0, (*Constructor_Data_Maybe_Just)(v2_2.UnsafePtr).V0, v1_1)
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
}))
	})
	return cache_Data_Foldable_foldableMaybe
}

var cache_Data_Foldable_foldableIdentity gopurs_runtime.Value
var once_Data_Foldable_foldableIdentity sync.Once
func Get_Data_Foldable_foldableIdentity() gopurs_runtime.Value {
	once_Data_Foldable_foldableIdentity.Do(func() {
		cache_Data_Foldable_foldableIdentity = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableIdentity
}

var cache_Data_Foldable_foldableEither gopurs_runtime.Value
var once_Data_Foldable_foldableEither sync.Once
func Get_Data_Foldable_foldableEither() gopurs_runtime.Value {
	once_Data_Foldable_foldableEither.Do(func() {
		cache_Data_Foldable_foldableEither = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply(v_2, (*Constructor_Data_Either_Right)(v1_3.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Either_Right)(v2_2.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t3 = gopurs_runtime.Apply2(v_0, (*Constructor_Data_Either_Right)(v2_2.UnsafePtr).V0, v1_1)
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
}))
	})
	return cache_Data_Foldable_foldableEither
}

var cache_Data_Foldable_foldableDual gopurs_runtime.Value
var once_Data_Foldable_foldableDual sync.Once
func Get_Data_Foldable_foldableDual() gopurs_runtime.Value {
	once_Data_Foldable_foldableDual.Do(func() {
		cache_Data_Foldable_foldableDual = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableDual
}

var cache_Data_Foldable_foldableDisj gopurs_runtime.Value
var once_Data_Foldable_foldableDisj sync.Once
func Get_Data_Foldable_foldableDisj() gopurs_runtime.Value {
	once_Data_Foldable_foldableDisj.Do(func() {
		cache_Data_Foldable_foldableDisj = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableDisj
}

var cache_Data_Foldable_foldableConst gopurs_runtime.Value
var once_Data_Foldable_foldableConst sync.Once
func Get_Data_Foldable_foldableConst() gopurs_runtime.Value {
	once_Data_Foldable_foldableConst.Do(func() {
		cache_Data_Foldable_foldableConst = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})
})
}))
	})
	return cache_Data_Foldable_foldableConst
}

var cache_Data_Foldable_foldableConj gopurs_runtime.Value
var once_Data_Foldable_foldableConj sync.Once
func Get_Data_Foldable_foldableConj() gopurs_runtime.Value {
	once_Data_Foldable_foldableConj.Do(func() {
		cache_Data_Foldable_foldableConj = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableConj
}

var cache_Data_Foldable_foldableAdditive gopurs_runtime.Value
var once_Data_Foldable_foldableAdditive sync.Once
func Get_Data_Foldable_foldableAdditive() gopurs_runtime.Value {
	once_Data_Foldable_foldableAdditive.Do(func() {
		cache_Data_Foldable_foldableAdditive = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableAdditive
}

var cache_Data_Foldable_foldMapDefaultR gopurs_runtime.Value
var once_Data_Foldable_foldMapDefaultR sync.Once
func Get_Data_Foldable_foldMapDefaultR() gopurs_runtime.Value {
	once_Data_Foldable_foldMapDefaultR.Do(func() {
		cache_Data_Foldable_foldMapDefaultR = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1_box))
})
	})
	return cache_Data_Foldable_foldMapDefaultR
}

var cache_Data_Foldable_foldableArray gopurs_runtime.Value
var once_Data_Foldable_foldableArray sync.Once
func Get_Data_Foldable_foldableArray() gopurs_runtime.Value {
	once_Data_Foldable_foldableArray.Do(func() {
		cache_Data_Foldable_foldableArray = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), Get_Data_Foldable_foldlArray(), Get_Data_Foldable_foldrArray())
	})
	return cache_Data_Foldable_foldableArray
}

var cache_Data_Foldable_foldableFreeMonoidTree gopurs_runtime.Value
var once_Data_Foldable_foldableFreeMonoidTree sync.Once
func Get_Data_Foldable_foldableFreeMonoidTree() gopurs_runtime.Value {
	once_Data_Foldable_foldableFreeMonoidTree.Do(func() {
		cache_Data_Foldable_foldableFreeMonoidTree = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFreeMonoidTree(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_0 gopurs_runtime.Value
go__go_1_2_0 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_2_0:
for {
if false { continue go__go_1_2_0 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t7 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, acc_2, (*Constructor_Data_Foldable_Node)(lhs_3.UnsafePtr).V0)
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_0
__t7 = gopurs_runtime.Value{}
goto end_branch_7
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2812549951) {
var __t5 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V0
rhs_4_loop = rhs_4
continue go__go_1_2_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
var __t3 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V0
rhs_4_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V1
continue go__go_1_2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V0
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Append{1, (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V1, rhs_4})}
continue go__go_1_2_0
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t5 = __t3
}
end_branch_5:
__t7 = __t5
goto end_branch_7
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
var __t6 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
__t6 = acc_2
goto end_branch_6
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_0
__t6 = gopurs_runtime.Value{}
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
}()
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_2_0, a_2, b_3, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
})
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_8_1 gopurs_runtime.Value
go__go_1_8_1 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_8_1:
for {
if false { continue go__go_1_8_1 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t13 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, (*Constructor_Data_Foldable_Node)(rhs_4.UnsafePtr).V0, acc_2)
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_1
__t13 = gopurs_runtime.Value{}
goto end_branch_13
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2812549951) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V1
continue go__go_1_8_1
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
var __t9 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V0
rhs_4_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V1
continue go__go_1_8_1
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Append{1, lhs_3, (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V0})}
rhs_4_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V1
continue go__go_1_8_1
__t9 = gopurs_runtime.Value{}
}
end_branch_9:
__t11 = __t9
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
var __t12 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
__t12 = acc_2
goto end_branch_12
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_1
__t12 = gopurs_runtime.Value{}
}
end_branch_12:
__t13 = __t12
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
}()
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_8_1, a_2, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}, b_3)
})
})
}))
	})
	return cache_Data_Foldable_foldableFreeMonoidTree
}

var cache_Data_Foldable_foldMapDefaultL gopurs_runtime.Value
var once_Data_Foldable_foldMapDefaultL sync.Once
func Get_Data_Foldable_foldMapDefaultL() gopurs_runtime.Value {
	once_Data_Foldable_foldMapDefaultL.Do(func() {
		cache_Data_Foldable_foldMapDefaultL = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMapDefaultL(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1_box))
})
	})
	return cache_Data_Foldable_foldMapDefaultL
}

var cache_Data_Foldable_foldMap gopurs_runtime.Value
var once_Data_Foldable_foldMap sync.Once
func Get_Data_Foldable_foldMap() gopurs_runtime.Value {
	once_Data_Foldable_foldMap.Do(func() {
		cache_Data_Foldable_foldMap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap
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
		cache_Data_Foldable_foldableFirst = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_3 == nil) {
__t5 = z_1
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_6 == nil) {
__t8 = z_1
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0, z_1)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
})
}))
	})
	return cache_Data_Foldable_foldableFirst
}

var cache_Data_Foldable_foldableLast gopurs_runtime.Value
var once_Data_Foldable_foldableLast sync.Once
func Get_Data_Foldable_foldableLast() gopurs_runtime.Value {
	once_Data_Foldable_foldableLast.Do(func() {
		cache_Data_Foldable_foldableLast = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_3 == nil) {
__t5 = z_1
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_6 == nil) {
__t8 = z_1
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0, z_1)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
})
}))
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
return Call_Data_Foldable_foldlDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_Data_Foldable_foldlDefault
}

var cache_Data_Foldable_foldrDefault gopurs_runtime.Value
var once_Data_Foldable_foldrDefault sync.Once
func Get_Data_Foldable_foldrDefault() gopurs_runtime.Value {
	once_Data_Foldable_foldrDefault.Do(func() {
		cache_Data_Foldable_foldrDefault = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_Data_Foldable_foldrDefault
}

var cache_Data_Foldable_lookup gopurs_runtime.Value
var once_Data_Foldable_lookup sync.Once
func Get_Data_Foldable_lookup() gopurs_runtime.Value {
	once_Data_Foldable_lookup.Do(func() {
		cache_Data_Foldable_lookup = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_lookup(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_1_box), a_2_box)
})
	})
	return cache_Data_Foldable_lookup
}

var cache_Data_Foldable_surroundMap gopurs_runtime.Value
var once_Data_Foldable_surroundMap sync.Once
func Get_Data_Foldable_surroundMap() gopurs_runtime.Value {
	once_Data_Foldable_surroundMap.Do(func() {
		cache_Data_Foldable_surroundMap = gopurs_runtime.Func5(func(dictFoldable_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value, t_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_surroundMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_1_box), d_2_box, t_3_box, f_4_box)
})
	})
	return cache_Data_Foldable_surroundMap
}

var cache_Data_Foldable_surround gopurs_runtime.Value
var once_Data_Foldable_surround sync.Once
func Get_Data_Foldable_surround() gopurs_runtime.Value {
	once_Data_Foldable_surround.Do(func() {
		cache_Data_Foldable_surround = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_surround(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_1_box), d_2_box, f_3_box)
})
	})
	return cache_Data_Foldable_surround
}

var cache_Data_Foldable_foldM gopurs_runtime.Value
var once_Data_Foldable_foldM sync.Once
func Get_Data_Foldable_foldM() gopurs_runtime.Value {
	once_Data_Foldable_foldM.Do(func() {
		cache_Data_Foldable_foldM = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldM(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_1_box))
})
	})
	return cache_Data_Foldable_foldM
}

var cache_Data_Foldable_fold gopurs_runtime.Value
var once_Data_Foldable_fold sync.Once
func Get_Data_Foldable_fold() gopurs_runtime.Value {
	once_Data_Foldable_fold.Do(func() {
		cache_Data_Foldable_fold = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_fold(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1_box))
})
	})
	return cache_Data_Foldable_fold
}

var cache_Data_Foldable_findMap gopurs_runtime.Value
var once_Data_Foldable_findMap sync.Once
func Get_Data_Foldable_findMap() gopurs_runtime.Value {
	once_Data_Foldable_findMap.Do(func() {
		cache_Data_Foldable_findMap = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_findMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), p_1_box)
})
	})
	return cache_Data_Foldable_findMap
}

var cache_Data_Foldable_find gopurs_runtime.Value
var once_Data_Foldable_find sync.Once
func Get_Data_Foldable_find() gopurs_runtime.Value {
	once_Data_Foldable_find.Do(func() {
		cache_Data_Foldable_find = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_find(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), p_1_box)
})
	})
	return cache_Data_Foldable_find
}

var cache_Data_Foldable_any gopurs_runtime.Value
var once_Data_Foldable_any sync.Once
func Get_Data_Foldable_any() gopurs_runtime.Value {
	once_Data_Foldable_any.Do(func() {
		cache_Data_Foldable_any = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_any(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dictHeytingAlgebra_1_box))
})
	})
	return cache_Data_Foldable_any
}

var cache_Data_Foldable_elem gopurs_runtime.Value
var once_Data_Foldable_elem sync.Once
func Get_Data_Foldable_elem() gopurs_runtime.Value {
	once_Data_Foldable_elem.Do(func() {
		cache_Data_Foldable_elem = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_elem(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_Foldable_elem
}

var cache_Data_Foldable_notElem gopurs_runtime.Value
var once_Data_Foldable_notElem sync.Once
func Get_Data_Foldable_notElem() gopurs_runtime.Value {
	once_Data_Foldable_notElem.Do(func() {
		cache_Data_Foldable_notElem = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_notElem(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_1_box), x_2_box)
})
	})
	return cache_Data_Foldable_notElem
}

var cache_Data_Foldable_or gopurs_runtime.Value
var once_Data_Foldable_or sync.Once
func Get_Data_Foldable_or() gopurs_runtime.Value {
	once_Data_Foldable_or.Do(func() {
		cache_Data_Foldable_or = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_or(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dictHeytingAlgebra_1_box))
})
	})
	return cache_Data_Foldable_or
}

var cache_Data_Foldable_all gopurs_runtime.Value
var once_Data_Foldable_all sync.Once
func Get_Data_Foldable_all() gopurs_runtime.Value {
	once_Data_Foldable_all.Do(func() {
		cache_Data_Foldable_all = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_all(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dictHeytingAlgebra_1_box))
})
	})
	return cache_Data_Foldable_all
}

var cache_Data_Foldable_and gopurs_runtime.Value
var once_Data_Foldable_and sync.Once
func Get_Data_Foldable_and() gopurs_runtime.Value {
	once_Data_Foldable_and.Do(func() {
		cache_Data_Foldable_and = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_and(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dictHeytingAlgebra_1_box))
})
	})
	return cache_Data_Foldable_and
}

var cache_Data_Foldable_all__4179648253 gopurs_runtime.Value
var once_Data_Foldable_all__4179648253 sync.Once
func Get_Data_Foldable_all__4179648253() gopurs_runtime.Value {
	once_Data_Foldable_all__4179648253.Do(func() {
		cache_Data_Foldable_all__4179648253 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_all__4179648253(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dictHeytingAlgebra_1_box))
})
	})
	return cache_Data_Foldable_all__4179648253
}

var cache_Data_Foldable_any__4179648253 gopurs_runtime.Value
var once_Data_Foldable_any__4179648253 sync.Once
func Get_Data_Foldable_any__4179648253() gopurs_runtime.Value {
	once_Data_Foldable_any__4179648253.Do(func() {
		cache_Data_Foldable_any__4179648253 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_any__4179648253(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dictHeytingAlgebra_1_box))
})
	})
	return cache_Data_Foldable_any__4179648253
}

var cache_Data_Foldable_any__4041742601 gopurs_runtime.Value
var once_Data_Foldable_any__4041742601 sync.Once
func Get_Data_Foldable_any__4041742601() gopurs_runtime.Value {
	once_Data_Foldable_any__4041742601.Do(func() {
		cache_Data_Foldable_any__4041742601 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_any__4041742601(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Foldable_any__4041742601
}

var cache_Data_Foldable_any__1385259145 gopurs_runtime.Value
var once_Data_Foldable_any__1385259145 sync.Once
func Get_Data_Foldable_any__1385259145() gopurs_runtime.Value {
	once_Data_Foldable_any__1385259145.Do(func() {
		cache_Data_Foldable_any__1385259145 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_any__1385259145(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Foldable_any__1385259145
}

var cache_Data_Foldable_any__842931401 gopurs_runtime.Value
var once_Data_Foldable_any__842931401 sync.Once
func Get_Data_Foldable_any__842931401() gopurs_runtime.Value {
	once_Data_Foldable_any__842931401.Do(func() {
		cache_Data_Foldable_any__842931401 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_any__842931401(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Foldable_any__842931401
}

var cache_Data_Foldable_elem__2343844090 gopurs_runtime.Value
var once_Data_Foldable_elem__2343844090 sync.Once
func Get_Data_Foldable_elem__2343844090() gopurs_runtime.Value {
	once_Data_Foldable_elem__2343844090.Do(func() {
		cache_Data_Foldable_elem__2343844090 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_elem__2343844090(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_Foldable_elem__2343844090
}

var cache_Data_Foldable_fold__910331789 gopurs_runtime.Value
var once_Data_Foldable_fold__910331789 sync.Once
func Get_Data_Foldable_fold__910331789() gopurs_runtime.Value {
	once_Data_Foldable_fold__910331789.Do(func() {
		cache_Data_Foldable_fold__910331789 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_fold__910331789(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1_box))
})
	})
	return cache_Data_Foldable_fold__910331789
}

var cache_Data_Foldable_foldMap__4098395794 gopurs_runtime.Value
var once_Data_Foldable_foldMap__4098395794 sync.Once
func Get_Data_Foldable_foldMap__4098395794() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__4098395794.Do(func() {
		cache_Data_Foldable_foldMap__4098395794 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__4098395794(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap__4098395794
}

var cache_Data_Foldable_foldMap__4151846418 gopurs_runtime.Value
var once_Data_Foldable_foldMap__4151846418 sync.Once
func Get_Data_Foldable_foldMap__4151846418() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__4151846418.Do(func() {
		cache_Data_Foldable_foldMap__4151846418 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__4151846418(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap__4151846418
}

var cache_Data_Foldable_foldMap__2966595236 gopurs_runtime.Value
var once_Data_Foldable_foldMap__2966595236 sync.Once
func Get_Data_Foldable_foldMap__2966595236() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__2966595236.Do(func() {
		cache_Data_Foldable_foldMap__2966595236 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__2966595236(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap__2966595236
}

var cache_Data_Foldable_foldMap__4130609395 gopurs_runtime.Value
var once_Data_Foldable_foldMap__4130609395 sync.Once
func Get_Data_Foldable_foldMap__4130609395() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__4130609395.Do(func() {
		cache_Data_Foldable_foldMap__4130609395 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__4130609395(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap__4130609395
}

var cache_Data_Foldable_foldMap__1118659089 gopurs_runtime.Value
var once_Data_Foldable_foldMap__1118659089 sync.Once
func Get_Data_Foldable_foldMap__1118659089() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__1118659089.Do(func() {
		cache_Data_Foldable_foldMap__1118659089 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__1118659089(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap__1118659089
}

var cache_Data_Foldable_foldMap__1315280116 gopurs_runtime.Value
var once_Data_Foldable_foldMap__1315280116 sync.Once
func Get_Data_Foldable_foldMap__1315280116() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__1315280116.Do(func() {
		cache_Data_Foldable_foldMap__1315280116 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__1315280116(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap__1315280116
}

var cache_Data_Foldable_foldMap__1811898306 gopurs_runtime.Value
var once_Data_Foldable_foldMap__1811898306 sync.Once
func Get_Data_Foldable_foldMap__1811898306() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__1811898306.Do(func() {
		cache_Data_Foldable_foldMap__1811898306 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__1811898306(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap__1811898306
}

var cache_Data_Foldable_foldMap__193737345 gopurs_runtime.Value
var once_Data_Foldable_foldMap__193737345 sync.Once
func Get_Data_Foldable_foldMap__193737345() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__193737345.Do(func() {
		cache_Data_Foldable_foldMap__193737345 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__193737345(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Foldable_foldMap__193737345
}

var cache_Data_Foldable_foldMap__3661646260 gopurs_runtime.Value
var once_Data_Foldable_foldMap__3661646260 sync.Once
func Get_Data_Foldable_foldMap__3661646260() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__3661646260.Do(func() {
		cache_Data_Foldable_foldMap__3661646260 = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__3661646260(dictMonoid_0_box)
})
	})
	return cache_Data_Foldable_foldMap__3661646260
}

var cache_Data_Foldable_foldMap__4073832436 gopurs_runtime.Value
var once_Data_Foldable_foldMap__4073832436 sync.Once
func Get_Data_Foldable_foldMap__4073832436() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__4073832436.Do(func() {
		cache_Data_Foldable_foldMap__4073832436 = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__4073832436(dictMonoid_0_box)
})
	})
	return cache_Data_Foldable_foldMap__4073832436
}

var cache_Data_Foldable_foldMap__3562626100 gopurs_runtime.Value
var once_Data_Foldable_foldMap__3562626100 sync.Once
func Get_Data_Foldable_foldMap__3562626100() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__3562626100.Do(func() {
		cache_Data_Foldable_foldMap__3562626100 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__3562626100(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Foldable_foldMap__3562626100
}

var cache_Data_Foldable_foldMap__2350611220 gopurs_runtime.Value
var once_Data_Foldable_foldMap__2350611220 sync.Once
func Get_Data_Foldable_foldMap__2350611220() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__2350611220.Do(func() {
		cache_Data_Foldable_foldMap__2350611220 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__2350611220(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Foldable_foldMap__2350611220
}

var cache_Data_Foldable_foldMap__3557693114 gopurs_runtime.Value
var once_Data_Foldable_foldMap__3557693114 sync.Once
func Get_Data_Foldable_foldMap__3557693114() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__3557693114.Do(func() {
		cache_Data_Foldable_foldMap__3557693114 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__3557693114(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap__3557693114
}

var cache_Data_Foldable_foldMap__2556604300 gopurs_runtime.Value
var once_Data_Foldable_foldMap__2556604300 sync.Once
func Get_Data_Foldable_foldMap__2556604300() gopurs_runtime.Value {
	once_Data_Foldable_foldMap__2556604300.Do(func() {
		cache_Data_Foldable_foldMap__2556604300 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldMap__2556604300(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldMap__2556604300
}

var cache_Data_Foldable_foldableAdditive__1841171440 gopurs_runtime.Value
var once_Data_Foldable_foldableAdditive__1841171440 sync.Once
func Get_Data_Foldable_foldableAdditive__1841171440() gopurs_runtime.Value {
	once_Data_Foldable_foldableAdditive__1841171440.Do(func() {
		cache_Data_Foldable_foldableAdditive__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableAdditive__1841171440
}

var cache_Data_Foldable_foldableArray__2950015754 gopurs_runtime.Value
var once_Data_Foldable_foldableArray__2950015754 sync.Once
func Get_Data_Foldable_foldableArray__2950015754() gopurs_runtime.Value {
	once_Data_Foldable_foldableArray__2950015754.Do(func() {
		cache_Data_Foldable_foldableArray__2950015754 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), Get_Data_Foldable_foldlArray(), Get_Data_Foldable_foldrArray())
	})
	return cache_Data_Foldable_foldableArray__2950015754
}

var cache_Data_Foldable_foldableArray__3859409398 gopurs_runtime.Value
var once_Data_Foldable_foldableArray__3859409398 sync.Once
func Get_Data_Foldable_foldableArray__3859409398() gopurs_runtime.Value {
	once_Data_Foldable_foldableArray__3859409398.Do(func() {
		cache_Data_Foldable_foldableArray__3859409398 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), Get_Data_Foldable_foldlArray(), Get_Data_Foldable_foldrArray())
	})
	return cache_Data_Foldable_foldableArray__3859409398
}

var cache_Data_Foldable_foldableConj__1841171440 gopurs_runtime.Value
var once_Data_Foldable_foldableConj__1841171440 sync.Once
func Get_Data_Foldable_foldableConj__1841171440() gopurs_runtime.Value {
	once_Data_Foldable_foldableConj__1841171440.Do(func() {
		cache_Data_Foldable_foldableConj__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableConj__1841171440
}

var cache_Data_Foldable_foldableConst__943899702 gopurs_runtime.Value
var once_Data_Foldable_foldableConst__943899702 sync.Once
func Get_Data_Foldable_foldableConst__943899702() gopurs_runtime.Value {
	once_Data_Foldable_foldableConst__943899702.Do(func() {
		cache_Data_Foldable_foldableConst__943899702 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})
})
}))
	})
	return cache_Data_Foldable_foldableConst__943899702
}

var cache_Data_Foldable_foldableDisj__1841171440 gopurs_runtime.Value
var once_Data_Foldable_foldableDisj__1841171440 sync.Once
func Get_Data_Foldable_foldableDisj__1841171440() gopurs_runtime.Value {
	once_Data_Foldable_foldableDisj__1841171440.Do(func() {
		cache_Data_Foldable_foldableDisj__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableDisj__1841171440
}

var cache_Data_Foldable_foldableDual__1841171440 gopurs_runtime.Value
var once_Data_Foldable_foldableDual__1841171440 sync.Once
func Get_Data_Foldable_foldableDual__1841171440() gopurs_runtime.Value {
	once_Data_Foldable_foldableDual__1841171440.Do(func() {
		cache_Data_Foldable_foldableDual__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableDual__1841171440
}

var cache_Data_Foldable_foldableEither__1622911640 gopurs_runtime.Value
var once_Data_Foldable_foldableEither__1622911640 sync.Once
func Get_Data_Foldable_foldableEither__1622911640() gopurs_runtime.Value {
	once_Data_Foldable_foldableEither__1622911640.Do(func() {
		cache_Data_Foldable_foldableEither__1622911640 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply(v_2, (*Constructor_Data_Either_Right)(v1_3.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Either_Right)(v2_2.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t3 = gopurs_runtime.Apply2(v_0, (*Constructor_Data_Either_Right)(v2_2.UnsafePtr).V0, v1_1)
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
}))
	})
	return cache_Data_Foldable_foldableEither__1622911640
}

var cache_Data_Foldable_foldableFirst__2831137713 gopurs_runtime.Value
var once_Data_Foldable_foldableFirst__2831137713 sync.Once
func Get_Data_Foldable_foldableFirst__2831137713() gopurs_runtime.Value {
	once_Data_Foldable_foldableFirst__2831137713.Do(func() {
		cache_Data_Foldable_foldableFirst__2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_3 == nil) {
__t5 = z_1
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_6 == nil) {
__t8 = z_1
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0, z_1)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
})
}))
	})
	return cache_Data_Foldable_foldableFirst__2831137713
}

var cache_Data_Foldable_foldableFreeMonoidTree__2832280077 gopurs_runtime.Value
var once_Data_Foldable_foldableFreeMonoidTree__2832280077 sync.Once
func Get_Data_Foldable_foldableFreeMonoidTree__2832280077() gopurs_runtime.Value {
	once_Data_Foldable_foldableFreeMonoidTree__2832280077.Do(func() {
		cache_Data_Foldable_foldableFreeMonoidTree__2832280077 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFreeMonoidTree(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_2 gopurs_runtime.Value
go__go_1_2_2 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_2_2:
for {
if false { continue go__go_1_2_2 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t7 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, acc_2, (*Constructor_Data_Foldable_Node)(lhs_3.UnsafePtr).V0)
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_2
__t7 = gopurs_runtime.Value{}
goto end_branch_7
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2812549951) {
var __t5 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V1
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V0
rhs_4_loop = rhs_4
continue go__go_1_2_2
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
var __t3 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V0
rhs_4_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V1
continue go__go_1_2_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V0
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Append{1, (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V1, rhs_4})}
continue go__go_1_2_2
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t5 = __t3
}
end_branch_5:
__t7 = __t5
goto end_branch_7
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
var __t6 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
__t6 = acc_2
goto end_branch_6
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_2
__t6 = gopurs_runtime.Value{}
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
}()
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_2_2, a_2, b_3, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
})
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_8_3 gopurs_runtime.Value
go__go_1_8_3 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_8_3:
for {
if false { continue go__go_1_8_3 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t13 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, (*Constructor_Data_Foldable_Node)(rhs_4.UnsafePtr).V0, acc_2)
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_3
__t13 = gopurs_runtime.Value{}
goto end_branch_13
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2812549951) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V0
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V1
continue go__go_1_8_3
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
var __t9 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V0
rhs_4_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V1
continue go__go_1_8_3
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Append{1, lhs_3, (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V0})}
rhs_4_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V1
continue go__go_1_8_3
__t9 = gopurs_runtime.Value{}
}
end_branch_9:
__t11 = __t9
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
var __t12 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
__t12 = acc_2
goto end_branch_12
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_3
__t12 = gopurs_runtime.Value{}
}
end_branch_12:
__t13 = __t12
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
}()
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_8_3, a_2, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}, b_3)
})
})
}))
	})
	return cache_Data_Foldable_foldableFreeMonoidTree__2832280077
}

var cache_Data_Foldable_foldableIdentity__1841171440 gopurs_runtime.Value
var once_Data_Foldable_foldableIdentity__1841171440 sync.Once
func Get_Data_Foldable_foldableIdentity__1841171440() gopurs_runtime.Value {
	once_Data_Foldable_foldableIdentity__1841171440.Do(func() {
		cache_Data_Foldable_foldableIdentity__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableIdentity__1841171440
}

var cache_Data_Foldable_foldableLast__2831137713 gopurs_runtime.Value
var once_Data_Foldable_foldableLast__2831137713 sync.Once
func Get_Data_Foldable_foldableLast__2831137713() gopurs_runtime.Value {
	once_Data_Foldable_foldableLast__2831137713.Do(func() {
		cache_Data_Foldable_foldableLast__2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_3 == nil) {
__t5 = z_1
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_6 == nil) {
__t8 = z_1
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.Apply2(f_0, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0, z_1)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
})
}))
	})
	return cache_Data_Foldable_foldableLast__2831137713
}

var cache_Data_Foldable_foldableMaybe__3653484922 gopurs_runtime.Value
var once_Data_Foldable_foldableMaybe__3653484922 sync.Once
func Get_Data_Foldable_foldableMaybe__3653484922() gopurs_runtime.Value {
	once_Data_Foldable_foldableMaybe__3653484922.Do(func() {
		cache_Data_Foldable_foldableMaybe__3653484922 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t1 = mempty_1_0
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just)(v1_3.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t2 = v1_1
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Maybe_Just)(v2_2.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t3 = v1_1
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(v_0, (*Constructor_Data_Maybe_Just)(v2_2.UnsafePtr).V0, v1_1)
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
}))
	})
	return cache_Data_Foldable_foldableMaybe__3653484922
}

var cache_Data_Foldable_foldableMaybe__2831137713 gopurs_runtime.Value
var once_Data_Foldable_foldableMaybe__2831137713 sync.Once
func Get_Data_Foldable_foldableMaybe__2831137713() gopurs_runtime.Value {
	once_Data_Foldable_foldableMaybe__2831137713.Do(func() {
		cache_Data_Foldable_foldableMaybe__2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t1 = mempty_1_0
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just)(v1_3.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t2 = v1_1
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Maybe_Just)(v2_2.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t3 = v1_1
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(v_0, (*Constructor_Data_Maybe_Just)(v2_2.UnsafePtr).V0, v1_1)
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
}))
	})
	return cache_Data_Foldable_foldableMaybe__2831137713
}

var cache_Data_Foldable_foldableMultiplicative__1841171440 gopurs_runtime.Value
var once_Data_Foldable_foldableMultiplicative__1841171440 sync.Once
func Get_Data_Foldable_foldableMultiplicative__1841171440() gopurs_runtime.Value {
	once_Data_Foldable_foldableMultiplicative__1841171440.Do(func() {
		cache_Data_Foldable_foldableMultiplicative__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableMultiplicative__1841171440
}

var cache_Data_Foldable_foldableTuple__1455669080 gopurs_runtime.Value
var once_Data_Foldable_foldableTuple__1455669080 sync.Once
func Get_Data_Foldable_foldableTuple__1455669080() gopurs_runtime.Value {
	once_Data_Foldable_foldableTuple__1455669080.Do(func() {
		cache_Data_Foldable_foldableTuple__1455669080 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1, z_1)
})
})
}))
	})
	return cache_Data_Foldable_foldableTuple__1455669080
}

var cache_Data_Foldable_foldl__1422885860 gopurs_runtime.Value
var once_Data_Foldable_foldl__1422885860 sync.Once
func Get_Data_Foldable_foldl__1422885860() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1422885860.Do(func() {
		cache_Data_Foldable_foldl__1422885860 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1422885860(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__1422885860
}

var cache_Data_Foldable_foldl__3850309840 gopurs_runtime.Value
var once_Data_Foldable_foldl__3850309840 sync.Once
func Get_Data_Foldable_foldl__3850309840() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3850309840.Do(func() {
		cache_Data_Foldable_foldl__3850309840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3850309840(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__3850309840
}

var cache_Data_Foldable_foldl__2111289130 gopurs_runtime.Value
var once_Data_Foldable_foldl__2111289130 sync.Once
func Get_Data_Foldable_foldl__2111289130() gopurs_runtime.Value {
	once_Data_Foldable_foldl__2111289130.Do(func() {
		cache_Data_Foldable_foldl__2111289130 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__2111289130(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__2111289130
}

var cache_Data_Foldable_foldl__94807652 gopurs_runtime.Value
var once_Data_Foldable_foldl__94807652 sync.Once
func Get_Data_Foldable_foldl__94807652() gopurs_runtime.Value {
	once_Data_Foldable_foldl__94807652.Do(func() {
		cache_Data_Foldable_foldl__94807652 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__94807652(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__94807652
}

var cache_Data_Foldable_foldl__2699291984 gopurs_runtime.Value
var once_Data_Foldable_foldl__2699291984 sync.Once
func Get_Data_Foldable_foldl__2699291984() gopurs_runtime.Value {
	once_Data_Foldable_foldl__2699291984.Do(func() {
		cache_Data_Foldable_foldl__2699291984 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__2699291984(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__2699291984
}

var cache_Data_Foldable_foldl__3041692656 gopurs_runtime.Value
var once_Data_Foldable_foldl__3041692656 sync.Once
func Get_Data_Foldable_foldl__3041692656() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3041692656.Do(func() {
		cache_Data_Foldable_foldl__3041692656 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3041692656(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__3041692656
}

var cache_Data_Foldable_foldl__66388714 gopurs_runtime.Value
var once_Data_Foldable_foldl__66388714 sync.Once
func Get_Data_Foldable_foldl__66388714() gopurs_runtime.Value {
	once_Data_Foldable_foldl__66388714.Do(func() {
		cache_Data_Foldable_foldl__66388714 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__66388714(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__66388714
}

var cache_Data_Foldable_foldl__506543652 gopurs_runtime.Value
var once_Data_Foldable_foldl__506543652 sync.Once
func Get_Data_Foldable_foldl__506543652() gopurs_runtime.Value {
	once_Data_Foldable_foldl__506543652.Do(func() {
		cache_Data_Foldable_foldl__506543652 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__506543652(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__506543652
}

var cache_Data_Foldable_foldl__1671904522 gopurs_runtime.Value
var once_Data_Foldable_foldl__1671904522 sync.Once
func Get_Data_Foldable_foldl__1671904522() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1671904522.Do(func() {
		cache_Data_Foldable_foldl__1671904522 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1671904522(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__1671904522
}

var cache_Data_Foldable_foldl__2602334544 gopurs_runtime.Value
var once_Data_Foldable_foldl__2602334544 sync.Once
func Get_Data_Foldable_foldl__2602334544() gopurs_runtime.Value {
	once_Data_Foldable_foldl__2602334544.Do(func() {
		cache_Data_Foldable_foldl__2602334544 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__2602334544(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__2602334544
}

var cache_Data_Foldable_foldl__4192477084 gopurs_runtime.Value
var once_Data_Foldable_foldl__4192477084 sync.Once
func Get_Data_Foldable_foldl__4192477084() gopurs_runtime.Value {
	once_Data_Foldable_foldl__4192477084.Do(func() {
		cache_Data_Foldable_foldl__4192477084 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__4192477084(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__4192477084
}

var cache_Data_Foldable_foldl__3272087748 gopurs_runtime.Value
var once_Data_Foldable_foldl__3272087748 sync.Once
func Get_Data_Foldable_foldl__3272087748() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3272087748.Do(func() {
		cache_Data_Foldable_foldl__3272087748 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3272087748(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__3272087748
}

var cache_Data_Foldable_foldl__165683952 gopurs_runtime.Value
var once_Data_Foldable_foldl__165683952 sync.Once
func Get_Data_Foldable_foldl__165683952() gopurs_runtime.Value {
	once_Data_Foldable_foldl__165683952.Do(func() {
		cache_Data_Foldable_foldl__165683952 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__165683952(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__165683952
}

var cache_Data_Foldable_foldl__267332164 gopurs_runtime.Value
var once_Data_Foldable_foldl__267332164 sync.Once
func Get_Data_Foldable_foldl__267332164() gopurs_runtime.Value {
	once_Data_Foldable_foldl__267332164.Do(func() {
		cache_Data_Foldable_foldl__267332164 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__267332164(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__267332164
}

var cache_Data_Foldable_foldl__371433392 gopurs_runtime.Value
var once_Data_Foldable_foldl__371433392 sync.Once
func Get_Data_Foldable_foldl__371433392() gopurs_runtime.Value {
	once_Data_Foldable_foldl__371433392.Do(func() {
		cache_Data_Foldable_foldl__371433392 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__371433392(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__371433392
}

var cache_Data_Foldable_foldl__3234403824 gopurs_runtime.Value
var once_Data_Foldable_foldl__3234403824 sync.Once
func Get_Data_Foldable_foldl__3234403824() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3234403824.Do(func() {
		cache_Data_Foldable_foldl__3234403824 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3234403824(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__3234403824
}

var cache_Data_Foldable_foldl__1656262032 gopurs_runtime.Value
var once_Data_Foldable_foldl__1656262032 sync.Once
func Get_Data_Foldable_foldl__1656262032() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1656262032.Do(func() {
		cache_Data_Foldable_foldl__1656262032 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1656262032(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__1656262032
}

var cache_Data_Foldable_foldl__3559959056 gopurs_runtime.Value
var once_Data_Foldable_foldl__3559959056 sync.Once
func Get_Data_Foldable_foldl__3559959056() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3559959056.Do(func() {
		cache_Data_Foldable_foldl__3559959056 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3559959056(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__3559959056
}

var cache_Data_Foldable_foldl__3131354468 gopurs_runtime.Value
var once_Data_Foldable_foldl__3131354468 sync.Once
func Get_Data_Foldable_foldl__3131354468() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3131354468.Do(func() {
		cache_Data_Foldable_foldl__3131354468 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3131354468(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__3131354468
}

var cache_Data_Foldable_foldl__1601164432 gopurs_runtime.Value
var once_Data_Foldable_foldl__1601164432 sync.Once
func Get_Data_Foldable_foldl__1601164432() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1601164432.Do(func() {
		cache_Data_Foldable_foldl__1601164432 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1601164432(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__1601164432
}

var cache_Data_Foldable_foldl__3893253828 gopurs_runtime.Value
var once_Data_Foldable_foldl__3893253828 sync.Once
func Get_Data_Foldable_foldl__3893253828() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3893253828.Do(func() {
		cache_Data_Foldable_foldl__3893253828 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3893253828(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__3893253828
}

var cache_Data_Foldable_foldl__1148906672 gopurs_runtime.Value
var once_Data_Foldable_foldl__1148906672 sync.Once
func Get_Data_Foldable_foldl__1148906672() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1148906672.Do(func() {
		cache_Data_Foldable_foldl__1148906672 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1148906672(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__1148906672
}

var cache_Data_Foldable_foldl__2656621979 gopurs_runtime.Value
var once_Data_Foldable_foldl__2656621979 sync.Once
func Get_Data_Foldable_foldl__2656621979() gopurs_runtime.Value {
	once_Data_Foldable_foldl__2656621979.Do(func() {
		cache_Data_Foldable_foldl__2656621979 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__2656621979(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__2656621979
}

var cache_Data_Foldable_foldl__3422238939 gopurs_runtime.Value
var once_Data_Foldable_foldl__3422238939 sync.Once
func Get_Data_Foldable_foldl__3422238939() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3422238939.Do(func() {
		cache_Data_Foldable_foldl__3422238939 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3422238939(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__3422238939
}

var cache_Data_Foldable_foldl__2151204251 gopurs_runtime.Value
var once_Data_Foldable_foldl__2151204251 sync.Once
func Get_Data_Foldable_foldl__2151204251() gopurs_runtime.Value {
	once_Data_Foldable_foldl__2151204251.Do(func() {
		cache_Data_Foldable_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__2151204251(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__2151204251
}

var cache_Data_Foldable_foldl__4099223803 gopurs_runtime.Value
var once_Data_Foldable_foldl__4099223803 sync.Once
func Get_Data_Foldable_foldl__4099223803() gopurs_runtime.Value {
	once_Data_Foldable_foldl__4099223803.Do(func() {
		cache_Data_Foldable_foldl__4099223803 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__4099223803(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_Foldable_foldl__4099223803
}

var cache_Data_Foldable_foldl__2123442907 gopurs_runtime.Value
var once_Data_Foldable_foldl__2123442907 sync.Once
func Get_Data_Foldable_foldl__2123442907() gopurs_runtime.Value {
	once_Data_Foldable_foldl__2123442907.Do(func() {
		cache_Data_Foldable_foldl__2123442907 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__2123442907(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__2123442907
}

var cache_Data_Foldable_foldl__524683195 gopurs_runtime.Value
var once_Data_Foldable_foldl__524683195 sync.Once
func Get_Data_Foldable_foldl__524683195() gopurs_runtime.Value {
	once_Data_Foldable_foldl__524683195.Do(func() {
		cache_Data_Foldable_foldl__524683195 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__524683195(op_0_box)
})
	})
	return cache_Data_Foldable_foldl__524683195
}

var cache_Data_Foldable_foldl__1712912315 gopurs_runtime.Value
var once_Data_Foldable_foldl__1712912315 sync.Once
func Get_Data_Foldable_foldl__1712912315() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1712912315.Do(func() {
		cache_Data_Foldable_foldl__1712912315 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1712912315(f_0_box)
})
	})
	return cache_Data_Foldable_foldl__1712912315
}

var cache_Data_Foldable_foldl__2138619643 gopurs_runtime.Value
var once_Data_Foldable_foldl__2138619643 sync.Once
func Get_Data_Foldable_foldl__2138619643() gopurs_runtime.Value {
	once_Data_Foldable_foldl__2138619643.Do(func() {
		cache_Data_Foldable_foldl__2138619643 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__2138619643(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__2138619643
}

var cache_Data_Foldable_foldl__4056605371 gopurs_runtime.Value
var once_Data_Foldable_foldl__4056605371 sync.Once
func Get_Data_Foldable_foldl__4056605371() gopurs_runtime.Value {
	once_Data_Foldable_foldl__4056605371.Do(func() {
		cache_Data_Foldable_foldl__4056605371 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__4056605371(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__4056605371
}

var cache_Data_Foldable_foldl__53736539 gopurs_runtime.Value
var once_Data_Foldable_foldl__53736539 sync.Once
func Get_Data_Foldable_foldl__53736539() gopurs_runtime.Value {
	once_Data_Foldable_foldl__53736539.Do(func() {
		cache_Data_Foldable_foldl__53736539 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__53736539(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__53736539
}

var cache_Data_Foldable_foldl__22573083 gopurs_runtime.Value
var once_Data_Foldable_foldl__22573083 sync.Once
func Get_Data_Foldable_foldl__22573083() gopurs_runtime.Value {
	once_Data_Foldable_foldl__22573083.Do(func() {
		cache_Data_Foldable_foldl__22573083 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__22573083(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__22573083
}

var cache_Data_Foldable_foldl__393765499 gopurs_runtime.Value
var once_Data_Foldable_foldl__393765499 sync.Once
func Get_Data_Foldable_foldl__393765499() gopurs_runtime.Value {
	once_Data_Foldable_foldl__393765499.Do(func() {
		cache_Data_Foldable_foldl__393765499 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__393765499(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__393765499
}

var cache_Data_Foldable_foldl__3306117403 gopurs_runtime.Value
var once_Data_Foldable_foldl__3306117403 sync.Once
func Get_Data_Foldable_foldl__3306117403() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3306117403.Do(func() {
		cache_Data_Foldable_foldl__3306117403 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3306117403(op_0_box)
})
	})
	return cache_Data_Foldable_foldl__3306117403
}

var cache_Data_Foldable_foldl__2159564571 gopurs_runtime.Value
var once_Data_Foldable_foldl__2159564571 sync.Once
func Get_Data_Foldable_foldl__2159564571() gopurs_runtime.Value {
	once_Data_Foldable_foldl__2159564571.Do(func() {
		cache_Data_Foldable_foldl__2159564571 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__2159564571(f_0_box)
})
	})
	return cache_Data_Foldable_foldl__2159564571
}

var cache_Data_Foldable_foldl__22791451 gopurs_runtime.Value
var once_Data_Foldable_foldl__22791451 sync.Once
func Get_Data_Foldable_foldl__22791451() gopurs_runtime.Value {
	once_Data_Foldable_foldl__22791451.Do(func() {
		cache_Data_Foldable_foldl__22791451 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__22791451(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__22791451
}

var cache_Data_Foldable_foldl__3736093275 gopurs_runtime.Value
var once_Data_Foldable_foldl__3736093275 sync.Once
func Get_Data_Foldable_foldl__3736093275() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3736093275.Do(func() {
		cache_Data_Foldable_foldl__3736093275 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3736093275(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__3736093275
}

var cache_Data_Foldable_foldl__3785384859 gopurs_runtime.Value
var once_Data_Foldable_foldl__3785384859 sync.Once
func Get_Data_Foldable_foldl__3785384859() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3785384859.Do(func() {
		cache_Data_Foldable_foldl__3785384859 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3785384859(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__3785384859
}

var cache_Data_Foldable_foldl__446290811 gopurs_runtime.Value
var once_Data_Foldable_foldl__446290811 sync.Once
func Get_Data_Foldable_foldl__446290811() gopurs_runtime.Value {
	once_Data_Foldable_foldl__446290811.Do(func() {
		cache_Data_Foldable_foldl__446290811 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__446290811(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__446290811
}

var cache_Data_Foldable_foldl__1459781277 gopurs_runtime.Value
var once_Data_Foldable_foldl__1459781277 sync.Once
func Get_Data_Foldable_foldl__1459781277() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1459781277.Do(func() {
		cache_Data_Foldable_foldl__1459781277 = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1459781277(fn_0_box)
})
	})
	return cache_Data_Foldable_foldl__1459781277
}

var cache_Data_Foldable_foldl__3288778237 gopurs_runtime.Value
var once_Data_Foldable_foldl__3288778237 sync.Once
func Get_Data_Foldable_foldl__3288778237() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3288778237.Do(func() {
		cache_Data_Foldable_foldl__3288778237 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3288778237(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_Data_Foldable_foldl__3288778237
}

var cache_Data_Foldable_foldl__3504930205 gopurs_runtime.Value
var once_Data_Foldable_foldl__3504930205 sync.Once
func Get_Data_Foldable_foldl__3504930205() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3504930205.Do(func() {
		cache_Data_Foldable_foldl__3504930205 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3504930205(op_0_box)
})
	})
	return cache_Data_Foldable_foldl__3504930205
}

var cache_Data_Foldable_foldl__3379885725 gopurs_runtime.Value
var once_Data_Foldable_foldl__3379885725 sync.Once
func Get_Data_Foldable_foldl__3379885725() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3379885725.Do(func() {
		cache_Data_Foldable_foldl__3379885725 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3379885725(op_0_box)
})
	})
	return cache_Data_Foldable_foldl__3379885725
}

var cache_Data_Foldable_foldl__3737487037 gopurs_runtime.Value
var once_Data_Foldable_foldl__3737487037 sync.Once
func Get_Data_Foldable_foldl__3737487037() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3737487037.Do(func() {
		cache_Data_Foldable_foldl__3737487037 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3737487037(op_0_box)
})
	})
	return cache_Data_Foldable_foldl__3737487037
}

var cache_Data_Foldable_foldl__1985071933 gopurs_runtime.Value
var once_Data_Foldable_foldl__1985071933 sync.Once
func Get_Data_Foldable_foldl__1985071933() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1985071933.Do(func() {
		cache_Data_Foldable_foldl__1985071933 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1985071933(op_0_box)
})
	})
	return cache_Data_Foldable_foldl__1985071933
}

var cache_Data_Foldable_foldl__536153533 gopurs_runtime.Value
var once_Data_Foldable_foldl__536153533 sync.Once
func Get_Data_Foldable_foldl__536153533() gopurs_runtime.Value {
	once_Data_Foldable_foldl__536153533.Do(func() {
		cache_Data_Foldable_foldl__536153533 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__536153533(op_0_box)
})
	})
	return cache_Data_Foldable_foldl__536153533
}

var cache_Data_Foldable_foldl__4234493053 gopurs_runtime.Value
var once_Data_Foldable_foldl__4234493053 sync.Once
func Get_Data_Foldable_foldl__4234493053() gopurs_runtime.Value {
	once_Data_Foldable_foldl__4234493053.Do(func() {
		cache_Data_Foldable_foldl__4234493053 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__4234493053(op_0_box)
})
	})
	return cache_Data_Foldable_foldl__4234493053
}

var cache_Data_Foldable_foldl__176907901 gopurs_runtime.Value
var once_Data_Foldable_foldl__176907901 sync.Once
func Get_Data_Foldable_foldl__176907901() gopurs_runtime.Value {
	once_Data_Foldable_foldl__176907901.Do(func() {
		cache_Data_Foldable_foldl__176907901 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__176907901(op_0_box)
})
	})
	return cache_Data_Foldable_foldl__176907901
}

var cache_Data_Foldable_foldl__170252797 gopurs_runtime.Value
var once_Data_Foldable_foldl__170252797 sync.Once
func Get_Data_Foldable_foldl__170252797() gopurs_runtime.Value {
	once_Data_Foldable_foldl__170252797.Do(func() {
		cache_Data_Foldable_foldl__170252797 = gopurs_runtime.Func(func(op_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__170252797(op_0_box)
})
	})
	return cache_Data_Foldable_foldl__170252797
}

var cache_Data_Foldable_foldl__1754241693 gopurs_runtime.Value
var once_Data_Foldable_foldl__1754241693 sync.Once
func Get_Data_Foldable_foldl__1754241693() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1754241693.Do(func() {
		cache_Data_Foldable_foldl__1754241693 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1754241693(f_0_box)
})
	})
	return cache_Data_Foldable_foldl__1754241693
}

var cache_Data_Foldable_foldl__3943124669 gopurs_runtime.Value
var once_Data_Foldable_foldl__3943124669 sync.Once
func Get_Data_Foldable_foldl__3943124669() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3943124669.Do(func() {
		cache_Data_Foldable_foldl__3943124669 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3943124669(f_0_box)
})
	})
	return cache_Data_Foldable_foldl__3943124669
}

var cache_Data_Foldable_foldl__396932925 gopurs_runtime.Value
var once_Data_Foldable_foldl__396932925 sync.Once
func Get_Data_Foldable_foldl__396932925() gopurs_runtime.Value {
	once_Data_Foldable_foldl__396932925.Do(func() {
		cache_Data_Foldable_foldl__396932925 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__396932925(f_0_box)
})
	})
	return cache_Data_Foldable_foldl__396932925
}

var cache_Data_Foldable_foldl__2928402749 gopurs_runtime.Value
var once_Data_Foldable_foldl__2928402749 sync.Once
func Get_Data_Foldable_foldl__2928402749() gopurs_runtime.Value {
	once_Data_Foldable_foldl__2928402749.Do(func() {
		cache_Data_Foldable_foldl__2928402749 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__2928402749(f_0_box)
})
	})
	return cache_Data_Foldable_foldl__2928402749
}

var cache_Data_Foldable_foldl__255626813 gopurs_runtime.Value
var once_Data_Foldable_foldl__255626813 sync.Once
func Get_Data_Foldable_foldl__255626813() gopurs_runtime.Value {
	once_Data_Foldable_foldl__255626813.Do(func() {
		cache_Data_Foldable_foldl__255626813 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__255626813(f_0_box)
})
	})
	return cache_Data_Foldable_foldl__255626813
}

var cache_Data_Foldable_foldl__3915700701 gopurs_runtime.Value
var once_Data_Foldable_foldl__3915700701 sync.Once
func Get_Data_Foldable_foldl__3915700701() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3915700701.Do(func() {
		cache_Data_Foldable_foldl__3915700701 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3915700701(f_0_box)
})
	})
	return cache_Data_Foldable_foldl__3915700701
}

var cache_Data_Foldable_foldl__3459294429 gopurs_runtime.Value
var once_Data_Foldable_foldl__3459294429 sync.Once
func Get_Data_Foldable_foldl__3459294429() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3459294429.Do(func() {
		cache_Data_Foldable_foldl__3459294429 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3459294429(f_0_box)
})
	})
	return cache_Data_Foldable_foldl__3459294429
}

var cache_Data_Foldable_foldl__512483965 gopurs_runtime.Value
var once_Data_Foldable_foldl__512483965 sync.Once
func Get_Data_Foldable_foldl__512483965() gopurs_runtime.Value {
	once_Data_Foldable_foldl__512483965.Do(func() {
		cache_Data_Foldable_foldl__512483965 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__512483965(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_1_box))
})
	})
	return cache_Data_Foldable_foldl__512483965
}

var cache_Data_Foldable_foldl__3016550397 gopurs_runtime.Value
var once_Data_Foldable_foldl__3016550397 sync.Once
func Get_Data_Foldable_foldl__3016550397() gopurs_runtime.Value {
	once_Data_Foldable_foldl__3016550397.Do(func() {
		cache_Data_Foldable_foldl__3016550397 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__3016550397(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Foldable_foldl__3016550397
}

var cache_Data_Foldable_foldl__1714316381 gopurs_runtime.Value
var once_Data_Foldable_foldl__1714316381 sync.Once
func Get_Data_Foldable_foldl__1714316381() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1714316381.Do(func() {
		cache_Data_Foldable_foldl__1714316381 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1714316381(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__1714316381
}

var cache_Data_Foldable_foldl__380919197 gopurs_runtime.Value
var once_Data_Foldable_foldl__380919197 sync.Once
func Get_Data_Foldable_foldl__380919197() gopurs_runtime.Value {
	once_Data_Foldable_foldl__380919197.Do(func() {
		cache_Data_Foldable_foldl__380919197 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__380919197(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__380919197
}

var cache_Data_Foldable_foldl__992072381 gopurs_runtime.Value
var once_Data_Foldable_foldl__992072381 sync.Once
func Get_Data_Foldable_foldl__992072381() gopurs_runtime.Value {
	once_Data_Foldable_foldl__992072381.Do(func() {
		cache_Data_Foldable_foldl__992072381 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__992072381(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldl__992072381
}

var cache_Data_Foldable_foldl__2188030845 gopurs_runtime.Value
var once_Data_Foldable_foldl__2188030845 sync.Once
func Get_Data_Foldable_foldl__2188030845() gopurs_runtime.Value {
	once_Data_Foldable_foldl__2188030845.Do(func() {
		cache_Data_Foldable_foldl__2188030845 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__2188030845(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_Foldable_foldl__2188030845
}

var cache_Data_Foldable_foldl__1444272061 gopurs_runtime.Value
var once_Data_Foldable_foldl__1444272061 sync.Once
func Get_Data_Foldable_foldl__1444272061() gopurs_runtime.Value {
	once_Data_Foldable_foldl__1444272061.Do(func() {
		cache_Data_Foldable_foldl__1444272061 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldl__1444272061(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_Foldable_foldl__1444272061
}

var cache_Data_Foldable_foldr__1038841770 gopurs_runtime.Value
var once_Data_Foldable_foldr__1038841770 sync.Once
func Get_Data_Foldable_foldr__1038841770() gopurs_runtime.Value {
	once_Data_Foldable_foldr__1038841770.Do(func() {
		cache_Data_Foldable_foldr__1038841770 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__1038841770(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__1038841770
}

var cache_Data_Foldable_foldr__2858227716 gopurs_runtime.Value
var once_Data_Foldable_foldr__2858227716 sync.Once
func Get_Data_Foldable_foldr__2858227716() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2858227716.Do(func() {
		cache_Data_Foldable_foldr__2858227716 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2858227716(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2858227716
}

var cache_Data_Foldable_foldr__3728540540 gopurs_runtime.Value
var once_Data_Foldable_foldr__3728540540 sync.Once
func Get_Data_Foldable_foldr__3728540540() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3728540540.Do(func() {
		cache_Data_Foldable_foldr__3728540540 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3728540540(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__3728540540
}

var cache_Data_Foldable_foldr__2111289130 gopurs_runtime.Value
var once_Data_Foldable_foldr__2111289130 sync.Once
func Get_Data_Foldable_foldr__2111289130() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2111289130.Do(func() {
		cache_Data_Foldable_foldr__2111289130 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2111289130(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2111289130
}

var cache_Data_Foldable_foldr__208886460 gopurs_runtime.Value
var once_Data_Foldable_foldr__208886460 sync.Once
func Get_Data_Foldable_foldr__208886460() gopurs_runtime.Value {
	once_Data_Foldable_foldr__208886460.Do(func() {
		cache_Data_Foldable_foldr__208886460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__208886460(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__208886460
}

var cache_Data_Foldable_foldr__926146538 gopurs_runtime.Value
var once_Data_Foldable_foldr__926146538 sync.Once
func Get_Data_Foldable_foldr__926146538() gopurs_runtime.Value {
	once_Data_Foldable_foldr__926146538.Do(func() {
		cache_Data_Foldable_foldr__926146538 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__926146538(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__926146538
}

var cache_Data_Foldable_foldr__2512763050 gopurs_runtime.Value
var once_Data_Foldable_foldr__2512763050 sync.Once
func Get_Data_Foldable_foldr__2512763050() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2512763050.Do(func() {
		cache_Data_Foldable_foldr__2512763050 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2512763050(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2512763050
}

var cache_Data_Foldable_foldr__3673994608 gopurs_runtime.Value
var once_Data_Foldable_foldr__3673994608 sync.Once
func Get_Data_Foldable_foldr__3673994608() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3673994608.Do(func() {
		cache_Data_Foldable_foldr__3673994608 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3673994608(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__3673994608
}

var cache_Data_Foldable_foldr__919612668 gopurs_runtime.Value
var once_Data_Foldable_foldr__919612668 sync.Once
func Get_Data_Foldable_foldr__919612668() gopurs_runtime.Value {
	once_Data_Foldable_foldr__919612668.Do(func() {
		cache_Data_Foldable_foldr__919612668 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__919612668(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__919612668
}

var cache_Data_Foldable_foldr__2232849770 gopurs_runtime.Value
var once_Data_Foldable_foldr__2232849770 sync.Once
func Get_Data_Foldable_foldr__2232849770() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2232849770.Do(func() {
		cache_Data_Foldable_foldr__2232849770 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2232849770(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2232849770
}

var cache_Data_Foldable_foldr__3630705947 gopurs_runtime.Value
var once_Data_Foldable_foldr__3630705947 sync.Once
func Get_Data_Foldable_foldr__3630705947() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3630705947.Do(func() {
		cache_Data_Foldable_foldr__3630705947 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3630705947(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__3630705947
}

var cache_Data_Foldable_foldr__2151204251 gopurs_runtime.Value
var once_Data_Foldable_foldr__2151204251 sync.Once
func Get_Data_Foldable_foldr__2151204251() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2151204251.Do(func() {
		cache_Data_Foldable_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2151204251(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2151204251
}

var cache_Data_Foldable_foldr__3675782427 gopurs_runtime.Value
var once_Data_Foldable_foldr__3675782427 sync.Once
func Get_Data_Foldable_foldr__3675782427() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3675782427.Do(func() {
		cache_Data_Foldable_foldr__3675782427 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3675782427(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__3675782427
}

var cache_Data_Foldable_foldr__2403185435 gopurs_runtime.Value
var once_Data_Foldable_foldr__2403185435 sync.Once
func Get_Data_Foldable_foldr__2403185435() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2403185435.Do(func() {
		cache_Data_Foldable_foldr__2403185435 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2403185435(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2403185435
}

var cache_Data_Foldable_foldr__2829803163 gopurs_runtime.Value
var once_Data_Foldable_foldr__2829803163 sync.Once
func Get_Data_Foldable_foldr__2829803163() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2829803163.Do(func() {
		cache_Data_Foldable_foldr__2829803163 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2829803163(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2829803163
}

var cache_Data_Foldable_foldr__4105571355 gopurs_runtime.Value
var once_Data_Foldable_foldr__4105571355 sync.Once
func Get_Data_Foldable_foldr__4105571355() gopurs_runtime.Value {
	once_Data_Foldable_foldr__4105571355.Do(func() {
		cache_Data_Foldable_foldr__4105571355 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__4105571355(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__4105571355
}

var cache_Data_Foldable_foldr__3591001499 gopurs_runtime.Value
var once_Data_Foldable_foldr__3591001499 sync.Once
func Get_Data_Foldable_foldr__3591001499() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3591001499.Do(func() {
		cache_Data_Foldable_foldr__3591001499 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3591001499(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__3591001499
}

var cache_Data_Foldable_foldr__2492367323 gopurs_runtime.Value
var once_Data_Foldable_foldr__2492367323 sync.Once
func Get_Data_Foldable_foldr__2492367323() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2492367323.Do(func() {
		cache_Data_Foldable_foldr__2492367323 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2492367323(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2492367323
}

var cache_Data_Foldable_foldr__391354971 gopurs_runtime.Value
var once_Data_Foldable_foldr__391354971 sync.Once
func Get_Data_Foldable_foldr__391354971() gopurs_runtime.Value {
	once_Data_Foldable_foldr__391354971.Do(func() {
		cache_Data_Foldable_foldr__391354971 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__391354971(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__391354971
}

var cache_Data_Foldable_foldr__2671482779 gopurs_runtime.Value
var once_Data_Foldable_foldr__2671482779 sync.Once
func Get_Data_Foldable_foldr__2671482779() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2671482779.Do(func() {
		cache_Data_Foldable_foldr__2671482779 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2671482779(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2671482779
}

var cache_Data_Foldable_foldr__1030499675 gopurs_runtime.Value
var once_Data_Foldable_foldr__1030499675 sync.Once
func Get_Data_Foldable_foldr__1030499675() gopurs_runtime.Value {
	once_Data_Foldable_foldr__1030499675.Do(func() {
		cache_Data_Foldable_foldr__1030499675 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__1030499675(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_Foldable_foldr__1030499675
}

var cache_Data_Foldable_foldr__3948834331 gopurs_runtime.Value
var once_Data_Foldable_foldr__3948834331 sync.Once
func Get_Data_Foldable_foldr__3948834331() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3948834331.Do(func() {
		cache_Data_Foldable_foldr__3948834331 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3948834331(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__3948834331
}

var cache_Data_Foldable_foldr__1687192379 gopurs_runtime.Value
var once_Data_Foldable_foldr__1687192379 sync.Once
func Get_Data_Foldable_foldr__1687192379() gopurs_runtime.Value {
	once_Data_Foldable_foldr__1687192379.Do(func() {
		cache_Data_Foldable_foldr__1687192379 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__1687192379(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__1687192379
}

var cache_Data_Foldable_foldr__3482737755 gopurs_runtime.Value
var once_Data_Foldable_foldr__3482737755 sync.Once
func Get_Data_Foldable_foldr__3482737755() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3482737755.Do(func() {
		cache_Data_Foldable_foldr__3482737755 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3482737755(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__3482737755
}

var cache_Data_Foldable_foldr__1459781277 gopurs_runtime.Value
var once_Data_Foldable_foldr__1459781277 sync.Once
func Get_Data_Foldable_foldr__1459781277() gopurs_runtime.Value {
	once_Data_Foldable_foldr__1459781277.Do(func() {
		cache_Data_Foldable_foldr__1459781277 = gopurs_runtime.Func(func(fn_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__1459781277(fn_0_box)
})
	})
	return cache_Data_Foldable_foldr__1459781277
}

var cache_Data_Foldable_foldr__3288778237 gopurs_runtime.Value
var once_Data_Foldable_foldr__3288778237 sync.Once
func Get_Data_Foldable_foldr__3288778237() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3288778237.Do(func() {
		cache_Data_Foldable_foldr__3288778237 = gopurs_runtime.Func3(func(x_0_box gopurs_runtime.Value, u_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3288778237(x_0_box, u_1_box, xs_2_box)
})
	})
	return cache_Data_Foldable_foldr__3288778237
}

var cache_Data_Foldable_foldr__1985071933 gopurs_runtime.Value
var once_Data_Foldable_foldr__1985071933 sync.Once
func Get_Data_Foldable_foldr__1985071933() gopurs_runtime.Value {
	once_Data_Foldable_foldr__1985071933.Do(func() {
		cache_Data_Foldable_foldr__1985071933 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__1985071933(op_0_box, z_1_box, xs_2_box)
})
	})
	return cache_Data_Foldable_foldr__1985071933
}

var cache_Data_Foldable_foldr__3192890333 gopurs_runtime.Value
var once_Data_Foldable_foldr__3192890333 sync.Once
func Get_Data_Foldable_foldr__3192890333() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3192890333.Do(func() {
		cache_Data_Foldable_foldr__3192890333 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3192890333(op_0_box, z_1_box, xs_2_box)
})
	})
	return cache_Data_Foldable_foldr__3192890333
}

var cache_Data_Foldable_foldr__2389967549 gopurs_runtime.Value
var once_Data_Foldable_foldr__2389967549 sync.Once
func Get_Data_Foldable_foldr__2389967549() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2389967549.Do(func() {
		cache_Data_Foldable_foldr__2389967549 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2389967549(op_0_box, z_1_box, xs_2_box)
})
	})
	return cache_Data_Foldable_foldr__2389967549
}

var cache_Data_Foldable_foldr__1278383325 gopurs_runtime.Value
var once_Data_Foldable_foldr__1278383325 sync.Once
func Get_Data_Foldable_foldr__1278383325() gopurs_runtime.Value {
	once_Data_Foldable_foldr__1278383325.Do(func() {
		cache_Data_Foldable_foldr__1278383325 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__1278383325(op_0_box, z_1_box, xs_2_box)
})
	})
	return cache_Data_Foldable_foldr__1278383325
}

var cache_Data_Foldable_foldr__2492628765 gopurs_runtime.Value
var once_Data_Foldable_foldr__2492628765 sync.Once
func Get_Data_Foldable_foldr__2492628765() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2492628765.Do(func() {
		cache_Data_Foldable_foldr__2492628765 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2492628765(op_0_box, z_1_box, xs_2_box)
})
	})
	return cache_Data_Foldable_foldr__2492628765
}

var cache_Data_Foldable_foldr__3433277981 gopurs_runtime.Value
var once_Data_Foldable_foldr__3433277981 sync.Once
func Get_Data_Foldable_foldr__3433277981() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3433277981.Do(func() {
		cache_Data_Foldable_foldr__3433277981 = gopurs_runtime.Func3(func(op_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Foldable_foldr__3433277981(op_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](z_1_box), xs_2_box))}
})
	})
	return cache_Data_Foldable_foldr__3433277981
}

var cache_Data_Foldable_foldr__3943124669 gopurs_runtime.Value
var once_Data_Foldable_foldr__3943124669 sync.Once
func Get_Data_Foldable_foldr__3943124669() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3943124669.Do(func() {
		cache_Data_Foldable_foldr__3943124669 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3943124669(f_0_box, b_1_box)
})
	})
	return cache_Data_Foldable_foldr__3943124669
}

var cache_Data_Foldable_foldr__2979608669 gopurs_runtime.Value
var once_Data_Foldable_foldr__2979608669 sync.Once
func Get_Data_Foldable_foldr__2979608669() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2979608669.Do(func() {
		cache_Data_Foldable_foldr__2979608669 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2979608669(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_1_box))
})
	})
	return cache_Data_Foldable_foldr__2979608669
}

var cache_Data_Foldable_foldr__4137485405 gopurs_runtime.Value
var once_Data_Foldable_foldr__4137485405 sync.Once
func Get_Data_Foldable_foldr__4137485405() gopurs_runtime.Value {
	once_Data_Foldable_foldr__4137485405.Do(func() {
		cache_Data_Foldable_foldr__4137485405 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__4137485405(f_0_box, b_1_box)
})
	})
	return cache_Data_Foldable_foldr__4137485405
}

var cache_Data_Foldable_foldr__3489910557 gopurs_runtime.Value
var once_Data_Foldable_foldr__3489910557 sync.Once
func Get_Data_Foldable_foldr__3489910557() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3489910557.Do(func() {
		cache_Data_Foldable_foldr__3489910557 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3489910557(f_0_box, b_1_box)
})
	})
	return cache_Data_Foldable_foldr__3489910557
}

var cache_Data_Foldable_foldr__3234921885 gopurs_runtime.Value
var once_Data_Foldable_foldr__3234921885 sync.Once
func Get_Data_Foldable_foldr__3234921885() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3234921885.Do(func() {
		cache_Data_Foldable_foldr__3234921885 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3234921885(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](b_1_box))
})
	})
	return cache_Data_Foldable_foldr__3234921885
}

var cache_Data_Foldable_foldr__3235634269 gopurs_runtime.Value
var once_Data_Foldable_foldr__3235634269 sync.Once
func Get_Data_Foldable_foldr__3235634269() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3235634269.Do(func() {
		cache_Data_Foldable_foldr__3235634269 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3235634269(f_0_box, x_1_box)
})
	})
	return cache_Data_Foldable_foldr__3235634269
}

var cache_Data_Foldable_foldr__530094749 gopurs_runtime.Value
var once_Data_Foldable_foldr__530094749 sync.Once
func Get_Data_Foldable_foldr__530094749() gopurs_runtime.Value {
	once_Data_Foldable_foldr__530094749.Do(func() {
		cache_Data_Foldable_foldr__530094749 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__530094749(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](z_1_box))
})
	})
	return cache_Data_Foldable_foldr__530094749
}

var cache_Data_Foldable_foldr__4254578461 gopurs_runtime.Value
var once_Data_Foldable_foldr__4254578461 sync.Once
func Get_Data_Foldable_foldr__4254578461() gopurs_runtime.Value {
	once_Data_Foldable_foldr__4254578461.Do(func() {
		cache_Data_Foldable_foldr__4254578461 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__4254578461(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_1_box))
})
	})
	return cache_Data_Foldable_foldr__4254578461
}

var cache_Data_Foldable_foldr__2178954717 gopurs_runtime.Value
var once_Data_Foldable_foldr__2178954717 sync.Once
func Get_Data_Foldable_foldr__2178954717() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2178954717.Do(func() {
		cache_Data_Foldable_foldr__2178954717 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2178954717(f_0_box, x_1_box)
})
	})
	return cache_Data_Foldable_foldr__2178954717
}

var cache_Data_Foldable_foldr__3016550397 gopurs_runtime.Value
var once_Data_Foldable_foldr__3016550397 sync.Once
func Get_Data_Foldable_foldr__3016550397() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3016550397.Do(func() {
		cache_Data_Foldable_foldr__3016550397 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3016550397(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v2_2_box))
})
	})
	return cache_Data_Foldable_foldr__3016550397
}

var cache_Data_Foldable_foldr__2147034525 gopurs_runtime.Value
var once_Data_Foldable_foldr__2147034525 sync.Once
func Get_Data_Foldable_foldr__2147034525() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2147034525.Do(func() {
		cache_Data_Foldable_foldr__2147034525 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2147034525(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2147034525
}

var cache_Data_Foldable_foldr__2737211997 gopurs_runtime.Value
var once_Data_Foldable_foldr__2737211997 sync.Once
func Get_Data_Foldable_foldr__2737211997() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2737211997.Do(func() {
		cache_Data_Foldable_foldr__2737211997 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2737211997(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__2737211997
}

var cache_Data_Foldable_foldr__2188030845 gopurs_runtime.Value
var once_Data_Foldable_foldr__2188030845 sync.Once
func Get_Data_Foldable_foldr__2188030845() gopurs_runtime.Value {
	once_Data_Foldable_foldr__2188030845.Do(func() {
		cache_Data_Foldable_foldr__2188030845 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__2188030845(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_Foldable_foldr__2188030845
}

var cache_Data_Foldable_foldr__344102461 gopurs_runtime.Value
var once_Data_Foldable_foldr__344102461 sync.Once
func Get_Data_Foldable_foldr__344102461() gopurs_runtime.Value {
	once_Data_Foldable_foldr__344102461.Do(func() {
		cache_Data_Foldable_foldr__344102461 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__344102461(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dict_0_box))
})
	})
	return cache_Data_Foldable_foldr__344102461
}

var cache_Data_Foldable_foldr__3749276701 gopurs_runtime.Value
var once_Data_Foldable_foldr__3749276701 sync.Once
func Get_Data_Foldable_foldr__3749276701() gopurs_runtime.Value {
	once_Data_Foldable_foldr__3749276701.Do(func() {
		cache_Data_Foldable_foldr__3749276701 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldr__3749276701(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_Foldable_foldr__3749276701
}

var cache_Data_Foldable_foldrDefault__2858227716 gopurs_runtime.Value
var once_Data_Foldable_foldrDefault__2858227716 sync.Once
func Get_Data_Foldable_foldrDefault__2858227716() gopurs_runtime.Value {
	once_Data_Foldable_foldrDefault__2858227716.Do(func() {
		cache_Data_Foldable_foldrDefault__2858227716 = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldrDefault__2858227716(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_Data_Foldable_foldrDefault__2858227716
}

var cache_Data_Foldable_foldrDefault__2151204251 gopurs_runtime.Value
var once_Data_Foldable_foldrDefault__2151204251 sync.Once
func Get_Data_Foldable_foldrDefault__2151204251() gopurs_runtime.Value {
	once_Data_Foldable_foldrDefault__2151204251.Do(func() {
		cache_Data_Foldable_foldrDefault__2151204251 = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldrDefault__2151204251(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_Data_Foldable_foldrDefault__2151204251
}

var cache_Data_Foldable_foldrDefault__3288778237 gopurs_runtime.Value
var once_Data_Foldable_foldrDefault__3288778237 sync.Once
func Get_Data_Foldable_foldrDefault__3288778237() gopurs_runtime.Value {
	once_Data_Foldable_foldrDefault__3288778237.Do(func() {
		cache_Data_Foldable_foldrDefault__3288778237 = gopurs_runtime.Func3(func(c_0_box gopurs_runtime.Value, u_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_foldrDefault__3288778237(c_0_box, u_1_box, xs_2_box)
})
	})
	return cache_Data_Foldable_foldrDefault__3288778237
}

var cache_Data_Foldable_intercalate__3813868388 gopurs_runtime.Value
var once_Data_Foldable_intercalate__3813868388 sync.Once
func Get_Data_Foldable_intercalate__3813868388() gopurs_runtime.Value {
	once_Data_Foldable_intercalate__3813868388.Do(func() {
		cache_Data_Foldable_intercalate__3813868388 = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, sep_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Foldable_intercalate__3813868388(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), sep_1_box.StrVal(), xs_2_box))
})
	})
	return cache_Data_Foldable_intercalate__3813868388
}

var cache_Data_Foldable_intercalate__3939234276 gopurs_runtime.Value
var once_Data_Foldable_intercalate__3939234276 sync.Once
func Get_Data_Foldable_intercalate__3939234276() gopurs_runtime.Value {
	once_Data_Foldable_intercalate__3939234276.Do(func() {
		cache_Data_Foldable_intercalate__3939234276 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_intercalate__3939234276(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1_box))
})
	})
	return cache_Data_Foldable_intercalate__3939234276
}

var cache_Data_Foldable_intercalate__2937349250 gopurs_runtime.Value
var once_Data_Foldable_intercalate__2937349250 sync.Once
func Get_Data_Foldable_intercalate__2937349250() gopurs_runtime.Value {
	once_Data_Foldable_intercalate__2937349250.Do(func() {
		cache_Data_Foldable_intercalate__2937349250 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_intercalate__2937349250(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Foldable_intercalate__2937349250
}

var cache_Data_Foldable_length__854370588 gopurs_runtime.Value
var once_Data_Foldable_length__854370588 sync.Once
func Get_Data_Foldable_length__854370588() gopurs_runtime.Value {
	once_Data_Foldable_length__854370588.Do(func() {
		cache_Data_Foldable_length__854370588 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_length__854370588(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_Foldable_length__854370588
}

var cache_Data_Foldable_length__1958096179 gopurs_runtime.Value
var once_Data_Foldable_length__1958096179 sync.Once
func Get_Data_Foldable_length__1958096179() gopurs_runtime.Value {
	once_Data_Foldable_length__1958096179.Do(func() {
		cache_Data_Foldable_length__1958096179 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_length__1958096179(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_Foldable_length__1958096179
}

var cache_Data_Foldable_length__949294460 gopurs_runtime.Value
var once_Data_Foldable_length__949294460 sync.Once
func Get_Data_Foldable_length__949294460() gopurs_runtime.Value {
	once_Data_Foldable_length__949294460.Do(func() {
		cache_Data_Foldable_length__949294460 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_length__949294460(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_length__949294460
}

var cache_Data_Foldable_length__1822702871 gopurs_runtime.Value
var once_Data_Foldable_length__1822702871 sync.Once
func Get_Data_Foldable_length__1822702871() gopurs_runtime.Value {
	once_Data_Foldable_length__1822702871.Do(func() {
		cache_Data_Foldable_length__1822702871 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_length__1822702871(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_length__1822702871
}

var cache_Data_Foldable_length__3154555538 gopurs_runtime.Value
var once_Data_Foldable_length__3154555538 sync.Once
func Get_Data_Foldable_length__3154555538() gopurs_runtime.Value {
	once_Data_Foldable_length__3154555538.Do(func() {
		cache_Data_Foldable_length__3154555538 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_length__3154555538(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_length__3154555538
}

var cache_Data_Foldable_length__2422281689 gopurs_runtime.Value
var once_Data_Foldable_length__2422281689 sync.Once
func Get_Data_Foldable_length__2422281689() gopurs_runtime.Value {
	once_Data_Foldable_length__2422281689.Do(func() {
		cache_Data_Foldable_length__2422281689 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_length__2422281689(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_length__2422281689
}

var cache_Data_Foldable_length__4007820284 gopurs_runtime.Value
var once_Data_Foldable_length__4007820284 sync.Once
func Get_Data_Foldable_length__4007820284() gopurs_runtime.Value {
	once_Data_Foldable_length__4007820284.Do(func() {
		cache_Data_Foldable_length__4007820284 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_length__4007820284(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dictSemiring_1_box))
})
	})
	return cache_Data_Foldable_length__4007820284
}

var cache_Data_Foldable_maximumBy__110571494 gopurs_runtime.Value
var once_Data_Foldable_maximumBy__110571494 sync.Once
func Get_Data_Foldable_maximumBy__110571494() gopurs_runtime.Value {
	once_Data_Foldable_maximumBy__110571494.Do(func() {
		cache_Data_Foldable_maximumBy__110571494 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_maximumBy__110571494(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), cmp_1_box)
})
	})
	return cache_Data_Foldable_maximumBy__110571494
}

var cache_Data_Foldable_minimumBy__110571494 gopurs_runtime.Value
var once_Data_Foldable_minimumBy__110571494 sync.Once
func Get_Data_Foldable_minimumBy__110571494() gopurs_runtime.Value {
	once_Data_Foldable_minimumBy__110571494.Do(func() {
		cache_Data_Foldable_minimumBy__110571494 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_minimumBy__110571494(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), cmp_1_box)
})
	})
	return cache_Data_Foldable_minimumBy__110571494
}

var cache_Data_Foldable_monoidFreeMonoidTree__2615096836 gopurs_runtime.Value
var once_Data_Foldable_monoidFreeMonoidTree__2615096836 sync.Once
func Get_Data_Foldable_monoidFreeMonoidTree__2615096836() gopurs_runtime.Value {
	once_Data_Foldable_monoidFreeMonoidTree__2615096836.Do(func() {
		cache_Data_Foldable_monoidFreeMonoidTree__2615096836 = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_semigroupFreeMonoidTree()
}), gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
	})
	return cache_Data_Foldable_monoidFreeMonoidTree__2615096836
}

var cache_Data_Foldable_oneOfMap__3719016818 gopurs_runtime.Value
var once_Data_Foldable_oneOfMap__3719016818 sync.Once
func Get_Data_Foldable_oneOfMap__3719016818() gopurs_runtime.Value {
	once_Data_Foldable_oneOfMap__3719016818.Do(func() {
		cache_Data_Foldable_oneOfMap__3719016818 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_oneOfMap__3719016818(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](dictPlus_1_box))
})
	})
	return cache_Data_Foldable_oneOfMap__3719016818
}

var cache_Data_Foldable_oneOfMap__1349369970 gopurs_runtime.Value
var once_Data_Foldable_oneOfMap__1349369970 sync.Once
func Get_Data_Foldable_oneOfMap__1349369970() gopurs_runtime.Value {
	once_Data_Foldable_oneOfMap__1349369970.Do(func() {
		cache_Data_Foldable_oneOfMap__1349369970 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_oneOfMap__1349369970(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Plus_Plus](dictPlus_1_box))
})
	})
	return cache_Data_Foldable_oneOfMap__1349369970
}

var cache_Data_Foldable_semigroupFreeMonoidTree__2398658907 gopurs_runtime.Value
var once_Data_Foldable_semigroupFreeMonoidTree__2398658907 sync.Once
func Get_Data_Foldable_semigroupFreeMonoidTree__2398658907() gopurs_runtime.Value {
	once_Data_Foldable_semigroupFreeMonoidTree__2398658907.Do(func() {
		cache_Data_Foldable_semigroupFreeMonoidTree__2398658907 = gopurs_runtime.RecordDict1("append", Get_Data_Foldable_Append())
	})
	return cache_Data_Foldable_semigroupFreeMonoidTree__2398658907
}

var cache_Data_Foldable_surroundMap__3689038427 gopurs_runtime.Value
var once_Data_Foldable_surroundMap__3689038427 sync.Once
func Get_Data_Foldable_surroundMap__3689038427() gopurs_runtime.Value {
	once_Data_Foldable_surroundMap__3689038427.Do(func() {
		cache_Data_Foldable_surroundMap__3689038427 = gopurs_runtime.Func5(func(dictFoldable_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value, t_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_surroundMap__3689038427(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_1_box), d_2_box, t_3_box, f_4_box)
})
	})
	return cache_Data_Foldable_surroundMap__3689038427
}

var cache_Data_Foldable_traverse___996968168 gopurs_runtime.Value
var once_Data_Foldable_traverse___996968168 sync.Once
func Get_Data_Foldable_traverse___996968168() gopurs_runtime.Value {
	once_Data_Foldable_traverse___996968168.Do(func() {
		cache_Data_Foldable_traverse___996968168 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_traverse___996968168(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Data_Foldable_traverse___996968168
}

var cache_Data_Foldable_traverse___1507800296 gopurs_runtime.Value
var once_Data_Foldable_traverse___1507800296 sync.Once
func Get_Data_Foldable_traverse___1507800296() gopurs_runtime.Value {
	once_Data_Foldable_traverse___1507800296.Do(func() {
		cache_Data_Foldable_traverse___1507800296 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_traverse___1507800296(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Data_Foldable_traverse___1507800296
}

var cache_Data_Foldable_traverse___1229293625 gopurs_runtime.Value
var once_Data_Foldable_traverse___1229293625 sync.Once
func Get_Data_Foldable_traverse___1229293625() gopurs_runtime.Value {
	once_Data_Foldable_traverse___1229293625.Do(func() {
		cache_Data_Foldable_traverse___1229293625 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_traverse___1229293625(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Foldable_traverse___1229293625
}

var cache_Data_Foldable_traverse___3630450585 gopurs_runtime.Value
var once_Data_Foldable_traverse___3630450585 sync.Once
func Get_Data_Foldable_traverse___3630450585() gopurs_runtime.Value {
	once_Data_Foldable_traverse___3630450585.Do(func() {
		cache_Data_Foldable_traverse___3630450585 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_traverse___3630450585(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Foldable_traverse___3630450585
}

var cache_Data_Foldable_traverse___3202958137 gopurs_runtime.Value
var once_Data_Foldable_traverse___3202958137 sync.Once
func Get_Data_Foldable_traverse___3202958137() gopurs_runtime.Value {
	once_Data_Foldable_traverse___3202958137.Do(func() {
		cache_Data_Foldable_traverse___3202958137 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_traverse___3202958137(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Foldable_traverse___3202958137
}

var cache_Data_Foldable_traverse___398118279 gopurs_runtime.Value
var once_Data_Foldable_traverse___398118279 sync.Once
func Get_Data_Foldable_traverse___398118279() gopurs_runtime.Value {
	once_Data_Foldable_traverse___398118279.Do(func() {
		cache_Data_Foldable_traverse___398118279 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Foldable_traverse___398118279(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Data_Foldable_traverse___398118279
}

type Constructor_Data_Foldable_Empty struct {
	Rc uint32
}


type Constructor_Data_Foldable_Node struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


type Constructor_Data_Foldable_Append struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Data_Foldable_Foldable struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4280266298] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Foldable_Foldable)(ptr)
		_ = c
		switch key {
		case "foldMap": return gopurs_runtime.Box(c.V0)
		case "foldl": return gopurs_runtime.Box(c.V1)
		case "foldr": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Foldable_Foldable: " + key)
		}
	}
}


func Call_Data_Foldable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Foldable_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Foldable_identity2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Foldable_Foldable_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Foldable_foldr(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_indexr(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, idx_1_loop int64) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var idx_1 int64 = idx_1_loop
_ = idx_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(cursor_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.RecordGet(cursor_3, "elem")
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
__t3 = cursor_3
goto end_branch_3
} else {

}
}
{
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(cursor_3, "pos").IntVal) == (idx_1) {
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_2})}, gopurs_runtime.Int(gopurs_runtime.RecordGet(cursor_3, "pos").IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.RecordGet(cursor_3, "elem")))}, gopurs_runtime.Int((gopurs_runtime.RecordGet(cursor_3, "pos").IntVal) + (1)))
}
end_branch_1:
__t3 = __t1
}
end_branch_3:
return __t3
})
}), gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(0)))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_2_0, x_3), "elem")))}
})
}

func Call_Data_Foldable_null(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
}), gopurs_runtime.Bool(true))
}

func Call_Data_Foldable_oneOf(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictPlus_1_loop *Constructor_Control_Plus_Plus) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *Constructor_Control_Plus_Plus = dictPlus_1_loop
_ = dictPlus_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictPlus_1.V0), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Box(dictPlus_1.V1))
}

func Call_Data_Foldable_oneOfMap(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictPlus_1_loop *Constructor_Control_Plus_Plus) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *Constructor_Control_Plus_Plus = dictPlus_1_loop
_ = dictPlus_1
// TAST (Let): alt_2_0 -> gopurs_runtime.Value
alt_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictPlus_1.V0), gopurs_runtime.Value{}), "alt")
_ = alt_2_0
// TAST (Let): empty_3_1 -> gopurs_runtime.Value
empty_3_1 := gopurs_runtime.Box(dictPlus_1.V1)
_ = empty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_2_0, gopurs_runtime.Apply(f_4, x_5))
}), empty_3_1)
})
}

func Call_Data_Foldable_traverse_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 -> *Constructor_Data_Functor_Functor
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): applySecond_1_0 -> gopurs_runtime.Value
applySecond_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), a_3), b_4)
})
})
_ = applySecond_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(f_3, x_4))
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit()))
})
})
}

func Call_Data_Foldable_for_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 -> *Constructor_Data_Functor_Functor
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): applySecond_1_0 -> gopurs_runtime.Value
applySecond_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), a_3), b_4)
})
})
_ = applySecond_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(a_4, x_5))
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit()), b_3)
})
})
})
}

func Call_Data_Foldable_sequence_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, dictFoldable_1_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var dictFoldable_1 *Constructor_Data_Foldable_Foldable = dictFoldable_1_loop
_ = dictFoldable_1
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): applySecond_2_0 -> gopurs_runtime.Value
applySecond_2_0 := gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return x_7
})
}), a_4), b_5)
})
})
_ = applySecond_2_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_1.V2), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_2_0, x_3)
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit()))
}

func Call_Data_Foldable_foldl(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_indexl(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, idx_1_loop int64) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var idx_1 int64 = idx_1_loop
_ = idx_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(cursor_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.RecordGet(cursor_2, "elem")
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
__t3 = cursor_2
goto end_branch_3
} else {

}
}
{
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(cursor_2, "pos").IntVal) == (idx_1) {
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_3})}, gopurs_runtime.Int(gopurs_runtime.RecordGet(cursor_2, "pos").IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.RecordGet(cursor_2, "elem")))}, gopurs_runtime.Int((gopurs_runtime.RecordGet(cursor_2, "pos").IntVal) + (1)))
}
end_branch_1:
__t3 = __t1
}
end_branch_3:
return __t3
})
}), gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(0)))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_2_0, x_3), "elem")))}
})
}

func Call_Data_Foldable_intercalate(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictMonoid_1_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 -> gopurs_runtime.Value
mempty_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = mempty_3_1
return gopurs_runtime.Func(func(sep_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_6, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_7, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.RecordGet(v_6, "acc"), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), sep_4, v1_7)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", mempty_3_1, gopurs_runtime.Bool(true)), xs_5), "acc")
})
})
}

func Call_Data_Foldable_length(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemiring_1_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemiring_1.V0), gopurs_runtime.Box(dictSemiring_1.V2), c_2)
})
}), gopurs_runtime.Box(dictSemiring_1.V3))
}

func Call_Data_Foldable_maximumBy(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t3 = &Constructor_Data_Maybe_Just{1, v1_3}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
// TAST (Let): __local_var_4_0 -> uint32
__local_var_4_0 := uint32(gopurs_runtime.Apply2(cmp_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0, v1_3).IntVal)
_ = __local_var_4_0
var __t1 bool
{
if (__local_var_4_0 == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 == 380165415) {
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
__t2 = (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = v1_3
}
end_branch_2:
__t3 = &Constructor_Data_Maybe_Just{1, __t2}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}

func Call_Data_Foldable_maximum(dictOrd_0_loop *Constructor_Data_Ord_Ord, dictFoldable_1_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 *Constructor_Data_Foldable_Foldable = dictFoldable_1_loop
_ = dictFoldable_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_1.V1), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t3 = &Constructor_Data_Maybe_Just{1, v1_3}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
// TAST (Let): __local_var_4_0 -> uint32
__local_var_4_0 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0, v1_3).IntVal)
_ = __local_var_4_0
var __t1 bool
{
if (__local_var_4_0 == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 == 380165415) {
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
__t2 = (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = v1_3
}
end_branch_2:
__t3 = &Constructor_Data_Maybe_Just{1, __t2}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}

func Call_Data_Foldable_minimumBy(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t3 = &Constructor_Data_Maybe_Just{1, v1_3}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
// TAST (Let): __local_var_4_0 -> uint32
__local_var_4_0 := uint32(gopurs_runtime.Apply2(cmp_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0, v1_3).IntVal)
_ = __local_var_4_0
var __t1 bool
{
if (__local_var_4_0 == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = v1_3
}
end_branch_2:
__t3 = &Constructor_Data_Maybe_Just{1, __t2}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}

func Call_Data_Foldable_minimum(dictOrd_0_loop *Constructor_Data_Ord_Ord, dictFoldable_1_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var dictFoldable_1 *Constructor_Data_Foldable_Foldable = dictFoldable_1_loop
_ = dictFoldable_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_1.V1), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t3 = &Constructor_Data_Maybe_Just{1, v1_3}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
// TAST (Let): __local_var_4_0 -> uint32
__local_var_4_0 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0, v1_3).IntVal)
_ = __local_var_4_0
var __t1 bool
{
if (__local_var_4_0 == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = v1_3
}
end_branch_2:
__t3 = &Constructor_Data_Maybe_Just{1, __t2}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}

func Call_Data_Foldable_product(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemiring_1_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Box(dictSemiring_1.V1), gopurs_runtime.Box(dictSemiring_1.V2))
}

func Call_Data_Foldable_sum(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemiring_1_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Box(dictSemiring_1.V0), gopurs_runtime.Box(dictSemiring_1.V3))
}

func Call_Data_Foldable_foldMapDefaultR(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictMonoid_1_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 -> gopurs_runtime.Value
mempty_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply(f_4, x_5), acc_6)
})
}), mempty_3_1)
})
}

func Call_Data_Foldable_foldMapDefaultL(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictMonoid_1_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 -> gopurs_runtime.Value
mempty_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), acc_5, gopurs_runtime.Apply(f_4, x_6))
})
}), mempty_3_1)
})
}

func Call_Data_Foldable_foldMap(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Foldable_foldableApp(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, f_2, v_3)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, i_2, v_3)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, i_2, v_3)
})
})
}))
}

func Call_Data_Foldable_foldableCompose(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_2))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_2))}, f_3), v_4)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2), i_3, v_4)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_0 -> gopurs_runtime.Value
__local_var_5_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2)
_ = __local_var_5_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_5_0, a_7, b_6)
})
}), i_3, v_4)
})
})
}))
}

func Call_Data_Foldable_foldableCoproduct(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_2))}, f_3)
_ = __local_var_4_0
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_2))}, f_3)
_ = __local_var_5_1
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply(__local_var_4_0, (*Constructor_Data_Either_Left)(v2_6.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply(__local_var_5_1, (*Constructor_Data_Either_Right)(v2_6.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, z_3)
_ = __local_var_4_3
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2, z_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t5 = gopurs_runtime.Apply(__local_var_4_3, (*Constructor_Data_Either_Left)(v2_6.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t5 = gopurs_runtime.Apply(__local_var_5_4, (*Constructor_Data_Either_Right)(v2_6.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_2, z_3)
_ = __local_var_4_6
// TAST (Let): __local_var_5_7 -> gopurs_runtime.Value
__local_var_5_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2, z_3)
_ = __local_var_5_7
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t8 = gopurs_runtime.Apply(__local_var_4_6, (*Constructor_Data_Either_Left)(v2_6.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t8 = gopurs_runtime.Apply(__local_var_5_7, (*Constructor_Data_Either_Right)(v2_6.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
})
}))
}

func Call_Data_Foldable_foldableProduct(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_0.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_2))}, f_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_2))}, f_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, z_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2, z_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0)
})
})
}))
}

func Call_Data_Foldable_foldlDefault(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFreeMonoidTree(), "foldl"), c_1, u_2, gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Foldable_monoidFreeMonoidTree()))}, Get_Data_Foldable_Node(), xs_3))
}

func Call_Data_Foldable_foldrDefault(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFreeMonoidTree(), "foldr"), c_1, u_2, gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Foldable_monoidFreeMonoidTree()))}, Get_Data_Foldable_Node(), xs_3))
}

func Call_Data_Foldable_lookup(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictEq_1_loop *Constructor_Data_Eq_Eq, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictEq_1 *Constructor_Data_Eq_Eq = dictEq_1_loop
_ = dictEq_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Maybe_First_monoidFirst()))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_1.V0), a_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1}
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
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, x_4)
})
}

func Call_Data_Foldable_surroundMap(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup, d_2_loop gopurs_runtime.Value, t_3_loop gopurs_runtime.Value, f_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_1_loop
_ = dictSemigroup_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
var t_3 gopurs_runtime.Value = t_3_loop
_ = t_3
var f_4 gopurs_runtime.Value = f_4_loop
_ = f_4
// TAST (Let): semigroupEndo1_5_0 -> gopurs_runtime.Value
semigroupEndo1_5_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_5, gopurs_runtime.Apply(v1_6, x_7))
})
})
}))
_ = semigroupEndo1_5_0
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_5_0
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
}))))}, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), d_2, gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), gopurs_runtime.Apply(t_3, a_5), m_6))
})
}), f_4, d_2)
}

func Call_Data_Foldable_surround(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup, d_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_1_loop
_ = dictSemigroup_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
// TAST (Let): semigroupEndo1_4_0 -> gopurs_runtime.Value
semigroupEndo1_4_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Apply(v1_5, x_6))
})
})
}))
_ = semigroupEndo1_4_0
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_4_0
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
}))))}, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), d_2, gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), a_4, m_5))
})
}), f_3, d_2)
}

func Call_Data_Foldable_foldM(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictMonad_1_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonad_1 *Constructor_Control_Monad_Monad = dictMonad_1_loop
_ = dictMonad_1
// TAST (Let): Bind1_2_0 -> *Constructor_Control_Bind_Bind
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_1.V1), gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): Applicative0_3_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_1.V0), gopurs_runtime.Value{}))
_ = Applicative0_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b0_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), b_6, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_4, a_8, a_7)
}))
})
}), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), b0_5))
})
})
}

func Call_Data_Foldable_fold(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictMonoid_1_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}

func Call_Data_Foldable_findMap(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(p_1, v1_3))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}

func Call_Data_Foldable_find(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if ((v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil)) && ((gopurs_runtime.Apply(p_1, v1_3).IntVal) != (0)) {
__t0 = &Constructor_Data_Maybe_Just{1, v1_3}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}

func Call_Data_Foldable_any(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): semigroupDisj1_2_0 -> gopurs_runtime.Value
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictHeytingAlgebra_1.V1), v_2, v1_3)
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), gopurs_runtime.Box(dictHeytingAlgebra_1.V2)))
}

func Call_Data_Foldable_elem(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): semigroupDisj1_1_1 -> gopurs_runtime.Value
semigroupDisj1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_1.IntVal) != (0)) || ((v1_2.IntVal) != (0)))
})
}))
_ = semigroupDisj1_1_1
// TAST (Let): any1_1_0 -> gopurs_runtime.Value
any1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_1
}), gopurs_runtime.Bool(false)))
_ = any1_1_0
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(any1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq_2, "eq"), x_3))
})
})
}

func Call_Data_Foldable_notElem(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictEq_1_loop *Constructor_Data_Eq_Eq, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictEq_1 *Constructor_Data_Eq_Eq = dictEq_1_loop
_ = dictEq_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
// TAST (Let): semigroupDisj1_3_1 -> gopurs_runtime.Value
semigroupDisj1_3_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_3.IntVal) != (0)) || ((v1_4.IntVal) != (0)))
})
}))
_ = semigroupDisj1_3_1
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_3_1
}), gopurs_runtime.Bool(false)), gopurs_runtime.Apply(gopurs_runtime.Box(dictEq_1.V0), x_2))
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(__local_var_3_0, x_4).IntVal) != (0)) != (true))
})
}

func Call_Data_Foldable_or(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): semigroupDisj1_2_0 -> gopurs_runtime.Value
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictHeytingAlgebra_1.V1), v_2, v1_3)
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), gopurs_runtime.Box(dictHeytingAlgebra_1.V2)), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}

func Call_Data_Foldable_all(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): semigroupConj1_2_0 -> gopurs_runtime.Value
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictHeytingAlgebra_1.V0), v_2, v1_3)
})
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
}), gopurs_runtime.Box(dictHeytingAlgebra_1.V5)))
}

func Call_Data_Foldable_and(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): semigroupConj1_2_0 -> gopurs_runtime.Value
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictHeytingAlgebra_1.V0), v_2, v1_3)
})
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
}), gopurs_runtime.Box(dictHeytingAlgebra_1.V5)), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}

func Call_Data_Foldable_all__4179648253(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): semigroupConj1_2_0 -> gopurs_runtime.Value
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictHeytingAlgebra_1.V0), v_2, v1_3)
})
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
}), gopurs_runtime.Box(dictHeytingAlgebra_1.V5)))
}

func Call_Data_Foldable_any__4179648253(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): semigroupDisj1_2_0 -> gopurs_runtime.Value
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictHeytingAlgebra_1.V1), v_2, v1_3)
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), gopurs_runtime.Box(dictHeytingAlgebra_1.V2)))
}

func Call_Data_Foldable_any__4041742601(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
// TAST (Let): semigroupDisj1_2_0 -> gopurs_runtime.Value
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_2.IntVal) != (0)) || ((v1_3.IntVal) != (0)))
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), gopurs_runtime.Bool(false)), __eta0_0, __eta1_1)
}

func Call_Data_Foldable_any__1385259145(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
// TAST (Let): semigroupDisj1_2_0 -> gopurs_runtime.Value
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_2.IntVal) != (0)) || ((v1_3.IntVal) != (0)))
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), gopurs_runtime.Bool(false)), __eta0_0, __eta1_1)
}

func Call_Data_Foldable_any__842931401(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
// TAST (Let): semigroupDisj1_2_0 -> gopurs_runtime.Value
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_2.IntVal) != (0)) || ((v1_3.IntVal) != (0)))
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), gopurs_runtime.Bool(false)), __eta0_0, __eta1_1)
}

func Call_Data_Foldable_elem__2343844090(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): semigroupDisj1_1_1 -> gopurs_runtime.Value
semigroupDisj1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((v_1.IntVal) != (0)) || ((v1_2.IntVal) != (0)))
})
}))
_ = semigroupDisj1_1_1
// TAST (Let): any1_1_0 -> gopurs_runtime.Value
any1_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_1
}), gopurs_runtime.Bool(false)))
_ = any1_1_0
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(any1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq_2, "eq"), x_3))
})
})
}

func Call_Data_Foldable_fold__910331789(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictMonoid_1_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}

func Call_Data_Foldable_foldMap__4098395794(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Foldable_foldMap__4151846418(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Foldable_foldMap__2966595236(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Foldable_monoidFreeMonoidTree()))})
}

func Call_Data_Foldable_foldMap__4130609395(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Maybe_First_monoidFirst()))})
}

func Call_Data_Foldable_foldMap__1118659089(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
// TAST (Let): semigroupEndo1_1_0 -> gopurs_runtime.Value
semigroupEndo1_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(v1_2, x_3))
})
})
}))
_ = semigroupEndo1_1_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))))})
}

func Call_Data_Foldable_foldMap__1315280116(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Foldable_foldMap__1811898306(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Foldable_monoidFreeMonoidTree()))})
}

func Call_Data_Foldable_foldMap__193737345(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Interval_Duration_Iso_monoidAdditive()).V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Float(acc_3.FloatVal()))
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(__eta0_0, x_5))
})
}), gopurs_runtime.Float(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Interval_Duration_Iso_monoidAdditive()).V1).FloatVal()), __eta1_1)
}

func Call_Data_Foldable_foldMap__3661646260(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Semigroup0_1_0.V0), acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}

func Call_Data_Foldable_foldMap__4073832436(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t1 = mempty_1_0
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just)(v1_3.UnsafePtr).V0)
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
})
}

func Call_Data_Foldable_foldMap__3562626100(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V0), __eta0_0, __eta1_1)
}

func Call_Data_Foldable_foldMap__2350611220(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V0), __eta0_0, __eta1_1)
}

func Call_Data_Foldable_foldMap__3557693114(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Foldable_foldMap__2556604300(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Foldable_foldl__1422885860(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3850309840(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__2111289130(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__94807652(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__2699291984(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3041692656(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__66388714(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__506543652(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__1671904522(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__2602334544(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__4192477084(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3272087748(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__165683952(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__267332164(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__371433392(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3234403824(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__1656262032(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3559959056(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3131354468(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__1601164432(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3893253828(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__1148906672(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__2656621979(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3422238939(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__2151204251(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__4099223803(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return func() gopurs_runtime.Value {
arr_val_foldlArray0 := __eta2_2
_ = arr_val_foldlArray0
res_go_foldlArray0 := __eta1_1
_ = res_go_foldlArray0
arr_go_foldlArray0 := (*[]gopurs_runtime.Value)(arr_val_foldlArray0.UnsafePtr)
_ = arr_go_foldlArray0
for _, v_foldlArray0 := range *arr_go_foldlArray0 {
res_go_foldlArray0 = gopurs_runtime.Apply2(__eta0_0, res_go_foldlArray0, v_foldlArray0)
}
return res_go_foldlArray0
}()
}

func Call_Data_Foldable_foldl__2123442907(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__524683195(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_4 gopurs_runtime.Value
go__go_1_0_4 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_4:
for {
if false { continue go__go_1_0_4 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (v_4_1).V0)
xs_3_loop = (v_4_1).V1
continue go__go_1_0_4
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
return go__go_1_0_4
}

func Call_Data_Foldable_foldl__1712912315(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_5 gopurs_runtime.Value
go__go_1_0_5 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_5:
for {
if false { continue go__go_1_0_5 }
var b_2 *Constructor_Data_List_Types_Cons = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 *Constructor_Data_List_Types_Cons
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_5
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
return go__go_1_0_5
}

func Call_Data_Foldable_foldl__2138619643(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__4056605371(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__53736539(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__22573083(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__393765499(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3306117403(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_6 gopurs_runtime.Value
go__go_1_0_6 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *Constructor_Data_Tuple_Tuple = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](b_2_loop_val)
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_6:
for {
if false { continue go__go_1_0_6 }
var b_2 *Constructor_Data_Tuple_Tuple = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
_ = v_4_1
var __t2 *Constructor_Data_Tuple_Tuple
{
if (v_4_1 == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (v_4_1 != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(op_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}, (v_4_1).V0))
xs_3_loop = (v_4_1).V1
continue go__go_1_0_6
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
return go__go_1_0_6
}

func Call_Data_Foldable_foldl__2159564571(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_7 gopurs_runtime.Value
go__go_1_0_7 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *Constructor_Data_Tuple_Tuple = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_7:
for {
if false { continue go__go_1_0_7 }
var b_2 *Constructor_Data_Tuple_Tuple = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 *Constructor_Data_Tuple_Tuple
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_7
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return go__go_1_0_7
}

func Call_Data_Foldable_foldl__22791451(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3736093275(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__3785384859(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__446290811(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__1459781277(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_8 gopurs_runtime.Value
go__go_1_0_8 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_0_8:
for {
if false { continue go__go_1_0_8 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t5 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, acc_2, (*Constructor_Data_Foldable_Node)(lhs_3.UnsafePtr).V0)
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_0_8
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2812549951) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V0
rhs_4_loop = rhs_4
continue go__go_1_0_8
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
var __t1 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V0
rhs_4_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V1
continue go__go_1_0_8
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V0
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Append{1, (*Constructor_Data_Foldable_Append)(lhs_3.UnsafePtr).V1, rhs_4})}
continue go__go_1_0_8
__t1 = gopurs_runtime.Value{}
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
continue go__go_1_0_8
__t4 = gopurs_runtime.Value{}
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
}()
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_0_8, a_2, b_3, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
})
})
}

func Call_Data_Foldable_foldl__3288778237(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 237113226) {
__t0 = gopurs_runtime.Apply2(v_0, gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd)(v2_2.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 1992629780) {
__t0 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_DurationEnd)(v2_2.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2020675835) {
__t0 = gopurs_runtime.Apply2(v_0, v1_1, (*Constructor_Data_Interval_StartDuration)(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = v1_1
}
end_branch_0:
return __t0
}

func Call_Data_Foldable_foldl__3504930205(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_9 gopurs_runtime.Value
go__go_1_0_9 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop int64 = b_2_loop_val.IntVal
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_9:
for {
if false { continue go__go_1_0_9 }
var b_2 int64 = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
_ = v_4_1
var __t2 int64
{
if (v_4_1 == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (v_4_1 != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, gopurs_runtime.Int(b_2), (v_4_1).V0).IntVal
xs_3_loop = (v_4_1).V1
continue go__go_1_0_9
__t2 = gopurs_runtime.Value{}.IntVal
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_2:
return gopurs_runtime.Int(__t2)
}
}()
})
})
return go__go_1_0_9
}

func Call_Data_Foldable_foldl__3379885725(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_10 gopurs_runtime.Value
go__go_1_0_10 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop int64 = b_2_loop_val.IntVal
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_10:
for {
if false { continue go__go_1_0_10 }
var b_2 int64 = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
_ = v_4_1
var __t2 int64
{
if (v_4_1 == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (v_4_1 != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, gopurs_runtime.Int(b_2), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_4_1).V0))}).IntVal
xs_3_loop = (v_4_1).V1
continue go__go_1_0_10
__t2 = gopurs_runtime.Value{}.IntVal
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_2:
return gopurs_runtime.Int(__t2)
}
}()
})
})
return go__go_1_0_10
}

func Call_Data_Foldable_foldl__3737487037(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_11 gopurs_runtime.Value
go__go_1_0_11 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop string = b_2_loop_val.StrVal()
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_11:
for {
if false { continue go__go_1_0_11 }
var b_2 string = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
_ = v_4_1
var __t2 string
{
if (v_4_1 == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (v_4_1 != nil) {
b_2_loop = gopurs_runtime.Apply2(op_0, gopurs_runtime.Str(b_2), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_4_1).V0))}).StrVal()
xs_3_loop = (v_4_1).V1
continue go__go_1_0_11
__t2 = gopurs_runtime.Value{}.StrVal()
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_2:
return gopurs_runtime.Str(__t2)
}
}()
})
})
return go__go_1_0_11
}

func Call_Data_Foldable_foldl__1985071933(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_12 gopurs_runtime.Value
go__go_1_0_12 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_12:
for {
if false { continue go__go_1_0_12 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (v_4_1).V0)
xs_3_loop = (v_4_1).V1
continue go__go_1_0_12
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
return go__go_1_0_12
}

func Call_Data_Foldable_foldl__536153533(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_13 gopurs_runtime.Value
go__go_1_0_13 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_13:
for {
if false { continue go__go_1_0_13 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_4_1).V0))})
xs_3_loop = (v_4_1).V1
continue go__go_1_0_13
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
return go__go_1_0_13
}

func Call_Data_Foldable_foldl__4234493053(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_14 gopurs_runtime.Value
go__go_1_0_14 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_14:
for {
if false { continue go__go_1_0_14 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, (v_4_1).V0)
xs_3_loop = (v_4_1).V1
continue go__go_1_0_14
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
return go__go_1_0_14
}

func Call_Data_Foldable_foldl__176907901(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_15 gopurs_runtime.Value
go__go_1_0_15 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_15:
for {
if false { continue go__go_1_0_15 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
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
b_2_loop = gopurs_runtime.Apply2(op_0, b_2, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_4_1).V0))})
xs_3_loop = (v_4_1).V1
continue go__go_1_0_15
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
return go__go_1_0_15
}

func Call_Data_Foldable_foldl__170252797(op_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var go__go_1_0_16 gopurs_runtime.Value
go__go_1_0_16 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *Constructor_Data_Tuple_Tuple = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](b_2_loop_val)
var xs_3_loop gopurs_runtime.Value = xs_3_loop_val
go__go_1_0_16:
for {
if false { continue go__go_1_0_16 }
var b_2 *Constructor_Data_Tuple_Tuple = b_2_loop
_ = b_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
// TAST (Let): v_4_1 -> *Constructor_Data_List_Lazy_Types_Cons
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), xs_3))
_ = v_4_1
var __t2 *Constructor_Data_Tuple_Tuple
{
if (v_4_1 == nil) {
__t2 = b_2
goto end_branch_2
} else {

}
}
{
if (v_4_1 != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(op_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons]((v_4_1).V0))}))
xs_3_loop = (v_4_1).V1
continue go__go_1_0_16
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t2)}
}
}()
})
})
return go__go_1_0_16
}

func Call_Data_Foldable_foldl__1754241693(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_17 gopurs_runtime.Value
go__go_1_0_17 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop int64 = b_2_loop_val.IntVal
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_17:
for {
if false { continue go__go_1_0_17 }
var b_2 int64 = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 int64
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, gopurs_runtime.Int(b_2), (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0).IntVal
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_17
__t1 = gopurs_runtime.Value{}.IntVal
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_1:
return gopurs_runtime.Int(__t1)
}
}()
})
})
return go__go_1_0_17
}

func Call_Data_Foldable_foldl__3943124669(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_18 gopurs_runtime.Value
go__go_1_0_18 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_18:
for {
if false { continue go__go_1_0_18 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_18
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
return go__go_1_0_18
}

func Call_Data_Foldable_foldl__396932925(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_19 gopurs_runtime.Value
go__go_1_0_19 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_19:
for {
if false { continue go__go_1_0_19 }
var b_2 *Constructor_Data_List_Types_Cons = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 *Constructor_Data_List_Types_Cons
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_19
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
return go__go_1_0_19
}

func Call_Data_Foldable_foldl__2928402749(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_20 gopurs_runtime.Value
go__go_1_0_20 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_20:
for {
if false { continue go__go_1_0_20 }
var b_2 *Constructor_Data_NonEmpty_NonEmpty = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 *Constructor_Data_NonEmpty_NonEmpty
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_20
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return go__go_1_0_20
}

func Call_Data_Foldable_foldl__255626813(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_21 gopurs_runtime.Value
go__go_1_0_21 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *Constructor_Data_Tuple_Tuple = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_21:
for {
if false { continue go__go_1_0_21 }
var b_2 *Constructor_Data_Tuple_Tuple = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 *Constructor_Data_Tuple_Tuple
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_2)}, (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_21
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t1)}
}
}()
})
})
return go__go_1_0_21
}

func Call_Data_Foldable_foldl__3915700701(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_22 gopurs_runtime.Value
go__go_1_0_22 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_22:
for {
if false { continue go__go_1_0_22 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_22
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
return go__go_1_0_22
}

func Call_Data_Foldable_foldl__3459294429(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_23 gopurs_runtime.Value
go__go_1_0_23 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_23:
for {
if false { continue go__go_1_0_23 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_23
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
return go__go_1_0_23
}

func Call_Data_Foldable_foldl__512483965(f_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 *Constructor_Data_Map_Internal_Node = x_1_loop
_ = x_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), f_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(x_1)})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(Get_Data_Set_toList(), x_3))
})
}

func Call_Data_Foldable_foldl__3016550397(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply2(v_0, v1_1, (v2_2).V0)
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

func Call_Data_Foldable_foldl__1714316381(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__380919197(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__992072381(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Foldable_foldl__2188030845(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V1), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_Foldable_foldl__1444272061(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V1), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_Foldable_foldr__1038841770(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2858227716(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__3728540540(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2111289130(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__208886460(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__926146538(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2512763050(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__3673994608(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__919612668(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2232849770(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__3630705947(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2151204251(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__3675782427(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2403185435(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2829803163(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__4105571355(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__3591001499(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2492367323(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__391354971(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2671482779(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__1030499675(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_Foldable_foldr__3948834331(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__1687192379(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__3482737755(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__1459781277(fn_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var fn_0 gopurs_runtime.Value = fn_0_loop
_ = fn_0
var go__go_1_0_24 gopurs_runtime.Value
go__go_1_0_24 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_0_24:
for {
if false { continue go__go_1_0_24 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t5 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, (*Constructor_Data_Foldable_Node)(rhs_4.UnsafePtr).V0, acc_2)
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_0_24
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2812549951) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V0
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V1
continue go__go_1_0_24
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
var __t1 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V0
rhs_4_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V1
continue go__go_1_0_24
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Data_Foldable_Append{1, lhs_3, (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V0})}
rhs_4_loop = (*Constructor_Data_Foldable_Append)(rhs_4.UnsafePtr).V1
continue go__go_1_0_24
__t1 = gopurs_runtime.Value{}
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
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
var __t4 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
__t4 = acc_2
goto end_branch_4
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_0_24
__t4 = gopurs_runtime.Value{}
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
}()
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_0_24, a_2, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}, b_3)
})
})
}

func Call_Data_Foldable_foldr__3288778237(x_0_loop gopurs_runtime.Value, u_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var u_1 gopurs_runtime.Value = u_1_loop
_ = u_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFreeMonoidTree(), "foldr"), x_0, u_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_foldableInterval(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Foldable_monoidFreeMonoidTree()))}, Get_Data_Foldable_Node(), xs_2))
}

func Call_Data_Foldable_foldr__1985071933(op_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 gopurs_runtime.Value = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_cons(), a_4, b_3)
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
}

func Call_Data_Foldable_foldr__3192890333(op_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 gopurs_runtime.Value = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_cons(), a_4, b_3)
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
}

func Call_Data_Foldable_foldr__2389967549(op_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 gopurs_runtime.Value = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_cons(), a_4, b_3)
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
}

func Call_Data_Foldable_foldr__1278383325(op_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 gopurs_runtime.Value = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_cons(), a_4, b_3)
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
}

func Call_Data_Foldable_foldr__2492628765(op_0_loop gopurs_runtime.Value, z_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 gopurs_runtime.Value = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), z_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_cons(), a_4, b_3)
})
}), Get_Data_List_Lazy_Types_nil(), xs_2))
}

func Call_Data_Foldable_foldr__3433277981(op_0_loop gopurs_runtime.Value, z_1_loop *Constructor_Data_Tuple_Tuple, xs_2_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var z_1 *Constructor_Data_Tuple_Tuple = z_1_loop
_ = z_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(op_0, a_4, b_3)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(z_1)}, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_cons(), a_4, b_3)
})
}), Get_Data_List_Lazy_Types_nil(), xs_2)))
}

func Call_Data_Foldable_foldr__3943124669(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_0
var go__go_3_2_25 gopurs_runtime.Value
go__go_3_2_25 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_25:
for {
if false { continue go__go_3_2_25 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_25
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
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_25, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](Get_Data_List_Types_Nil()))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_Data_Foldable_foldr__2979608669(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_List_Types_Cons = b_1_loop
_ = b_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_1)})
_ = __local_var_2_0
var go__go_3_2_26 gopurs_runtime.Value
go__go_3_2_26 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_26:
for {
if false { continue go__go_3_2_26 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_26
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
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_26, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](Get_Data_List_Types_Nil()))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_Data_Foldable_foldr__4137485405(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_0
var go__go_3_2_27 gopurs_runtime.Value
go__go_3_2_27 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_27:
for {
if false { continue go__go_3_2_27 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_27
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
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_27, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](Get_Data_List_Types_Nil()))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_Data_Foldable_foldr__3489910557(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_0
var go__go_3_2_28 gopurs_runtime.Value
go__go_3_2_28 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_28:
for {
if false { continue go__go_3_2_28 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), (*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_28
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
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_28, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](Get_Data_List_Types_Nil()))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_Data_Foldable_foldr__3234921885(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_Tuple_Tuple) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_Tuple_Tuple = b_1_loop
_ = b_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(b_1)})
_ = __local_var_2_0
var go__go_3_2_29 gopurs_runtime.Value
go__go_3_2_29 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_29:
for {
if false { continue go__go_3_2_29 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = v_4
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}))
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_2_29
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
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_29, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](Get_Data_List_Types_Nil()))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_Data_Foldable_foldr__3235634269(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(Get_Data_Set_toList(), x_3))
})
}

func Call_Data_Foldable_foldr__530094749(f_0_loop gopurs_runtime.Value, z_1_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var z_1 *Constructor_Data_List_Types_Cons = z_1_loop
_ = z_1
var go__go_2_0_30 gopurs_runtime.Value
_ = go__go_2_0_30
go__go_2_0_30 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_List_Types_Cons
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](__local_var_4)
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(f_0, (*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(__local_var_3.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](__local_var_4))})))})))}))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.UncurriedApp2(go__go_2_0_30, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](m_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(z_1)})))}
})
}

func Call_Data_Foldable_foldr__4254578461(f_0_loop gopurs_runtime.Value, x_1_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 *Constructor_Data_Map_Internal_Node = x_1_loop
_ = x_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), f_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(x_1)})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(Get_Data_Set_toList(), x_3))
})
}

func Call_Data_Foldable_foldr__2178954717(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldr"), f_0, x_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(Get_Data_Set_toList(), x_3))
})
}

func Call_Data_Foldable_foldr__3016550397(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *Constructor_Data_Maybe_Just = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2 == nil) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
if (v2_2 != nil) {
__t0 = gopurs_runtime.Apply2(v_0, (v2_2).V0, v1_1)
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

func Call_Data_Foldable_foldr__2147034525(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2737211997(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__2188030845(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V2), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_Foldable_foldr__344102461(dict_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Foldable_Foldable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Foldable_foldr__3749276701(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Lazy_Types_foldableNonEmpty()).V2), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_Foldable_foldrDefault__2858227716(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFreeMonoidTree(), "foldr"), c_1, u_2, gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Foldable_monoidFreeMonoidTree()))}, Get_Data_Foldable_Node(), xs_3))
}

func Call_Data_Foldable_foldrDefault__2151204251(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFreeMonoidTree(), "foldr"), c_1, u_2, gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Foldable_monoidFreeMonoidTree()))}, Get_Data_Foldable_Node(), xs_3))
}

func Call_Data_Foldable_foldrDefault__3288778237(c_0_loop gopurs_runtime.Value, u_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var c_0 gopurs_runtime.Value = c_0_loop
_ = c_0
var u_1 gopurs_runtime.Value = u_1_loop
_ = u_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFreeMonoidTree(), "foldr"), c_0, u_1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Interval_foldableInterval(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](Get_Data_Foldable_monoidFreeMonoidTree()))}, Get_Data_Foldable_Node(), xs_2))
}

func Call_Data_Foldable_intercalate__3813868388(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, sep_1_loop string, xs_2_loop gopurs_runtime.Value) string {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var sep_1 string = sep_1_loop
_ = sep_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_3, "init").IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Str(v1_4.StrVal()), gopurs_runtime.Bool(false))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Str(((gopurs_runtime.RecordGet(v_3, "acc").StrVal()) + (sep_1)) + (v1_4.StrVal())), gopurs_runtime.Bool(false))
}
end_branch_0:
return __t0
})
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Str(""), gopurs_runtime.Bool(true)), xs_2), "acc").StrVal()
}

func Call_Data_Foldable_intercalate__3939234276(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictMonoid_1_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 -> gopurs_runtime.Value
mempty_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = mempty_3_1
return gopurs_runtime.Func(func(sep_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_6, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_7, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.RecordGet(v_6, "acc"), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), sep_4, v1_7)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", mempty_3_1, gopurs_runtime.Bool(true)), xs_5), "acc")
})
})
}

func Call_Data_Foldable_intercalate__2937349250(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str(gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_2, "init").IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Str(v1_3.StrVal()), gopurs_runtime.Bool(false))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Str(((gopurs_runtime.RecordGet(v_2, "acc").StrVal()) + (__eta0_0.StrVal())) + (v1_3.StrVal())), gopurs_runtime.Bool(false))
}
end_branch_0:
return __t0
})
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Str(""), gopurs_runtime.Bool(true)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](__eta1_1))}), "acc").StrVal())
}

func Call_Data_Foldable_length__854370588(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(c_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_1.IntVal))
})
}), gopurs_runtime.Int(0))
}

func Call_Data_Foldable_length__1958096179(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(c_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_1.IntVal))
})
}), gopurs_runtime.Int(0))
}

func Call_Data_Foldable_length__949294460(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemiring_1_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemiring_1.V0), gopurs_runtime.Box(dictSemiring_1.V2), c_2)
})
}), gopurs_runtime.Box(dictSemiring_1.V3))
}

func Call_Data_Foldable_length__1822702871(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemiring_1_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemiring_1.V0), gopurs_runtime.Int(gopurs_runtime.Box(dictSemiring_1.V2).IntVal), gopurs_runtime.Int(c_2.IntVal)).IntVal)
})
}), gopurs_runtime.Int(gopurs_runtime.Box(dictSemiring_1.V3).IntVal))
}

func Call_Data_Foldable_length__3154555538(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemiring_1_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemiring_1.V0), gopurs_runtime.Box(dictSemiring_1.V2), c_2)
})
}), gopurs_runtime.Box(dictSemiring_1.V3))
}

func Call_Data_Foldable_length__2422281689(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemiring_1_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemiring_1.V0), gopurs_runtime.Int(gopurs_runtime.Box(dictSemiring_1.V2).IntVal), gopurs_runtime.Int(c_2.IntVal)).IntVal)
})
}), gopurs_runtime.Int(gopurs_runtime.Box(dictSemiring_1.V3).IntVal))
}

func Call_Data_Foldable_length__4007820284(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemiring_1_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *Constructor_Data_Semiring_Semiring = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemiring_1.V0), gopurs_runtime.Box(dictSemiring_1.V2), c_2)
})
}), gopurs_runtime.Box(dictSemiring_1.V3))
}

func Call_Data_Foldable_maximumBy__110571494(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t3 = &Constructor_Data_Maybe_Just{1, v1_3}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
// TAST (Let): __local_var_4_0 -> uint32
__local_var_4_0 := uint32(gopurs_runtime.Apply2(cmp_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0, v1_3).IntVal)
_ = __local_var_4_0
var __t1 bool
{
if (__local_var_4_0 == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 == 380165415) {
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
__t2 = (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = v1_3
}
end_branch_2:
__t3 = &Constructor_Data_Maybe_Just{1, __t2}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}

func Call_Data_Foldable_minimumBy__110571494(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t3 = &Constructor_Data_Maybe_Just{1, v1_3}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
// TAST (Let): __local_var_4_0 -> uint32
__local_var_4_0 := uint32(gopurs_runtime.Apply2(cmp_1, (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0, v1_3).IntVal)
_ = __local_var_4_0
var __t1 bool
{
if (__local_var_4_0 == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = v1_3
}
end_branch_2:
__t3 = &Constructor_Data_Maybe_Just{1, __t2}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
}

func Call_Data_Foldable_oneOfMap__3719016818(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictPlus_1_loop *Constructor_Control_Plus_Plus) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *Constructor_Control_Plus_Plus = dictPlus_1_loop
_ = dictPlus_1
// TAST (Let): alt_2_0 -> gopurs_runtime.Value
alt_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictPlus_1.V0), gopurs_runtime.Value{}), "alt")
_ = alt_2_0
// TAST (Let): empty_3_1 -> gopurs_runtime.Value
empty_3_1 := gopurs_runtime.Box(dictPlus_1.V1)
_ = empty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_2_0, gopurs_runtime.Apply(f_4, x_5))
}), empty_3_1)
})
}

func Call_Data_Foldable_oneOfMap__1349369970(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictPlus_1_loop *Constructor_Control_Plus_Plus) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *Constructor_Control_Plus_Plus = dictPlus_1_loop
_ = dictPlus_1
// TAST (Let): alt_2_0 -> gopurs_runtime.Value
alt_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictPlus_1.V0), gopurs_runtime.Value{}), "alt")
_ = alt_2_0
// TAST (Let): empty_3_1 -> gopurs_runtime.Value
empty_3_1 := gopurs_runtime.Box(dictPlus_1.V1)
_ = empty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable_0.V2), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_2_0, gopurs_runtime.Apply(f_4, x_5))
}), empty_3_1)
})
}

func Call_Data_Foldable_surroundMap__3689038427(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup, d_2_loop gopurs_runtime.Value, t_3_loop gopurs_runtime.Value, f_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_1_loop
_ = dictSemigroup_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
var t_3 gopurs_runtime.Value = t_3_loop
_ = t_3
var f_4 gopurs_runtime.Value = f_4_loop
_ = f_4
// TAST (Let): semigroupEndo1_5_0 -> gopurs_runtime.Value
semigroupEndo1_5_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_5, gopurs_runtime.Apply(v1_6, x_7))
})
})
}))
_ = semigroupEndo1_5_0
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_5_0
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
}))))}, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), d_2, gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), gopurs_runtime.Apply(t_3, a_5), m_6))
})
}), f_4, d_2)
}

func Call_Data_Foldable_traverse___996968168(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 -> *Constructor_Data_Functor_Functor
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): applySecond_1_0 -> gopurs_runtime.Value
applySecond_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), a_3), b_4)
})
})
_ = applySecond_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(f_3, x_4))
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit()))
})
})
}

func Call_Data_Foldable_traverse___1507800296(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 -> *Constructor_Data_Functor_Functor
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): applySecond_1_0 -> gopurs_runtime.Value
applySecond_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), a_3), b_4)
})
})
_ = applySecond_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(f_3, x_4))
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit()))
})
})
}

func Call_Data_Foldable_traverse___1229293625(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__eta0_0, "foldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(__eta1_1, x_2)
_ = __local_var_3_0
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
_ = __local_var_5_2
__local_var_5_1 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
_ = __local_var_5_1
__local_var_6_3 := gopurs_runtime.Apply(b_4, gopurs_runtime.Value{})
_ = __local_var_6_3
return gopurs_runtime.Apply(__local_var_5_1, __local_var_6_3)
})
})
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
}

func Call_Data_Foldable_traverse___3630450585(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__eta0_0, "foldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(__eta1_1, x_2)
_ = __local_var_3_0
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
_ = __local_var_5_2
__local_var_5_1 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
_ = __local_var_5_1
__local_var_6_3 := gopurs_runtime.Apply(b_4, gopurs_runtime.Value{})
_ = __local_var_6_3
return gopurs_runtime.Apply(__local_var_5_1, __local_var_6_3)
})
})
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
}

func Call_Data_Foldable_traverse___3202958137(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__eta0_0, "foldr"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 -> gopurs_runtime.Value
__local_var_3_0 := gopurs_runtime.Apply(__eta1_1, x_2)
_ = __local_var_3_0
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(__local_var_3_0, gopurs_runtime.Value{})
_ = __local_var_5_2
__local_var_5_1 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
_ = __local_var_5_1
__local_var_6_3 := gopurs_runtime.Apply(b_4, gopurs_runtime.Value{})
_ = __local_var_6_3
return gopurs_runtime.Apply(__local_var_5_1, __local_var_6_3)
})
})
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
}

func Call_Data_Foldable_traverse___398118279(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 -> *Constructor_Data_Functor_Functor
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): applySecond_1_0 -> gopurs_runtime.Value
applySecond_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), a_3), b_4)
})
})
_ = applySecond_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_2, "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(f_3, x_4))
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit()))
})
})
}

func Get_Data_Foldable_foldlArray() gopurs_runtime.Value {
	return _Gopurs_Data_Foldable_FoldlArray
}

func Get_Data_Foldable_foldrArray() gopurs_runtime.Value {
	return _Gopurs_Data_Foldable_FoldrArray
}

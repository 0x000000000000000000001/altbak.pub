package Data_Foldable

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Maybe_First "gopurs/output/Data.Maybe.First"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

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

var cache_monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		cache_monoidEndo = func() gopurs_runtime.Value {
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_1
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "compose"), v_1, v1_2)
})
}))
_ = semigroupEndo1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))))}
}()
	})
	return cache_monoidEndo
}

var cache_identity1 gopurs_runtime.Value
var once_identity1 sync.Once
func Get_identity1() gopurs_runtime.Value {
	once_identity1.Do(func() {
		cache_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity1(x_0_box)
})
	})
	return cache_identity1
}

var cache_identity2 gopurs_runtime.Value
var once_identity2 sync.Once
func Get_identity2() gopurs_runtime.Value {
	once_identity2.Do(func() {
		cache_identity2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity2(x_0_box)
})
	})
	return cache_identity2
}

var cache_Empty gopurs_runtime.Value
var once_Empty sync.Once
func Get_Empty() gopurs_runtime.Value {
	once_Empty.Do(func() {
		cache_Empty = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Empty
}

var cache_Node gopurs_runtime.Value
var once_Node sync.Once
func Get_Node() gopurs_runtime.Value {
	once_Node.Do(func() {
		cache_Node = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2421944209, UnsafePtr: unsafe.Pointer(&Constructor_Node[gopurs_runtime.Value]{1, value0})}
})
	})
	return cache_Node
}

var cache_Append gopurs_runtime.Value
var once_Append sync.Once
func Get_Append() gopurs_runtime.Value {
	once_Append.Do(func() {
		cache_Append = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Append[gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_Append
}

var cache_semigroupFreeMonoidTree gopurs_runtime.Value
var once_semigroupFreeMonoidTree sync.Once
func Get_semigroupFreeMonoidTree() gopurs_runtime.Value {
	once_semigroupFreeMonoidTree.Do(func() {
		cache_semigroupFreeMonoidTree = gopurs_runtime.RecordDict1("append", Get_Append())
	})
	return cache_semigroupFreeMonoidTree
}

var cache_semigroupFreeMonoidTree__gopurs_runtime_Value_2398658907 gopurs_runtime.Value
var once_semigroupFreeMonoidTree__gopurs_runtime_Value_2398658907 sync.Once
func Get_semigroupFreeMonoidTree__gopurs_runtime_Value_2398658907() gopurs_runtime.Value {
	once_semigroupFreeMonoidTree__gopurs_runtime_Value_2398658907.Do(func() {
		cache_semigroupFreeMonoidTree__gopurs_runtime_Value_2398658907 = gopurs_runtime.RecordDict1("append", Get_Append())
	})
	return cache_semigroupFreeMonoidTree__gopurs_runtime_Value_2398658907
}

var cache_monoidFreeMonoidTree gopurs_runtime.Value
var once_monoidFreeMonoidTree sync.Once
func Get_monoidFreeMonoidTree() gopurs_runtime.Value {
	once_monoidFreeMonoidTree.Do(func() {
		cache_monoidFreeMonoidTree = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupFreeMonoidTree()
}), gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
	})
	return cache_monoidFreeMonoidTree
}

var cache_monoidFreeMonoidTree__ptrData_Monoid_Constructor_Monoid_gopurs_runtime_Value__2615096836 gopurs_runtime.Value
var once_monoidFreeMonoidTree__ptrData_Monoid_Constructor_Monoid_gopurs_runtime_Value__2615096836 sync.Once
func Get_monoidFreeMonoidTree__ptrData_Monoid_Constructor_Monoid_gopurs_runtime_Value__2615096836() gopurs_runtime.Value {
	once_monoidFreeMonoidTree__ptrData_Monoid_Constructor_Monoid_gopurs_runtime_Value__2615096836.Do(func() {
		cache_monoidFreeMonoidTree__ptrData_Monoid_Constructor_Monoid_gopurs_runtime_Value__2615096836 = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupFreeMonoidTree()
}), gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}})}
	})
	return cache_monoidFreeMonoidTree__ptrData_Monoid_Constructor_Monoid_gopurs_runtime_Value__2615096836
}

var cache_foldr gopurs_runtime.Value
var once_foldr sync.Once
func Get_foldr() gopurs_runtime.Value {
	once_foldr.Do(func() {
		cache_foldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr
}

var cache_foldr__gopurs_runtime_Value_2151204251 gopurs_runtime.Value
var once_foldr__gopurs_runtime_Value_2151204251 sync.Once
func Get_foldr__gopurs_runtime_Value_2151204251() gopurs_runtime.Value {
	once_foldr__gopurs_runtime_Value_2151204251.Do(func() {
		cache_foldr__gopurs_runtime_Value_2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__gopurs_runtime_Value_2151204251(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__gopurs_runtime_Value_2151204251
}

var cache_indexr gopurs_runtime.Value
var once_indexr sync.Once
func Get_indexr() gopurs_runtime.Value {
	once_indexr.Do(func() {
		cache_indexr = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, idx_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_indexr(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), idx_1_box.IntVal)
})
	})
	return cache_indexr
}

var cache_null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		cache_null = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_null(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_null
}

var cache_oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		cache_oneOf = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOf(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]](dictPlus_1_box))
})
	})
	return cache_oneOf
}

var cache_oneOfMap gopurs_runtime.Value
var once_oneOfMap sync.Once
func Get_oneOfMap() gopurs_runtime.Value {
	once_oneOfMap.Do(func() {
		cache_oneOfMap = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOfMap(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]](dictPlus_1_box))
})
	})
	return cache_oneOfMap
}

var cache_oneOfMap__gopurs_runtime_Value_3719016818 gopurs_runtime.Value
var once_oneOfMap__gopurs_runtime_Value_3719016818 sync.Once
func Get_oneOfMap__gopurs_runtime_Value_3719016818() gopurs_runtime.Value {
	once_oneOfMap__gopurs_runtime_Value_3719016818.Do(func() {
		cache_oneOfMap__gopurs_runtime_Value_3719016818 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictPlus_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOfMap__gopurs_runtime_Value_3719016818(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]](dictPlus_1_box))
})
	})
	return cache_oneOfMap__gopurs_runtime_Value_3719016818
}

var cache_traverse_ gopurs_runtime.Value
var once_traverse_ sync.Once
func Get_traverse_() gopurs_runtime.Value {
	once_traverse_.Do(func() {
		cache_traverse_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse_(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_traverse_
}

var cache_traverse___gopurs_runtime_Value_996968168 gopurs_runtime.Value
var once_traverse___gopurs_runtime_Value_996968168 sync.Once
func Get_traverse___gopurs_runtime_Value_996968168() gopurs_runtime.Value {
	once_traverse___gopurs_runtime_Value_996968168.Do(func() {
		cache_traverse___gopurs_runtime_Value_996968168 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse___gopurs_runtime_Value_996968168(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_traverse___gopurs_runtime_Value_996968168
}

var cache_for_ gopurs_runtime.Value
var once_for_ sync.Once
func Get_for_() gopurs_runtime.Value {
	once_for_.Do(func() {
		cache_for_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_for_(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_for_
}

var cache_sequence_ gopurs_runtime.Value
var once_sequence_ sync.Once
func Get_sequence_() gopurs_runtime.Value {
	once_sequence_.Do(func() {
		cache_sequence_ = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictFoldable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence_(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_1_box))
})
	})
	return cache_sequence_
}

var cache_foldl gopurs_runtime.Value
var once_foldl sync.Once
func Get_foldl() gopurs_runtime.Value {
	once_foldl.Do(func() {
		cache_foldl = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl
}

var cache_foldl__gopurs_runtime_Value_2151204251 gopurs_runtime.Value
var once_foldl__gopurs_runtime_Value_2151204251 sync.Once
func Get_foldl__gopurs_runtime_Value_2151204251() gopurs_runtime.Value {
	once_foldl__gopurs_runtime_Value_2151204251.Do(func() {
		cache_foldl__gopurs_runtime_Value_2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__gopurs_runtime_Value_2151204251(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__gopurs_runtime_Value_2151204251
}

var cache_indexl gopurs_runtime.Value
var once_indexl sync.Once
func Get_indexl() gopurs_runtime.Value {
	once_indexl.Do(func() {
		cache_indexl = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, idx_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_indexl(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), idx_1_box.IntVal)
})
	})
	return cache_indexl
}

var cache_intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		cache_intercalate = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_intercalate
}

var cache_intercalate__gopurs_runtime_Value_3939234276 gopurs_runtime.Value
var once_intercalate__gopurs_runtime_Value_3939234276 sync.Once
func Get_intercalate__gopurs_runtime_Value_3939234276() gopurs_runtime.Value {
	once_intercalate__gopurs_runtime_Value_3939234276.Do(func() {
		cache_intercalate__gopurs_runtime_Value_3939234276 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate__gopurs_runtime_Value_3939234276(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_intercalate__gopurs_runtime_Value_3939234276
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dictSemiring_1_box))
})
	})
	return cache_length
}

var cache_length__gopurs_runtime_Value_949294460 gopurs_runtime.Value
var once_length__gopurs_runtime_Value_949294460 sync.Once
func Get_length__gopurs_runtime_Value_949294460() gopurs_runtime.Value {
	once_length__gopurs_runtime_Value_949294460.Do(func() {
		cache_length__gopurs_runtime_Value_949294460 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length__gopurs_runtime_Value_949294460(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dictSemiring_1_box))
})
	})
	return cache_length__gopurs_runtime_Value_949294460
}

var cache_maximumBy gopurs_runtime.Value
var once_maximumBy sync.Once
func Get_maximumBy() gopurs_runtime.Value {
	once_maximumBy.Do(func() {
		cache_maximumBy = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximumBy(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), cmp_1_box)
})
	})
	return cache_maximumBy
}

var cache_maximumBy__gopurs_runtime_Value_110571494 gopurs_runtime.Value
var once_maximumBy__gopurs_runtime_Value_110571494 sync.Once
func Get_maximumBy__gopurs_runtime_Value_110571494() gopurs_runtime.Value {
	once_maximumBy__gopurs_runtime_Value_110571494.Do(func() {
		cache_maximumBy__gopurs_runtime_Value_110571494 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximumBy__gopurs_runtime_Value_110571494(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), cmp_1_box)
})
	})
	return cache_maximumBy__gopurs_runtime_Value_110571494
}

var cache_maximum gopurs_runtime.Value
var once_maximum sync.Once
func Get_maximum() gopurs_runtime.Value {
	once_maximum.Do(func() {
		cache_maximum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximum(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_maximum
}

var cache_minimumBy gopurs_runtime.Value
var once_minimumBy sync.Once
func Get_minimumBy() gopurs_runtime.Value {
	once_minimumBy.Do(func() {
		cache_minimumBy = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimumBy(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), cmp_1_box)
})
	})
	return cache_minimumBy
}

var cache_minimumBy__gopurs_runtime_Value_110571494 gopurs_runtime.Value
var once_minimumBy__gopurs_runtime_Value_110571494 sync.Once
func Get_minimumBy__gopurs_runtime_Value_110571494() gopurs_runtime.Value {
	once_minimumBy__gopurs_runtime_Value_110571494.Do(func() {
		cache_minimumBy__gopurs_runtime_Value_110571494 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimumBy__gopurs_runtime_Value_110571494(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), cmp_1_box)
})
	})
	return cache_minimumBy__gopurs_runtime_Value_110571494
}

var cache_minimum gopurs_runtime.Value
var once_minimum sync.Once
func Get_minimum() gopurs_runtime.Value {
	once_minimum.Do(func() {
		cache_minimum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimum(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_minimum
}

var cache_product gopurs_runtime.Value
var once_product sync.Once
func Get_product() gopurs_runtime.Value {
	once_product.Do(func() {
		cache_product = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_product(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dictSemiring_1_box))
})
	})
	return cache_product
}

var cache_sum gopurs_runtime.Value
var once_sum sync.Once
func Get_sum() gopurs_runtime.Value {
	once_sum.Do(func() {
		cache_sum = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sum(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dictSemiring_1_box))
})
	})
	return cache_sum
}

var cache_foldableTuple gopurs_runtime.Value
var once_foldableTuple sync.Once
func Get_foldableTuple() gopurs_runtime.Value {
	once_foldableTuple.Do(func() {
		cache_foldableTuple = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
})
})
}))
	})
	return cache_foldableTuple
}

var cache_foldableTuple__gopurs_runtime_Value_1455669080 gopurs_runtime.Value
var once_foldableTuple__gopurs_runtime_Value_1455669080 sync.Once
func Get_foldableTuple__gopurs_runtime_Value_1455669080() gopurs_runtime.Value {
	once_foldableTuple__gopurs_runtime_Value_1455669080.Do(func() {
		cache_foldableTuple__gopurs_runtime_Value_1455669080 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
})
})
}))
	})
	return cache_foldableTuple__gopurs_runtime_Value_1455669080
}

var cache_foldableMultiplicative gopurs_runtime.Value
var once_foldableMultiplicative sync.Once
func Get_foldableMultiplicative() gopurs_runtime.Value {
	once_foldableMultiplicative.Do(func() {
		cache_foldableMultiplicative = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableMultiplicative
}

var cache_foldableMultiplicative__gopurs_runtime_Value_1841171440 gopurs_runtime.Value
var once_foldableMultiplicative__gopurs_runtime_Value_1841171440 sync.Once
func Get_foldableMultiplicative__gopurs_runtime_Value_1841171440() gopurs_runtime.Value {
	once_foldableMultiplicative__gopurs_runtime_Value_1841171440.Do(func() {
		cache_foldableMultiplicative__gopurs_runtime_Value_1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableMultiplicative__gopurs_runtime_Value_1841171440
}

var cache_foldableMaybe gopurs_runtime.Value
var once_foldableMaybe sync.Once
func Get_foldableMaybe() gopurs_runtime.Value {
	once_foldableMaybe.Do(func() {
		cache_foldableMaybe = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply(v_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
__t3 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
	return cache_foldableMaybe
}

var cache_foldableMaybe__ptrConstructor_Foldable_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___3653484922 gopurs_runtime.Value
var once_foldableMaybe__ptrConstructor_Foldable_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___3653484922 sync.Once
func Get_foldableMaybe__ptrConstructor_Foldable_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___3653484922() gopurs_runtime.Value {
	once_foldableMaybe__ptrConstructor_Foldable_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___3653484922.Do(func() {
		cache_foldableMaybe__ptrConstructor_Foldable_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___3653484922 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Foldable[*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply(v_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
__t3 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
})})}
	})
	return cache_foldableMaybe__ptrConstructor_Foldable_ptrData_Maybe_Constructor_Just_gopurs_runtime_Value___3653484922
}

var cache_foldableMaybe__gopurs_runtime_Value_2831137713 gopurs_runtime.Value
var once_foldableMaybe__gopurs_runtime_Value_2831137713 sync.Once
func Get_foldableMaybe__gopurs_runtime_Value_2831137713() gopurs_runtime.Value {
	once_foldableMaybe__gopurs_runtime_Value_2831137713.Do(func() {
		cache_foldableMaybe__gopurs_runtime_Value_2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply(v_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
__t3 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
	return cache_foldableMaybe__gopurs_runtime_Value_2831137713
}

var cache_foldableIdentity gopurs_runtime.Value
var once_foldableIdentity sync.Once
func Get_foldableIdentity() gopurs_runtime.Value {
	once_foldableIdentity.Do(func() {
		cache_foldableIdentity = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableIdentity
}

var cache_foldableIdentity__gopurs_runtime_Value_1841171440 gopurs_runtime.Value
var once_foldableIdentity__gopurs_runtime_Value_1841171440 sync.Once
func Get_foldableIdentity__gopurs_runtime_Value_1841171440() gopurs_runtime.Value {
	once_foldableIdentity__gopurs_runtime_Value_1841171440.Do(func() {
		cache_foldableIdentity__gopurs_runtime_Value_1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableIdentity__gopurs_runtime_Value_1841171440
}

var cache_foldableEither gopurs_runtime.Value
var once_foldableEither sync.Once
func Get_foldableEither() gopurs_runtime.Value {
	once_foldableEither.Do(func() {
		cache_foldableEither = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply(v_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
__t3 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
	return cache_foldableEither
}

var cache_foldableEither__gopurs_runtime_Value_1622911640 gopurs_runtime.Value
var once_foldableEither__gopurs_runtime_Value_1622911640 sync.Once
func Get_foldableEither__gopurs_runtime_Value_1622911640() gopurs_runtime.Value {
	once_foldableEither__gopurs_runtime_Value_1622911640.Do(func() {
		cache_foldableEither__gopurs_runtime_Value_1622911640 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply(v_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
__t3 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
	return cache_foldableEither__gopurs_runtime_Value_1622911640
}

var cache_foldableDual gopurs_runtime.Value
var once_foldableDual sync.Once
func Get_foldableDual() gopurs_runtime.Value {
	once_foldableDual.Do(func() {
		cache_foldableDual = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableDual
}

var cache_foldableDual__gopurs_runtime_Value_1841171440 gopurs_runtime.Value
var once_foldableDual__gopurs_runtime_Value_1841171440 sync.Once
func Get_foldableDual__gopurs_runtime_Value_1841171440() gopurs_runtime.Value {
	once_foldableDual__gopurs_runtime_Value_1841171440.Do(func() {
		cache_foldableDual__gopurs_runtime_Value_1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableDual__gopurs_runtime_Value_1841171440
}

var cache_foldableDisj gopurs_runtime.Value
var once_foldableDisj sync.Once
func Get_foldableDisj() gopurs_runtime.Value {
	once_foldableDisj.Do(func() {
		cache_foldableDisj = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableDisj
}

var cache_foldableDisj__gopurs_runtime_Value_1841171440 gopurs_runtime.Value
var once_foldableDisj__gopurs_runtime_Value_1841171440 sync.Once
func Get_foldableDisj__gopurs_runtime_Value_1841171440() gopurs_runtime.Value {
	once_foldableDisj__gopurs_runtime_Value_1841171440.Do(func() {
		cache_foldableDisj__gopurs_runtime_Value_1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableDisj__gopurs_runtime_Value_1841171440
}

var cache_foldableConst gopurs_runtime.Value
var once_foldableConst sync.Once
func Get_foldableConst() gopurs_runtime.Value {
	once_foldableConst.Do(func() {
		cache_foldableConst = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableConst
}

var cache_foldableConst__gopurs_runtime_Value_943899702 gopurs_runtime.Value
var once_foldableConst__gopurs_runtime_Value_943899702 sync.Once
func Get_foldableConst__gopurs_runtime_Value_943899702() gopurs_runtime.Value {
	once_foldableConst__gopurs_runtime_Value_943899702.Do(func() {
		cache_foldableConst__gopurs_runtime_Value_943899702 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableConst__gopurs_runtime_Value_943899702
}

var cache_foldableConj gopurs_runtime.Value
var once_foldableConj sync.Once
func Get_foldableConj() gopurs_runtime.Value {
	once_foldableConj.Do(func() {
		cache_foldableConj = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableConj
}

var cache_foldableConj__gopurs_runtime_Value_1841171440 gopurs_runtime.Value
var once_foldableConj__gopurs_runtime_Value_1841171440 sync.Once
func Get_foldableConj__gopurs_runtime_Value_1841171440() gopurs_runtime.Value {
	once_foldableConj__gopurs_runtime_Value_1841171440.Do(func() {
		cache_foldableConj__gopurs_runtime_Value_1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableConj__gopurs_runtime_Value_1841171440
}

var cache_foldableAdditive gopurs_runtime.Value
var once_foldableAdditive sync.Once
func Get_foldableAdditive() gopurs_runtime.Value {
	once_foldableAdditive.Do(func() {
		cache_foldableAdditive = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableAdditive
}

var cache_foldableAdditive__gopurs_runtime_Value_1841171440 gopurs_runtime.Value
var once_foldableAdditive__gopurs_runtime_Value_1841171440 sync.Once
func Get_foldableAdditive__gopurs_runtime_Value_1841171440() gopurs_runtime.Value {
	once_foldableAdditive__gopurs_runtime_Value_1841171440.Do(func() {
		cache_foldableAdditive__gopurs_runtime_Value_1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableAdditive__gopurs_runtime_Value_1841171440
}

var cache_foldMapDefaultR gopurs_runtime.Value
var once_foldMapDefaultR sync.Once
func Get_foldMapDefaultR() gopurs_runtime.Value {
	once_foldMapDefaultR.Do(func() {
		cache_foldMapDefaultR = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_foldMapDefaultR
}

var cache_foldableArray_ADT_Data_Foldable_Foldable_ADT_Prim_Array gopurs_runtime.Value
var once_foldableArray_ADT_Data_Foldable_Foldable_ADT_Prim_Array sync.Once
func Get_foldableArray_ADT_Data_Foldable_Foldable_ADT_Prim_Array() gopurs_runtime.Value {
	once_foldableArray_ADT_Data_Foldable_Foldable_ADT_Prim_Array.Do(func() {
		cache_foldableArray_ADT_Data_Foldable_Foldable_ADT_Prim_Array = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), Get_foldlArray(), Get_foldrArray())
	})
	return cache_foldableArray_ADT_Data_Foldable_Foldable_ADT_Prim_Array
}

var cache_foldableArray_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any gopurs_runtime.Value
var once_foldableArray_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any sync.Once
func Get_foldableArray_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any() gopurs_runtime.Value {
	once_foldableArray_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any.Do(func() {
		cache_foldableArray_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), Get_foldlArray(), Get_foldrArray())
	})
	return cache_foldableArray_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any
}

var cache_foldableArray gopurs_runtime.Value
var once_foldableArray sync.Once
func Get_foldableArray() gopurs_runtime.Value {
	once_foldableArray.Do(func() {
		cache_foldableArray = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), Get_foldlArray(), Get_foldrArray())
	})
	return cache_foldableArray
}

var cache_foldableArray__ptrConstructor_Foldable_gopurs_runtime_Value__2950015754 gopurs_runtime.Value
var once_foldableArray__ptrConstructor_Foldable_gopurs_runtime_Value__2950015754 sync.Once
func Get_foldableArray__ptrConstructor_Foldable_gopurs_runtime_Value__2950015754() gopurs_runtime.Value {
	once_foldableArray__ptrConstructor_Foldable_gopurs_runtime_Value__2950015754.Do(func() {
		cache_foldableArray__ptrConstructor_Foldable_gopurs_runtime_Value__2950015754 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), Get_foldlArray(), Get_foldrArray()})}
	})
	return cache_foldableArray__ptrConstructor_Foldable_gopurs_runtime_Value__2950015754
}

var cache_foldableArray__gopurs_runtime_Value_3859409398 gopurs_runtime.Value
var once_foldableArray__gopurs_runtime_Value_3859409398 sync.Once
func Get_foldableArray__gopurs_runtime_Value_3859409398() gopurs_runtime.Value {
	once_foldableArray__gopurs_runtime_Value_3859409398.Do(func() {
		cache_foldableArray__gopurs_runtime_Value_3859409398 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), Get_foldlArray(), Get_foldrArray())
	})
	return cache_foldableArray__gopurs_runtime_Value_3859409398
}

var cache_foldableFreeMonoidTree_ADT_Data_Foldable_Foldable_ADT_Data_Foldable_FreeMonoidTree gopurs_runtime.Value
var once_foldableFreeMonoidTree_ADT_Data_Foldable_Foldable_ADT_Data_Foldable_FreeMonoidTree sync.Once
func Get_foldableFreeMonoidTree_ADT_Data_Foldable_Foldable_ADT_Data_Foldable_FreeMonoidTree() gopurs_runtime.Value {
	once_foldableFreeMonoidTree_ADT_Data_Foldable_Foldable_ADT_Data_Foldable_FreeMonoidTree.Do(func() {
		cache_foldableFreeMonoidTree_ADT_Data_Foldable_Foldable_ADT_Data_Foldable_FreeMonoidTree = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
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
var __t3 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, acc_2, (*Constructor_Node[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0)
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2812549951) {
var __t5 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = rhs_4
continue go__go_1_2_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
var __t4 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
continue go__go_1_2_0
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Append[gopurs_runtime.Value]{1, (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1, rhs_4})}
continue go__go_1_2_0
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
var __t7 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
__t7 = acc_2
goto end_branch_7
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_0
__t7 = gopurs_runtime.Value{}
}
end_branch_7:
__t3 = __t7
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
var __t9 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, (*Constructor_Node[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0, acc_2)
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_1
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2812549951) {
var __t11 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_1
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
var __t10 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_1
__t10 = gopurs_runtime.Value{}
goto end_branch_10
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Append[gopurs_runtime.Value]{1, lhs_3, (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0})}
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_1
__t10 = gopurs_runtime.Value{}
}
end_branch_10:
__t11 = __t10
}
end_branch_11:
__t9 = __t11
goto end_branch_9
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
var __t13 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
__t13 = acc_2
goto end_branch_13
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_1
__t13 = gopurs_runtime.Value{}
}
end_branch_13:
__t9 = __t13
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
	return cache_foldableFreeMonoidTree_ADT_Data_Foldable_Foldable_ADT_Data_Foldable_FreeMonoidTree
}

var cache_foldableFreeMonoidTree gopurs_runtime.Value
var once_foldableFreeMonoidTree sync.Once
func Get_foldableFreeMonoidTree() gopurs_runtime.Value {
	once_foldableFreeMonoidTree.Do(func() {
		cache_foldableFreeMonoidTree = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
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
var __t3 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, acc_2, (*Constructor_Node[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0)
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_2
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2812549951) {
var __t5 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = rhs_4
continue go__go_1_2_2
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
var __t4 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
continue go__go_1_2_2
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Append[gopurs_runtime.Value]{1, (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1, rhs_4})}
continue go__go_1_2_2
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
var __t7 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
__t7 = acc_2
goto end_branch_7
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_2
__t7 = gopurs_runtime.Value{}
}
end_branch_7:
__t3 = __t7
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
var __t9 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, (*Constructor_Node[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0, acc_2)
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_3
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2812549951) {
var __t11 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_3
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
var __t10 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_3
__t10 = gopurs_runtime.Value{}
goto end_branch_10
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Append[gopurs_runtime.Value]{1, lhs_3, (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0})}
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_3
__t10 = gopurs_runtime.Value{}
}
end_branch_10:
__t11 = __t10
}
end_branch_11:
__t9 = __t11
goto end_branch_9
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
var __t13 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
__t13 = acc_2
goto end_branch_13
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_3
__t13 = gopurs_runtime.Value{}
}
end_branch_13:
__t9 = __t13
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
	return cache_foldableFreeMonoidTree
}

var cache_foldableFreeMonoidTree__ptrConstructor_Foldable_gopurs_runtime_Value__2832280077 gopurs_runtime.Value
var once_foldableFreeMonoidTree__ptrConstructor_Foldable_gopurs_runtime_Value__2832280077 sync.Once
func Get_foldableFreeMonoidTree__ptrConstructor_Foldable_gopurs_runtime_Value__2832280077() gopurs_runtime.Value {
	once_foldableFreeMonoidTree__ptrConstructor_Foldable_gopurs_runtime_Value__2832280077.Do(func() {
		cache_foldableFreeMonoidTree__ptrConstructor_Foldable_gopurs_runtime_Value__2832280077 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(&Constructor_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_2_4 gopurs_runtime.Value
go__go_1_2_4 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_2_4:
for {
if false { continue go__go_1_2_4 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t3 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, acc_2, (*Constructor_Node[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0)
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_4
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2812549951) {
var __t5 gopurs_runtime.Value
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = rhs_4
continue go__go_1_2_4
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
var __t4 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1
continue go__go_1_2_4
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V0
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Append[gopurs_runtime.Value]{1, (*Constructor_Append[gopurs_runtime.Value])(lhs_3.UnsafePtr).V1, rhs_4})}
continue go__go_1_2_4
__t4 = gopurs_runtime.Value{}
}
end_branch_4:
__t5 = __t4
}
end_branch_5:
__t3 = __t5
goto end_branch_3
} else {

}
}
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
var __t7 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
__t7 = acc_2
goto end_branch_7
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = rhs_4
rhs_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
continue go__go_1_2_4
__t7 = gopurs_runtime.Value{}
}
end_branch_7:
__t3 = __t7
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
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_2_4, a_2, b_3, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)})
})
})
}), gopurs_runtime.Func(func(fn_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_8_5 gopurs_runtime.Value
go__go_1_8_5 = gopurs_runtime.Func(func(acc_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(lhs_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rhs_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var acc_2_loop gopurs_runtime.Value = acc_2_loop_val
var lhs_3_loop gopurs_runtime.Value = lhs_3_loop_val
var rhs_4_loop gopurs_runtime.Value = rhs_4_loop_val
go__go_1_8_5:
for {
if false { continue go__go_1_8_5 }
var acc_2 gopurs_runtime.Value = acc_2_loop
_ = acc_2
var lhs_3 gopurs_runtime.Value = lhs_3_loop
_ = lhs_3
var rhs_4 gopurs_runtime.Value = rhs_4_loop
_ = rhs_4
var __t9 gopurs_runtime.Value
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2421944209) {
acc_2_loop = gopurs_runtime.Apply2(fn_0, (*Constructor_Node[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0, acc_2)
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_5
__t9 = gopurs_runtime.Value{}
goto end_branch_9
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2812549951) {
var __t11 gopurs_runtime.Value
{
var __t_tag_12 gopurs_runtime.Value = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = lhs_3
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_5
__t11 = gopurs_runtime.Value{}
goto end_branch_11
} else {

}
}
{
var __t10 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
acc_2_loop = acc_2
lhs_3_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_5
__t10 = gopurs_runtime.Value{}
goto end_branch_10
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2812549951, UnsafePtr: unsafe.Pointer(&Constructor_Append[gopurs_runtime.Value]{1, lhs_3, (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V0})}
rhs_4_loop = (*Constructor_Append[gopurs_runtime.Value])(rhs_4.UnsafePtr).V1
continue go__go_1_8_5
__t10 = gopurs_runtime.Value{}
}
end_branch_10:
__t11 = __t10
}
end_branch_11:
__t9 = __t11
goto end_branch_9
} else {

}
}
{
if (rhs_4.Type == 9 && rhs_4.IntVal == 2065045956) {
var __t13 gopurs_runtime.Value
{
if (lhs_3.Type == 9 && lhs_3.IntVal == 2065045956) {
__t13 = acc_2
goto end_branch_13
} else {

}
}
{
acc_2_loop = acc_2
lhs_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}
rhs_4_loop = lhs_3
continue go__go_1_8_5
__t13 = gopurs_runtime.Value{}
}
end_branch_13:
__t9 = __t13
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
}()
})
})
})
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(go__go_1_8_5, a_2, gopurs_runtime.Value{Type: 9, IntVal: 2065045956, UnsafePtr: unsafe.Pointer(nil)}, b_3)
})
})
})})}
	})
	return cache_foldableFreeMonoidTree__ptrConstructor_Foldable_gopurs_runtime_Value__2832280077
}

var cache_foldMapDefaultL gopurs_runtime.Value
var once_foldMapDefaultL sync.Once
func Get_foldMapDefaultL() gopurs_runtime.Value {
	once_foldMapDefaultL.Do(func() {
		cache_foldMapDefaultL = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapDefaultL(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_foldMapDefaultL
}

var cache_foldMap gopurs_runtime.Value
var once_foldMap sync.Once
func Get_foldMap() gopurs_runtime.Value {
	once_foldMap.Do(func() {
		cache_foldMap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap
}

var cache_foldMap__gopurs_runtime_Value_4098395794 gopurs_runtime.Value
var once_foldMap__gopurs_runtime_Value_4098395794 sync.Once
func Get_foldMap__gopurs_runtime_Value_4098395794() gopurs_runtime.Value {
	once_foldMap__gopurs_runtime_Value_4098395794.Do(func() {
		cache_foldMap__gopurs_runtime_Value_4098395794 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__gopurs_runtime_Value_4098395794(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap__gopurs_runtime_Value_4098395794
}

var cache_foldableApp gopurs_runtime.Value
var once_foldableApp sync.Once
func Get_foldableApp() gopurs_runtime.Value {
	once_foldableApp.Do(func() {
		cache_foldableApp = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableApp(dictFoldable_0_box)
})
	})
	return cache_foldableApp
}

var cache_foldableCompose gopurs_runtime.Value
var once_foldableCompose sync.Once
func Get_foldableCompose() gopurs_runtime.Value {
	once_foldableCompose.Do(func() {
		cache_foldableCompose = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableCompose(dictFoldable_0_box, dictFoldable1_1_box)
})
	})
	return cache_foldableCompose
}

var cache_foldableCoproduct gopurs_runtime.Value
var once_foldableCoproduct sync.Once
func Get_foldableCoproduct() gopurs_runtime.Value {
	once_foldableCoproduct.Do(func() {
		cache_foldableCoproduct = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableCoproduct(dictFoldable_0_box, dictFoldable1_1_box)
})
	})
	return cache_foldableCoproduct
}

var cache_foldableFirst gopurs_runtime.Value
var once_foldableFirst sync.Once
func Get_foldableFirst() gopurs_runtime.Value {
	once_foldableFirst.Do(func() {
		cache_foldableFirst = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldMap"), dictMonoid_0, f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldl"), f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldr"), f_0, z_1, v_2)
})
})
}))
	})
	return cache_foldableFirst
}

var cache_foldableFirst__gopurs_runtime_Value_2831137713 gopurs_runtime.Value
var once_foldableFirst__gopurs_runtime_Value_2831137713 sync.Once
func Get_foldableFirst__gopurs_runtime_Value_2831137713() gopurs_runtime.Value {
	once_foldableFirst__gopurs_runtime_Value_2831137713.Do(func() {
		cache_foldableFirst__gopurs_runtime_Value_2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldMap"), dictMonoid_0, f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldl"), f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldr"), f_0, z_1, v_2)
})
})
}))
	})
	return cache_foldableFirst__gopurs_runtime_Value_2831137713
}

var cache_foldableLast gopurs_runtime.Value
var once_foldableLast sync.Once
func Get_foldableLast() gopurs_runtime.Value {
	once_foldableLast.Do(func() {
		cache_foldableLast = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldMap"), dictMonoid_0, f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldl"), f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldr"), f_0, z_1, v_2)
})
})
}))
	})
	return cache_foldableLast
}

var cache_foldableLast__gopurs_runtime_Value_2831137713 gopurs_runtime.Value
var once_foldableLast__gopurs_runtime_Value_2831137713 sync.Once
func Get_foldableLast__gopurs_runtime_Value_2831137713() gopurs_runtime.Value {
	once_foldableLast__gopurs_runtime_Value_2831137713.Do(func() {
		cache_foldableLast__gopurs_runtime_Value_2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldMap"), dictMonoid_0, f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldl"), f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableMaybe(), "foldr"), f_0, z_1, v_2)
})
})
}))
	})
	return cache_foldableLast__gopurs_runtime_Value_2831137713
}

var cache_foldableProduct gopurs_runtime.Value
var once_foldableProduct sync.Once
func Get_foldableProduct() gopurs_runtime.Value {
	once_foldableProduct.Do(func() {
		cache_foldableProduct = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictFoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableProduct(dictFoldable_0_box, dictFoldable1_1_box)
})
	})
	return cache_foldableProduct
}

var cache_foldlDefault gopurs_runtime.Value
var once_foldlDefault sync.Once
func Get_foldlDefault() gopurs_runtime.Value {
	once_foldlDefault.Do(func() {
		cache_foldlDefault = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlDefault(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_foldlDefault
}

var cache_foldrDefault gopurs_runtime.Value
var once_foldrDefault sync.Once
func Get_foldrDefault() gopurs_runtime.Value {
	once_foldrDefault.Do(func() {
		cache_foldrDefault = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_foldrDefault
}

var cache_foldrDefault__gopurs_runtime_Value_2151204251 gopurs_runtime.Value
var once_foldrDefault__gopurs_runtime_Value_2151204251 sync.Once
func Get_foldrDefault__gopurs_runtime_Value_2151204251() gopurs_runtime.Value {
	once_foldrDefault__gopurs_runtime_Value_2151204251.Do(func() {
		cache_foldrDefault__gopurs_runtime_Value_2151204251 = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrDefault__gopurs_runtime_Value_2151204251(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_foldrDefault__gopurs_runtime_Value_2151204251
}

var cache_lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		cache_lookup = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookup(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_1_box), a_2_box)
})
	})
	return cache_lookup
}

var cache_surroundMap gopurs_runtime.Value
var once_surroundMap sync.Once
func Get_surroundMap() gopurs_runtime.Value {
	once_surroundMap.Do(func() {
		cache_surroundMap = gopurs_runtime.Func5(func(dictFoldable_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value, t_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_surroundMap(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box), d_2_box, t_3_box, f_4_box)
})
	})
	return cache_surroundMap
}

var cache_surroundMap__gopurs_runtime_Value_3689038427 gopurs_runtime.Value
var once_surroundMap__gopurs_runtime_Value_3689038427 sync.Once
func Get_surroundMap__gopurs_runtime_Value_3689038427() gopurs_runtime.Value {
	once_surroundMap__gopurs_runtime_Value_3689038427.Do(func() {
		cache_surroundMap__gopurs_runtime_Value_3689038427 = gopurs_runtime.Func5(func(dictFoldable_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value, t_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_surroundMap__gopurs_runtime_Value_3689038427(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box), d_2_box, t_3_box, f_4_box)
})
	})
	return cache_surroundMap__gopurs_runtime_Value_3689038427
}

var cache_surround gopurs_runtime.Value
var once_surround sync.Once
func Get_surround() gopurs_runtime.Value {
	once_surround.Do(func() {
		cache_surround = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_surround(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box), d_2_box)
})
	})
	return cache_surround
}

var cache_foldM gopurs_runtime.Value
var once_foldM sync.Once
func Get_foldM() gopurs_runtime.Value {
	once_foldM.Do(func() {
		cache_foldM = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldM(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]](dictMonad_1_box))
})
	})
	return cache_foldM
}

var cache_fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		cache_fold = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_fold
}

var cache_fold__gopurs_runtime_Value_910331789 gopurs_runtime.Value
var once_fold__gopurs_runtime_Value_910331789 sync.Once
func Get_fold__gopurs_runtime_Value_910331789() gopurs_runtime.Value {
	once_fold__gopurs_runtime_Value_910331789.Do(func() {
		cache_fold__gopurs_runtime_Value_910331789 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold__gopurs_runtime_Value_910331789(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_fold__gopurs_runtime_Value_910331789
}

var cache_findMap gopurs_runtime.Value
var once_findMap sync.Once
func Get_findMap() gopurs_runtime.Value {
	once_findMap.Do(func() {
		cache_findMap = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findMap(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), p_1_box)
})
	})
	return cache_findMap
}

var cache_find gopurs_runtime.Value
var once_find sync.Once
func Get_find() gopurs_runtime.Value {
	once_find.Do(func() {
		cache_find = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_find(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), p_1_box)
})
	})
	return cache_find
}

var cache_any gopurs_runtime.Value
var once_any sync.Once
func Get_any() gopurs_runtime.Value {
	once_any.Do(func() {
		cache_any = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_any(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_any
}

var cache_any__gopurs_runtime_Value_4179648253 gopurs_runtime.Value
var once_any__gopurs_runtime_Value_4179648253 sync.Once
func Get_any__gopurs_runtime_Value_4179648253() gopurs_runtime.Value {
	once_any__gopurs_runtime_Value_4179648253.Do(func() {
		cache_any__gopurs_runtime_Value_4179648253 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_any__gopurs_runtime_Value_4179648253(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_any__gopurs_runtime_Value_4179648253
}

var cache_elem gopurs_runtime.Value
var once_elem sync.Once
func Get_elem() gopurs_runtime.Value {
	once_elem.Do(func() {
		cache_elem = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elem(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_elem
}

var cache_elem__gopurs_runtime_Value_2343844090 gopurs_runtime.Value
var once_elem__gopurs_runtime_Value_2343844090 sync.Once
func Get_elem__gopurs_runtime_Value_2343844090() gopurs_runtime.Value {
	once_elem__gopurs_runtime_Value_2343844090.Do(func() {
		cache_elem__gopurs_runtime_Value_2343844090 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elem__gopurs_runtime_Value_2343844090(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_elem__gopurs_runtime_Value_2343844090
}

var cache_notElem gopurs_runtime.Value
var once_notElem sync.Once
func Get_notElem() gopurs_runtime.Value {
	once_notElem.Do(func() {
		cache_notElem = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_notElem(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_1_box), x_2_box)
})
	})
	return cache_notElem
}

var cache_or gopurs_runtime.Value
var once_or sync.Once
func Get_or() gopurs_runtime.Value {
	once_or.Do(func() {
		cache_or = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_or(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_or
}

var cache_all gopurs_runtime.Value
var once_all sync.Once
func Get_all() gopurs_runtime.Value {
	once_all.Do(func() {
		cache_all = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_all(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_all
}

var cache_all__gopurs_runtime_Value_4179648253 gopurs_runtime.Value
var once_all__gopurs_runtime_Value_4179648253 sync.Once
func Get_all__gopurs_runtime_Value_4179648253() gopurs_runtime.Value {
	once_all__gopurs_runtime_Value_4179648253.Do(func() {
		cache_all__gopurs_runtime_Value_4179648253 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_all__gopurs_runtime_Value_4179648253(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_all__gopurs_runtime_Value_4179648253
}

var cache_and gopurs_runtime.Value
var once_and sync.Once
func Get_and() gopurs_runtime.Value {
	once_and.Do(func() {
		cache_and = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_and(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_and
}

type Constructor_Empty[T_a any] struct {
	Rc uint32
}


type Constructor_Node[T_a any] struct {
	Rc uint32
	V0 T_a
}


type Constructor_Append[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Foldable[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4280266298] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Foldable[gopurs_runtime.Value])(ptr)
		switch key {
		case "foldMap": return c.V0
		case "foldl": return c.V1
		case "foldr": return c.V2
		default: panic("Key not found in dictionary Constructor_Foldable: " + key)
		}
	}
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_foldr(dict_0_loop *Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__gopurs_runtime_Value_2151204251(dict_0_loop *Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_indexr(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], idx_1_loop int64) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var idx_1 int64 = idx_1_loop
_ = idx_1
__local_var_2_0 := gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(cursor_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.RecordGet(cursor_3, "elem")
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t2 = cursor_3
goto end_branch_2
} else {

}
}
{
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(cursor_3, "pos").IntVal) == (idx_1) {
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, a_2})}))}, gopurs_runtime.Int(gopurs_runtime.RecordGet(cursor_3, "pos").IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(cursor_3, "elem")))}, gopurs_runtime.Int((gopurs_runtime.RecordGet(cursor_3, "pos").IntVal) + (1)))
}
end_branch_1:
__t2 = __t1
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}, gopurs_runtime.Int(0)))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_2_0, x_3), "elem")
})
}

func Call_null(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
})
}), gopurs_runtime.Bool(true))
}

func Call_oneOf(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictPlus_1_loop *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value] = dictPlus_1_loop
_ = dictPlus_1
return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictPlus_1.V0, gopurs_runtime.Value{}), "alt"), dictPlus_1.V1)
}

func Call_oneOfMap(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictPlus_1_loop *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value] = dictPlus_1_loop
_ = dictPlus_1
alt_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictPlus_1.V0, gopurs_runtime.Value{}), "alt")
_ = alt_2_0
empty_3_1 := dictPlus_1.V1
_ = empty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_2_0, gopurs_runtime.Apply(f_4, x_5))
}), empty_3_1)
})
}

func Call_oneOfMap__gopurs_runtime_Value_3719016818(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictPlus_1_loop *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictPlus_1 *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value] = dictPlus_1_loop
_ = dictPlus_1
alt_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictPlus_1.V0, gopurs_runtime.Value{}), "alt")
_ = alt_2_0
empty_3_1 := dictPlus_1.V1
_ = empty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(alt_2_0, gopurs_runtime.Apply(f_4, x_5))
}), empty_3_1)
})
}

func Call_traverse_(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_1 := gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{})
_ = __local_var_1_1
Functor0_2_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
applySecond_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(Functor0_2_2.V0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
}), gopurs_runtime.Apply(dictApplicative_0.V1, pkg_Data_Unit.Get_unit()))
})
})
}

func Call_traverse___gopurs_runtime_Value_996968168(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_1 := gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{})
_ = __local_var_1_1
Functor0_2_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
applySecond_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), gopurs_runtime.Apply2(Functor0_2_2.V0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
}), gopurs_runtime.Apply(dictApplicative_0.V1, pkg_Data_Unit.Get_unit()))
})
})
}

func Call_for_(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
traverse_1_1_0 := Call_traverse_(dictApplicative_0)
_ = traverse_1_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(traverse_1_1_0, dictFoldable_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_3_1, a_5, b_4)
})
})
})
}

func Call_sequence_(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictFoldable_1_loop *Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var dictFoldable_1 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_1_loop
_ = dictFoldable_1
return gopurs_runtime.Apply2(Call_traverse_(dictApplicative_0), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_1)}, Get_identity())
}

func Call_foldl(dict_0_loop *Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__gopurs_runtime_Value_2151204251(dict_0_loop *Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_indexl(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], idx_1_loop int64) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var idx_1 int64 = idx_1_loop
_ = idx_1
__local_var_2_0 := gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(cursor_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.RecordGet(cursor_2, "elem")
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr != nil) {
__t2 = cursor_2
goto end_branch_2
} else {

}
}
{
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(cursor_2, "pos").IntVal) == (idx_1) {
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, a_3})}))}, gopurs_runtime.Int(gopurs_runtime.RecordGet(cursor_2, "pos").IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.RecordGet(cursor_2, "elem")))}, gopurs_runtime.Int((gopurs_runtime.RecordGet(cursor_2, "pos").IntVal) + (1)))
}
end_branch_1:
__t2 = __t1
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("elem", "pos", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}, gopurs_runtime.Int(0)))
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(__local_var_2_0, x_3), "elem")
})
}

func Call_intercalate(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
mempty_3_1 := dictMonoid_1.V1
_ = mempty_3_1
return gopurs_runtime.Func(func(sep_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(dictFoldable_0.V1, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(Semigroup0_2_0.V0, gopurs_runtime.RecordGet(v_6, "acc"), gopurs_runtime.Apply2(Semigroup0_2_0.V0, sep_4, v1_7)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", mempty_3_1, gopurs_runtime.Bool(true)), xs_5), "acc")
})
})
}

func Call_intercalate__gopurs_runtime_Value_3939234276(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
mempty_3_1 := dictMonoid_1.V1
_ = mempty_3_1
return gopurs_runtime.Func(func(sep_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(dictFoldable_0.V1, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(Semigroup0_2_0.V0, gopurs_runtime.RecordGet(v_6, "acc"), gopurs_runtime.Apply2(Semigroup0_2_0.V0, sep_4, v1_7)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", mempty_3_1, gopurs_runtime.Bool(true)), xs_5), "acc")
})
})
}

func Call_length(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictSemiring_1_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemiring_1.V0, dictSemiring_1.V2, c_2)
})
}), dictSemiring_1.V3)
}

func Call_length__gopurs_runtime_Value_949294460(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictSemiring_1_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemiring_1.V0, dictSemiring_1.V2, c_2)
})
}), dictSemiring_1.V3)
}

func Call_maximumBy(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v1_3})}))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3), gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __t1})}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_maximumBy__gopurs_runtime_Value_110571494(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v1_3})}))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3), gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __t1})}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_maximum(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximumBy(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_2), compare_1_0)
})
}

func Call_minimumBy(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v1_3})}))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}).IntVal) != (0) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __t1})}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_minimumBy__gopurs_runtime_Value_110571494(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v1_3})}))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), gopurs_runtime.Apply2(cmp_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0, v1_3), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}).IntVal) != (0) {
__t1 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = v1_3
}
end_branch_1:
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __t1})}))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_minimum(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
compare_1_0 := dictOrd_0.V1
_ = compare_1_0
return gopurs_runtime.Func(func(dictFoldable_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimumBy(gopurs_runtime.CoerceToStruct[Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_2), compare_1_0)
})
}

func Call_product(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictSemiring_1_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, dictSemiring_1.V1, dictSemiring_1.V2)
}

func Call_sum(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictSemiring_1_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, dictSemiring_1.V0, dictSemiring_1.V3)
}

func Call_foldMapDefaultR(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
mempty_3_1 := dictMonoid_1.V1
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V2, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_2_0.V0, gopurs_runtime.Apply(f_4, x_5), acc_6)
})
}), mempty_3_1)
})
}

func Call_foldMapDefaultL(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
mempty_3_1 := dictMonoid_1.V1
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_2_0.V0, acc_5, gopurs_runtime.Apply(f_4, x_6))
})
}), mempty_3_1)
})
}

func Call_foldMap(dict_0_loop *Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldMap__gopurs_runtime_Value_4098395794(dict_0_loop *Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldableApp(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1, f_2, v_3)
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

func Call_foldableCompose(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), dictMonoid_2, f_3), v_4)
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

func Call_foldableCoproduct(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_2, f_3)
_ = __local_var_4_0
__local_var_5_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), dictMonoid_2, f_3)
_ = __local_var_5_1
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply(__local_var_4_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply(__local_var_5_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
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
__local_var_4_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, z_3)
_ = __local_var_4_3
__local_var_5_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2, z_3)
_ = __local_var_5_4
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t5 = gopurs_runtime.Apply(__local_var_4_3, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t5 = gopurs_runtime.Apply(__local_var_5_4, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
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
__local_var_4_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_2, z_3)
_ = __local_var_4_6
__local_var_5_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2, z_3)
_ = __local_var_5_7
return gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t8 = gopurs_runtime.Apply(__local_var_4_6, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t8 = gopurs_runtime.Apply(__local_var_5_7, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_6.UnsafePtr).V0)
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

func Call_foldableProduct(dictFoldable_0_loop gopurs_runtime.Value, dictFoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var dictFoldable1_1 gopurs_runtime.Value = dictFoldable1_1_loop
_ = dictFoldable1_1
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_3_0.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_2, f_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldMap"), dictMonoid_2, f_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldl"), f_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, z_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_2, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable1_1, "foldr"), f_2, z_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
})
})
}))
}

func Call_foldlDefault(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldl"), c_1, u_2, gopurs_runtime.Apply3(dictFoldable_0.V0, Get_monoidFreeMonoidTree(), Get_Node(), xs_3))
}

func Call_foldrDefault(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), c_1, u_2, gopurs_runtime.Apply3(dictFoldable_0.V0, Get_monoidFreeMonoidTree(), Get_Node(), xs_3))
}

func Call_foldrDefault__gopurs_runtime_Value_2151204251(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_foldableFreeMonoidTree(), "foldr"), c_1, u_2, gopurs_runtime.Apply3(dictFoldable_0.V0, Get_monoidFreeMonoidTree(), Get_Node(), xs_3))
}

func Call_lookup(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictEq_1_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictEq_1 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_1_loop
_ = dictEq_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
__local_var_3_0 := gopurs_runtime.Apply2(dictFoldable_0.V0, pkg_Data_Maybe_First.Get_monoidFirst(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(dictEq_1.V0, a_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1})}))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t1))}
}))
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_0, x_4)
})
}

func Call_surroundMap(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictSemigroup_1_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value], d_2_loop gopurs_runtime.Value, t_3_loop gopurs_runtime.Value, f_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemigroup_1 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
var t_3 gopurs_runtime.Value = t_3_loop
_ = t_3
var f_4 gopurs_runtime.Value = f_4_loop
_ = f_4
return gopurs_runtime.Apply4(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidEndo()))}, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemigroup_1.V0, d_2, gopurs_runtime.Apply2(dictSemigroup_1.V0, gopurs_runtime.Apply(t_3, a_5), m_6))
})
}), f_4, d_2)
}

func Call_surroundMap__gopurs_runtime_Value_3689038427(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictSemigroup_1_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value], d_2_loop gopurs_runtime.Value, t_3_loop gopurs_runtime.Value, f_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemigroup_1 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
var t_3 gopurs_runtime.Value = t_3_loop
_ = t_3
var f_4 gopurs_runtime.Value = f_4_loop
_ = f_4
return gopurs_runtime.Apply4(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidEndo()))}, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemigroup_1.V0, d_2, gopurs_runtime.Apply2(dictSemigroup_1.V0, gopurs_runtime.Apply(t_3, a_5), m_6))
})
}), f_4, d_2)
}

func Call_surround(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictSemigroup_1_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value], d_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemigroup_1 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
return gopurs_runtime.Apply4(Get_surroundMap(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_1)}, d_2, Get_identity1())
}

func Call_foldM(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictMonad_1_loop *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonad_1 *pkg_Control_Monad.Constructor_Monad[gopurs_runtime.Value] = dictMonad_1_loop
_ = dictMonad_1
Bind1_2_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_1.V1, gopurs_runtime.Value{}))
_ = Bind1_2_0
Applicative0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_1.V0, gopurs_runtime.Value{}))
_ = Applicative0_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b0_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_0.V1, b_6, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_4, a_8, a_7)
}))
})
}), gopurs_runtime.Apply(Applicative0_3_1.V1, b0_5))
})
})
}

func Call_fold(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply2(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, Get_identity1())
}

func Call_fold__gopurs_runtime_Value_910331789(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply2(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, Get_identity1())
}

func Call_findMap(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(p_1, v1_3)))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_find(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil)) && ((gopurs_runtime.Apply(p_1, v1_3).IntVal) != (0)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, v1_3})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_any(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictHeytingAlgebra_1_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_1.V1, v_2, v1_3)
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply(dictFoldable_0.V0, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), dictHeytingAlgebra_1.V2))
}

func Call_any__gopurs_runtime_Value_4179648253(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictHeytingAlgebra_1_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_1.V1, v_2, v1_3)
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply(dictFoldable_0.V0, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), dictHeytingAlgebra_1.V2))
}

func Call_elem(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
semigroupDisj1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_1, v1_2)
})
}))
_ = semigroupDisj1_1_1
any1_1_0 := gopurs_runtime.Apply(dictFoldable_0.V0, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_1
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")))
_ = any1_1_0
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(any1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq_2, "eq"), x_3))
})
})
}

func Call_elem__gopurs_runtime_Value_2343844090(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
semigroupDisj1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), v_1, v1_2)
})
}))
_ = semigroupDisj1_1_1
any1_1_0 := gopurs_runtime.Apply(dictFoldable_0.V0, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_1_1
}), gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")))
_ = any1_1_0
return gopurs_runtime.Func(func(dictEq_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(any1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq_2, "eq"), x_3))
})
})
}

func Call_notElem(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictEq_1_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictEq_1 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_1_loop
_ = dictEq_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
__local_var_3_0 := gopurs_runtime.Apply2(Call_elem(dictFoldable_0), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_1)}, x_2)
_ = __local_var_3_0
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(__local_var_3_0, x_4))
})
}

func Call_or(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictHeytingAlgebra_1_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupDisj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_1.V1, v_2, v1_3)
})
}))
_ = semigroupDisj1_2_0
return gopurs_runtime.Apply2(dictFoldable_0.V0, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_0
}), dictHeytingAlgebra_1.V2), Get_identity2())
}

func Call_all(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictHeytingAlgebra_1_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_1.V0, v_2, v1_3)
})
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply(dictFoldable_0.V0, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
}), dictHeytingAlgebra_1.V5))
}

func Call_all__gopurs_runtime_Value_4179648253(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictHeytingAlgebra_1_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_1.V0, v_2, v1_3)
})
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply(dictFoldable_0.V0, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
}), dictHeytingAlgebra_1.V5))
}

func Call_and(dictFoldable_0_loop *Constructor_Foldable[gopurs_runtime.Value], dictHeytingAlgebra_1_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictHeytingAlgebra_1 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupConj1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictHeytingAlgebra_1.V0, v_2, v1_3)
})
}))
_ = semigroupConj1_2_0
return gopurs_runtime.Apply2(dictFoldable_0.V0, gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_0
}), dictHeytingAlgebra_1.V5), Get_identity2())
}

func Get_foldlArray() gopurs_runtime.Value {
	return _Gopurs_FoldlArray
}

func Get_foldrArray() gopurs_runtime.Value {
	return _Gopurs_FoldrArray
}

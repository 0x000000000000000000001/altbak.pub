package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_FoldableWithIndex_foldr gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr sync.Once
func Get_Data_FoldableWithIndex_foldr() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr.Do(func() {
		cache_Data_FoldableWithIndex_foldr = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMultiplicative(), "foldr")
	})
	return cache_Data_FoldableWithIndex_foldr
}

var cache_Data_FoldableWithIndex_foldl gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl sync.Once
func Get_Data_FoldableWithIndex_foldl() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl.Do(func() {
		cache_Data_FoldableWithIndex_foldl = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMultiplicative(), "foldl")
	})
	return cache_Data_FoldableWithIndex_foldl
}

var cache_Data_FoldableWithIndex_foldMap gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap sync.Once
func Get_Data_FoldableWithIndex_foldMap() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap.Do(func() {
		cache_Data_FoldableWithIndex_foldMap = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMultiplicative(), "foldMap")
	})
	return cache_Data_FoldableWithIndex_foldMap
}

var cache_Data_FoldableWithIndex_foldr1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr1 sync.Once
func Get_Data_FoldableWithIndex_foldr1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr1.Do(func() {
		cache_Data_FoldableWithIndex_foldr1 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMaybe(), "foldr")
	})
	return cache_Data_FoldableWithIndex_foldr1
}

var cache_Data_FoldableWithIndex_foldl1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl1 sync.Once
func Get_Data_FoldableWithIndex_foldl1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl1.Do(func() {
		cache_Data_FoldableWithIndex_foldl1 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMaybe(), "foldl")
	})
	return cache_Data_FoldableWithIndex_foldl1
}

var cache_Data_FoldableWithIndex_foldMap1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap1 sync.Once
func Get_Data_FoldableWithIndex_foldMap1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap1.Do(func() {
		cache_Data_FoldableWithIndex_foldMap1 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMaybe(), "foldMap")
	})
	return cache_Data_FoldableWithIndex_foldMap1
}

var cache_Data_FoldableWithIndex_foldr2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr2 sync.Once
func Get_Data_FoldableWithIndex_foldr2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr2.Do(func() {
		cache_Data_FoldableWithIndex_foldr2 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableLast(), "foldr")
	})
	return cache_Data_FoldableWithIndex_foldr2
}

var cache_Data_FoldableWithIndex_foldl2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl2 sync.Once
func Get_Data_FoldableWithIndex_foldl2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl2.Do(func() {
		cache_Data_FoldableWithIndex_foldl2 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableLast(), "foldl")
	})
	return cache_Data_FoldableWithIndex_foldl2
}

var cache_Data_FoldableWithIndex_foldMap2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap2 sync.Once
func Get_Data_FoldableWithIndex_foldMap2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap2.Do(func() {
		cache_Data_FoldableWithIndex_foldMap2 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableLast(), "foldMap")
	})
	return cache_Data_FoldableWithIndex_foldMap2
}

var cache_Data_FoldableWithIndex_foldr3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr3 sync.Once
func Get_Data_FoldableWithIndex_foldr3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr3.Do(func() {
		cache_Data_FoldableWithIndex_foldr3 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFirst(), "foldr")
	})
	return cache_Data_FoldableWithIndex_foldr3
}

var cache_Data_FoldableWithIndex_foldl3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl3 sync.Once
func Get_Data_FoldableWithIndex_foldl3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl3.Do(func() {
		cache_Data_FoldableWithIndex_foldl3 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFirst(), "foldl")
	})
	return cache_Data_FoldableWithIndex_foldl3
}

var cache_Data_FoldableWithIndex_foldMap3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap3 sync.Once
func Get_Data_FoldableWithIndex_foldMap3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap3.Do(func() {
		cache_Data_FoldableWithIndex_foldMap3 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFirst(), "foldMap")
	})
	return cache_Data_FoldableWithIndex_foldMap3
}

var cache_Data_FoldableWithIndex_foldr4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr4 sync.Once
func Get_Data_FoldableWithIndex_foldr4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr4.Do(func() {
		cache_Data_FoldableWithIndex_foldr4 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDual(), "foldr")
	})
	return cache_Data_FoldableWithIndex_foldr4
}

var cache_Data_FoldableWithIndex_foldl4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl4 sync.Once
func Get_Data_FoldableWithIndex_foldl4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl4.Do(func() {
		cache_Data_FoldableWithIndex_foldl4 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDual(), "foldl")
	})
	return cache_Data_FoldableWithIndex_foldl4
}

var cache_Data_FoldableWithIndex_foldMap4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap4 sync.Once
func Get_Data_FoldableWithIndex_foldMap4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap4.Do(func() {
		cache_Data_FoldableWithIndex_foldMap4 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDual(), "foldMap")
	})
	return cache_Data_FoldableWithIndex_foldMap4
}

var cache_Data_FoldableWithIndex_foldr5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr5 sync.Once
func Get_Data_FoldableWithIndex_foldr5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr5.Do(func() {
		cache_Data_FoldableWithIndex_foldr5 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDisj(), "foldr")
	})
	return cache_Data_FoldableWithIndex_foldr5
}

var cache_Data_FoldableWithIndex_foldl5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl5 sync.Once
func Get_Data_FoldableWithIndex_foldl5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl5.Do(func() {
		cache_Data_FoldableWithIndex_foldl5 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDisj(), "foldl")
	})
	return cache_Data_FoldableWithIndex_foldl5
}

var cache_Data_FoldableWithIndex_foldMap5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap5 sync.Once
func Get_Data_FoldableWithIndex_foldMap5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap5.Do(func() {
		cache_Data_FoldableWithIndex_foldMap5 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDisj(), "foldMap")
	})
	return cache_Data_FoldableWithIndex_foldMap5
}

var cache_Data_FoldableWithIndex_foldr6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr6 sync.Once
func Get_Data_FoldableWithIndex_foldr6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr6.Do(func() {
		cache_Data_FoldableWithIndex_foldr6 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableConj(), "foldr")
	})
	return cache_Data_FoldableWithIndex_foldr6
}

var cache_Data_FoldableWithIndex_foldl6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl6 sync.Once
func Get_Data_FoldableWithIndex_foldl6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl6.Do(func() {
		cache_Data_FoldableWithIndex_foldl6 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableConj(), "foldl")
	})
	return cache_Data_FoldableWithIndex_foldl6
}

var cache_Data_FoldableWithIndex_foldMap6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap6 sync.Once
func Get_Data_FoldableWithIndex_foldMap6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap6.Do(func() {
		cache_Data_FoldableWithIndex_foldMap6 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableConj(), "foldMap")
	})
	return cache_Data_FoldableWithIndex_foldMap6
}

var cache_Data_FoldableWithIndex_foldr7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr7 sync.Once
func Get_Data_FoldableWithIndex_foldr7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr7.Do(func() {
		cache_Data_FoldableWithIndex_foldr7 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableAdditive(), "foldr")
	})
	return cache_Data_FoldableWithIndex_foldr7
}

var cache_Data_FoldableWithIndex_foldl7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl7 sync.Once
func Get_Data_FoldableWithIndex_foldl7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl7.Do(func() {
		cache_Data_FoldableWithIndex_foldl7 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableAdditive(), "foldl")
	})
	return cache_Data_FoldableWithIndex_foldl7
}

var cache_Data_FoldableWithIndex_foldMap7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap7 sync.Once
func Get_Data_FoldableWithIndex_foldMap7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap7.Do(func() {
		cache_Data_FoldableWithIndex_foldMap7 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableAdditive(), "foldMap")
	})
	return cache_Data_FoldableWithIndex_foldMap7
}

var cache_Data_FoldableWithIndex_monoidDual gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidDual sync.Once
func Get_Data_FoldableWithIndex_monoidDual() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidDual.Do(func() {
		cache_Data_FoldableWithIndex_monoidDual = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Category_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_1
// TAST (Let): semigroupEndo1_0_0 -> gopurs_runtime.Value
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "compose"), v_1, v1_2)
})
}))
_ = semigroupEndo1_0_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
_ = __local_var_1_2
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_4
// TAST (Let): semigroupDual1_2_3 -> gopurs_runtime.Value
semigroupDual1_2_3 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_4, "append"), v1_4, v_3)
})
}))
_ = semigroupDual1_2_3
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_2_3
}), gopurs_runtime.RecordGet(__local_var_1_2, "mempty"))))}
}()
	})
	return cache_Data_FoldableWithIndex_monoidDual
}

var cache_Data_FoldableWithIndex_monoidEndo gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidEndo sync.Once
func Get_Data_FoldableWithIndex_monoidEndo() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidEndo.Do(func() {
		cache_Data_FoldableWithIndex_monoidEndo = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Category_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_1
// TAST (Let): semigroupEndo1_0_0 -> gopurs_runtime.Value
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "compose"), v_1, v1_2)
})
}))
_ = semigroupEndo1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))))}
}()
	})
	return cache_Data_FoldableWithIndex_monoidEndo
}

var cache_Data_FoldableWithIndex_monoidEndo1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidEndo1 sync.Once
func Get_Data_FoldableWithIndex_monoidEndo1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidEndo1.Do(func() {
		cache_Data_FoldableWithIndex_monoidEndo1 = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Category_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_1
// TAST (Let): semigroupEndo1_0_0 -> gopurs_runtime.Value
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "compose"), v_1, v1_2)
})
}))
_ = semigroupEndo1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))))}
}()
	})
	return cache_Data_FoldableWithIndex_monoidEndo1
}

var cache_Data_FoldableWithIndex_unwrap gopurs_runtime.Value
var once_Data_FoldableWithIndex_unwrap sync.Once
func Get_Data_FoldableWithIndex_unwrap() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_unwrap.Do(func() {
		cache_Data_FoldableWithIndex_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_FoldableWithIndex_unwrap
}

var cache_Data_FoldableWithIndex_FoldableWithIndex_dollarDict gopurs_runtime.Value
var once_Data_FoldableWithIndex_FoldableWithIndex_dollarDict sync.Once
func Get_Data_FoldableWithIndex_FoldableWithIndex_dollarDict() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_FoldableWithIndex_dollarDict.Do(func() {
		cache_Data_FoldableWithIndex_FoldableWithIndex_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_FoldableWithIndex_dollarDict(x_0_box)
})
	})
	return cache_Data_FoldableWithIndex_FoldableWithIndex_dollarDict
}

var cache_Data_FoldableWithIndex_foldrWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex
}

var cache_Data_FoldableWithIndex_traverseWithIndex_ gopurs_runtime.Value
var once_Data_FoldableWithIndex_traverseWithIndex_ sync.Once
func Get_Data_FoldableWithIndex_traverseWithIndex_() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_traverseWithIndex_.Do(func() {
		cache_Data_FoldableWithIndex_traverseWithIndex_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_traverseWithIndex_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Data_FoldableWithIndex_traverseWithIndex_
}

var cache_Data_FoldableWithIndex_forWithIndex_ gopurs_runtime.Value
var once_Data_FoldableWithIndex_forWithIndex_ sync.Once
func Get_Data_FoldableWithIndex_forWithIndex_() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_forWithIndex_.Do(func() {
		cache_Data_FoldableWithIndex_forWithIndex_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_forWithIndex_(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Data_FoldableWithIndex_forWithIndex_
}

var cache_Data_FoldableWithIndex_foldrDefault gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrDefault sync.Once
func Get_Data_FoldableWithIndex_foldrDefault() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrDefault.Do(func() {
		cache_Data_FoldableWithIndex_foldrDefault = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), f_1_box)
})
	})
	return cache_Data_FoldableWithIndex_foldrDefault
}

var cache_Data_FoldableWithIndex_foldlWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex
}

var cache_Data_FoldableWithIndex_foldlDefault gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlDefault sync.Once
func Get_Data_FoldableWithIndex_foldlDefault() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlDefault.Do(func() {
		cache_Data_FoldableWithIndex_foldlDefault = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), f_1_box)
})
	})
	return cache_Data_FoldableWithIndex_foldlDefault
}

var cache_Data_FoldableWithIndex_foldableWithIndexTuple gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexTuple sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexTuple() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexTuple.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexTuple = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableTuple()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), z_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
})
})
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexTuple
}

var cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexMultiplicative sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexMultiplicative() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexMultiplicative.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMultiplicative()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMultiplicative(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMultiplicative(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMultiplicative(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative
}

var cache_Data_FoldableWithIndex_foldableWithIndexMaybe gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexMaybe sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexMaybe() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexMaybe.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexMaybe = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMaybe()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMaybe(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMaybe(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMaybe(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexMaybe
}

var cache_Data_FoldableWithIndex_foldableWithIndexLast gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexLast sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexLast() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexLast.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexLast = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableLast()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableLast(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableLast(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableLast(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexLast
}

var cache_Data_FoldableWithIndex_foldableWithIndexIdentity gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexIdentity sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexIdentity() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexIdentity.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexIdentity = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableIdentity()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), v_2, z_1)
})
})
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexIdentity
}

var cache_Data_FoldableWithIndex_foldableWithIndexFirst gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexFirst sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexFirst() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexFirst.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexFirst = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableFirst()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFirst(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFirst(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFirst(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexFirst
}

var cache_Data_FoldableWithIndex_foldableWithIndexEither gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexEither sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexEither() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexEither.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexEither = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableEither()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply2(v_2, Get_Data_Unit_unit(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
__t2 = gopurs_runtime.Apply3(v_0, Get_Data_Unit_unit(), v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
__t3 = gopurs_runtime.Apply3(v_0, Get_Data_Unit_unit(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
	return cache_Data_FoldableWithIndex_foldableWithIndexEither
}

var cache_Data_FoldableWithIndex_foldableWithIndexDual gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexDual sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexDual() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexDual.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexDual = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableDual()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDual(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDual(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDual(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexDual
}

var cache_Data_FoldableWithIndex_foldableWithIndexDisj gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexDisj sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexDisj() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexDisj.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexDisj = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableDisj()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDisj(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDisj(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDisj(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexDisj
}

var cache_Data_FoldableWithIndex_foldableWithIndexConst gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexConst sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexConst() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexConst.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexConst = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableConst()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_FoldableWithIndex_foldableWithIndexConst
}

var cache_Data_FoldableWithIndex_foldableWithIndexConj gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexConj sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexConj() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexConj.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexConj = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableConj()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableConj(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableConj(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableConj(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexConj
}

var cache_Data_FoldableWithIndex_foldableWithIndexAdditive gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexAdditive sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexAdditive() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexAdditive.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexAdditive = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableAdditive()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableAdditive(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableAdditive(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableAdditive(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexAdditive
}

var cache_Data_FoldableWithIndex_foldWithIndexM gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldWithIndexM sync.Once
func Get_Data_FoldableWithIndex_foldWithIndexM() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldWithIndexM.Do(func() {
		cache_Data_FoldableWithIndex_foldWithIndexM = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldWithIndexM(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_1_box))
})
	})
	return cache_Data_FoldableWithIndex_foldWithIndexM
}

var cache_Data_FoldableWithIndex_foldMapWithIndexDefaultR gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndexDefaultR sync.Once
func Get_Data_FoldableWithIndex_foldMapWithIndexDefaultR() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndexDefaultR.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndexDefaultR = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapWithIndexDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndexDefaultR
}

var cache_Data_FoldableWithIndex_foldableWithIndexArray gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexArray sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexArray() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexArray.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexArray = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableArray()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_FoldableWithIndex_foldableWithIndexArray(), "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(f_3, i_4, x_5), acc_6)
})
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl"), gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal), y_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
}), z_1)
_ = __local_var_2_2
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_FunctorWithIndex_functorWithIndexArray(), "mapWithIndex"), Get_Data_Tuple_Tuple())
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(__local_var_3_3, x_4))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldr"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_3_5
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_4_6
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_3_5.IntVal), __local_var_4_6, y_5)
})
}), z_1)
_ = __local_var_2_4
// TAST (Let): __local_var_3_7 -> gopurs_runtime.Value
__local_var_3_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_FunctorWithIndex_functorWithIndexArray(), "mapWithIndex"), Get_Data_Tuple_Tuple())
_ = __local_var_3_7
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(__local_var_3_7, x_4))
})
})
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexArray
}

var cache_Data_FoldableWithIndex_foldMapWithIndexDefaultL gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndexDefaultL sync.Once
func Get_Data_FoldableWithIndex_foldMapWithIndexDefaultL() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndexDefaultL.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndexDefaultL = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapWithIndexDefaultL(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndexDefaultL
}

var cache_Data_FoldableWithIndex_foldMapWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndex sync.Once
func Get_Data_FoldableWithIndex_foldMapWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndex
}

var cache_Data_FoldableWithIndex_foldableWithIndexApp gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexApp sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexApp() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexApp.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexApp = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldableWithIndexApp(dictFoldableWithIndex_0_box)
})
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexApp
}

var cache_Data_FoldableWithIndex_foldableWithIndexCompose gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexCompose sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexCompose() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexCompose.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexCompose = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldableWithIndexCompose(dictFoldableWithIndex_0_box)
})
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexCompose
}

var cache_Data_FoldableWithIndex_foldableWithIndexCoproduct gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexCoproduct sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexCoproduct() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexCoproduct.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexCoproduct = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldableWithIndexCoproduct(dictFoldableWithIndex_0_box)
})
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexCoproduct
}

var cache_Data_FoldableWithIndex_foldableWithIndexProduct gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexProduct sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexProduct() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexProduct.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexProduct = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldableWithIndexProduct(dictFoldableWithIndex_0_box)
})
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexProduct
}

var cache_Data_FoldableWithIndex_foldlWithIndexDefault gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndexDefault sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndexDefault() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndexDefault.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndexDefault = gopurs_runtime.Func4(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndexDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndexDefault
}

var cache_Data_FoldableWithIndex_foldrWithIndexDefault gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndexDefault sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndexDefault() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndexDefault.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndexDefault = gopurs_runtime.Func4(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, c_1_box gopurs_runtime.Value, u_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndexDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), c_1_box, u_2_box, xs_3_box)
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndexDefault
}

var cache_Data_FoldableWithIndex_surroundMapWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_surroundMapWithIndex sync.Once
func Get_Data_FoldableWithIndex_surroundMapWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_surroundMapWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_surroundMapWithIndex = gopurs_runtime.Func5(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value, d_2_box gopurs_runtime.Value, t_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_surroundMapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box), d_2_box, t_3_box, f_4_box)
})
	})
	return cache_Data_FoldableWithIndex_surroundMapWithIndex
}

var cache_Data_FoldableWithIndex_foldMapDefault gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapDefault sync.Once
func Get_Data_FoldableWithIndex_foldMapDefault() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapDefault.Do(func() {
		cache_Data_FoldableWithIndex_foldMapDefault = gopurs_runtime.Func3(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box), f_2_box)
})
	})
	return cache_Data_FoldableWithIndex_foldMapDefault
}

var cache_Data_FoldableWithIndex_findWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_findWithIndex sync.Once
func Get_Data_FoldableWithIndex_findWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_findWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_findWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_findWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), p_1_box)
})
	})
	return cache_Data_FoldableWithIndex_findWithIndex
}

var cache_Data_FoldableWithIndex_findMapWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_findMapWithIndex sync.Once
func Get_Data_FoldableWithIndex_findMapWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_findMapWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_findMapWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_findMapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), f_1_box)
})
	})
	return cache_Data_FoldableWithIndex_findMapWithIndex
}

var cache_Data_FoldableWithIndex_anyWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_anyWithIndex sync.Once
func Get_Data_FoldableWithIndex_anyWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_anyWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_anyWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_anyWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_Data_FoldableWithIndex_anyWithIndex
}

var cache_Data_FoldableWithIndex_allWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_allWithIndex sync.Once
func Get_Data_FoldableWithIndex_allWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_allWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_allWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_allWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_1_box))
})
	})
	return cache_Data_FoldableWithIndex_allWithIndex
}

var cache_Data_FoldableWithIndex_foldMapWithIndex__2292551140 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndex__2292551140 sync.Once
func Get_Data_FoldableWithIndex_foldMapWithIndex__2292551140() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndex__2292551140.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndex__2292551140 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapWithIndex__2292551140(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndex__2292551140
}

var cache_Data_FoldableWithIndex_foldMapWithIndex__3459474788 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndex__3459474788 sync.Once
func Get_Data_FoldableWithIndex_foldMapWithIndex__3459474788() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndex__3459474788.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndex__3459474788 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapWithIndex__3459474788(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndex__3459474788
}

var cache_Data_FoldableWithIndex_foldMapWithIndex__1757753703 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndex__1757753703 sync.Once
func Get_Data_FoldableWithIndex_foldMapWithIndex__1757753703() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndex__1757753703.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndex__1757753703 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapWithIndex__1757753703(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndex__1757753703
}

var cache_Data_FoldableWithIndex_foldMapWithIndex__1722031522 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndex__1722031522 sync.Once
func Get_Data_FoldableWithIndex_foldMapWithIndex__1722031522() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndex__1722031522.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndex__1722031522 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapWithIndex__1722031522(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndex__1722031522
}

var cache_Data_FoldableWithIndex_foldMapWithIndex__852526914 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndex__852526914 sync.Once
func Get_Data_FoldableWithIndex_foldMapWithIndex__852526914() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndex__852526914.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndex__852526914 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapWithIndex__852526914(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndex__852526914
}

var cache_Data_FoldableWithIndex_foldMapWithIndex__2880267906 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMapWithIndex__2880267906 sync.Once
func Get_Data_FoldableWithIndex_foldMapWithIndex__2880267906() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMapWithIndex__2880267906.Do(func() {
		cache_Data_FoldableWithIndex_foldMapWithIndex__2880267906 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapWithIndex__2880267906(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_FoldableWithIndex_foldMapWithIndex__2880267906
}

var cache_Data_FoldableWithIndex_foldableWithIndexAdditive__3548846699 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexAdditive__3548846699 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexAdditive__3548846699() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexAdditive__3548846699.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexAdditive__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableAdditive()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableAdditive(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableAdditive(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableAdditive(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexAdditive__3548846699
}

var cache_Data_FoldableWithIndex_foldableWithIndexArray__740253118 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexArray__740253118 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexArray__740253118() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexArray__740253118.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexArray__740253118 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableArray()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_FoldableWithIndex_foldableWithIndexArray(), "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(f_3, i_4, x_5), acc_6)
})
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl"), gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal), y_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
}), z_1)
_ = __local_var_2_2
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_FunctorWithIndex_functorWithIndexArray(), "mapWithIndex"), Get_Data_Tuple_Tuple())
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(__local_var_3_3, x_4))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldr"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_3_5
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_4_6
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_3_5.IntVal), __local_var_4_6, y_5)
})
}), z_1)
_ = __local_var_2_4
// TAST (Let): __local_var_3_7 -> gopurs_runtime.Value
__local_var_3_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_FunctorWithIndex_functorWithIndexArray(), "mapWithIndex"), Get_Data_Tuple_Tuple())
_ = __local_var_3_7
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(__local_var_3_7, x_4))
})
})
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexArray__740253118
}

var cache_Data_FoldableWithIndex_foldableWithIndexConj__3548846699 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexConj__3548846699 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexConj__3548846699() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexConj__3548846699.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexConj__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableConj()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableConj(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableConj(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableConj(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexConj__3548846699
}

var cache_Data_FoldableWithIndex_foldableWithIndexConst__1906290739 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexConst__1906290739 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexConst__1906290739() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexConst__1906290739.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexConst__1906290739 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableConst()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_FoldableWithIndex_foldableWithIndexConst__1906290739
}

var cache_Data_FoldableWithIndex_foldableWithIndexDisj__3548846699 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexDisj__3548846699 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexDisj__3548846699() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexDisj__3548846699.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexDisj__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableDisj()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDisj(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDisj(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDisj(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexDisj__3548846699
}

var cache_Data_FoldableWithIndex_foldableWithIndexDual__3548846699 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexDual__3548846699 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexDual__3548846699() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexDual__3548846699.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexDual__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableDual()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDual(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDual(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableDual(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexDual__3548846699
}

var cache_Data_FoldableWithIndex_foldableWithIndexEither__162584107 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexEither__162584107 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexEither__162584107() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexEither__162584107.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexEither__162584107 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableEither()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply2(v_2, Get_Data_Unit_unit(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
__t2 = gopurs_runtime.Apply3(v_0, Get_Data_Unit_unit(), v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
__t3 = gopurs_runtime.Apply3(v_0, Get_Data_Unit_unit(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
	return cache_Data_FoldableWithIndex_foldableWithIndexEither__162584107
}

var cache_Data_FoldableWithIndex_foldableWithIndexFirst__1642852683 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexFirst__1642852683 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexFirst__1642852683() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexFirst__1642852683.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexFirst__1642852683 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableFirst()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFirst(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFirst(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableFirst(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexFirst__1642852683
}

var cache_Data_FoldableWithIndex_foldableWithIndexIdentity__3548846699 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexIdentity__3548846699 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexIdentity__3548846699() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexIdentity__3548846699.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexIdentity__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableIdentity()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), v_2, z_1)
})
})
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexIdentity__3548846699
}

var cache_Data_FoldableWithIndex_foldableWithIndexLast__1642852683 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexLast__1642852683 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexLast__1642852683() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexLast__1642852683.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexLast__1642852683 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableLast()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableLast(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableLast(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableLast(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexLast__1642852683
}

var cache_Data_FoldableWithIndex_foldableWithIndexMaybe__1642852683 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexMaybe__1642852683 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexMaybe__1642852683() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexMaybe__1642852683.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexMaybe__1642852683 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMaybe()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMaybe(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMaybe(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMaybe(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexMaybe__1642852683
}

var cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative__3548846699 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexMultiplicative__3548846699 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexMultiplicative__3548846699() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexMultiplicative__3548846699.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMultiplicative()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 -> gopurs_runtime.Value
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMultiplicative(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMultiplicative(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableMultiplicative(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative__3548846699
}

var cache_Data_FoldableWithIndex_foldableWithIndexTuple__4170851755 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexTuple__4170851755 sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexTuple__4170851755() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexTuple__4170851755.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexTuple__4170851755 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableTuple()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), z_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
})
})
}))
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexTuple__4170851755
}

var cache_Data_FoldableWithIndex_foldlWithIndex__2972270123 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__2972270123 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__2972270123() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__2972270123.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__2972270123 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__2972270123(f_0_box, acc_1_box)
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__2972270123
}

var cache_Data_FoldableWithIndex_foldlWithIndex__234438827 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__234438827 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__234438827() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__234438827.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__234438827 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__234438827(f_0_box, acc_1_box)
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__234438827
}

var cache_Data_FoldableWithIndex_foldlWithIndex__2808220203 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__2808220203 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__2808220203() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__2808220203.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__2808220203 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__2808220203(f_0_box, acc_1_box)
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__2808220203
}

var cache_Data_FoldableWithIndex_foldlWithIndex__2764250251 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__2764250251 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__2764250251() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__2764250251.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__2764250251 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, acc_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__2764250251(f_0_box, acc_1_box)
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__2764250251
}

var cache_Data_FoldableWithIndex_foldlWithIndex__2986161357 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__2986161357 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__2986161357() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__2986161357.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__2986161357(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__2986161357
}

var cache_Data_FoldableWithIndex_foldlWithIndex__2942277133 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__2942277133 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__2942277133() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__2942277133.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__2942277133 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__2942277133(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__2942277133
}

var cache_Data_FoldableWithIndex_foldlWithIndex__2499716749 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__2499716749 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__2499716749() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__2499716749.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__2499716749 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__2499716749(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__2499716749
}

var cache_Data_FoldableWithIndex_foldlWithIndex__1917751149 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__1917751149 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__1917751149() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__1917751149.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__1917751149 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__1917751149(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__1917751149
}

var cache_Data_FoldableWithIndex_foldlWithIndex__1224542477 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__1224542477 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__1224542477() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__1224542477.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__1224542477 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__1224542477(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__1224542477
}

var cache_Data_FoldableWithIndex_foldlWithIndex__3618272333 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__3618272333 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__3618272333() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__3618272333.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__3618272333 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__3618272333(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__3618272333
}

var cache_Data_FoldableWithIndex_foldlWithIndex__3610348555 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__3610348555 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__3610348555() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__3610348555.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__3610348555 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__3610348555(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__3610348555
}

var cache_Data_FoldableWithIndex_foldlWithIndex__446277963 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__446277963 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__446277963() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__446277963.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__446277963 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__446277963(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__446277963
}

var cache_Data_FoldableWithIndex_foldlWithIndex__1651851147 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldlWithIndex__1651851147 sync.Once
func Get_Data_FoldableWithIndex_foldlWithIndex__1651851147() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldlWithIndex__1651851147.Do(func() {
		cache_Data_FoldableWithIndex_foldlWithIndex__1651851147 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldlWithIndex__1651851147(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_FoldableWithIndex_foldlWithIndex__1651851147
}

var cache_Data_FoldableWithIndex_foldrWithIndex__2972270123 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__2972270123 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__2972270123() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__2972270123.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__2972270123 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__2972270123(f_0_box, b_1_box, xs_2_box)
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__2972270123
}

var cache_Data_FoldableWithIndex_foldrWithIndex__3735894283 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__3735894283 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__3735894283() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__3735894283.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__3735894283 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__3735894283(f_0_box, b_1_box, xs_2_box)
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__3735894283
}

var cache_Data_FoldableWithIndex_foldrWithIndex__500807083 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__500807083 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__500807083() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__500807083.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__500807083 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__500807083(f_0_box, b_1_box, xs_2_box)
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__500807083
}

var cache_Data_FoldableWithIndex_foldrWithIndex__2808220203 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__2808220203 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__2808220203() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__2808220203.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__2808220203 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__2808220203(f_0_box, b_1_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2_box))
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__2808220203
}

var cache_Data_FoldableWithIndex_foldrWithIndex__2439396107 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__2439396107 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__2439396107() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__2439396107.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__2439396107 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_FoldableWithIndex_foldrWithIndex__2439396107(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](xs_2_box)))}
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__2439396107
}

var cache_Data_FoldableWithIndex_foldrWithIndex__2986161357 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__2986161357 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__2986161357() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__2986161357.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__2986161357(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__2986161357
}

var cache_Data_FoldableWithIndex_foldrWithIndex__2143732941 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__2143732941 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__2143732941() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__2143732941.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__2143732941 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__2143732941(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__2143732941
}

var cache_Data_FoldableWithIndex_foldrWithIndex__119840077 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__119840077 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__119840077() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__119840077.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__119840077 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__119840077(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__119840077
}

var cache_Data_FoldableWithIndex_foldrWithIndex__3511467915 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__3511467915 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__3511467915() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__3511467915.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__3511467915 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__3511467915(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](z_1_box))
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__3511467915
}

var cache_Data_FoldableWithIndex_foldrWithIndex__3610348555 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__3610348555 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__3610348555() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__3610348555.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__3610348555 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__3610348555(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__3610348555
}

var cache_Data_FoldableWithIndex_foldrWithIndex__63302635 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__63302635 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__63302635() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__63302635.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__63302635 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__63302635(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__63302635
}

var cache_Data_FoldableWithIndex_foldrWithIndex__979136683 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__979136683 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__979136683() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__979136683.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__979136683 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__979136683(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__979136683
}

var cache_Data_FoldableWithIndex_foldrWithIndex__3896323563 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldrWithIndex__3896323563 sync.Once
func Get_Data_FoldableWithIndex_foldrWithIndex__3896323563() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldrWithIndex__3896323563.Do(func() {
		cache_Data_FoldableWithIndex_foldrWithIndex__3896323563 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, z_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldrWithIndex__3896323563(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](z_1_box))
})
	})
	return cache_Data_FoldableWithIndex_foldrWithIndex__3896323563
}

type Constructor_Data_FoldableWithIndex_FoldableWithIndex[T_i any, T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[74250362] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Foldable0": return gopurs_runtime.Box(c.V0)
		case "foldMapWithIndex": return gopurs_runtime.Box(c.V1)
		case "foldlWithIndex": return gopurs_runtime.Box(c.V2)
		case "foldrWithIndex": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_FoldableWithIndex_FoldableWithIndex: " + key)
		}
	}
}


func Call_Data_FoldableWithIndex_FoldableWithIndex_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_FoldableWithIndex_foldrWithIndex(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_FoldableWithIndex_traverseWithIndex_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
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
return gopurs_runtime.Func(func(dictFoldableWithIndex_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_2, "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(f_3, i_4)
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit()))
})
})
}

func Call_Data_FoldableWithIndex_forWithIndex_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): traverseWithIndex_1_1_0 -> gopurs_runtime.Value
traverseWithIndex_1_1_0 := Call_Data_FoldableWithIndex_traverseWithIndex_(dictApplicative_0)
_ = traverseWithIndex_1_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(traverseWithIndex_1_1_0, dictFoldableWithIndex_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_3_1, a_5, b_4)
})
})
})
}

func Call_Data_FoldableWithIndex_foldrDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldableWithIndex_0.V3), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
}

func Call_Data_FoldableWithIndex_foldlWithIndex(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_FoldableWithIndex_foldlDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldableWithIndex_0.V2), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
}

func Call_Data_FoldableWithIndex_foldWithIndexM(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonad_1_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonad_1 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_1_loop
_ = dictMonad_1
// TAST (Let): Bind1_2_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_1.V1), gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): Applicative0_3_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_1.V0), gopurs_runtime.Value{}))
_ = Applicative0_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a0_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V2), gopurs_runtime.Func(func(i_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ma_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_2 -> gopurs_runtime.Value
__local_var_9_2 := gopurs_runtime.Apply(f_4, i_6)
_ = __local_var_9_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), ma_7, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_9_2, a_10, b_8)
}))
})
})
}), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), a0_5))
})
})
}

func Call_Data_FoldableWithIndex_foldMapWithIndexDefaultR(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 -> gopurs_runtime.Value
mempty_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V3), gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply2(f_4, i_5, x_6), acc_7)
})
})
}), mempty_3_1)
})
}

func Call_Data_FoldableWithIndex_foldMapWithIndexDefaultL(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 -> gopurs_runtime.Value
mempty_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V2), gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), acc_6, gopurs_runtime.Apply2(f_4, i_5, x_7))
})
})
}), mempty_3_1)
})
}

func Call_Data_FoldableWithIndex_foldMapWithIndex(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_FoldableWithIndex_foldableWithIndexApp(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): foldableApp_1_0 -> gopurs_runtime.Value
foldableApp_1_0 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, f_3, v_4)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldl"), f_2, i_3, v_4)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldr"), f_2, i_3, v_4)
})
})
}))
_ = foldableApp_1_0
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableApp_1_0
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, f_3, v_4)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), f_2, z_3, v_4)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), f_2, z_3, v_4)
})
})
}))
}

func Call_Data_FoldableWithIndex_foldableWithIndexCompose(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): foldableCompose1_3_1 -> gopurs_runtime.Value
foldableCompose1_3_1 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, f_5), v_6)
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "foldl"), f_4), i_5, v_6)
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "foldr"), f_4)
_ = __local_var_7_3
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldr"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_7_3, a_9, b_8)
})
}), i_5, v_6)
})
})
}))
_ = foldableCompose1_3_1
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCompose1_3_1
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMapWithIndex2_5_4 -> gopurs_runtime.Value
foldMapWithIndex2_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), dictMonoid_4)
_ = foldMapWithIndex2_5_4
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMapWithIndex2_5_4, gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_8, b_9})})
}))
}), v_7)
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_7, b_8})})
}))
}), i_5, v_6)
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_7, b_8})})
}))
_ = __local_var_8_5
return gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_8_5, a_10, b_9)
})
})
}), i_5, v_6)
})
})
}))
})
}

func Call_Data_FoldableWithIndex_foldableWithIndexCoproduct(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): foldableCoproduct_1_0 -> gopurs_runtime.Value
foldableCoproduct_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableCoproduct_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldableCoproduct1_3_1 -> gopurs_runtime.Value
foldableCoproduct1_3_1 := gopurs_runtime.Apply(foldableCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableCoproduct1_3_1
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCoproduct1_3_1
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_6})})
}))
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_7})})
}))
_ = __local_var_7_3
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t4 = gopurs_runtime.Apply(__local_var_6_2, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t4 = gopurs_runtime.Apply(__local_var_7_3, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_6})})
}), z_5)
_ = __local_var_6_5
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_7})})
}), z_5)
_ = __local_var_7_6
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(__local_var_6_5, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t7 = gopurs_runtime.Apply(__local_var_7_6, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_8 -> gopurs_runtime.Value
__local_var_6_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_6})})
}), z_5)
_ = __local_var_6_8
// TAST (Let): __local_var_7_9 -> gopurs_runtime.Value
__local_var_7_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_7})})
}), z_5)
_ = __local_var_7_9
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(__local_var_6_8, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t10 = gopurs_runtime.Apply(__local_var_7_9, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
})
}))
})
}

func Call_Data_FoldableWithIndex_foldableWithIndexProduct(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): foldableProduct_1_0 -> gopurs_runtime.Value
foldableProduct_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableProduct_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldableProduct1_3_1 -> gopurs_runtime.Value
foldableProduct1_3_1 := gopurs_runtime.Apply(foldableProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableProduct1_3_1
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableProduct1_3_1
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_2 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_2
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_2.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_8})})
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_8})})
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_7})})
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_7})})
}), z_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_7})})
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_7})})
}), z_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
})
})
}))
})
}

func Call_Data_FoldableWithIndex_foldlWithIndexDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_FoldableWithIndex_monoidDual()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(c_1, i_4)
_ = __local_var_5_2
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_5_2, a_7, b_6)
})
})
_ = __local_var_5_1
// TAST (Let): __local_var_5_0 -> gopurs_runtime.Value
__local_var_5_0 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, x_6)
})
_ = __local_var_5_0
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_0, x_6)
})
}), xs_3, u_2)
}

func Call_Data_FoldableWithIndex_foldrWithIndexDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], c_1_loop gopurs_runtime.Value, u_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var c_1 gopurs_runtime.Value = c_1_loop
_ = c_1
var u_2 gopurs_runtime.Value = u_2_loop
_ = u_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_FoldableWithIndex_monoidEndo()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_0 -> gopurs_runtime.Value
__local_var_5_0 := gopurs_runtime.Apply(c_1, i_4)
_ = __local_var_5_0
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_0, x_6)
})
}), xs_3, u_2)
}

func Call_Data_FoldableWithIndex_surroundMapWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value], d_2_loop gopurs_runtime.Value, t_3_loop gopurs_runtime.Value, f_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
var d_2 gopurs_runtime.Value = d_2_loop
_ = d_2
var t_3 gopurs_runtime.Value = t_3_loop
_ = t_3
var f_4 gopurs_runtime.Value = f_4_loop
_ = f_4
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_FoldableWithIndex_monoidEndo1()))}, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), d_2, gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), gopurs_runtime.Apply2(t_3, i_5, a_6), m_7))
})
})
}), f_4, d_2)
}

func Call_Data_FoldableWithIndex_foldMapDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value], f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return f_2
}))
}

func Call_Data_FoldableWithIndex_findWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V2), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil)) && ((gopurs_runtime.Apply2(p_1, v_2, v2_4).IntVal) != (0)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("index", "value", v_2, v2_4)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t0))}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
}

func Call_Data_FoldableWithIndex_findMapWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V2), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_1, v_2, v2_4))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
}

func Call_Data_FoldableWithIndex_anyWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): semigroupDisj1_2_1 -> gopurs_runtime.Value
semigroupDisj1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictHeytingAlgebra_1.V1), v_2, v1_3)
})
}))
_ = semigroupDisj1_2_1
// TAST (Let): monoidDisj_2_0 -> *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]
monoidDisj_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_1
}), gopurs_runtime.Box(dictHeytingAlgebra_1.V2)))
_ = monoidDisj_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(monoidDisj_2_0)}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(t_3, i_4)
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, x_6)
})
}))
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, x_5)
})
})
}

func Call_Data_FoldableWithIndex_allWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): semigroupConj1_2_1 -> gopurs_runtime.Value
semigroupConj1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictHeytingAlgebra_1.V0), v_2, v1_3)
})
}))
_ = semigroupConj1_2_1
// TAST (Let): monoidConj_2_0 -> *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]
monoidConj_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_1
}), gopurs_runtime.Box(dictHeytingAlgebra_1.V5)))
_ = monoidConj_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(monoidConj_2_0)}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(t_3, i_4)
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, x_6)
})
}))
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, x_5)
})
})
}

func Call_Data_FoldableWithIndex_foldMapWithIndex__2292551140(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_FoldableWithIndex_foldMapWithIndex__3459474788(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_FoldableWithIndex_foldMapWithIndex__1757753703(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_FoldableWithIndex_monoidDual()))})
}

func Call_Data_FoldableWithIndex_foldMapWithIndex__1722031522(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V1), __eta0_0, __eta1_1)
}

func Call_Data_FoldableWithIndex_foldMapWithIndex__852526914(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V1), __eta0_0, __eta1_1)
}

func Call_Data_FoldableWithIndex_foldMapWithIndex__2880267906(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexNonEmpty()).V1), __eta0_0, __eta1_1)
}

func Call_Data_FoldableWithIndex_foldlWithIndex__2972270123(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_2.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_2.IntVal), __local_var_3_1, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_0, x_3).UnsafePtr).V1
})
}

func Call_Data_FoldableWithIndex_foldlWithIndex__234438827(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_2.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_2.IntVal), __local_var_3_1, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](a_5))})})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_0, x_3).UnsafePtr).V1
})
}

func Call_Data_FoldableWithIndex_foldlWithIndex__2808220203(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_2.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_2.IntVal), __local_var_3_1, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_0, x_3).UnsafePtr).V1
})
}

func Call_Data_FoldableWithIndex_foldlWithIndex__2764250251(f_0_loop gopurs_runtime.Value, acc_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var acc_1 gopurs_runtime.Value = acc_1_loop
_ = acc_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_4_2
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_4_2.IntVal) + (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_4_2.IntVal), __local_var_3_1, a_5)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), acc_1})})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(__local_var_2_0, x_3).UnsafePtr).V1
})
}

func Call_Data_FoldableWithIndex_foldlWithIndex__2986161357(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_FoldableWithIndex_foldlWithIndex__2942277133(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_FoldableWithIndex_foldlWithIndex__2499716749(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_FoldableWithIndex_foldlWithIndex__1917751149(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_FoldableWithIndex_foldlWithIndex__1224542477(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_FoldableWithIndex_foldlWithIndex__3618272333(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_FoldableWithIndex_foldlWithIndex__3610348555(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V2), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_FoldableWithIndex_foldlWithIndex__446277963(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V2), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_FoldableWithIndex_foldlWithIndex__1651851147(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexNonEmpty()).V2), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_FoldableWithIndex_foldrWithIndex__2972270123(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): v_3_0 -> *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]
v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_2.IntVal) + (1)), gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_cons(), a_6, __local_var_4_1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_2))
_ = v_3_0
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), a_7, __local_var_5_3)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0.IntVal), b_1})}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V1).UnsafePtr).V1
}

func Call_Data_FoldableWithIndex_foldrWithIndex__3735894283(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): v_3_0 -> *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]
v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_2.IntVal) + (1)), gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](a_6))}, __local_var_4_1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_2))
_ = v_3_0
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](a_7))}, __local_var_5_3)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0.IntVal), b_1})}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V1).UnsafePtr).V1
}

func Call_Data_FoldableWithIndex_foldrWithIndex__500807083(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): v_3_0 -> *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]
v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_2.IntVal) + (1)), gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](a_6))}, __local_var_4_1)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_Data_List_Lazy_Types_nil()})}, xs_2))
_ = v_3_0
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons[gopurs_runtime.Value]](a_7))}, __local_var_5_3)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0.IntVal), b_1})}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V1).UnsafePtr).V1
}

func Call_Data_FoldableWithIndex_foldrWithIndex__2808220203(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, xs_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var xs_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_2_loop
_ = xs_2
// TAST (Let): v_3_0 -> *Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_2.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), a_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_4_1))})))}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_Data_List_Types_Nil()})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_2)}))
_ = v_3_0
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), a_7, __local_var_5_3)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0.IntVal), b_1})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V1))}).UnsafePtr).V1
}

func Call_Data_FoldableWithIndex_foldrWithIndex__2439396107(f_0_loop gopurs_runtime.Value, b_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], xs_2_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = b_1_loop
_ = b_1
var xs_2 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = xs_2_loop
_ = xs_2
// TAST (Let): v_3_0 -> *Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]
v_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1
_ = __local_var_4_1
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0
_ = __local_var_5_2
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_5_2.IntVal) + (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), a_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_4_1))})))}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(0), Get_Data_List_Types_Nil()})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_2)}))
_ = v_3_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V1
_ = __local_var_5_3
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_4.UnsafePtr).V0
_ = __local_var_6_4
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((__local_var_6_4.IntVal) - (1)), a_7, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_5_3))})))}})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_1)}})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V1))}).UnsafePtr).V1)
}

func Call_Data_FoldableWithIndex_foldrWithIndex__2986161357(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_FoldableWithIndex_foldrWithIndex__2143732941(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_FoldableWithIndex_foldrWithIndex__119840077(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_FoldableWithIndex_foldrWithIndex__3511467915(f_0_loop gopurs_runtime.Value, z_1_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var z_1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = z_1_loop
_ = z_1
var go__go_2_0_0 gopurs_runtime.Value
_ = go__go_2_0_0
go__go_2_0_0 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_4))}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_2_0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_2_0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__local_var_4))})))})))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](__t1))}
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp2(go__go_2_0_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](m_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(z_1)})))}
})
}

func Call_Data_FoldableWithIndex_foldrWithIndex__3610348555(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V3), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_FoldableWithIndex_foldrWithIndex__63302635(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_List_Lazy_Types_foldableWithIndexNonEmpty()).V3), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_FoldableWithIndex_foldrWithIndex__979136683(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Maybe_Just[int64], *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](Get_Data_List_Types_foldableWithIndexNonEmpty()).V3), __eta0_0, __eta1_1, __eta2_2)
}

func Call_Data_FoldableWithIndex_foldrWithIndex__3896323563(f_0_loop gopurs_runtime.Value, z_1_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var z_1 *Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] = z_1_loop
_ = z_1
var go__go_2_0_1 gopurs_runtime.Value
_ = go__go_2_0_1
go__go_2_0_1 = gopurs_runtime.Func2(func(__local_var_3 gopurs_runtime.Value, __local_var_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](__local_var_4))}
goto end_branch_1
} else {

}
}
{
if (__local_var_3.Type == 9 && __local_var_3.IntVal == 324739070 && __local_var_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](gopurs_runtime.UncurriedApp2(go__go_2_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply3(f_0, (*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](gopurs_runtime.UncurriedApp2(go__go_2_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3.UnsafePtr).V5)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](__local_var_4))})))})))})))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](__t1))}
})
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](gopurs_runtime.UncurriedApp2(go__go_2_0_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](m_3))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(z_1)})))}
})
}



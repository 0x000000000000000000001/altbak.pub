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
		cache_Data_FoldableWithIndex_foldr = Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()))
	})
	return cache_Data_FoldableWithIndex_foldr
}

var cache_Data_FoldableWithIndex_foldl gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl sync.Once
func Get_Data_FoldableWithIndex_foldl() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl.Do(func() {
		cache_Data_FoldableWithIndex_foldl = Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()))
	})
	return cache_Data_FoldableWithIndex_foldl
}

var cache_Data_FoldableWithIndex_foldMap gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap sync.Once
func Get_Data_FoldableWithIndex_foldMap() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap.Do(func() {
		cache_Data_FoldableWithIndex_foldMap = Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()))
	})
	return cache_Data_FoldableWithIndex_foldMap
}

var cache_Data_FoldableWithIndex_foldr1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr1 sync.Once
func Get_Data_FoldableWithIndex_foldr1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr1.Do(func() {
		cache_Data_FoldableWithIndex_foldr1 = Call_Data_Foldable_foldr(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe()))))
	})
	return cache_Data_FoldableWithIndex_foldr1
}

var cache_Data_FoldableWithIndex_foldl1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl1 sync.Once
func Get_Data_FoldableWithIndex_foldl1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl1.Do(func() {
		cache_Data_FoldableWithIndex_foldl1 = Call_Data_Foldable_foldl(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe()))))
	})
	return cache_Data_FoldableWithIndex_foldl1
}

var cache_Data_FoldableWithIndex_foldMap1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap1 sync.Once
func Get_Data_FoldableWithIndex_foldMap1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap1.Do(func() {
		cache_Data_FoldableWithIndex_foldMap1 = Call_Data_Foldable_foldMap(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe()))))
	})
	return cache_Data_FoldableWithIndex_foldMap1
}

var cache_Data_FoldableWithIndex_foldr2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr2 sync.Once
func Get_Data_FoldableWithIndex_foldr2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr2.Do(func() {
		cache_Data_FoldableWithIndex_foldr2 = Call_Data_Foldable_foldr(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast()))))
	})
	return cache_Data_FoldableWithIndex_foldr2
}

var cache_Data_FoldableWithIndex_foldl2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl2 sync.Once
func Get_Data_FoldableWithIndex_foldl2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl2.Do(func() {
		cache_Data_FoldableWithIndex_foldl2 = Call_Data_Foldable_foldl(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast()))))
	})
	return cache_Data_FoldableWithIndex_foldl2
}

var cache_Data_FoldableWithIndex_foldMap2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap2 sync.Once
func Get_Data_FoldableWithIndex_foldMap2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap2.Do(func() {
		cache_Data_FoldableWithIndex_foldMap2 = Call_Data_Foldable_foldMap(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast()))))
	})
	return cache_Data_FoldableWithIndex_foldMap2
}

var cache_Data_FoldableWithIndex_foldr3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr3 sync.Once
func Get_Data_FoldableWithIndex_foldr3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr3.Do(func() {
		cache_Data_FoldableWithIndex_foldr3 = Call_Data_Foldable_foldr(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst()))))
	})
	return cache_Data_FoldableWithIndex_foldr3
}

var cache_Data_FoldableWithIndex_foldl3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl3 sync.Once
func Get_Data_FoldableWithIndex_foldl3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl3.Do(func() {
		cache_Data_FoldableWithIndex_foldl3 = Call_Data_Foldable_foldl(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst()))))
	})
	return cache_Data_FoldableWithIndex_foldl3
}

var cache_Data_FoldableWithIndex_foldMap3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap3 sync.Once
func Get_Data_FoldableWithIndex_foldMap3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap3.Do(func() {
		cache_Data_FoldableWithIndex_foldMap3 = Call_Data_Foldable_foldMap(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst()))))
	})
	return cache_Data_FoldableWithIndex_foldMap3
}

var cache_Data_FoldableWithIndex_foldr4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr4 sync.Once
func Get_Data_FoldableWithIndex_foldr4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr4.Do(func() {
		cache_Data_FoldableWithIndex_foldr4 = Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()))
	})
	return cache_Data_FoldableWithIndex_foldr4
}

var cache_Data_FoldableWithIndex_foldl4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl4 sync.Once
func Get_Data_FoldableWithIndex_foldl4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl4.Do(func() {
		cache_Data_FoldableWithIndex_foldl4 = Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()))
	})
	return cache_Data_FoldableWithIndex_foldl4
}

var cache_Data_FoldableWithIndex_foldMap4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap4 sync.Once
func Get_Data_FoldableWithIndex_foldMap4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap4.Do(func() {
		cache_Data_FoldableWithIndex_foldMap4 = Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()))
	})
	return cache_Data_FoldableWithIndex_foldMap4
}

var cache_Data_FoldableWithIndex_foldr5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr5 sync.Once
func Get_Data_FoldableWithIndex_foldr5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr5.Do(func() {
		cache_Data_FoldableWithIndex_foldr5 = Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()))
	})
	return cache_Data_FoldableWithIndex_foldr5
}

var cache_Data_FoldableWithIndex_foldl5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl5 sync.Once
func Get_Data_FoldableWithIndex_foldl5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl5.Do(func() {
		cache_Data_FoldableWithIndex_foldl5 = Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()))
	})
	return cache_Data_FoldableWithIndex_foldl5
}

var cache_Data_FoldableWithIndex_foldMap5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap5 sync.Once
func Get_Data_FoldableWithIndex_foldMap5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap5.Do(func() {
		cache_Data_FoldableWithIndex_foldMap5 = Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()))
	})
	return cache_Data_FoldableWithIndex_foldMap5
}

var cache_Data_FoldableWithIndex_foldr6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr6 sync.Once
func Get_Data_FoldableWithIndex_foldr6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr6.Do(func() {
		cache_Data_FoldableWithIndex_foldr6 = Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()))
	})
	return cache_Data_FoldableWithIndex_foldr6
}

var cache_Data_FoldableWithIndex_foldl6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl6 sync.Once
func Get_Data_FoldableWithIndex_foldl6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl6.Do(func() {
		cache_Data_FoldableWithIndex_foldl6 = Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()))
	})
	return cache_Data_FoldableWithIndex_foldl6
}

var cache_Data_FoldableWithIndex_foldMap6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap6 sync.Once
func Get_Data_FoldableWithIndex_foldMap6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap6.Do(func() {
		cache_Data_FoldableWithIndex_foldMap6 = Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()))
	})
	return cache_Data_FoldableWithIndex_foldMap6
}

var cache_Data_FoldableWithIndex_foldr7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr7 sync.Once
func Get_Data_FoldableWithIndex_foldr7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr7.Do(func() {
		cache_Data_FoldableWithIndex_foldr7 = Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()))
	})
	return cache_Data_FoldableWithIndex_foldr7
}

var cache_Data_FoldableWithIndex_foldl7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl7 sync.Once
func Get_Data_FoldableWithIndex_foldl7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl7.Do(func() {
		cache_Data_FoldableWithIndex_foldl7 = Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()))
	})
	return cache_Data_FoldableWithIndex_foldl7
}

var cache_Data_FoldableWithIndex_foldMap7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap7 sync.Once
func Get_Data_FoldableWithIndex_foldMap7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap7.Do(func() {
		cache_Data_FoldableWithIndex_foldMap7 = Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()))
	})
	return cache_Data_FoldableWithIndex_foldMap7
}

var cache_Data_FoldableWithIndex_monoidDual gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidDual sync.Once
func Get_Data_FoldableWithIndex_monoidDual() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidDual.Do(func() {
		cache_Data_FoldableWithIndex_monoidDual = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Dual_monoidDual(Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))))}
	})
	return cache_Data_FoldableWithIndex_monoidDual
}

var cache_Data_FoldableWithIndex_monoidEndo gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidEndo sync.Once
func Get_Data_FoldableWithIndex_monoidEndo() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidEndo.Do(func() {
		cache_Data_FoldableWithIndex_monoidEndo = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}
	})
	return cache_Data_FoldableWithIndex_monoidEndo
}

var cache_Data_FoldableWithIndex_monoidEndo1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidEndo1 sync.Once
func Get_Data_FoldableWithIndex_monoidEndo1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidEndo1.Do(func() {
		cache_Data_FoldableWithIndex_monoidEndo1 = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}
	})
	return cache_Data_FoldableWithIndex_monoidEndo1
}

var cache_Data_FoldableWithIndex_unwrap gopurs_runtime.Value
var once_Data_FoldableWithIndex_unwrap sync.Once
func Get_Data_FoldableWithIndex_unwrap() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_unwrap.Do(func() {
		cache_Data_FoldableWithIndex_unwrap = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_FoldableWithIndex_unwrap
}

var cache_Data_FoldableWithIndex_unwrap1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_unwrap1 sync.Once
func Get_Data_FoldableWithIndex_unwrap1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_unwrap1.Do(func() {
		cache_Data_FoldableWithIndex_unwrap1 = Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{}))
	})
	return cache_Data_FoldableWithIndex_unwrap1
}

var cache_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict gopurs_runtime.Value
var once_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict sync.Once
func Get_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict.Do(func() {
		cache_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Call_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict(func() struct{
	Foldable0 gopurs_runtime.Value
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex gopurs_runtime.Value
	foldrWithIndex gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Foldable0 gopurs_runtime.Value
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex gopurs_runtime.Value
	foldrWithIndex gopurs_runtime.Value
}{}
					clone.Foldable0 = gopurs_runtime.RecordGet(orig, "Foldable0")
					clone.foldMapWithIndex = gopurs_runtime.RecordGet(orig, "foldMapWithIndex")
					clone.foldlWithIndex = gopurs_runtime.RecordGet(orig, "foldlWithIndex")
					clone.foldrWithIndex = gopurs_runtime.RecordGet(orig, "foldrWithIndex")
					return clone
				}()))}
})
	})
	return cache_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict
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
		cache_Data_FoldableWithIndex_foldableWithIndexTuple = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_3374046885_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_4173511203_1680800814(Rebox_Data_FoldableWithIndex_1680800814_4173511203(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableTuple()))))}
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), z_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
})})))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexTuple
}

var cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexMultiplicative sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexMultiplicative() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexMultiplicative.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope73)] (TypeVar m$scope74)), (TypeVar a$scope73)] (TypeVar m$scope74))
foldMap8_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative())), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative
}

var cache_Data_FoldableWithIndex_foldableWithIndexMaybe gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexMaybe sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexMaybe() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexMaybe.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexMaybe = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_2286084809_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe()))))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope87)] (TypeVar m$scope88)), (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope87)])] (TypeVar m$scope88))
foldMap8_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_foldMap(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe())))), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldl(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe())))), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldr(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMaybe())))), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
})})))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexMaybe
}

var cache_Data_FoldableWithIndex_foldableWithIndexLast gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexLast sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexLast() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexLast.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexLast = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_2286084809_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast()))))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope101)] (TypeVar m$scope102)), (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope101)])] (TypeVar m$scope102))
foldMap8_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_foldMap(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast())))), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldl(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast())))), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldr(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableLast())))), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
})})))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexLast
}

var cache_Data_FoldableWithIndex_foldableWithIndexIdentity gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexIdentity sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexIdentity() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexIdentity.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexIdentity = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableIdentity()))}
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, Get_Data_Unit_unit(), v_2, z_1)
})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexIdentity
}

var cache_Data_FoldableWithIndex_foldableWithIndexFirst gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexFirst sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexFirst() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexFirst.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexFirst = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_2286084809_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst()))))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope129)] (TypeVar m$scope130)), (ADT ["Data","Maybe","Maybe"] [(TypeVar a$scope129)])] (TypeVar m$scope130))
foldMap8_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_foldMap(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst())))), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldl(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst())))), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldr(Rebox_Data_FoldableWithIndex_1146820559_1680800814(Rebox_Data_FoldableWithIndex_1680800814_1146820559(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableFirst())))), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
})})))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexFirst
}

var cache_Data_FoldableWithIndex_foldableWithIndexEither gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexEither sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexEither() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexEither.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexEither = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableEither()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope146)
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
})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexEither
}

var cache_Data_FoldableWithIndex_foldableWithIndexDual gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexDual sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexDual() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexDual.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexDual = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope159)] (TypeVar m$scope160)), (TypeVar a$scope159)] (TypeVar m$scope160))
foldMap8_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual())), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexDual
}

var cache_Data_FoldableWithIndex_foldableWithIndexDisj gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexDisj sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexDisj() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexDisj.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexDisj = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope173)] (TypeVar m$scope174)), (TypeVar a$scope173)] (TypeVar m$scope174))
foldMap8_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj())), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexDisj
}

var cache_Data_FoldableWithIndex_foldableWithIndexConst gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexConst sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexConst() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexConst.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexConst = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConst()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_0 shape=App(Var) bindingType=(TypeVar m$scope190)
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
	return cache_Data_FoldableWithIndex_foldableWithIndexConst
}

var cache_Data_FoldableWithIndex_foldableWithIndexConj gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexConj sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexConj() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexConj.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexConj = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope203)] (TypeVar m$scope204)), (TypeVar a$scope203)] (TypeVar m$scope204))
foldMap8_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj())), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexConj
}

var cache_Data_FoldableWithIndex_foldableWithIndexAdditive gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexAdditive sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexAdditive() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexAdditive.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexAdditive = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap8_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope217)] (TypeVar m$scope218)), (TypeVar a$scope217)] (TypeVar m$scope218))
foldMap8_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive())), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
})}))}
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
		cache_Data_FoldableWithIndex_foldableWithIndexArray = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_2491554675_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FoldableWithIndex_foldMapWithIndexDefaultR(Rebox_Data_FoldableWithIndex_2491554675_3725484264(Rebox_Data_FoldableWithIndex_3725484264_2491554675(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexArray()))), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0))
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Get_Data_Foldable_foldlArray(), gopurs_runtime.Func2(func(y_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal), y_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
}), z_1), gopurs_runtime.Apply(Get_Data_FunctorWithIndex_mapWithIndexArray(), Get_Data_Tuple_Tuple()))
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply2(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_0 shape=Other bindingType=Any
__local_var_3_0 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_3_0
// TAST (Let): __local_var_4_1 shape=Other bindingType=Any
__local_var_4_1 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_4_1
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_3_0.IntVal), __local_var_4_1, y_5)
})
}), z_1), gopurs_runtime.Apply(Get_Data_FunctorWithIndex_mapWithIndexArray(), Get_Data_Tuple_Tuple()))
})})))}
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


func Call_Data_FoldableWithIndex_FoldableWithIndex_dollar_Dict(x_0_loop struct{
	Foldable0 gopurs_runtime.Value
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex gopurs_runtime.Value
	foldrWithIndex gopurs_runtime.Value
}) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	Foldable0 gopurs_runtime.Value
	foldMapWithIndex gopurs_runtime.Value
	foldlWithIndex gopurs_runtime.Value
	foldrWithIndex gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", orig.Foldable0, orig.foldMapWithIndex, orig.foldlWithIndex, orig.foldrWithIndex)
				}())
}

func Call_Data_FoldableWithIndex_foldrWithIndex(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_Data_FoldableWithIndex_traverseWithIndex_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): applySecond_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m$scope13) [(TypeVar b$scope11)]), (TypeApp (TypeVar m$scope13) [Unit])] (TypeApp (TypeVar m$scope13) [Unit]))
applySecond_1_0 := Call_Control_Apply_applySecond(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApplicative_0.V0, gopurs_runtime.Value{})))
_ = applySecond_1_0
return gopurs_runtime.Func2(func(dictFoldableWithIndex_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_2, "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), applySecond_1_0, gopurs_runtime.Apply(f_3, i_4))
}), gopurs_runtime.Apply(dictApplicative_0.V1, Get_Data_Unit_unit()))
})
}

func Call_Data_FoldableWithIndex_forWithIndex_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): traverseWithIndex_1_1_0 shape=App(Var) bindingType=Any
traverseWithIndex_1_1_0 := Call_Data_FoldableWithIndex_traverseWithIndex_(dictApplicative_0)
_ = traverseWithIndex_1_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
__local_var_3_1 := gopurs_runtime.Apply(traverseWithIndex_1_1_0, dictFoldableWithIndex_2)
_ = __local_var_3_1
return gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_3_1, a_5, b_4)
})
})
}

func Call_Data_FoldableWithIndex_foldrDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(dictFoldableWithIndex_0.V3, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
}

func Call_Data_FoldableWithIndex_foldlWithIndex(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_Data_FoldableWithIndex_foldlDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(dictFoldableWithIndex_0.V2, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
}

func Call_Data_FoldableWithIndex_foldWithIndexM(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonad_1_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonad_1 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_1_loop
_ = dictMonad_1
// TAST (Let): Bind1_2_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m$scope226)])
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_1.V1, gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): Applicative0_3_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m$scope226)])
Applicative0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonad_1.V0, gopurs_runtime.Value{}))
_ = Applicative0_3_1
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, a0_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V2, gopurs_runtime.Func3(func(i_6 gopurs_runtime.Value, ma_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_2 shape=App(Other) bindingType=(Func [(TypeVar a$scope227), (TypeVar b$scope228)] (TypeApp (TypeVar m$scope226) [(TypeVar a$scope227)]))
__local_var_9_2 := gopurs_runtime.Apply(f_4, i_6)
_ = __local_var_9_2
return gopurs_runtime.Apply2(Bind1_2_0.V1, ma_7, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_9_2, a_10, b_8)
}))
}), gopurs_runtime.Apply(Applicative0_3_1.V1, a0_5))
})
}

func Call_Data_FoldableWithIndex_foldMapWithIndexDefaultR(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope236)])
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 shape=App(Var) bindingType=(TypeVar m$scope236)
mempty_3_1 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)})
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V3, gopurs_runtime.Func3(func(i_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value, acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_2_0.V0, gopurs_runtime.Apply2(f_4, i_5, x_6), acc_7)
}), mempty_3_1)
})
}

func Call_Data_FoldableWithIndex_foldMapWithIndexDefaultL(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope264)])
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 shape=App(Var) bindingType=(TypeVar m$scope264)
mempty_3_1 := Call_Data_Monoid_mempty(gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)})
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V2, gopurs_runtime.Func3(func(i_5 gopurs_runtime.Value, acc_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_2_0.V0, acc_6, gopurs_runtime.Apply2(f_4, i_5, x_7))
}), mempty_3_1)
})
}

func Call_Data_FoldableWithIndex_foldMapWithIndex(dict_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_Data_FoldableWithIndex_foldableWithIndexApp(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): foldableApp_1_0 shape=App(Var) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f$scope272)])])
foldableApp_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Call_Data_Foldable_foldableApp(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})))
_ = foldableApp_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableApp_1_0)}
}), gopurs_runtime.Func3(func(dictMonoid_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, f_3, v_4)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), f_2, z_3, v_4)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, z_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), f_2, z_3, v_4)
})}))}
}

func Call_Data_FoldableWithIndex_foldableWithIndexCompose(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): foldableCompose_1_0 shape=App(Var) bindingType=Any
foldableCompose_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldableCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableCompose_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldlWithIndex1_3_1 shape=App(Var) bindingType=(Func [(Func [(TypeVar b$scope293), (TypeVar b$scope306), (TypeVar a$scope305)] (TypeVar b$scope306)), (TypeVar b$scope306), (TypeApp (TypeVar g$scope294) [(TypeVar a$scope305)])] (TypeVar b$scope306))
foldlWithIndex1_3_1 := Call_Data_FoldableWithIndex_foldlWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex1_2))
_ = foldlWithIndex1_3_1
// TAST (Let): foldMapWithIndex1_4_2 shape=App(Var) bindingType=Any
foldMapWithIndex1_4_2 := Call_Data_FoldableWithIndex_foldMapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex1_2))
_ = foldMapWithIndex1_4_2
// TAST (Let): foldableCompose1_5_3 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope292), (TypeVar g$scope294)])])
foldableCompose1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableCompose_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{})))
_ = foldableCompose1_5_3
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_139552293_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableCompose1_5_3)}
}), gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMapWithIndex2_7_4 shape=App(Other) bindingType=(Func [(Func [(TypeVar b$scope293), (TypeVar a$scope308)] (TypeVar m$scope309)), (TypeApp (TypeVar g$scope294) [(TypeVar a$scope308)])] (TypeVar m$scope309))
foldMapWithIndex2_7_4 := gopurs_runtime.Apply(foldMapWithIndex1_4_2, dictMonoid_6)
_ = foldMapWithIndex2_7_4
return gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_6))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), foldMapWithIndex2_7_4, gopurs_runtime.Func2(func(a_10 gopurs_runtime.Value, b_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_10, b_11}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})
})), v_9)
})
}), gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, i_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), foldlWithIndex1_3_1, gopurs_runtime.Func2(func(a_9 gopurs_runtime.Value, b_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_9, b_10}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})
})), i_7, v_8)
}), gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, i_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_5 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeVar b))
__local_var_10_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_9, b_10}
				return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})}
			}()))})
}))
_ = __local_var_10_5
return gopurs_runtime.Func2(func(b_11 gopurs_runtime.Value, a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_10_5, a_12, b_11)
})
}), i_7, v_8)
})})))}
})
}

func Call_Data_FoldableWithIndex_foldableWithIndexCoproduct(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): foldableCoproduct_1_0 shape=App(Var) bindingType=Any
foldableCoproduct_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableCoproduct_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldableCoproduct1_3_1 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope316), (TypeVar g$scope318)])])
foldableCoproduct1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{})))
_ = foldableCoproduct1_3_1
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableCoproduct1_3_1)}
}), gopurs_runtime.Func2(func(dictMonoid_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_5, Get_Data_Either_Left())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_5, Get_Data_Either_Right())))
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Left()), z_5), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Right()), z_5))
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Functor_Coproduct_coproduct(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Left()), z_5), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Right()), z_5))
})}))}
})
}

func Call_Data_FoldableWithIndex_foldableWithIndexProduct(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): foldableProduct_1_0 shape=App(Var) bindingType=Any
foldableProduct_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableProduct_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldableProduct1_3_1 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f$scope359), (TypeVar g$scope361)])])
foldableProduct1_3_1 := Rebox_Data_FoldableWithIndex_1680800814_4173511203(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{}))))
_ = foldableProduct1_3_1
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_3374046885_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_4173511203_1680800814(foldableProduct1_3_1))}
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_2 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m$scope387)])
Semigroup0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_2
return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_5_2.V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_6, Get_Data_Either_Left()), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_6, Get_Data_Either_Right()), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Right()), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Left()), z_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Left()), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), f_4, Get_Data_Either_Right()), z_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
})})))}
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
return gopurs_runtime.Apply2(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply3(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Dual_monoidDual(Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}))))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_0 shape=App(Other) bindingType=(Func [(TypeVar b$scope401), (TypeVar a$scope400)] (TypeVar b$scope401))
__local_var_5_0 := gopurs_runtime.Apply(c_1, i_4)
_ = __local_var_5_0
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Monoid_Dual_Dual(), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Monoid_Endo_Endo(), gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_5_0, a_7, b_6)
})))
}), xs_3)), u_2)
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
return gopurs_runtime.Apply2(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply3(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Monoid_Endo_Endo(), gopurs_runtime.Apply(c_1, i_4))
}), xs_3), u_2)
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
return gopurs_runtime.Apply2(Call_Safe_Coerce_coerce(gopurs_runtime.Value{}), gopurs_runtime.Apply3(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Endo_monoidEndo(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})))}, gopurs_runtime.Func3(func(i_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemigroup_1.V0, d_2, gopurs_runtime.Apply2(dictSemigroup_1.V0, gopurs_runtime.Apply2(t_3, i_5, a_6), m_7))
}), f_4), d_2)
}

func Call_Data_FoldableWithIndex_foldMapDefault(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value], f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return f_2
}))
}

func Call_Data_FoldableWithIndex_findWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V2, gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
}] = Rebox_Data_FoldableWithIndex_3094389156_516634048(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3))
_ = __t_tag_0
if ((__t_tag_0 == nil)) && ((gopurs_runtime.Apply2(p_1, v_2, v2_4).IntVal) != (0)) {
__t1 = Rebox_Data_FoldableWithIndex_3094389156_516634048(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("index", "value", v_2, v2_4), true}
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
__t1 = Rebox_Data_FoldableWithIndex_3094389156_516634048(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3))
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_516634048_3094389156(__t1))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_516634048_3094389156(Rebox_Data_FoldableWithIndex_3094389156_516634048(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))})
}

func Call_Data_FoldableWithIndex_findMapWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(dictFoldableWithIndex_0.V2, gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
_ = __t_tag_0
if (__t_tag_0 == nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_1, v_2, v2_4))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
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

func Call_Data_FoldableWithIndex_anyWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): monoidDisj_2_0 shape=App(Var) bindingType=(ADT ["Data","Monoid","Monoid"] [(TypeVar b$scope462)])
monoidDisj_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Disj_monoidDisj(gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(dictHeytingAlgebra_1)}))
_ = monoidDisj_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply2(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(monoidDisj_2_0)}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Monoid_Disj_Disj(), gopurs_runtime.Apply(t_3, i_4))
})))
})
}

func Call_Data_FoldableWithIndex_allWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): monoidConj_2_0 shape=App(Var) bindingType=(ADT ["Data","Monoid","Monoid"] [(TypeVar b$scope473)])
monoidConj_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Call_Data_Monoid_Conj_monoidConj(gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(dictHeytingAlgebra_1)}))
_ = monoidConj_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_Newtype_unwrap(gopurs_runtime.CoerceToStruct[Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{})), gopurs_runtime.Apply2(dictFoldableWithIndex_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(monoidConj_2_0)}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Monoid_Conj_Conj(), gopurs_runtime.Apply(t_3, i_4))
})))
})
}

func Rebox_Data_FoldableWithIndex_1146820559_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_FoldableWithIndex_139552293_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_FoldableWithIndex_1680800814_1146820559(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_FoldableWithIndex_1680800814_4173511203(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_FoldableWithIndex_2286084809_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_FoldableWithIndex_2491554675_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_FoldableWithIndex_3094389156_516634048(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
}] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
}]{}
		out.V0 = func() struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := in.V0
					_ = orig
					clone := struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.index = gopurs_runtime.RecordGet(orig, "index")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
	return out
}

func Rebox_Data_FoldableWithIndex_3374046885_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_FoldableWithIndex_3725484264_2491554675(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_FoldableWithIndex_4173511203_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_FoldableWithIndex_516634048_3094389156(in *Constructor_Data_Maybe_Just[struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict2("index", "value", orig.index, orig.value)
				}()
	return out
}



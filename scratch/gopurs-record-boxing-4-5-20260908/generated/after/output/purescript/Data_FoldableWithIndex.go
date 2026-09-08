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
		cache_Data_FoldableWithIndex_foldr = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()).V2)
	})
	return cache_Data_FoldableWithIndex_foldr
}

var cache_Data_FoldableWithIndex_foldl gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl sync.Once
func Get_Data_FoldableWithIndex_foldl() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl.Do(func() {
		cache_Data_FoldableWithIndex_foldl = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()).V1)
	})
	return cache_Data_FoldableWithIndex_foldl
}

var cache_Data_FoldableWithIndex_foldMap gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap sync.Once
func Get_Data_FoldableWithIndex_foldMap() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap.Do(func() {
		cache_Data_FoldableWithIndex_foldMap = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()).V0)
	})
	return cache_Data_FoldableWithIndex_foldMap
}

var cache_Data_FoldableWithIndex_foldr1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr1 sync.Once
func Get_Data_FoldableWithIndex_foldr1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr1.Do(func() {
		cache_Data_FoldableWithIndex_foldr1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableMaybe()).V2)
	})
	return cache_Data_FoldableWithIndex_foldr1
}

var cache_Data_FoldableWithIndex_foldl1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl1 sync.Once
func Get_Data_FoldableWithIndex_foldl1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl1.Do(func() {
		cache_Data_FoldableWithIndex_foldl1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableMaybe()).V1)
	})
	return cache_Data_FoldableWithIndex_foldl1
}

var cache_Data_FoldableWithIndex_foldMap1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap1 sync.Once
func Get_Data_FoldableWithIndex_foldMap1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap1.Do(func() {
		cache_Data_FoldableWithIndex_foldMap1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableMaybe()).V0)
	})
	return cache_Data_FoldableWithIndex_foldMap1
}

var cache_Data_FoldableWithIndex_foldr2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr2 sync.Once
func Get_Data_FoldableWithIndex_foldr2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr2.Do(func() {
		cache_Data_FoldableWithIndex_foldr2 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableLast()).V2)
	})
	return cache_Data_FoldableWithIndex_foldr2
}

var cache_Data_FoldableWithIndex_foldl2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl2 sync.Once
func Get_Data_FoldableWithIndex_foldl2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl2.Do(func() {
		cache_Data_FoldableWithIndex_foldl2 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableLast()).V1)
	})
	return cache_Data_FoldableWithIndex_foldl2
}

var cache_Data_FoldableWithIndex_foldMap2 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap2 sync.Once
func Get_Data_FoldableWithIndex_foldMap2() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap2.Do(func() {
		cache_Data_FoldableWithIndex_foldMap2 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableLast()).V0)
	})
	return cache_Data_FoldableWithIndex_foldMap2
}

var cache_Data_FoldableWithIndex_foldr3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr3 sync.Once
func Get_Data_FoldableWithIndex_foldr3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr3.Do(func() {
		cache_Data_FoldableWithIndex_foldr3 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableFirst()).V2)
	})
	return cache_Data_FoldableWithIndex_foldr3
}

var cache_Data_FoldableWithIndex_foldl3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl3 sync.Once
func Get_Data_FoldableWithIndex_foldl3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl3.Do(func() {
		cache_Data_FoldableWithIndex_foldl3 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableFirst()).V1)
	})
	return cache_Data_FoldableWithIndex_foldl3
}

var cache_Data_FoldableWithIndex_foldMap3 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap3 sync.Once
func Get_Data_FoldableWithIndex_foldMap3() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap3.Do(func() {
		cache_Data_FoldableWithIndex_foldMap3 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableFirst()).V0)
	})
	return cache_Data_FoldableWithIndex_foldMap3
}

var cache_Data_FoldableWithIndex_foldr4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr4 sync.Once
func Get_Data_FoldableWithIndex_foldr4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr4.Do(func() {
		cache_Data_FoldableWithIndex_foldr4 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()).V2)
	})
	return cache_Data_FoldableWithIndex_foldr4
}

var cache_Data_FoldableWithIndex_foldl4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl4 sync.Once
func Get_Data_FoldableWithIndex_foldl4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl4.Do(func() {
		cache_Data_FoldableWithIndex_foldl4 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()).V1)
	})
	return cache_Data_FoldableWithIndex_foldl4
}

var cache_Data_FoldableWithIndex_foldMap4 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap4 sync.Once
func Get_Data_FoldableWithIndex_foldMap4() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap4.Do(func() {
		cache_Data_FoldableWithIndex_foldMap4 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()).V0)
	})
	return cache_Data_FoldableWithIndex_foldMap4
}

var cache_Data_FoldableWithIndex_foldr5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr5 sync.Once
func Get_Data_FoldableWithIndex_foldr5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr5.Do(func() {
		cache_Data_FoldableWithIndex_foldr5 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()).V2)
	})
	return cache_Data_FoldableWithIndex_foldr5
}

var cache_Data_FoldableWithIndex_foldl5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl5 sync.Once
func Get_Data_FoldableWithIndex_foldl5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl5.Do(func() {
		cache_Data_FoldableWithIndex_foldl5 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()).V1)
	})
	return cache_Data_FoldableWithIndex_foldl5
}

var cache_Data_FoldableWithIndex_foldMap5 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap5 sync.Once
func Get_Data_FoldableWithIndex_foldMap5() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap5.Do(func() {
		cache_Data_FoldableWithIndex_foldMap5 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()).V0)
	})
	return cache_Data_FoldableWithIndex_foldMap5
}

var cache_Data_FoldableWithIndex_foldr6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr6 sync.Once
func Get_Data_FoldableWithIndex_foldr6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr6.Do(func() {
		cache_Data_FoldableWithIndex_foldr6 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()).V2)
	})
	return cache_Data_FoldableWithIndex_foldr6
}

var cache_Data_FoldableWithIndex_foldl6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl6 sync.Once
func Get_Data_FoldableWithIndex_foldl6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl6.Do(func() {
		cache_Data_FoldableWithIndex_foldl6 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()).V1)
	})
	return cache_Data_FoldableWithIndex_foldl6
}

var cache_Data_FoldableWithIndex_foldMap6 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap6 sync.Once
func Get_Data_FoldableWithIndex_foldMap6() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap6.Do(func() {
		cache_Data_FoldableWithIndex_foldMap6 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()).V0)
	})
	return cache_Data_FoldableWithIndex_foldMap6
}

var cache_Data_FoldableWithIndex_foldr7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldr7 sync.Once
func Get_Data_FoldableWithIndex_foldr7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldr7.Do(func() {
		cache_Data_FoldableWithIndex_foldr7 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()).V2)
	})
	return cache_Data_FoldableWithIndex_foldr7
}

var cache_Data_FoldableWithIndex_foldl7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldl7 sync.Once
func Get_Data_FoldableWithIndex_foldl7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldl7.Do(func() {
		cache_Data_FoldableWithIndex_foldl7 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()).V1)
	})
	return cache_Data_FoldableWithIndex_foldl7
}

var cache_Data_FoldableWithIndex_foldMap7 gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldMap7 sync.Once
func Get_Data_FoldableWithIndex_foldMap7() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldMap7.Do(func() {
		cache_Data_FoldableWithIndex_foldMap7 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()).V0)
	})
	return cache_Data_FoldableWithIndex_foldMap7
}

var cache_Data_FoldableWithIndex_monoidDual gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidDual sync.Once
func Get_Data_FoldableWithIndex_monoidDual() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidDual.Do(func() {
		cache_Data_FoldableWithIndex_monoidDual = func() gopurs_runtime.Value {
// TAST (Let): semigroupEndo1_0_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(TypeVar c), (TypeVar a)])])
semigroupEndo1_0_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(v1_1, x_2))
})})
_ = semigroupEndo1_0_0
// TAST (Let): __local_var_1_1 shape=LitRecord bindingType=(ADT ["Data","Monoid","Monoid"] [(TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(TypeVar c), (TypeVar a)])])
__local_var_1_1 := (&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupEndo1_0_0)}
}), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_1.V0), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): semigroupDual1_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupDual1_2_2 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "append"), v1_4, v_3)
})})
_ = semigroupDual1_2_2
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupDual1_2_2)}
}), gopurs_runtime.Box(__local_var_1_1.V1)}))}
}()
	})
	return cache_Data_FoldableWithIndex_monoidDual
}

var cache_Data_FoldableWithIndex_monoidEndo gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidEndo sync.Once
func Get_Data_FoldableWithIndex_monoidEndo() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidEndo.Do(func() {
		cache_Data_FoldableWithIndex_monoidEndo = func() gopurs_runtime.Value {
// TAST (Let): semigroupEndo1_0_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(TypeVar c), (TypeVar a)])])
semigroupEndo1_0_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(v1_1, x_2))
})})
_ = semigroupEndo1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupEndo1_0_0)}
}), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)}))}
}()
	})
	return cache_Data_FoldableWithIndex_monoidEndo
}

var cache_Data_FoldableWithIndex_monoidEndo1 gopurs_runtime.Value
var once_Data_FoldableWithIndex_monoidEndo1 sync.Once
func Get_Data_FoldableWithIndex_monoidEndo1() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_monoidEndo1.Do(func() {
		cache_Data_FoldableWithIndex_monoidEndo1 = func() gopurs_runtime.Value {
// TAST (Let): semigroupEndo1_0_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(TypeVar c), (TypeVar a)])])
semigroupEndo1_0_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(v1_1, x_2))
})})
_ = semigroupEndo1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupEndo1_0_0)}
}), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)}))}
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
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), v_2)
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeVar a)] (TypeVar b))
__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_0
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, z_2, v_3)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar a), (TypeVar b)] (TypeVar b))
__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_1
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_1, v_3, z_2)
})
})}))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexMultiplicative
}

var cache_Data_FoldableWithIndex_foldableWithIndexMaybe gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexMaybe sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexMaybe() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexMaybe.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexMaybe = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_2286084809_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_1146820559_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableMaybe())))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): mempty_1_1 shape=Other bindingType=(TypeVar m)
mempty_1_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_1
// TAST (Let): foldMap8_1_0 shape=Let(Abs(Abs(Branch(Other, App(Other), def=Other)))) bindingType=(Func [(Func [(TypeVar a)] (TypeVar m)), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])] (TypeVar m))
foldMap8_1_0 := gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
if (__t_tag_2 == nil) {
__t4 = mempty_1_1
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
if (__t_tag_3 != nil) {
__t4 = gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_5 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeVar a)] (TypeVar b))
__local_var_1_5 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_5
return gopurs_runtime.Func2(func(v1_2 gopurs_runtime.Value, v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3)
if (__t_tag_6 == nil) {
__t8 = v1_2
goto end_branch_8
} else {

}
}
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3)
if (__t_tag_7 != nil) {
__t8 = gopurs_runtime.Apply2(__local_var_1_5, v1_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_9 shape=App(Other) bindingType=(Func [(TypeVar a), (TypeVar b)] (TypeVar b))
__local_var_1_9 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_9
return gopurs_runtime.Func2(func(v1_2 gopurs_runtime.Value, v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
var __t_tag_10 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3)
if (__t_tag_10 == nil) {
__t12 = v1_2
goto end_branch_12
} else {

}
}
{
var __t_tag_11 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v2_3)
if (__t_tag_11 != nil) {
__t12 = gopurs_runtime.Apply2(__local_var_1_9, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v2_3.UnsafePtr).V0, v1_2)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
})})))}
	})
	return cache_Data_FoldableWithIndex_foldableWithIndexMaybe
}

var cache_Data_FoldableWithIndex_foldableWithIndexLast gopurs_runtime.Value
var once_Data_FoldableWithIndex_foldableWithIndexLast sync.Once
func Get_Data_FoldableWithIndex_foldableWithIndexLast() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_foldableWithIndexLast.Do(func() {
		cache_Data_FoldableWithIndex_foldableWithIndexLast = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_2286084809_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_1146820559_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableLast())))}
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeVar a)] (TypeVar b))
__local_var_1_3 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_3
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_4 == nil) {
__t6 = z_2
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.Apply2(__local_var_1_3, z_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_7 shape=App(Other) bindingType=(Func [(TypeVar a), (TypeVar b)] (TypeVar b))
__local_var_1_7 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_7
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_8 == nil) {
__t10 = z_2
goto end_branch_10
} else {

}
}
{
var __t_tag_9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_9 != nil) {
__t10 = gopurs_runtime.Apply2(__local_var_1_7, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0, z_2)
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
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_1146820559_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableFirst())))}
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_0 == nil) {
__t2 = gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
goto end_branch_2
} else {

}
}
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_1 != nil) {
__t2 = gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_3 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeVar a)] (TypeVar b))
__local_var_1_3 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_3
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_4 == nil) {
__t6 = z_2
goto end_branch_6
} else {

}
}
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_5 != nil) {
__t6 = gopurs_runtime.Apply2(__local_var_1_3, z_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_7 shape=App(Other) bindingType=(Func [(TypeVar a), (TypeVar b)] (TypeVar b))
__local_var_1_7 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_7
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_8 == nil) {
__t10 = z_2
goto end_branch_10
} else {

}
}
{
var __t_tag_9 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_9 != nil) {
__t10 = gopurs_runtime.Apply2(__local_var_1_7, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0, z_2)
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
// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar m)
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
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
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), v_2)
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeVar a)] (TypeVar b))
__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_0
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, z_2, v_3)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar a), (TypeVar b)] (TypeVar b))
__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_1
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_1, v_3, z_2)
})
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
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), v_2)
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeVar a)] (TypeVar b))
__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_0
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, z_2, v_3)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar a), (TypeVar b)] (TypeVar b))
__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_1
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_1, v_3, z_2)
})
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
// TAST (Let): mempty_1_0 shape=Other bindingType=(TypeVar m)
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
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
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), v_2)
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeVar a)] (TypeVar b))
__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_0
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, z_2, v_3)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar a), (TypeVar b)] (TypeVar b))
__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_1
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_1, v_3, z_2)
})
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
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, Get_Data_Unit_unit(), v_2)
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeVar b), (TypeVar a)] (TypeVar b))
__local_var_1_0 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_0
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, z_2, v_3)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=(Func [(TypeVar a), (TypeVar b)] (TypeVar b))
__local_var_1_1 := gopurs_runtime.Apply(f_0, Get_Data_Unit_unit())
_ = __local_var_1_1
return gopurs_runtime.Func2(func(z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_1, v_3, z_2)
})
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
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexArray()).V3), gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value, acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply2(f_3, i_4, x_5), acc_6)
}), mempty_2_1)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 shape=App(Var) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply2(Get_Data_Foldable_foldlArray(), gopurs_runtime.Func2(func(y_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal), y_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
}), z_1)
_ = __local_var_2_2
// TAST (Let): __local_var_3_3 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (ADT ["Data","Tuple","Tuple"] [Int, (TypeVar a)])))
__local_var_3_3 := gopurs_runtime.Apply(Get_Data_FunctorWithIndex_mapWithIndexArray(), Get_Data_Tuple_Tuple())
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(__local_var_3_3, x_4))
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_4 shape=App(Var) bindingType=Any
__local_var_2_4 := gopurs_runtime.Apply2(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_5 shape=Other bindingType=Any
__local_var_3_5 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_3_5
// TAST (Let): __local_var_4_6 shape=Other bindingType=Any
__local_var_4_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_4_6
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_3_5.IntVal), __local_var_4_6, y_5)
})
}), z_1)
_ = __local_var_2_4
// TAST (Let): __local_var_3_7 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] (Array (ADT ["Data","Tuple","Tuple"] [Int, (TypeVar a)])))
__local_var_3_7 := gopurs_runtime.Apply(Get_Data_FunctorWithIndex_mapWithIndexArray(), Get_Data_Tuple_Tuple())
_ = __local_var_3_7
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(__local_var_3_7, x_4))
})
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
return func() gopurs_runtime.Value {
				_v := Call_Data_FoldableWithIndex_findWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), p_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_FoldableWithIndex_findWithIndex
}

var cache_Data_FoldableWithIndex_findMapWithIndex gopurs_runtime.Value
var once_Data_FoldableWithIndex_findMapWithIndex sync.Once
func Get_Data_FoldableWithIndex_findMapWithIndex() gopurs_runtime.Value {
	once_Data_FoldableWithIndex_findMapWithIndex.Do(func() {
		cache_Data_FoldableWithIndex_findMapWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_FoldableWithIndex_findMapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictFoldableWithIndex_0_box), f_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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
		c := (*Constructor_Data_FoldableWithIndex_FoldableWithIndex[any, any])(ptr)
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
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_FoldableWithIndex_traverseWithIndex_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): applySecond_1_0 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar m) [(TypeVar b)]), (TypeApp (TypeVar m) [Unit])] (TypeApp (TypeVar m) [Unit]))
applySecond_1_0 := gopurs_runtime.Apply(Get_Control_Apply_applySecond(), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}))
_ = applySecond_1_0
return gopurs_runtime.Func2(func(dictFoldableWithIndex_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_2, "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Other) bindingType=(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)]))
__local_var_5_1 := gopurs_runtime.Apply(f_3, i_4)
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_1_0, gopurs_runtime.Apply(__local_var_5_1, x_6))
})
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V1), Get_Data_Unit_unit()))
})
}

func Call_Data_FoldableWithIndex_forWithIndex_(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): traverseWithIndex_1__193435443_1_0 shape=App(Var) bindingType=Any
traverseWithIndex_1__193435443_1_0 := Call_Data_FoldableWithIndex_traverseWithIndex_(dictApplicative_0)
_ = traverseWithIndex_1__193435443_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
__local_var_3_1 := gopurs_runtime.Apply(traverseWithIndex_1__193435443_1_0, dictFoldableWithIndex_2)
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
// TAST (Let): Bind1_2_0 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_1.V1), gopurs_runtime.Value{}))
_ = Bind1_2_0
// TAST (Let): Applicative0_3_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_1.V0), gopurs_runtime.Value{}))
_ = Applicative0_3_1
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, a0_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V2), gopurs_runtime.Func3(func(i_6 gopurs_runtime.Value, ma_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_2 shape=App(Other) bindingType=Any
__local_var_9_2 := gopurs_runtime.Apply(f_4, i_6)
_ = __local_var_9_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), ma_7, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_9_2, a_10, b_8)
}))
}), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), a0_5))
})
}

func Call_Data_FoldableWithIndex_foldMapWithIndexDefaultR(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 shape=Other bindingType=(TypeVar m)
mempty_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V3), gopurs_runtime.Func3(func(i_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value, acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), gopurs_runtime.Apply2(f_4, i_5, x_6), acc_7)
}), mempty_3_1)
})
}

func Call_Data_FoldableWithIndex_foldMapWithIndexDefaultL(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 shape=Other bindingType=(TypeVar m)
mempty_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V2), gopurs_runtime.Func3(func(i_5 gopurs_runtime.Value, acc_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), acc_6, gopurs_runtime.Apply2(f_4, i_5, x_7))
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
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): foldableApp_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f)])])
foldableApp_1_0 := (&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, f_3, v_4)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldl"), f_2, i_3, v_4)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldr"), f_2, i_3, v_4)
})})
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
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): foldableCompose__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
foldableCompose__193435443_1_0 := gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_3))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_3))}, f_4), v_5)
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, i_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "foldl"), f_3), i_4, v_5)
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, i_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=App(Other) bindingType=Any
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_2, "foldr"), f_3)
_ = __local_var_6_2
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldr"), gopurs_runtime.Func2(func(b_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_6_2, a_8, b_7)
}), i_4, v_5)
})}))}
})
_ = foldableCompose__193435443_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldableCompose1_3_3 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
foldableCompose1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableCompose__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{})))
_ = foldableCompose1_3_3
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_139552293_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableCompose1_3_3)}
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMapWithIndex2_5_4 shape=App(Other) bindingType=(Func [(Func [(TypeVar b), (TypeVar a)] (TypeVar m)), (TypeApp (TypeVar g) [(TypeVar a)])] (TypeVar m))
foldMapWithIndex2_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), dictMonoid_4)
_ = foldMapWithIndex2_5_4
return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMapWithIndex2_5_4, gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{x_8, b_9}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
}), v_7)
})
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, i_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{x_7, b_8}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
}), i_5, v_6)
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, i_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 shape=App(Other) bindingType=Any
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_7, b_8}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
_ = __local_var_8_5
return gopurs_runtime.Func2(func(b_9 gopurs_runtime.Value, a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_8_5, a_10, b_9)
})
}), i_5, v_6)
})})))}
})
}

func Call_Data_FoldableWithIndex_foldableWithIndexCoproduct(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): foldableCoproduct__193435443_1_0 shape=App(Var) bindingType=Any
foldableCoproduct__193435443_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableCoproduct__193435443_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldableCoproduct1_3_1 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
foldableCoproduct1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableCoproduct__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{})))
_ = foldableCoproduct1_3_1
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableCoproduct1_3_1)}
}), gopurs_runtime.Func2(func(dictMonoid_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=App(Other) bindingType=Any
__local_var_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_6, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeVar m))
__local_var_7_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_7, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
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
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 shape=App(Other) bindingType=Any
__local_var_6_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_6, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}), z_5)
_ = __local_var_6_5
// TAST (Let): __local_var_7_6 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeVar b))
__local_var_7_6 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_7, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
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
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_8 shape=App(Other) bindingType=Any
__local_var_6_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_6, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}), z_5)
_ = __local_var_6_8
// TAST (Let): __local_var_7_9 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeVar b))
__local_var_7_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_7, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
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
})}))}
})
}

func Call_Data_FoldableWithIndex_foldableWithIndexProduct(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): foldableProduct__193435443_1_0 shape=App(Var) bindingType=Any
foldableProduct__193435443_1_0 := gopurs_runtime.Apply(Get_Data_Foldable_foldableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableProduct__193435443_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldableProduct1_3_1 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
foldableProduct1_3_1 := Rebox_Data_FoldableWithIndex_1680800814_4173511203(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableProduct__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{}))))
_ = foldableProduct1_3_1
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_3374046885_3725484264((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_4173511203_1680800814(foldableProduct1_3_1))}
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_2 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_2
return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_2.V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_8, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_8, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_7, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_7, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}), z_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_7, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_7, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}), z_5, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
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
// TAST (Let): semigroupEndo1_4_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(TypeVar c), (TypeVar a)])])
semigroupEndo1_4_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Apply(v1_5, x_6))
})})
_ = semigroupEndo1_4_0
// TAST (Let): __local_var_5_1 shape=LitRecord bindingType=(ADT ["Data","Monoid","Monoid"] [(TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(TypeVar c), (TypeVar a)])])
__local_var_5_1 := (&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupEndo1_4_0)}
}), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)})
_ = __local_var_5_1
// TAST (Let): __local_var_6_3 shape=App(Other) bindingType=Any
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_1.V0), gopurs_runtime.Value{})
_ = __local_var_6_3
// TAST (Let): semigroupDual1_6_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupDual1_6_2 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_3, "append"), v1_8, v_7)
})})
_ = semigroupDual1_6_2
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupDual1_6_2)}
}), gopurs_runtime.Box(__local_var_5_1.V1)}))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 shape=App(Other) bindingType=Any
__local_var_5_6 := gopurs_runtime.Apply(c_1, i_4)
_ = __local_var_5_6
// TAST (Let): __local_var_5_5 shape=Let(Abs(Abs(App(Other)))) bindingType=(Func [(TypeVar a), (TypeVar b)] (TypeVar b))
__local_var_5_5 := gopurs_runtime.Func2(func(b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_5_6, a_7, b_6)
})
_ = __local_var_5_5
// TAST (Let): __local_var_5_4 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeVar a)] (TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(ADT ["Prim","Function"] []), (TypeVar b)]))
__local_var_5_4 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_5, x_6)
})
_ = __local_var_5_4
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_4, x_6)
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
// TAST (Let): semigroupEndo1_4_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(TypeVar c), (TypeVar a)])])
semigroupEndo1_4_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Apply(v1_5, x_6))
})})
_ = semigroupEndo1_4_0
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupEndo1_4_0)}
}), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)}))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Other) bindingType=(Func [(TypeVar a), (TypeVar b)] (TypeVar b))
__local_var_5_1 := gopurs_runtime.Apply(c_1, i_4)
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_1, x_6)
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
// TAST (Let): semigroupEndo1_5_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeApp (TypeVar c) [(TypeVar a), (TypeVar a)]) [(TypeVar c), (TypeVar a)])])
semigroupEndo1_5_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_5, gopurs_runtime.Apply(v1_6, x_7))
})})
_ = semigroupEndo1_5_0
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupEndo1_5_0)}
}), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)}))}, gopurs_runtime.Func3(func(i_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), d_2, gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), gopurs_runtime.Apply2(t_3, i_5, a_6), m_7))
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

func Call_Data_FoldableWithIndex_findWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], p_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V2), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
}]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
}] = Rebox_Data_FoldableWithIndex_3094389156_516634048(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3))
if ((__t_tag_0 == nil)) && ((gopurs_runtime.Apply2(p_1, v_2, v2_4).IntVal) != (0)) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("index", "value", v_2, v2_4), true}
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
__t1 = Rebox_Data_FoldableWithIndex_3094389156_516634048(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3))
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_516634048_3094389156(__t1))}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_FoldableWithIndex_516634048_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	index gopurs_runtime.Value
	value gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_FoldableWithIndex_findMapWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V2), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
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
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))})
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_FoldableWithIndex_anyWithIndex(dictFoldableWithIndex_0_loop *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictHeytingAlgebra_1_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldableWithIndex_0 *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictHeytingAlgebra_1 *Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value] = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
// TAST (Let): semigroupDisj1_2_1 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupDisj1_2_1 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictHeytingAlgebra_1.V1), v_2, v1_3)
})})
_ = semigroupDisj1_2_1
// TAST (Let): monoidDisj_2_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Monoid","Monoid"] [(TypeVar b)])
monoidDisj_2_0 := (&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupDisj1_2_1)}
}), gopurs_runtime.Box(dictHeytingAlgebra_1.V2)})
_ = monoidDisj_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (TypeVar b))
__local_var_4_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(monoidDisj_2_0)}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 shape=App(Other) bindingType=(Func [(TypeVar a)] (TypeVar b))
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
// TAST (Let): semigroupConj1_2_1 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupConj1_2_1 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictHeytingAlgebra_1.V0), v_2, v1_3)
})})
_ = semigroupConj1_2_1
// TAST (Let): monoidConj_2_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Monoid","Monoid"] [(TypeVar b)])
monoidConj_2_0 := (&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupConj1_2_1)}
}), gopurs_runtime.Box(dictHeytingAlgebra_1.V5)})
_ = monoidConj_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (TypeVar b))
__local_var_4_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldableWithIndex_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(monoidConj_2_0)}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 shape=App(Other) bindingType=(Func [(TypeVar a)] (TypeVar b))
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



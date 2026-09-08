package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_TraversableWithIndex_traverse gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse sync.Once
func Get_Data_TraversableWithIndex_traverse() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse.Do(func() {
		cache_Data_TraversableWithIndex_traverse = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMultiplicative()).V3)
	})
	return cache_Data_TraversableWithIndex_traverse
}

var cache_Data_TraversableWithIndex_traverse1 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse1 sync.Once
func Get_Data_TraversableWithIndex_traverse1() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse1.Do(func() {
		cache_Data_TraversableWithIndex_traverse1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Traversable_traversableMaybe()).V3)
	})
	return cache_Data_TraversableWithIndex_traverse1
}

var cache_Data_TraversableWithIndex_traverse2 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse2 sync.Once
func Get_Data_TraversableWithIndex_traverse2() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse2.Do(func() {
		cache_Data_TraversableWithIndex_traverse2 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Traversable_traversableLast()).V3)
	})
	return cache_Data_TraversableWithIndex_traverse2
}

var cache_Data_TraversableWithIndex_traverse3 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse3 sync.Once
func Get_Data_TraversableWithIndex_traverse3() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse3.Do(func() {
		cache_Data_TraversableWithIndex_traverse3 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Traversable_traversableFirst()).V3)
	})
	return cache_Data_TraversableWithIndex_traverse3
}

var cache_Data_TraversableWithIndex_traverse4 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse4 sync.Once
func Get_Data_TraversableWithIndex_traverse4() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse4.Do(func() {
		cache_Data_TraversableWithIndex_traverse4 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableDual()).V3)
	})
	return cache_Data_TraversableWithIndex_traverse4
}

var cache_Data_TraversableWithIndex_traverse5 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse5 sync.Once
func Get_Data_TraversableWithIndex_traverse5() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse5.Do(func() {
		cache_Data_TraversableWithIndex_traverse5 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableDisj()).V3)
	})
	return cache_Data_TraversableWithIndex_traverse5
}

var cache_Data_TraversableWithIndex_traverse6 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse6 sync.Once
func Get_Data_TraversableWithIndex_traverse6() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse6.Do(func() {
		cache_Data_TraversableWithIndex_traverse6 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableConj()).V3)
	})
	return cache_Data_TraversableWithIndex_traverse6
}

var cache_Data_TraversableWithIndex_traverse7 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse7 sync.Once
func Get_Data_TraversableWithIndex_traverse7() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse7.Do(func() {
		cache_Data_TraversableWithIndex_traverse7 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableAdditive()).V3)
	})
	return cache_Data_TraversableWithIndex_traverse7
}

var cache_Data_TraversableWithIndex_TraversableWithIndex_dollar_Dict gopurs_runtime.Value
var once_Data_TraversableWithIndex_TraversableWithIndex_dollar_Dict sync.Once
func Get_Data_TraversableWithIndex_TraversableWithIndex_dollar_Dict() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_TraversableWithIndex_dollar_Dict.Do(func() {
		cache_Data_TraversableWithIndex_TraversableWithIndex_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Call_Data_TraversableWithIndex_TraversableWithIndex_dollar_Dict(func() struct{
	FoldableWithIndex1 gopurs_runtime.Value
	FunctorWithIndex0 gopurs_runtime.Value
	Traversable2 gopurs_runtime.Value
	traverseWithIndex gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	FoldableWithIndex1 gopurs_runtime.Value
	FunctorWithIndex0 gopurs_runtime.Value
	Traversable2 gopurs_runtime.Value
	traverseWithIndex gopurs_runtime.Value
}{}
					clone.FoldableWithIndex1 = gopurs_runtime.RecordGet(orig, "FoldableWithIndex1")
					clone.FunctorWithIndex0 = gopurs_runtime.RecordGet(orig, "FunctorWithIndex0")
					clone.Traversable2 = gopurs_runtime.RecordGet(orig, "Traversable2")
					clone.traverseWithIndex = gopurs_runtime.RecordGet(orig, "traverseWithIndex")
					return clone
				}()))}
})
	})
	return cache_Data_TraversableWithIndex_TraversableWithIndex_dollar_Dict
}

var cache_Data_TraversableWithIndex_traverseWithIndexDefault gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverseWithIndexDefault sync.Once
func Get_Data_TraversableWithIndex_traverseWithIndexDefault() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverseWithIndexDefault.Do(func() {
		cache_Data_TraversableWithIndex_traverseWithIndexDefault = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traverseWithIndexDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box))
})
	})
	return cache_Data_TraversableWithIndex_traverseWithIndexDefault
}

var cache_Data_TraversableWithIndex_traverseWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverseWithIndex sync.Once
func Get_Data_TraversableWithIndex_traverseWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverseWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_traverseWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traverseWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_TraversableWithIndex_traverseWithIndex
}

var cache_Data_TraversableWithIndex_traverseDefault gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverseDefault sync.Once
func Get_Data_TraversableWithIndex_traverseDefault() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverseDefault.Do(func() {
		cache_Data_TraversableWithIndex_traverseDefault = gopurs_runtime.Func3(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traverseDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), f_2_box)
})
	})
	return cache_Data_TraversableWithIndex_traverseDefault
}

var cache_Data_TraversableWithIndex_traversableWithIndexTuple gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexTuple sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexTuple() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexTuple.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexTuple = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_2190796645_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_3374046885_3725484264(Rebox_Data_TraversableWithIndex_3725484264_3374046885(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexTuple()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_935379557_2412140840(Rebox_Data_TraversableWithIndex_2412140840_935379557(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexTuple()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_3543431075_3043886126(Rebox_Data_TraversableWithIndex_3043886126_3543431075(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableTuple()))))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(f_2, Get_Data_Unit_unit(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
})})))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexTuple
}

var cache_Data_TraversableWithIndex_traversableWithIndexProduct gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexProduct sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexProduct() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexProduct.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexProduct = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traversableWithIndexProduct(dictTraversableWithIndex_0_box)
})
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexProduct
}

var cache_Data_TraversableWithIndex_traversableWithIndexMultiplicative gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexMultiplicative sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexMultiplicative() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexMultiplicative.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexMultiplicative()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexMultiplicative()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMultiplicative()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)]))
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableMultiplicative()).V3), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
})}))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexMultiplicative
}

var cache_Data_TraversableWithIndex_traversableWithIndexMaybe gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexMaybe sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexMaybe() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexMaybe.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexMaybe = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_998104713_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_2286084809_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_FoldableWithIndex_foldableWithIndexMaybe())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_3136246921_2412140840(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_FunctorWithIndex_functorWithIndexMaybe())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_3188237647_3043886126(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Traversable_traversableMaybe())))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])]))
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Traversable_traversableMaybe()).V3), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
})})))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexMaybe
}

var cache_Data_TraversableWithIndex_traversableWithIndexLast gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexLast sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexLast() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexLast.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexLast = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_998104713_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_2286084809_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_FoldableWithIndex_foldableWithIndexLast())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_3136246921_2412140840(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_FunctorWithIndex_functorWithIndexLast())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_3188237647_3043886126(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Traversable_traversableLast())))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])]))
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Traversable_traversableLast()).V3), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
})})))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexLast
}

var cache_Data_TraversableWithIndex_traversableWithIndexIdentity gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexIdentity sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexIdentity() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexIdentity.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexIdentity = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexIdentity()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexIdentity()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableIdentity()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Identity_Identity(), gopurs_runtime.Apply2(f_2, Get_Data_Unit_unit(), v_3))
})
})}))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexIdentity
}

var cache_Data_TraversableWithIndex_traversableWithIndexFirst gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexFirst sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexFirst() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexFirst.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexFirst = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_998104713_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_2286084809_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_FoldableWithIndex_foldableWithIndexFirst())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_3136246921_2412140840(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_FunctorWithIndex_functorWithIndexFirst())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_3188237647_3043886126(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Traversable_traversableFirst())))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])] (TypeApp (TypeVar m) [(ADT ["Data","Maybe","Maybe"] [(TypeVar b)])]))
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Traversable_traversableFirst()).V3), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
})})))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexFirst
}

var cache_Data_TraversableWithIndex_traversableWithIndexEither gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexEither sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexEither() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexEither.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexEither = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexEither()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexEither()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableEither()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Right(), gopurs_runtime.Apply2(v_2, Get_Data_Unit_unit(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
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
})}))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexEither
}

var cache_Data_TraversableWithIndex_traversableWithIndexDual gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexDual sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexDual() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexDual.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexDual = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexDual()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexDual()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableDual()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)]))
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableDual()).V3), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
})}))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexDual
}

var cache_Data_TraversableWithIndex_traversableWithIndexDisj gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexDisj sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexDisj() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexDisj.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexDisj = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexDisj()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexDisj()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableDisj()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)]))
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableDisj()).V3), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
})}))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexDisj
}

var cache_Data_TraversableWithIndex_traversableWithIndexCoproduct gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexCoproduct sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexCoproduct() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexCoproduct.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexCoproduct = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traversableWithIndexCoproduct(dictTraversableWithIndex_0_box)
})
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexCoproduct
}

var cache_Data_TraversableWithIndex_traversableWithIndexConst gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexConst sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexConst() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexConst.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexConst = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexConst()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexConst()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableConst()))}
}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v1_2)
})}))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexConst
}

var cache_Data_TraversableWithIndex_traversableWithIndexConj gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexConj sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexConj() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexConj.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexConj = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexConj()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexConj()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableConj()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)]))
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableConj()).V3), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
})}))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexConj
}

var cache_Data_TraversableWithIndex_traversableWithIndexCompose gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexCompose sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexCompose() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexCompose.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexCompose = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traversableWithIndexCompose(dictTraversableWithIndex_0_box)
})
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexCompose
}

var cache_Data_TraversableWithIndex_traversableWithIndexArray gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexArray sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexArray() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexArray.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexArray = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_2955889203_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_2491554675_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexArray())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_2773701683_2412140840(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexArray())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): FunctorWithIndex0_1_0 shape=App(Other) bindingType=(ADT ["Data","FunctorWithIndex","FunctorWithIndex"] [(TypeVar i), (TypeVar t)])
FunctorWithIndex0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray()).V1), gopurs_runtime.Value{}))
_ = FunctorWithIndex0_1_0
// TAST (Let): sequence1_2_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeApp (TypeVar m) [(TypeVar b)])])] (TypeApp (TypeVar m) [(TypeApp (TypeVar t) [(TypeVar b)])]))
sequence1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray()).V2), gopurs_runtime.Value{}), "sequence"), dictApplicative_0)
_ = sequence1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (TypeApp (TypeVar t) [(TypeApp (TypeVar m) [(TypeVar b)])]))
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(FunctorWithIndex0_1_0.V1), f_3)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_2_1, gopurs_runtime.Apply(__local_var_4_2, x_5))
})
})
})})))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexArray
}

var cache_Data_TraversableWithIndex_traversableWithIndexApp gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexApp sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexApp() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexApp.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexApp = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traversableWithIndexApp(dictTraversableWithIndex_0_box)
})
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexApp
}

var cache_Data_TraversableWithIndex_traversableWithIndexAdditive gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexAdditive sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexAdditive() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexAdditive.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexAdditive = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexAdditive()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexAdditive()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableAdditive()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 shape=App(Other) bindingType=(Func [(Func [(TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)]))
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableAdditive()).V3), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
})}))}
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexAdditive
}

var cache_Data_TraversableWithIndex_mapAccumRWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_mapAccumRWithIndex sync.Once
func Get_Data_TraversableWithIndex_mapAccumRWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_mapAccumRWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_mapAccumRWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_TraversableWithIndex_mapAccumRWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"accum", "value"}, []gopurs_runtime.Value{orig.accum, orig.value})
				}()
})
	})
	return cache_Data_TraversableWithIndex_mapAccumRWithIndex
}

var cache_Data_TraversableWithIndex_scanrWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_scanrWithIndex sync.Once
func Get_Data_TraversableWithIndex_scanrWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_scanrWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_scanrWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_scanrWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_Data_TraversableWithIndex_scanrWithIndex
}

var cache_Data_TraversableWithIndex_mapAccumLWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_mapAccumLWithIndex sync.Once
func Get_Data_TraversableWithIndex_mapAccumLWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_mapAccumLWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_mapAccumLWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_TraversableWithIndex_mapAccumLWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
				_ = orig
				return gopurs_runtime.RecordDict([]string{"accum", "value"}, []gopurs_runtime.Value{orig.accum, orig.value})
				}()
})
	})
	return cache_Data_TraversableWithIndex_mapAccumLWithIndex
}

var cache_Data_TraversableWithIndex_scanlWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_scanlWithIndex sync.Once
func Get_Data_TraversableWithIndex_scanlWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_scanlWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_scanlWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_scanlWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_Data_TraversableWithIndex_scanlWithIndex
}

var cache_Data_TraversableWithIndex_forWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_forWithIndex sync.Once
func Get_Data_TraversableWithIndex_forWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_forWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_forWithIndex = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictTraversableWithIndex_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_forWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_1_box))
})
	})
	return cache_Data_TraversableWithIndex_forWithIndex
}

type Constructor_Data_TraversableWithIndex_TraversableWithIndex[T_i any, T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2078610234] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_TraversableWithIndex_TraversableWithIndex[any, any])(ptr)
		_ = c
		switch key {
		case "FoldableWithIndex1": return gopurs_runtime.Box(c.V0)
		case "FunctorWithIndex0": return gopurs_runtime.Box(c.V1)
		case "Traversable2": return gopurs_runtime.Box(c.V2)
		case "traverseWithIndex": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_TraversableWithIndex_TraversableWithIndex: " + key)
		}
	}
}


func Call_Data_TraversableWithIndex_TraversableWithIndex_dollar_Dict(x_0_loop struct{
	FoldableWithIndex1 gopurs_runtime.Value
	FunctorWithIndex0 gopurs_runtime.Value
	Traversable2 gopurs_runtime.Value
	traverseWithIndex gopurs_runtime.Value
}) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	FoldableWithIndex1 gopurs_runtime.Value
	FunctorWithIndex0 gopurs_runtime.Value
	Traversable2 gopurs_runtime.Value
	traverseWithIndex gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex"}, []gopurs_runtime.Value{orig.FoldableWithIndex1, orig.FunctorWithIndex0, orig.Traversable2, orig.traverseWithIndex})
				}())
}

func Call_Data_TraversableWithIndex_traverseWithIndexDefault(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): sequence__193435443_1_0 shape=Other bindingType=Any
sequence__193435443_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversableWithIndex_0.V2), gopurs_runtime.Value{}), "sequence")
_ = sequence__193435443_1_0
// TAST (Let): FunctorWithIndex0_2_1 shape=App(Other) bindingType=(ADT ["Data","FunctorWithIndex","FunctorWithIndex"] [(TypeVar i), (TypeVar t)])
FunctorWithIndex0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversableWithIndex_0.V1), gopurs_runtime.Value{}))
_ = FunctorWithIndex0_2_1
return gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): sequence1_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeApp (TypeVar m) [(TypeVar b)])])] (TypeApp (TypeVar m) [(TypeApp (TypeVar t) [(TypeVar b)])]))
sequence1_4_2 := gopurs_runtime.Apply(sequence__193435443_1_0, dictApplicative_3)
_ = sequence1_4_2
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (TypeApp (TypeVar t) [(TypeApp (TypeVar m) [(TypeVar b)])]))
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Box(FunctorWithIndex0_2_1.V1), f_5)
_ = __local_var_6_3
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_4_2, gopurs_runtime.Apply(__local_var_6_3, x_7))
})
})
})
}

func Call_Data_TraversableWithIndex_traverseWithIndex(dict_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_TraversableWithIndex_traverseDefault(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return f_2
}))
}

func Call_Data_TraversableWithIndex_traversableWithIndexProduct(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): functorWithIndexProduct__193435443_1_0 shape=App(Var) bindingType=Any
functorWithIndexProduct__193435443_1_0 := gopurs_runtime.Apply(Get_Data_FunctorWithIndex_functorWithIndexProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = functorWithIndexProduct__193435443_1_0
// TAST (Let): foldableWithIndexProduct__193435443_2_1 shape=App(Var) bindingType=Any
foldableWithIndexProduct__193435443_2_1 := gopurs_runtime.Apply(Get_Data_FoldableWithIndex_foldableWithIndexProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexProduct__193435443_2_1
// TAST (Let): traversableProduct__193435443_3_2 shape=App(Var) bindingType=Any
traversableProduct__193435443_3_2 := gopurs_runtime.Apply(Get_Data_Traversable_traversableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableProduct__193435443_3_2
return gopurs_runtime.Func(func(dictTraversableWithIndex1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorWithIndexProduct1_5_3 shape=App(Other) bindingType=(ADT ["Data","FunctorWithIndex","FunctorWithIndex"] [(ADT ["Data","Either","Either"] [(TypeVar a), (TypeVar b)]), (TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorWithIndexProduct1_5_3 := Rebox_Data_TraversableWithIndex_2412140840_935379557(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(functorWithIndexProduct__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FunctorWithIndex0"), gopurs_runtime.Value{}))))
_ = functorWithIndexProduct1_5_3
// TAST (Let): foldableWithIndexProduct1_6_4 shape=App(Other) bindingType=(ADT ["Data","FoldableWithIndex","FoldableWithIndex"] [(ADT ["Data","Either","Either"] [(TypeVar a), (TypeVar b)]), (TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
foldableWithIndexProduct1_6_4 := Rebox_Data_TraversableWithIndex_3725484264_3374046885(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(foldableWithIndexProduct__193435443_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FoldableWithIndex1"), gopurs_runtime.Value{}))))
_ = foldableWithIndexProduct1_6_4
// TAST (Let): traversableProduct1_7_5 shape=App(Other) bindingType=(ADT ["Data","Traversable","Traversable"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
traversableProduct1_7_5 := Rebox_Data_TraversableWithIndex_3043886126_3543431075(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](gopurs_runtime.Apply(traversableProduct__193435443_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "Traversable2"), gopurs_runtime.Value{}))))
_ = traversableProduct1_7_5
return gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_2190796645_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_3374046885_3725484264(foldableWithIndexProduct1_6_4))}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_935379557_2412140840(functorWithIndexProduct1_5_3))}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_3543431075_3043886126(traversableProduct1_7_5))}
}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_9_6 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_6
return gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_6.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_9_6.V0), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Product_product(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_8))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_12, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_8))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_12, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1))
})
})})))}
})
}

func Call_Data_TraversableWithIndex_traversableWithIndexCoproduct(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): functorWithIndexCoproduct__193435443_1_0 shape=App(Var) bindingType=Any
functorWithIndexCoproduct__193435443_1_0 := gopurs_runtime.Apply(Get_Data_FunctorWithIndex_functorWithIndexCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = functorWithIndexCoproduct__193435443_1_0
// TAST (Let): foldableWithIndexCoproduct__193435443_2_1 shape=App(Var) bindingType=Any
foldableWithIndexCoproduct__193435443_2_1 := gopurs_runtime.Apply(Get_Data_FoldableWithIndex_foldableWithIndexCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCoproduct__193435443_2_1
// TAST (Let): traversableCoproduct__193435443_3_2 shape=App(Var) bindingType=Any
traversableCoproduct__193435443_3_2 := gopurs_runtime.Apply(Get_Data_Traversable_traversableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCoproduct__193435443_3_2
return gopurs_runtime.Func(func(dictTraversableWithIndex1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorWithIndexCoproduct1_5_3 shape=App(Other) bindingType=(ADT ["Data","FunctorWithIndex","FunctorWithIndex"] [(ADT ["Data","Either","Either"] [(TypeVar a), (TypeVar b)]), (TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorWithIndexCoproduct1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(functorWithIndexCoproduct__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FunctorWithIndex0"), gopurs_runtime.Value{})))
_ = functorWithIndexCoproduct1_5_3
// TAST (Let): foldableWithIndexCoproduct1_6_4 shape=App(Other) bindingType=(ADT ["Data","FoldableWithIndex","FoldableWithIndex"] [(ADT ["Data","Either","Either"] [(TypeVar a), (TypeVar b)]), (TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
foldableWithIndexCoproduct1_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(foldableWithIndexCoproduct__193435443_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FoldableWithIndex1"), gopurs_runtime.Value{})))
_ = foldableWithIndexCoproduct1_6_4
// TAST (Let): traversableCoproduct1_7_5 shape=App(Other) bindingType=(ADT ["Data","Traversable","Traversable"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
traversableCoproduct1_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](gopurs_runtime.Apply(traversableCoproduct__193435443_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "Traversable2"), gopurs_runtime.Value{})))
_ = traversableCoproduct1_7_5
return gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(foldableWithIndexCoproduct1_6_4)}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(functorWithIndexCoproduct1_5_3)}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(traversableCoproduct1_7_5)}
}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_9_6 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_6
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_7 shape=App(Other) bindingType=Any
__local_var_11_7 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_9_6.V0), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_11, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
}))
_ = __local_var_11_7
// TAST (Let): __local_var_12_9 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar m) [(TypeApp (TypeVar f) [(TypeVar b)])]))
__local_var_12_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_8))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_12, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
_ = __local_var_12_9
// TAST (Let): __local_var_12_8 shape=Let(Abs(App(Other))) bindingType=(TypeApp (TypeVar m) [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar b)])])
__local_var_12_8 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_7, gopurs_runtime.Apply(__local_var_12_9, x_13))
})
_ = __local_var_12_8
// TAST (Let): __local_var_13_11 shape=App(Other) bindingType=Any
__local_var_13_11 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_9_6.V0), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_13, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
}))
_ = __local_var_13_11
// TAST (Let): __local_var_14_12 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar m) [(TypeApp (TypeVar g) [(TypeVar b)])]))
__local_var_14_12 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_8))}, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_14, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
}))
_ = __local_var_14_12
// TAST (Let): __local_var_13_10 shape=Let(Let(Abs(App(Other)))) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar m) [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar b)])]))
__local_var_13_10 := gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_11, gopurs_runtime.Apply(__local_var_14_12, x_15))
})
_ = __local_var_13_10
return gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(__local_var_12_8, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t13 = gopurs_runtime.Apply(__local_var_13_10, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_14.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
})
})}))}
})
}

func Call_Data_TraversableWithIndex_traversableWithIndexCompose(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorCompose__193435443_2_2 shape=Let(Abs(LitRecord)) bindingType=Any
functorCompose__193435443_2_2 := gopurs_runtime.Func(func(dictFunctor1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_3, "map"), f_4), v_5)
})}))}
})
_ = functorCompose__193435443_2_2
// TAST (Let): functorWithIndexCompose__193435443_1_0 shape=Let(Let(Abs(Let(LitRecord)))) bindingType=Any
functorWithIndexCompose__193435443_1_0 := gopurs_runtime.Func(func(dictFunctorWithIndex1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCompose1_4_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCompose1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_3, "Functor0"), gopurs_runtime.Value{})))
_ = functorCompose1_4_4
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_1961781733_2412140840((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose1_4_4)}
}), gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mapWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_3, "mapWithIndex"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{x_7, b_8}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
}), v_6)
})})))}
})
_ = functorWithIndexCompose__193435443_1_0
// TAST (Let): foldableWithIndexCompose__193435443_2_5 shape=App(Var) bindingType=Any
foldableWithIndexCompose__193435443_2_5 := gopurs_runtime.Apply(Get_Data_FoldableWithIndex_foldableWithIndexCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCompose__193435443_2_5
// TAST (Let): traversableCompose__193435443_3_6 shape=App(Var) bindingType=Any
traversableCompose__193435443_3_6 := gopurs_runtime.Apply(Get_Data_Traversable_traversableCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCompose__193435443_3_6
return gopurs_runtime.Func(func(dictTraversableWithIndex1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorWithIndexCompose1_5_7 shape=App(Other) bindingType=(ADT ["Data","FunctorWithIndex","FunctorWithIndex"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)]), (TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorWithIndexCompose1_5_7 := Rebox_Data_TraversableWithIndex_2412140840_1961781733(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(functorWithIndexCompose__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FunctorWithIndex0"), gopurs_runtime.Value{}))))
_ = functorWithIndexCompose1_5_7
// TAST (Let): foldableWithIndexCompose1_6_8 shape=App(Other) bindingType=(ADT ["Data","FoldableWithIndex","FoldableWithIndex"] [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)]), (TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
foldableWithIndexCompose1_6_8 := Rebox_Data_TraversableWithIndex_3725484264_139552293(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(foldableWithIndexCompose__193435443_2_5, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FoldableWithIndex1"), gopurs_runtime.Value{}))))
_ = foldableWithIndexCompose1_6_8
// TAST (Let): traversableCompose1_7_9 shape=App(Other) bindingType=(ADT ["Data","Traversable","Traversable"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
traversableCompose1_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](gopurs_runtime.Apply(traversableCompose__193435443_3_6, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "Traversable2"), gopurs_runtime.Value{})))
_ = traversableCompose1_7_9
return gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_145886949_1812164904((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_139552293_3725484264(foldableWithIndexCompose1_6_8))}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_TraversableWithIndex_1961781733_2412140840(functorWithIndexCompose1_5_7))}
}), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(traversableCompose1_7_9)}
}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_9_10 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_9_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_10
// TAST (Let): traverseWithIndex2_10_11 shape=App(Other) bindingType=(Func [(Func [(TypeVar b), (TypeVar a)] (TypeApp (TypeVar m) [(TypeVar b)])), (TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar m) [(TypeApp (TypeVar g) [(TypeVar b)])]))
traverseWithIndex2_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), dictApplicative_8)
_ = traverseWithIndex2_10_11
return gopurs_runtime.Func2(func(f_11 gopurs_runtime.Value, v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_9_10.V0), Get_Data_Functor_Compose_Compose(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_8))}, gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverseWithIndex2_10_11, gopurs_runtime.Func(func(b_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_11, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{x_13, b_14}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
}), v_12))
})
})})))}
})
}

func Call_Data_TraversableWithIndex_traversableWithIndexApp(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorApp_2_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f)])])
functorApp_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = functorApp_2_2
// TAST (Let): functorWithIndexApp_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","FunctorWithIndex","FunctorWithIndex"] [(TypeVar a), (TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f)])])
functorWithIndexApp_1_0 := (&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorApp_2_2)}
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mapWithIndex"), f_3, v_4)
})})
_ = functorWithIndexApp_1_0
// TAST (Let): foldableWithIndexApp_2_3 shape=App(Var) bindingType=(ADT ["Data","FoldableWithIndex","FoldableWithIndex"] [(TypeVar a), (TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f)])])
foldableWithIndexApp_2_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_FoldableWithIndex_foldableWithIndexApp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{})))
_ = foldableWithIndexApp_2_3
// TAST (Let): traversableApp_3_4 shape=App(Var) bindingType=(ADT ["Data","Traversable","Traversable"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f)])])
traversableApp_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Traversable_traversableApp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{})))
_ = traversableApp_3_4
return gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(foldableWithIndexApp_2_3)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(functorWithIndexApp_1_0)}
}), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(traversableApp_3_4)}
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_5_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_5
return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_5.V0), Get_Data_Functor_App_App(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_4))}, f_6, v_7))
})
})}))}
}

func Call_Data_TraversableWithIndex_mapAccumRWithIndex(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return func() struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Traversable_Accum_Internal_applicativeStateR()))}, gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := func() struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
					_ = orig
					clone := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.accum = gopurs_runtime.RecordGet(orig, "accum")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"accum", "value"}, []gopurs_runtime.Value{orig.accum, orig.value})
				}()
}), xs_3, s0_2)
					_ = orig
					clone := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.accum = gopurs_runtime.RecordGet(orig, "accum")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
}

func Call_Data_TraversableWithIndex_scanrWithIndex(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return Call_Data_TraversableWithIndex_mapAccumRWithIndex(dictTraversableWithIndex_0, gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): b_prime__7_0 shape=App(Other) bindingType=(TypeVar b)
b_prime__7_0 := gopurs_runtime.Apply3(f_1, i_4, a_6, b_5)
_ = b_prime__7_0
return func() gopurs_runtime.Value {
				orig := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{b_prime__7_0, b_prime__7_0}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"accum", "value"}, []gopurs_runtime.Value{orig.accum, orig.value})
				}()
}), b0_2, xs_3).value
}

func Call_Data_TraversableWithIndex_mapAccumLWithIndex(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return func() struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Traversable_Accum_Internal_applicativeStateL()))}, gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := func() struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
					_ = orig
					clone := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.accum = gopurs_runtime.RecordGet(orig, "accum")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict([]string{"accum", "value"}, []gopurs_runtime.Value{orig.accum, orig.value})
				}()
}), xs_3, s0_2)
					_ = orig
					clone := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.accum = gopurs_runtime.RecordGet(orig, "accum")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
}

func Call_Data_TraversableWithIndex_scanlWithIndex(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return Call_Data_TraversableWithIndex_mapAccumLWithIndex(dictTraversableWithIndex_0, gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): b_prime__7_0 shape=App(Other) bindingType=(TypeVar b)
b_prime__7_0 := gopurs_runtime.Apply3(f_1, i_4, b_5, a_6)
_ = b_prime__7_0
return func() gopurs_runtime.Value {
				orig := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{b_prime__7_0, b_prime__7_0}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"accum", "value"}, []gopurs_runtime.Value{orig.accum, orig.value})
				}()
}), b0_2, xs_3).value
}

func Call_Data_TraversableWithIndex_forWithIndex(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictTraversableWithIndex_1_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var dictTraversableWithIndex_1 *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_1_loop
_ = dictTraversableWithIndex_1
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversableWithIndex_1.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)})
_ = __local_var_2_0
return gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_4, b_3)
})
}

func Rebox_Data_TraversableWithIndex_139552293_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_145886949_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_1961781733_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_TraversableWithIndex_2190796645_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_2286084809_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_2412140840_1961781733(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_TraversableWithIndex_2412140840_935379557(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_TraversableWithIndex_2491554675_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_2773701683_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_TraversableWithIndex_2955889203_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_3043886126_3543431075(in *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) *Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_3136246921_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_TraversableWithIndex_3188237647_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_3374046885_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_3543431075_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_3725484264_139552293(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_3725484264_3374046885(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_TraversableWithIndex_935379557_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_TraversableWithIndex_998104713_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, *Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}



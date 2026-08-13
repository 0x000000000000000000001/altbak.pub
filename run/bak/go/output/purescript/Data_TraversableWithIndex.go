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
		cache_Data_TraversableWithIndex_traverse = gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMultiplicative(), "traverse")
	})
	return cache_Data_TraversableWithIndex_traverse
}

var cache_Data_TraversableWithIndex_traverse1 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse1 sync.Once
func Get_Data_TraversableWithIndex_traverse1() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse1.Do(func() {
		cache_Data_TraversableWithIndex_traverse1 = gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMaybe(), "traverse")
	})
	return cache_Data_TraversableWithIndex_traverse1
}

var cache_Data_TraversableWithIndex_traverse2 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse2 sync.Once
func Get_Data_TraversableWithIndex_traverse2() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse2.Do(func() {
		cache_Data_TraversableWithIndex_traverse2 = gopurs_runtime.RecordGet(Get_Data_Traversable_traversableLast(), "traverse")
	})
	return cache_Data_TraversableWithIndex_traverse2
}

var cache_Data_TraversableWithIndex_traverse3 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse3 sync.Once
func Get_Data_TraversableWithIndex_traverse3() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse3.Do(func() {
		cache_Data_TraversableWithIndex_traverse3 = gopurs_runtime.RecordGet(Get_Data_Traversable_traversableFirst(), "traverse")
	})
	return cache_Data_TraversableWithIndex_traverse3
}

var cache_Data_TraversableWithIndex_traverse4 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse4 sync.Once
func Get_Data_TraversableWithIndex_traverse4() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse4.Do(func() {
		cache_Data_TraversableWithIndex_traverse4 = gopurs_runtime.RecordGet(Get_Data_Traversable_traversableDual(), "traverse")
	})
	return cache_Data_TraversableWithIndex_traverse4
}

var cache_Data_TraversableWithIndex_traverse5 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse5 sync.Once
func Get_Data_TraversableWithIndex_traverse5() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse5.Do(func() {
		cache_Data_TraversableWithIndex_traverse5 = gopurs_runtime.RecordGet(Get_Data_Traversable_traversableDisj(), "traverse")
	})
	return cache_Data_TraversableWithIndex_traverse5
}

var cache_Data_TraversableWithIndex_traverse6 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse6 sync.Once
func Get_Data_TraversableWithIndex_traverse6() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse6.Do(func() {
		cache_Data_TraversableWithIndex_traverse6 = gopurs_runtime.RecordGet(Get_Data_Traversable_traversableConj(), "traverse")
	})
	return cache_Data_TraversableWithIndex_traverse6
}

var cache_Data_TraversableWithIndex_traverse7 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverse7 sync.Once
func Get_Data_TraversableWithIndex_traverse7() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverse7.Do(func() {
		cache_Data_TraversableWithIndex_traverse7 = gopurs_runtime.RecordGet(Get_Data_Traversable_traversableAdditive(), "traverse")
	})
	return cache_Data_TraversableWithIndex_traverse7
}

var cache_Data_TraversableWithIndex_TraversableWithIndex_dollarDict gopurs_runtime.Value
var once_Data_TraversableWithIndex_TraversableWithIndex_dollarDict sync.Once
func Get_Data_TraversableWithIndex_TraversableWithIndex_dollarDict() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_TraversableWithIndex_dollarDict.Do(func() {
		cache_Data_TraversableWithIndex_TraversableWithIndex_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_TraversableWithIndex_dollarDict(x_0_box)
})
	})
	return cache_Data_TraversableWithIndex_TraversableWithIndex_dollarDict
}

var cache_Data_TraversableWithIndex_traverseWithIndexDefault gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverseWithIndexDefault sync.Once
func Get_Data_TraversableWithIndex_traverseWithIndexDefault() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverseWithIndexDefault.Do(func() {
		cache_Data_TraversableWithIndex_traverseWithIndexDefault = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traverseWithIndexDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_0_box))
})
	})
	return cache_Data_TraversableWithIndex_traverseWithIndexDefault
}

var cache_Data_TraversableWithIndex_traverseWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverseWithIndex sync.Once
func Get_Data_TraversableWithIndex_traverseWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverseWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_traverseWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traverseWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dict_0_box))
})
	})
	return cache_Data_TraversableWithIndex_traverseWithIndex
}

var cache_Data_TraversableWithIndex_traverseDefault gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverseDefault sync.Once
func Get_Data_TraversableWithIndex_traverseDefault() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverseDefault.Do(func() {
		cache_Data_TraversableWithIndex_traverseDefault = gopurs_runtime.Func3(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traverseDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_1_box), f_2_box)
})
	})
	return cache_Data_TraversableWithIndex_traverseDefault
}

var cache_Data_TraversableWithIndex_traversableWithIndexTuple gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexTuple sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexTuple() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexTuple.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexTuple = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableTuple()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(f_2, Get_Data_Unit_unit(), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1))
})
})
}))
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
		cache_Data_TraversableWithIndex_traversableWithIndexMultiplicative = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableMultiplicative()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 -> gopurs_runtime.Value
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMultiplicative(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}))
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexMultiplicative
}

var cache_Data_TraversableWithIndex_traversableWithIndexMaybe gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexMaybe sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexMaybe() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexMaybe.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexMaybe = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 -> gopurs_runtime.Value
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMaybe(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}))
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexMaybe
}

var cache_Data_TraversableWithIndex_traversableWithIndexLast gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexLast sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexLast() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexLast.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexLast = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexLast()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexLast()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableLast()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 -> gopurs_runtime.Value
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableLast(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}))
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexLast
}

var cache_Data_TraversableWithIndex_traversableWithIndexIdentity gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexIdentity sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexIdentity() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexIdentity.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexIdentity = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableIdentity()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Identity_Identity(), gopurs_runtime.Apply2(f_2, Get_Data_Unit_unit(), v_3))
})
})
}))
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexIdentity
}

var cache_Data_TraversableWithIndex_traversableWithIndexFirst gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexFirst sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexFirst() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexFirst.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexFirst = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexFirst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexFirst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableFirst()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 -> gopurs_runtime.Value
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableFirst(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}))
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexFirst
}

var cache_Data_TraversableWithIndex_traversableWithIndexEither gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexEither sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexEither() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexEither.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexEither = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableEither()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v1_3.UnsafePtr).V0})})
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Right(), gopurs_runtime.Apply2(v_2, Get_Data_Unit_unit(), (*Constructor_Data_Either_Right)(v1_3.UnsafePtr).V0))
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
}))
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexEither
}

var cache_Data_TraversableWithIndex_traversableWithIndexDual gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexDual sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexDual() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexDual.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexDual = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableDual()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 -> gopurs_runtime.Value
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableDual(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}))
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexDual
}

var cache_Data_TraversableWithIndex_traversableWithIndexDisj gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexDisj sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexDisj() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexDisj.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexDisj = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableDisj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 -> gopurs_runtime.Value
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableDisj(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}))
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
		cache_Data_TraversableWithIndex_traversableWithIndexConst = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableConst()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v1_2)
})
})
}))
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexConst
}

var cache_Data_TraversableWithIndex_traversableWithIndexConj gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexConj sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexConj() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexConj.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexConj = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableConj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 -> gopurs_runtime.Value
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableConj(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}))
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
		cache_Data_TraversableWithIndex_traversableWithIndexArray = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableArray()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): FunctorWithIndex0_1_0 -> *Constructor_Data_FunctorWithIndex_FunctorWithIndex
FunctorWithIndex0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_TraversableWithIndex_traversableWithIndexArray(), "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = FunctorWithIndex0_1_0
// TAST (Let): sequence1_2_1 -> gopurs_runtime.Value
sequence1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_TraversableWithIndex_traversableWithIndexArray(), "Traversable2"), gopurs_runtime.Value{}), "sequence"), dictApplicative_0)
_ = sequence1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(FunctorWithIndex0_1_0.V1), f_3)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_2_1, gopurs_runtime.Apply(__local_var_4_2, x_5))
})
})
}))
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
		cache_Data_TraversableWithIndex_traversableWithIndexAdditive = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableAdditive()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse8_1_0 -> gopurs_runtime.Value
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableAdditive(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}))
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexAdditive
}

var cache_Data_TraversableWithIndex_mapAccumRWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_mapAccumRWithIndex sync.Once
func Get_Data_TraversableWithIndex_mapAccumRWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_mapAccumRWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_mapAccumRWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_mapAccumRWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_TraversableWithIndex_mapAccumRWithIndex
}

var cache_Data_TraversableWithIndex_scanrWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_scanrWithIndex sync.Once
func Get_Data_TraversableWithIndex_scanrWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_scanrWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_scanrWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_scanrWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_Data_TraversableWithIndex_scanrWithIndex
}

var cache_Data_TraversableWithIndex_mapAccumLWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_mapAccumLWithIndex sync.Once
func Get_Data_TraversableWithIndex_mapAccumLWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_mapAccumLWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_mapAccumLWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_mapAccumLWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_TraversableWithIndex_mapAccumLWithIndex
}

var cache_Data_TraversableWithIndex_scanlWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_scanlWithIndex sync.Once
func Get_Data_TraversableWithIndex_scanlWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_scanlWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_scanlWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_scanlWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_Data_TraversableWithIndex_scanlWithIndex
}

var cache_Data_TraversableWithIndex_forWithIndex gopurs_runtime.Value
var once_Data_TraversableWithIndex_forWithIndex sync.Once
func Get_Data_TraversableWithIndex_forWithIndex() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_forWithIndex.Do(func() {
		cache_Data_TraversableWithIndex_forWithIndex = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictTraversableWithIndex_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_forWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_1_box))
})
	})
	return cache_Data_TraversableWithIndex_forWithIndex
}

var cache_Data_TraversableWithIndex_mapAccumLWithIndex__142050190 gopurs_runtime.Value
var once_Data_TraversableWithIndex_mapAccumLWithIndex__142050190 sync.Once
func Get_Data_TraversableWithIndex_mapAccumLWithIndex__142050190() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_mapAccumLWithIndex__142050190.Do(func() {
		cache_Data_TraversableWithIndex_mapAccumLWithIndex__142050190 = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_mapAccumLWithIndex__142050190(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_TraversableWithIndex_mapAccumLWithIndex__142050190
}

var cache_Data_TraversableWithIndex_mapAccumLWithIndex__1384596302 gopurs_runtime.Value
var once_Data_TraversableWithIndex_mapAccumLWithIndex__1384596302 sync.Once
func Get_Data_TraversableWithIndex_mapAccumLWithIndex__1384596302() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_mapAccumLWithIndex__1384596302.Do(func() {
		cache_Data_TraversableWithIndex_mapAccumLWithIndex__1384596302 = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_mapAccumLWithIndex__1384596302(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_TraversableWithIndex_mapAccumLWithIndex__1384596302
}

var cache_Data_TraversableWithIndex_mapAccumRWithIndex__142050190 gopurs_runtime.Value
var once_Data_TraversableWithIndex_mapAccumRWithIndex__142050190 sync.Once
func Get_Data_TraversableWithIndex_mapAccumRWithIndex__142050190() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_mapAccumRWithIndex__142050190.Do(func() {
		cache_Data_TraversableWithIndex_mapAccumRWithIndex__142050190 = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_mapAccumRWithIndex__142050190(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_TraversableWithIndex_mapAccumRWithIndex__142050190
}

var cache_Data_TraversableWithIndex_mapAccumRWithIndex__1384596302 gopurs_runtime.Value
var once_Data_TraversableWithIndex_mapAccumRWithIndex__1384596302 sync.Once
func Get_Data_TraversableWithIndex_mapAccumRWithIndex__1384596302() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_mapAccumRWithIndex__1384596302.Do(func() {
		cache_Data_TraversableWithIndex_mapAccumRWithIndex__1384596302 = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_mapAccumRWithIndex__1384596302(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_TraversableWithIndex_mapAccumRWithIndex__1384596302
}

var cache_Data_TraversableWithIndex_traversableWithIndexArray__1681559805 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traversableWithIndexArray__1681559805 sync.Once
func Get_Data_TraversableWithIndex_traversableWithIndexArray__1681559805() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traversableWithIndexArray__1681559805.Do(func() {
		cache_Data_TraversableWithIndex_traversableWithIndexArray__1681559805 = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FoldableWithIndex_foldableWithIndexArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_FunctorWithIndex_functorWithIndexArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableArray()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): sequence1_1_0 -> gopurs_runtime.Value
sequence1_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableArray(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
_ = sequence1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(Get_Data_FunctorWithIndex_mapWithIndexArray(), f_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
})
}))
	})
	return cache_Data_TraversableWithIndex_traversableWithIndexArray__1681559805
}

var cache_Data_TraversableWithIndex_traverseWithIndex__2726076659 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverseWithIndex__2726076659 sync.Once
func Get_Data_TraversableWithIndex_traverseWithIndex__2726076659() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverseWithIndex__2726076659.Do(func() {
		cache_Data_TraversableWithIndex_traverseWithIndex__2726076659 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traverseWithIndex__2726076659(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dict_0_box))
})
	})
	return cache_Data_TraversableWithIndex_traverseWithIndex__2726076659
}

var cache_Data_TraversableWithIndex_traverseWithIndex__2979276403 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverseWithIndex__2979276403 sync.Once
func Get_Data_TraversableWithIndex_traverseWithIndex__2979276403() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverseWithIndex__2979276403.Do(func() {
		cache_Data_TraversableWithIndex_traverseWithIndex__2979276403 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traverseWithIndex__2979276403(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](dict_0_box))
})
	})
	return cache_Data_TraversableWithIndex_traverseWithIndex__2979276403
}

var cache_Data_TraversableWithIndex_traverseWithIndex__2841069947 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverseWithIndex__2841069947 sync.Once
func Get_Data_TraversableWithIndex_traverseWithIndex__2841069947() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverseWithIndex__2841069947.Do(func() {
		cache_Data_TraversableWithIndex_traverseWithIndex__2841069947 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traverseWithIndex__2841069947(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_TraversableWithIndex_traverseWithIndex__2841069947
}

var cache_Data_TraversableWithIndex_traverseWithIndex__1901281819 gopurs_runtime.Value
var once_Data_TraversableWithIndex_traverseWithIndex__1901281819 sync.Once
func Get_Data_TraversableWithIndex_traverseWithIndex__1901281819() gopurs_runtime.Value {
	once_Data_TraversableWithIndex_traverseWithIndex__1901281819.Do(func() {
		cache_Data_TraversableWithIndex_traverseWithIndex__1901281819 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_TraversableWithIndex_traverseWithIndex__1901281819(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_TraversableWithIndex_traverseWithIndex__1901281819
}

type Constructor_Data_TraversableWithIndex_TraversableWithIndex struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2078610234] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_TraversableWithIndex_TraversableWithIndex)(ptr)
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


func Call_Data_TraversableWithIndex_TraversableWithIndex_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_TraversableWithIndex_traverseWithIndexDefault(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): FunctorWithIndex0_1_0 -> *Constructor_Data_FunctorWithIndex_FunctorWithIndex
FunctorWithIndex0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversableWithIndex_0.V1), gopurs_runtime.Value{}))
_ = FunctorWithIndex0_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): sequence1_3_1 -> gopurs_runtime.Value
sequence1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversableWithIndex_0.V2), gopurs_runtime.Value{}), "sequence"), dictApplicative_2)
_ = sequence1_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.Box(FunctorWithIndex0_1_0.V1), f_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_3_1, gopurs_runtime.Apply(__local_var_5_2, x_6))
})
})
})
}

func Call_Data_TraversableWithIndex_traverseWithIndex(dict_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex) gopurs_runtime.Value {
var dict_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_TraversableWithIndex_traverseDefault(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex, dictApplicative_1_loop *Constructor_Control_Applicative_Applicative, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative = dictApplicative_1_loop
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
// TAST (Let): functorWithIndexProduct_1_0 -> gopurs_runtime.Value
functorWithIndexProduct_1_0 := gopurs_runtime.Apply(Get_Data_FunctorWithIndex_functorWithIndexProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = functorWithIndexProduct_1_0
// TAST (Let): foldableWithIndexProduct_2_1 -> gopurs_runtime.Value
foldableWithIndexProduct_2_1 := gopurs_runtime.Apply(Get_Data_FoldableWithIndex_foldableWithIndexProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexProduct_2_1
// TAST (Let): traversableProduct_3_2 -> gopurs_runtime.Value
traversableProduct_3_2 := gopurs_runtime.Apply(Get_Data_Traversable_traversableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableProduct_3_2
return gopurs_runtime.Func(func(dictTraversableWithIndex1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorWithIndexProduct1_5_3 -> gopurs_runtime.Value
functorWithIndexProduct1_5_3 := gopurs_runtime.Apply(functorWithIndexProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = functorWithIndexProduct1_5_3
// TAST (Let): foldableWithIndexProduct1_6_4 -> gopurs_runtime.Value
foldableWithIndexProduct1_6_4 := gopurs_runtime.Apply(foldableWithIndexProduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexProduct1_6_4
// TAST (Let): traversableProduct1_7_5 -> gopurs_runtime.Value
traversableProduct1_7_5 := gopurs_runtime.Apply(traversableProduct_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableProduct1_7_5
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexProduct1_6_4
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexProduct1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableProduct1_7_5
}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_9_6 -> *Constructor_Control_Apply_Apply
Apply0_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_6
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_9_6.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_9_6.V0), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Product_product(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_8))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_12})})
}), (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_8))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_12})})
}), (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1))
})
})
}))
})
}

func Call_Data_TraversableWithIndex_traversableWithIndexCoproduct(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): functorWithIndexCoproduct_1_0 -> gopurs_runtime.Value
functorWithIndexCoproduct_1_0 := gopurs_runtime.Apply(Get_Data_FunctorWithIndex_functorWithIndexCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = functorWithIndexCoproduct_1_0
// TAST (Let): foldableWithIndexCoproduct_2_1 -> gopurs_runtime.Value
foldableWithIndexCoproduct_2_1 := gopurs_runtime.Apply(Get_Data_FoldableWithIndex_foldableWithIndexCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCoproduct_2_1
// TAST (Let): traversableCoproduct_3_2 -> gopurs_runtime.Value
traversableCoproduct_3_2 := gopurs_runtime.Apply(Get_Data_Traversable_traversableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCoproduct_3_2
return gopurs_runtime.Func(func(dictTraversableWithIndex1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorWithIndexCoproduct1_5_3 -> gopurs_runtime.Value
functorWithIndexCoproduct1_5_3 := gopurs_runtime.Apply(functorWithIndexCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = functorWithIndexCoproduct1_5_3
// TAST (Let): foldableWithIndexCoproduct1_6_4 -> gopurs_runtime.Value
foldableWithIndexCoproduct1_6_4 := gopurs_runtime.Apply(foldableWithIndexCoproduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCoproduct1_6_4
// TAST (Let): traversableCoproduct1_7_5 -> gopurs_runtime.Value
traversableCoproduct1_7_5 := gopurs_runtime.Apply(traversableCoproduct_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCoproduct1_7_5
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexCoproduct1_6_4
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexCoproduct1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableCoproduct1_7_5
}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_9_6 -> *Constructor_Data_Functor_Functor
Functor0_9_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_6
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_8 -> gopurs_runtime.Value
__local_var_11_8 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_9_6.V0), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_11})}
}))
_ = __local_var_11_8
// TAST (Let): __local_var_12_9 -> gopurs_runtime.Value
__local_var_12_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_8))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_12})})
}))
_ = __local_var_12_9
// TAST (Let): __local_var_11_7 -> gopurs_runtime.Value
__local_var_11_7 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_8, gopurs_runtime.Apply(__local_var_12_9, x_13))
})
_ = __local_var_11_7
// TAST (Let): __local_var_12_11 -> gopurs_runtime.Value
__local_var_12_11 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_9_6.V0), gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_12})}
}))
_ = __local_var_12_11
// TAST (Let): __local_var_13_12 -> gopurs_runtime.Value
__local_var_13_12 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_8))}, gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
}))
_ = __local_var_13_12
// TAST (Let): __local_var_12_10 -> gopurs_runtime.Value
__local_var_12_10 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_11, gopurs_runtime.Apply(__local_var_13_12, x_14))
})
_ = __local_var_12_10
return gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(__local_var_11_7, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t13 = gopurs_runtime.Apply(__local_var_12_10, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
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
}))
})
}

func Call_Data_TraversableWithIndex_traversableWithIndexCompose(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): foldableWithIndexCompose_3_2 -> gopurs_runtime.Value
foldableWithIndexCompose_3_2 := gopurs_runtime.Apply(Get_Data_FoldableWithIndex_foldableWithIndexCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCompose_3_2
// TAST (Let): traversableCompose_4_3 -> gopurs_runtime.Value
traversableCompose_4_3 := gopurs_runtime.Apply(Get_Data_Traversable_traversableCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCompose_4_3
return gopurs_runtime.Func(func(dictTraversableWithIndex1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_6_5
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): functorCompose1_7_6 -> gopurs_runtime.Value
functorCompose1_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "map"), f_8), v_9)
})
}))
_ = functorCompose1_7_6
// TAST (Let): functorWithIndexCompose1_6_4 -> gopurs_runtime.Value
functorWithIndexCompose1_6_4 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_7_6
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mapWithIndex"), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "mapWithIndex"), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_10, b_11})})
}))
}), v_9)
})
}))
_ = functorWithIndexCompose1_6_4
// TAST (Let): foldableWithIndexCompose1_7_8 -> gopurs_runtime.Value
foldableWithIndexCompose1_7_8 := gopurs_runtime.Apply(foldableWithIndexCompose_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCompose1_7_8
// TAST (Let): traversableCompose1_8_9 -> gopurs_runtime.Value
traversableCompose1_8_9 := gopurs_runtime.Apply(traversableCompose_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCompose1_8_9
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexCompose1_7_8
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexCompose1_6_4
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableCompose1_8_9
}), gopurs_runtime.Func(func(dictApplicative_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_10_10 -> *Constructor_Data_Functor_Functor
Functor0_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_9, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_10
// TAST (Let): traverseWithIndex2_11_11 -> gopurs_runtime.Value
traverseWithIndex2_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "traverseWithIndex"), dictApplicative_9)
_ = traverseWithIndex2_11_11
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_10_10.V0), Get_Data_Functor_Compose_Compose(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_9))}, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverseWithIndex2_11_11, gopurs_runtime.Func(func(b_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_12, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_14, b_15})})
}))
}), v_13))
})
})
}))
})
}

func Call_Data_TraversableWithIndex_traversableWithIndexApp(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorApp_2_2 -> gopurs_runtime.Value
functorApp_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = functorApp_2_2
// TAST (Let): functorWithIndexApp_1_0 -> gopurs_runtime.Value
functorWithIndexApp_1_0 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorApp_2_2
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mapWithIndex"), f_3, v_4)
})
}))
_ = functorWithIndexApp_1_0
// TAST (Let): foldableWithIndexApp_2_3 -> gopurs_runtime.Value
foldableWithIndexApp_2_3 := gopurs_runtime.Apply(Get_Data_FoldableWithIndex_foldableWithIndexApp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexApp_2_3
// TAST (Let): traversableApp_3_4 -> gopurs_runtime.Value
traversableApp_3_4 := gopurs_runtime.Apply(Get_Data_Traversable_traversableApp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableApp_3_4
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexApp_2_3
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexApp_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableApp_3_4
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_5_5 -> *Constructor_Data_Functor_Functor
Functor0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_5
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_5.V0), Get_Data_Functor_App_App(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_4))}, f_6, v_7))
})
})
}))
}

func Call_Data_TraversableWithIndex_mapAccumRWithIndex(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateR()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_Data_TraversableWithIndex_scanrWithIndex(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex, f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.RecordGet(gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateR()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): b_prime_7_0 -> gopurs_runtime.Value
b_prime_7_0 := gopurs_runtime.Apply3(f_1, i_4, a_5, s_6)
_ = b_prime_7_0
return gopurs_runtime.RecordDict2("accum", "value", b_prime_7_0, b_prime_7_0)
})
})
}), xs_3, b0_2), "value")
}

func Call_Data_TraversableWithIndex_mapAccumLWithIndex(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateL()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_Data_TraversableWithIndex_scanlWithIndex(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex, f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.RecordGet(gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateL()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): b_prime_7_0 -> gopurs_runtime.Value
b_prime_7_0 := gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
_ = b_prime_7_0
return gopurs_runtime.RecordDict2("accum", "value", b_prime_7_0, b_prime_7_0)
})
})
}), xs_3, b0_2), "value")
}

func Call_Data_TraversableWithIndex_forWithIndex(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, dictTraversableWithIndex_1_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var dictTraversableWithIndex_1 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_1_loop
_ = dictTraversableWithIndex_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversableWithIndex_1.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)})
_ = __local_var_2_0
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_4, b_3)
})
})
}

func Call_Data_TraversableWithIndex_mapAccumLWithIndex__142050190(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateL()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_Data_TraversableWithIndex_mapAccumLWithIndex__1384596302(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateL()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_Data_TraversableWithIndex_mapAccumRWithIndex__142050190(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateR()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_Data_TraversableWithIndex_mapAccumRWithIndex__1384596302(dictTraversableWithIndex_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversableWithIndex_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateR()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_Data_TraversableWithIndex_traverseWithIndex__2726076659(dict_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex) gopurs_runtime.Value {
var dict_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_TraversableWithIndex_traverseWithIndex__2979276403(dict_0_loop *Constructor_Data_TraversableWithIndex_TraversableWithIndex) gopurs_runtime.Value {
var dict_0 *Constructor_Data_TraversableWithIndex_TraversableWithIndex = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_TraversableWithIndex_traverseWithIndex__2841069947(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](Get_Data_List_Lazy_Types_traversableWithIndexNonEmpty()).V3), __eta0_0, __eta1_1)
}

func Call_Data_TraversableWithIndex_traverseWithIndex__1901281819(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](Get_Data_List_Types_traversableWithIndexNonEmpty()).V3), __eta0_0, __eta1_1)
}



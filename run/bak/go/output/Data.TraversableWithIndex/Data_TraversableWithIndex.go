package Data_TraversableWithIndex

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Functor_App "gopurs/output/Data.Functor.App"
	pkg_Data_Functor_Compose "gopurs/output/Data.Functor.Compose"
	pkg_Data_Functor_Product "gopurs/output/Data.Functor.Product"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Traversable_Accum_Internal "gopurs/output/Data.Traversable.Accum.Internal"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_traverseWithIndexDefault gopurs_runtime.Value
var once_traverseWithIndexDefault sync.Once
func Get_traverseWithIndexDefault() gopurs_runtime.Value {
	once_traverseWithIndexDefault.Do(func() {
		cache_traverseWithIndexDefault = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndexDefault(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box))
})
	})
	return cache_traverseWithIndexDefault
}

var cache_traverseWithIndex gopurs_runtime.Value
var once_traverseWithIndex sync.Once
func Get_traverseWithIndex() gopurs_runtime.Value {
	once_traverseWithIndex.Do(func() {
		cache_traverseWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverseWithIndex
}

var cache_traverseWithIndex__gopurs_runtime_Value_2726076659 gopurs_runtime.Value
var once_traverseWithIndex__gopurs_runtime_Value_2726076659 sync.Once
func Get_traverseWithIndex__gopurs_runtime_Value_2726076659() gopurs_runtime.Value {
	once_traverseWithIndex__gopurs_runtime_Value_2726076659.Do(func() {
		cache_traverseWithIndex__gopurs_runtime_Value_2726076659 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex__gopurs_runtime_Value_2726076659(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverseWithIndex__gopurs_runtime_Value_2726076659
}

var cache_traverseDefault gopurs_runtime.Value
var once_traverseDefault sync.Once
func Get_traverseDefault() gopurs_runtime.Value {
	once_traverseDefault.Do(func() {
		cache_traverseDefault = gopurs_runtime.Func3(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseDefault(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), f_2_box)
})
	})
	return cache_traverseDefault
}

var cache_traversableWithIndexTuple gopurs_runtime.Value
var once_traversableWithIndexTuple sync.Once
func Get_traversableWithIndexTuple() gopurs_runtime.Value {
	once_traversableWithIndexTuple.Do(func() {
		cache_traversableWithIndexTuple = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableTuple()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(f_2, pkg_Data_Unit.Get_unit(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
})
}))
	})
	return cache_traversableWithIndexTuple
}

var cache_traversableWithIndexProduct gopurs_runtime.Value
var once_traversableWithIndexProduct sync.Once
func Get_traversableWithIndexProduct() gopurs_runtime.Value {
	once_traversableWithIndexProduct.Do(func() {
		cache_traversableWithIndexProduct = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableWithIndexProduct(dictTraversableWithIndex_0_box)
})
	})
	return cache_traversableWithIndexProduct
}

var cache_traversableWithIndexMultiplicative gopurs_runtime.Value
var once_traversableWithIndexMultiplicative sync.Once
func Get_traversableWithIndexMultiplicative() gopurs_runtime.Value {
	once_traversableWithIndexMultiplicative.Do(func() {
		cache_traversableWithIndexMultiplicative = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableMultiplicative()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableMultiplicative(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}))
	})
	return cache_traversableWithIndexMultiplicative
}

var cache_traversableWithIndexMaybe gopurs_runtime.Value
var once_traversableWithIndexMaybe sync.Once
func Get_traversableWithIndexMaybe() gopurs_runtime.Value {
	once_traversableWithIndexMaybe.Do(func() {
		cache_traversableWithIndexMaybe = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableMaybe(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}))
	})
	return cache_traversableWithIndexMaybe
}

var cache_traversableWithIndexLast gopurs_runtime.Value
var once_traversableWithIndexLast sync.Once
func Get_traversableWithIndexLast() gopurs_runtime.Value {
	once_traversableWithIndexLast.Do(func() {
		cache_traversableWithIndexLast = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexLast()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexLast()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableLast()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableLast(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}))
	})
	return cache_traversableWithIndexLast
}

var cache_traversableWithIndexIdentity gopurs_runtime.Value
var once_traversableWithIndexIdentity sync.Once
func Get_traversableWithIndexIdentity() gopurs_runtime.Value {
	once_traversableWithIndexIdentity.Do(func() {
		cache_traversableWithIndexIdentity = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableIdentity()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply2(f_2, pkg_Data_Unit.Get_unit(), v_3))
})
})
}))
	})
	return cache_traversableWithIndexIdentity
}

var cache_traversableWithIndexFirst gopurs_runtime.Value
var once_traversableWithIndexFirst sync.Once
func Get_traversableWithIndexFirst() gopurs_runtime.Value {
	once_traversableWithIndexFirst.Do(func() {
		cache_traversableWithIndexFirst = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexFirst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexFirst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableFirst()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableFirst(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}))
	})
	return cache_traversableWithIndexFirst
}

var cache_traversableWithIndexEither gopurs_runtime.Value
var once_traversableWithIndexEither sync.Once
func Get_traversableWithIndexEither() gopurs_runtime.Value {
	once_traversableWithIndexEither.Do(func() {
		cache_traversableWithIndexEither = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableEither()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0})})
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Either.Get_Right(), gopurs_runtime.Apply2(v_2, pkg_Data_Unit.Get_unit(), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
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
	return cache_traversableWithIndexEither
}

var cache_traversableWithIndexDual gopurs_runtime.Value
var once_traversableWithIndexDual sync.Once
func Get_traversableWithIndexDual() gopurs_runtime.Value {
	once_traversableWithIndexDual.Do(func() {
		cache_traversableWithIndexDual = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableDual()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableDual(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}))
	})
	return cache_traversableWithIndexDual
}

var cache_traversableWithIndexDisj gopurs_runtime.Value
var once_traversableWithIndexDisj sync.Once
func Get_traversableWithIndexDisj() gopurs_runtime.Value {
	once_traversableWithIndexDisj.Do(func() {
		cache_traversableWithIndexDisj = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableDisj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableDisj(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}))
	})
	return cache_traversableWithIndexDisj
}

var cache_traversableWithIndexCoproduct gopurs_runtime.Value
var once_traversableWithIndexCoproduct sync.Once
func Get_traversableWithIndexCoproduct() gopurs_runtime.Value {
	once_traversableWithIndexCoproduct.Do(func() {
		cache_traversableWithIndexCoproduct = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableWithIndexCoproduct(dictTraversableWithIndex_0_box)
})
	})
	return cache_traversableWithIndexCoproduct
}

var cache_traversableWithIndexConst gopurs_runtime.Value
var once_traversableWithIndexConst sync.Once
func Get_traversableWithIndexConst() gopurs_runtime.Value {
	once_traversableWithIndexConst.Do(func() {
		cache_traversableWithIndexConst = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableConst()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v1_2)
})
})
}))
	})
	return cache_traversableWithIndexConst
}

var cache_traversableWithIndexConj gopurs_runtime.Value
var once_traversableWithIndexConj sync.Once
func Get_traversableWithIndexConj() gopurs_runtime.Value {
	once_traversableWithIndexConj.Do(func() {
		cache_traversableWithIndexConj = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableConj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableConj(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}))
	})
	return cache_traversableWithIndexConj
}

var cache_traversableWithIndexCompose gopurs_runtime.Value
var once_traversableWithIndexCompose sync.Once
func Get_traversableWithIndexCompose() gopurs_runtime.Value {
	once_traversableWithIndexCompose.Do(func() {
		cache_traversableWithIndexCompose = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableWithIndexCompose(dictTraversableWithIndex_0_box)
})
	})
	return cache_traversableWithIndexCompose
}

var cache_traversableWithIndexArray_Record_Row_traverseWithIndex_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Int_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_FunctorWithIndex0_Func_Record_Row__Any_Record_Row_mapWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Array_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Any_FoldableWithIndex1_Func_Record_Row__Any_Record_Row_foldrWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldlWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldMapWithIndex_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Int_Any_Any_Array_Any_Any_Foldable0_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Traversable2_Func_Record_Row__Any_Record_Row_traverse_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_sequence_ForAll_a_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Array_TypeApp_Any_Any_TypeApp_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Foldable1_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Any gopurs_runtime.Value
var once_traversableWithIndexArray_Record_Row_traverseWithIndex_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Int_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_FunctorWithIndex0_Func_Record_Row__Any_Record_Row_mapWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Array_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Any_FoldableWithIndex1_Func_Record_Row__Any_Record_Row_foldrWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldlWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldMapWithIndex_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Int_Any_Any_Array_Any_Any_Foldable0_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Traversable2_Func_Record_Row__Any_Record_Row_traverse_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_sequence_ForAll_a_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Array_TypeApp_Any_Any_TypeApp_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Foldable1_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Any sync.Once
func Get_traversableWithIndexArray_Record_Row_traverseWithIndex_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Int_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_FunctorWithIndex0_Func_Record_Row__Any_Record_Row_mapWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Array_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Any_FoldableWithIndex1_Func_Record_Row__Any_Record_Row_foldrWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldlWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldMapWithIndex_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Int_Any_Any_Array_Any_Any_Foldable0_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Traversable2_Func_Record_Row__Any_Record_Row_traverse_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_sequence_ForAll_a_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Array_TypeApp_Any_Any_TypeApp_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Foldable1_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Any() gopurs_runtime.Value {
	once_traversableWithIndexArray_Record_Row_traverseWithIndex_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Int_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_FunctorWithIndex0_Func_Record_Row__Any_Record_Row_mapWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Array_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Any_FoldableWithIndex1_Func_Record_Row__Any_Record_Row_foldrWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldlWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldMapWithIndex_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Int_Any_Any_Array_Any_Any_Foldable0_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Traversable2_Func_Record_Row__Any_Record_Row_traverse_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_sequence_ForAll_a_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Array_TypeApp_Any_Any_TypeApp_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Foldable1_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Any.Do(func() {
		cache_traversableWithIndexArray_Record_Row_traverseWithIndex_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Int_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_FunctorWithIndex0_Func_Record_Row__Any_Record_Row_mapWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Array_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Any_FoldableWithIndex1_Func_Record_Row__Any_Record_Row_foldrWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldlWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldMapWithIndex_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Int_Any_Any_Array_Any_Any_Foldable0_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Traversable2_Func_Record_Row__Any_Record_Row_traverse_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_sequence_ForAll_a_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Array_TypeApp_Any_Any_TypeApp_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Foldable1_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Any = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexArray__gopurs_runtime_Value_740253118()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexArray__gopurs_runtime_Value_490015842()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableArray__gopurs_runtime_Value_2643873085()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
FunctorWithIndex0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableWithIndexArray(), "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = FunctorWithIndex0_1_0
sequence1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableWithIndexArray(), "Traversable2"), gopurs_runtime.Value{}), "sequence"), dictApplicative_0)
_ = sequence1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(FunctorWithIndex0_1_0.V1, f_3)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_2_1, gopurs_runtime.Apply(__local_var_4_2, x_5))
})
})
}))
	})
	return cache_traversableWithIndexArray_Record_Row_traverseWithIndex_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Int_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_FunctorWithIndex0_Func_Record_Row__Any_Record_Row_mapWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Array_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Any_FoldableWithIndex1_Func_Record_Row__Any_Record_Row_foldrWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldlWithIndex_ForAll_a_b_Func_Func_Int_Any_Any_Any_Any_Array_Any_Any_foldMapWithIndex_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Int_Any_Any_Array_Any_Any_Foldable0_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Traversable2_Func_Record_Row__Any_Record_Row_traverse_ForAll_a_b_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Func_Any_TypeApp_Any_Any_Array_Any_TypeApp_Any_Array_Any_sequence_ForAll_a_m_ConstrainedType_Control_Applicative_Applicative_Any_Func_Array_TypeApp_Any_Any_TypeApp_Any_Array_Any_Functor0_Func_Record_Row__Any_Record_Row_map_ForAll_a_b_Func_Func_Any_Any_Array_Any_Array_Any_Any_Foldable1_Func_Record_Row__Any_Record_Row_foldr_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldl_ForAll_a_b_Func_Func_Any_Any_Any_Any_Array_Any_Any_foldMap_ForAll_a_m_ConstrainedType_Data_Monoid_Monoid_Any_Func_Func_Any_Any_Array_Any_Any_Any_Any_Any
}

var cache_traversableWithIndexArray gopurs_runtime.Value
var once_traversableWithIndexArray sync.Once
func Get_traversableWithIndexArray() gopurs_runtime.Value {
	once_traversableWithIndexArray.Do(func() {
		cache_traversableWithIndexArray = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableArray()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
FunctorWithIndex0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableWithIndexArray(), "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = FunctorWithIndex0_1_0
sequence1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableWithIndexArray(), "Traversable2"), gopurs_runtime.Value{}), "sequence"), dictApplicative_0)
_ = sequence1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(FunctorWithIndex0_1_0.V1, f_3)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_2_1, gopurs_runtime.Apply(__local_var_4_2, x_5))
})
})
}))
	})
	return cache_traversableWithIndexArray
}

var cache_traversableWithIndexArray__gopurs_runtime_Value_1681559805 gopurs_runtime.Value
var once_traversableWithIndexArray__gopurs_runtime_Value_1681559805 sync.Once
func Get_traversableWithIndexArray__gopurs_runtime_Value_1681559805() gopurs_runtime.Value {
	once_traversableWithIndexArray__gopurs_runtime_Value_1681559805.Do(func() {
		cache_traversableWithIndexArray__gopurs_runtime_Value_1681559805 = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexArray__gopurs_runtime_Value_740253118()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexArray__gopurs_runtime_Value_490015842()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableArray__gopurs_runtime_Value_2643873085()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
FunctorWithIndex0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableWithIndexArray(), "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = FunctorWithIndex0_1_0
sequence1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableWithIndexArray(), "Traversable2"), gopurs_runtime.Value{}), "sequence"), dictApplicative_0)
_ = sequence1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(FunctorWithIndex0_1_0.V1, f_3)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_2_1, gopurs_runtime.Apply(__local_var_4_2, x_5))
})
})
}))
	})
	return cache_traversableWithIndexArray__gopurs_runtime_Value_1681559805
}

var cache_traversableWithIndexApp gopurs_runtime.Value
var once_traversableWithIndexApp sync.Once
func Get_traversableWithIndexApp() gopurs_runtime.Value {
	once_traversableWithIndexApp.Do(func() {
		cache_traversableWithIndexApp = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableWithIndexApp(dictTraversableWithIndex_0_box)
})
	})
	return cache_traversableWithIndexApp
}

var cache_traversableWithIndexAdditive gopurs_runtime.Value
var once_traversableWithIndexAdditive sync.Once
func Get_traversableWithIndexAdditive() gopurs_runtime.Value {
	once_traversableWithIndexAdditive.Do(func() {
		cache_traversableWithIndexAdditive = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableAdditive()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableAdditive(), "traverse"), dictApplicative_0)
_ = traverse8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}))
	})
	return cache_traversableWithIndexAdditive
}

var cache_mapAccumRWithIndex gopurs_runtime.Value
var once_mapAccumRWithIndex sync.Once
func Get_mapAccumRWithIndex() gopurs_runtime.Value {
	once_mapAccumRWithIndex.Do(func() {
		cache_mapAccumRWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumRWithIndex(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumRWithIndex
}

var cache_mapAccumRWithIndex__gopurs_runtime_Value_142050190 gopurs_runtime.Value
var once_mapAccumRWithIndex__gopurs_runtime_Value_142050190 sync.Once
func Get_mapAccumRWithIndex__gopurs_runtime_Value_142050190() gopurs_runtime.Value {
	once_mapAccumRWithIndex__gopurs_runtime_Value_142050190.Do(func() {
		cache_mapAccumRWithIndex__gopurs_runtime_Value_142050190 = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumRWithIndex__gopurs_runtime_Value_142050190(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumRWithIndex__gopurs_runtime_Value_142050190
}

var cache_scanrWithIndex gopurs_runtime.Value
var once_scanrWithIndex sync.Once
func Get_scanrWithIndex() gopurs_runtime.Value {
	once_scanrWithIndex.Do(func() {
		cache_scanrWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanrWithIndex(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_scanrWithIndex
}

var cache_mapAccumLWithIndex gopurs_runtime.Value
var once_mapAccumLWithIndex sync.Once
func Get_mapAccumLWithIndex() gopurs_runtime.Value {
	once_mapAccumLWithIndex.Do(func() {
		cache_mapAccumLWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumLWithIndex(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumLWithIndex
}

var cache_mapAccumLWithIndex__gopurs_runtime_Value_142050190 gopurs_runtime.Value
var once_mapAccumLWithIndex__gopurs_runtime_Value_142050190 sync.Once
func Get_mapAccumLWithIndex__gopurs_runtime_Value_142050190() gopurs_runtime.Value {
	once_mapAccumLWithIndex__gopurs_runtime_Value_142050190.Do(func() {
		cache_mapAccumLWithIndex__gopurs_runtime_Value_142050190 = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumLWithIndex__gopurs_runtime_Value_142050190(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumLWithIndex__gopurs_runtime_Value_142050190
}

var cache_scanlWithIndex gopurs_runtime.Value
var once_scanlWithIndex sync.Once
func Get_scanlWithIndex() gopurs_runtime.Value {
	once_scanlWithIndex.Do(func() {
		cache_scanlWithIndex = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanlWithIndex(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_scanlWithIndex
}

var cache_forWithIndex gopurs_runtime.Value
var once_forWithIndex sync.Once
func Get_forWithIndex() gopurs_runtime.Value {
	once_forWithIndex.Do(func() {
		cache_forWithIndex = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictTraversableWithIndex_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_forWithIndex(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_1_box))
})
	})
	return cache_forWithIndex
}

type Constructor_TraversableWithIndex[T_i any, T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2078610234] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "FoldableWithIndex1": return c.V0
		case "FunctorWithIndex0": return c.V1
		case "Traversable2": return c.V2
		case "traverseWithIndex": return c.V3
		default: panic("Key not found in dictionary Constructor_TraversableWithIndex: " + key)
		}
	}
}


func Call_traverseWithIndexDefault(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
FunctorWithIndex0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(dictTraversableWithIndex_0.V1, gopurs_runtime.Value{}))
_ = FunctorWithIndex0_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictTraversableWithIndex_0.V2, gopurs_runtime.Value{}), "sequence"), dictApplicative_2)
_ = sequence1_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(FunctorWithIndex0_1_0.V1, f_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_3_1, gopurs_runtime.Apply(__local_var_5_2, x_6))
})
})
})
}

func Call_traverseWithIndex(dict_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverseWithIndex__gopurs_runtime_Value_2726076659(dict_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverseDefault(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(dictTraversableWithIndex_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return f_2
}))
}

func Call_traversableWithIndexProduct(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
functorWithIndexProduct_1_0 := gopurs_runtime.Apply(pkg_Data_FunctorWithIndex.Get_functorWithIndexProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = functorWithIndexProduct_1_0
foldableWithIndexProduct_2_1 := gopurs_runtime.Apply(pkg_Data_FoldableWithIndex.Get_foldableWithIndexProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexProduct_2_1
traversableProduct_3_2 := gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableProduct_3_2
return gopurs_runtime.Func(func(dictTraversableWithIndex1_4 gopurs_runtime.Value) gopurs_runtime.Value {
functorWithIndexProduct1_5_3 := gopurs_runtime.Apply(functorWithIndexProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = functorWithIndexProduct1_5_3
foldableWithIndexProduct1_6_4 := gopurs_runtime.Apply(foldableWithIndexProduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexProduct1_6_4
traversableProduct1_7_5 := gopurs_runtime.Apply(traversableProduct_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableProduct1_7_5
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexProduct1_6_4
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexProduct1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableProduct1_7_5
}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_9_6 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_9_6
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_9_6.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Apply0_9_6.V0, gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product.Get_product(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_8, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_12})})
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), dictApplicative_8, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_12})})
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1))
})
})
}))
})
}

func Call_traversableWithIndexCoproduct(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
functorWithIndexCoproduct_1_0 := gopurs_runtime.Apply(pkg_Data_FunctorWithIndex.Get_functorWithIndexCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = functorWithIndexCoproduct_1_0
foldableWithIndexCoproduct_2_1 := gopurs_runtime.Apply(pkg_Data_FoldableWithIndex.Get_foldableWithIndexCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCoproduct_2_1
traversableCoproduct_3_2 := gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCoproduct_3_2
return gopurs_runtime.Func(func(dictTraversableWithIndex1_4 gopurs_runtime.Value) gopurs_runtime.Value {
functorWithIndexCoproduct1_5_3 := gopurs_runtime.Apply(functorWithIndexCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = functorWithIndexCoproduct1_5_3
foldableWithIndexCoproduct1_6_4 := gopurs_runtime.Apply(foldableWithIndexCoproduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCoproduct1_6_4
traversableCoproduct1_7_5 := gopurs_runtime.Apply(traversableCoproduct_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCoproduct1_7_5
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexCoproduct1_6_4
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexCoproduct1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableCoproduct1_7_5
}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_9_6 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_9_6
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_11_7 := gopurs_runtime.Apply(Functor0_9_6.V0, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_11})}
}))
_ = __local_var_11_7
__local_var_12_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_8, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_12})})
}))
_ = __local_var_12_8
__local_var_13_9 := gopurs_runtime.Apply(Functor0_9_6.V0, gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_13})}
}))
_ = __local_var_13_9
__local_var_14_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), dictApplicative_8, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_14})})
}))
_ = __local_var_14_10
return gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t11 = gopurs_runtime.Apply(__local_var_11_7, gopurs_runtime.Apply(__local_var_12_8, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0))
goto end_branch_11
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t11 = gopurs_runtime.Apply(__local_var_13_9, gopurs_runtime.Apply(__local_var_14_10, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0))
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
})
})
}))
})
}

func Call_traversableWithIndexCompose(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
foldableWithIndexCompose_3_2 := gopurs_runtime.Apply(pkg_Data_FoldableWithIndex.Get_foldableWithIndexCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCompose_3_2
traversableCompose_4_3 := gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableCompose(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCompose_4_3
return gopurs_runtime.Func(func(dictTraversableWithIndex1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_6_5
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
functorCompose1_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "map"), f_8), v_9)
})
}))
_ = functorCompose1_7_6
functorWithIndexCompose1_6_4 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_7_6
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mapWithIndex"), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "mapWithIndex"), gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_10, b_11})})
}))
}), v_9)
})
}))
_ = functorWithIndexCompose1_6_4
foldableWithIndexCompose1_7_8 := gopurs_runtime.Apply(foldableWithIndexCompose_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCompose1_7_8
traversableCompose1_8_9 := gopurs_runtime.Apply(traversableCompose_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCompose1_8_9
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexCompose1_7_8
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexCompose1_6_4
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableCompose1_8_9
}), gopurs_runtime.Func(func(dictApplicative_9 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_10_10 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_9, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_10_10
traverseWithIndex2_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "traverseWithIndex"), dictApplicative_9)
_ = traverseWithIndex2_11_11
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_10_10.V0, pkg_Data_Functor_Compose.Get_Compose(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_9, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverseWithIndex2_11_11, gopurs_runtime.Func(func(b_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_12, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_14, b_15})})
}))
}), v_13))
})
})
}))
})
}

func Call_traversableWithIndexApp(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_1
functorApp_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = functorApp_2_2
functorWithIndexApp_1_0 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorApp_2_2
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mapWithIndex"), f_3, v_4)
})
}))
_ = functorWithIndexApp_1_0
foldableWithIndexApp_2_3 := gopurs_runtime.Apply(pkg_Data_FoldableWithIndex.Get_foldableWithIndexApp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexApp_2_3
traversableApp_3_4 := gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableApp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableApp_3_4
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexApp_2_3
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexApp_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableApp_3_4
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_5_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_5
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_5_5.V0, pkg_Data_Functor_App.Get_App(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_4, f_6, v_7))
})
})
}))
}

func Call_mapAccumRWithIndex(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversableWithIndex_0.V3, pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR(), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_mapAccumRWithIndex__gopurs_runtime_Value_142050190(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversableWithIndex_0.V3, pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR(), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_scanrWithIndex(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.RecordGet(Call_mapAccumRWithIndex(dictTraversableWithIndex_0, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_7_0 := gopurs_runtime.Apply3(f_1, i_4, a_6, b_5)
_ = b_prime_7_0
return gopurs_runtime.RecordDict2("accum", "value", b_prime_7_0, b_prime_7_0)
})
})
}), b0_2, xs_3), "value")
}

func Call_mapAccumLWithIndex(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversableWithIndex_0.V3, pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL(), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_mapAccumLWithIndex__gopurs_runtime_Value_142050190(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversableWithIndex_0.V3, pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL(), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_scanlWithIndex(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.RecordGet(Call_mapAccumLWithIndex(dictTraversableWithIndex_0, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_7_0 := gopurs_runtime.Apply3(f_1, i_4, b_5, a_6)
_ = b_prime_7_0
return gopurs_runtime.RecordDict2("accum", "value", b_prime_7_0, b_prime_7_0)
})
})
}), b0_2, xs_3), "value")
}

func Call_forWithIndex(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictTraversableWithIndex_1_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var dictTraversableWithIndex_1 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_1_loop
_ = dictTraversableWithIndex_1
__local_var_2_0 := gopurs_runtime.Apply(dictTraversableWithIndex_1.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)})
_ = __local_var_2_0
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_4, b_3)
})
})
}



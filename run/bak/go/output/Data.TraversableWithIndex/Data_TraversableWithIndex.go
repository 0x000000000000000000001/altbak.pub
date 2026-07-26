package Data_TraversableWithIndex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor_Product "gopurs/output/Data.Functor.Product"
	pkg_Data_Functor_Compose "gopurs/output/Data.Functor.Compose"
	pkg_Data_Functor_App "gopurs/output/Data.Functor.App"
	pkg_Data_Traversable_Accum_Internal "gopurs/output/Data.Traversable.Accum.Internal"
	unsafe "unsafe"
)

var cache_traverseWithIndexDefault gopurs_runtime.Value
var once_traverseWithIndexDefault sync.Once
func Get_traverseWithIndexDefault() gopurs_runtime.Value {
	once_traverseWithIndexDefault.Do(func() {
		cache_traverseWithIndexDefault = gopurs_runtime.Func2(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndexDefault(dictTraversableWithIndex_0_box, dictApplicative_1_box)
})
	})
	return cache_traverseWithIndexDefault
}

var cache_traverseWithIndex gopurs_runtime.Value
var once_traverseWithIndex sync.Once
func Get_traverseWithIndex() gopurs_runtime.Value {
	once_traverseWithIndex.Do(func() {
		cache_traverseWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex(dict_0_box)
})
	})
	return cache_traverseWithIndex
}

var cache_traverseDefault gopurs_runtime.Value
var once_traverseDefault sync.Once
func Get_traverseDefault() gopurs_runtime.Value {
	once_traverseDefault.Do(func() {
		cache_traverseDefault = gopurs_runtime.Func2(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseDefault(dictTraversableWithIndex_0_box, dictApplicative_1_box)
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
}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(f_1, pkg_Data_Unit.Get_unit(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))
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
}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply2(f_1, pkg_Data_Unit.Get_unit(), v_2))
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
}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0})})
goto end_branch_0
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Either.Get_Right(), gopurs_runtime.Apply2(v_1, pkg_Data_Unit.Get_unit(), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
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
}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v1_2)
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
sequence1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableWithIndexArray(), "Traversable2"), gopurs_runtime.Value{}), "sequence"), dictApplicative_0)
_ = sequence1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableWithIndexArray(), "FunctorWithIndex0"), gopurs_runtime.Value{}), "mapWithIndex"), f_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_1_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
})
}))
	})
	return cache_traversableWithIndexArray
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
		cache_mapAccumRWithIndex = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumRWithIndex(dictTraversableWithIndex_0_box)
})
	})
	return cache_mapAccumRWithIndex
}

var cache_scanrWithIndex gopurs_runtime.Value
var once_scanrWithIndex sync.Once
func Get_scanrWithIndex() gopurs_runtime.Value {
	once_scanrWithIndex.Do(func() {
		cache_scanrWithIndex = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanrWithIndex(dictTraversableWithIndex_0_box)
})
	})
	return cache_scanrWithIndex
}

var cache_mapAccumLWithIndex gopurs_runtime.Value
var once_mapAccumLWithIndex sync.Once
func Get_mapAccumLWithIndex() gopurs_runtime.Value {
	once_mapAccumLWithIndex.Do(func() {
		cache_mapAccumLWithIndex = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumLWithIndex(dictTraversableWithIndex_0_box)
})
	})
	return cache_mapAccumLWithIndex
}

var cache_scanlWithIndex gopurs_runtime.Value
var once_scanlWithIndex sync.Once
func Get_scanlWithIndex() gopurs_runtime.Value {
	once_scanlWithIndex.Do(func() {
		cache_scanlWithIndex = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanlWithIndex(dictTraversableWithIndex_0_box)
})
	})
	return cache_scanlWithIndex
}

var cache_forWithIndex gopurs_runtime.Value
var once_forWithIndex sync.Once
func Get_forWithIndex() gopurs_runtime.Value {
	once_forWithIndex.Do(func() {
		cache_forWithIndex = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictTraversableWithIndex_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_forWithIndex(dictApplicative_0_box, dictTraversableWithIndex_1_box)
})
	})
	return cache_forWithIndex
}

func Call_traverseWithIndexDefault(dictTraversableWithIndex_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
sequence1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}), "sequence"), dictApplicative_1)
_ = sequence1_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{}), "mapWithIndex"), f_3)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_2_0, gopurs_runtime.Apply(__local_var_4_1, x_5))
})
})
}

func Call_traverseWithIndex(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "traverseWithIndex")
}

func Call_traverseDefault(dictTraversableWithIndex_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
traverseWithIndex2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_1)
_ = traverseWithIndex2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverseWithIndex2_2_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return f_3
}))
})
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
__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_9_6
traverseWithIndex3_10_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_8)
_ = traverseWithIndex3_10_7
traverseWithIndex4_11_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), dictApplicative_8)
_ = traverseWithIndex4_11_8
return gopurs_runtime.Func2(func(f_12 gopurs_runtime.Value, v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_6, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_6, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product.Get_product(), gopurs_runtime.Apply2(traverseWithIndex3_10_7, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_12, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{x_14})})
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V0)), gopurs_runtime.Apply2(traverseWithIndex4_11_8, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_12, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{x_14})})
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_13.UnsafePtr).V1))
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
__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_6
traverseWithIndex3_10_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_8)
_ = traverseWithIndex3_10_7
traverseWithIndex4_11_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), dictApplicative_8)
_ = traverseWithIndex4_11_8
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_13_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_6, "map"), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{x_13})}
}))
_ = __local_var_13_9
__local_var_14_10 := gopurs_runtime.Apply(traverseWithIndex3_10_7, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_12, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{x_14})})
}))
_ = __local_var_14_10
__local_var_15_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_6, "map"), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{x_15})}
}))
_ = __local_var_15_11
__local_var_16_12 := gopurs_runtime.Apply(traverseWithIndex4_11_8, gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_12, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{x_16})})
}))
_ = __local_var_16_12
return gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(__local_var_13_9, gopurs_runtime.Apply(__local_var_14_10, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_17.UnsafePtr).V0))
goto end_branch_13
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t13 = gopurs_runtime.Apply(__local_var_15_11, gopurs_runtime.Apply(__local_var_16_12, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_17.UnsafePtr).V0))
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
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_4, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_5
functorCompose1_8_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_5, "map"), f_8), v_9)
}))
_ = functorCompose1_8_7
functorWithIndexCompose1_8_6 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_8_7
}), gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mapWithIndex"), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_4, "mapWithIndex"), gopurs_runtime.Func(func(b_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_9, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{x_11, b_12})})
}))
}), v_10)
}))
_ = functorWithIndexCompose1_8_6
foldableWithIndexCompose1_9_8 := gopurs_runtime.Apply(foldableWithIndexCompose_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexCompose1_9_8
traversableCompose1_10_9 := gopurs_runtime.Apply(traversableCompose_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableCompose1_10_9
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexCompose1_9_8
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexCompose1_8_6
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableCompose1_10_9
}), gopurs_runtime.Func(func(dictApplicative_11 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex3_12_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_11)
_ = traverseWithIndex3_12_10
traverseWithIndex4_13_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex1_5, "traverseWithIndex"), dictApplicative_11)
_ = traverseWithIndex4_13_11
return gopurs_runtime.Func2(func(f_14 gopurs_runtime.Value, v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_11, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Compose.Get_Compose(), gopurs_runtime.Apply2(traverseWithIndex3_12_10, gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverseWithIndex4_13_11, gopurs_runtime.Func(func(b_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_14, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{x_16, b_17})})
}))
}), v_15))
})
}))
})
}

func Call_traversableWithIndexApp(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorWithIndexApp_3_2 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_2_1
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mapWithIndex"), f_3, v_4)
}))
_ = functorWithIndexApp_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{})
_ = __local_var_4_3
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_5_4
foldableApp_6_6 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "foldMap"), dictMonoid_6)
}), gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, i_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_4, "foldl"), f_6, i_7, v_8)
}), gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, i_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_4, "foldr"), f_6, i_7, v_8)
}))
_ = foldableApp_6_6
foldableWithIndexApp_6_5 := gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableApp_6_6
}), gopurs_runtime.Func(func(dictMonoid_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "foldMapWithIndex"), dictMonoid_7)
}), gopurs_runtime.Func3(func(f_7 gopurs_runtime.Value, z_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_3, "foldlWithIndex"), f_7, z_8, v_9)
}), gopurs_runtime.Func3(func(f_7 gopurs_runtime.Value, z_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_3, "foldrWithIndex"), f_7, z_8, v_9)
}))
_ = foldableWithIndexApp_6_5
traversableApp_7_7 := gopurs_runtime.Apply(pkg_Data_Traversable.Get_traversableApp(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableApp_7_7
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexApp_6_5
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexApp_3_2
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableApp_7_7
}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex2_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_8)
_ = traverseWithIndex2_9_8
return gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_App.Get_App(), gopurs_runtime.Apply2(traverseWithIndex2_9_8, f_10, v_11))
})
}))
}

func Call_mapAccumRWithIndex(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
traverseWithIndex1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR())
_ = traverseWithIndex1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, s0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(traverseWithIndex1_1_0, gopurs_runtime.Func3(func(i_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_2, i_5, s_7, a_6)
}), xs_4, s0_3)
})
}

func Call_scanrWithIndex(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
mapAccumRWithIndex1_1_0 := Call_mapAccumRWithIndex(dictTraversableWithIndex_0)
_ = mapAccumRWithIndex1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(mapAccumRWithIndex1_1_0, gopurs_runtime.Func3(func(i_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_8_1 := gopurs_runtime.Apply3(f_2, i_5, a_7, b_6)
_ = b_prime_8_1
return gopurs_runtime.RecordDict2("accum", "value", b_prime_8_1, b_prime_8_1)
}), b0_3, xs_4), "value")
})
}

func Call_mapAccumLWithIndex(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
traverseWithIndex1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL())
_ = traverseWithIndex1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, s0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(traverseWithIndex1_1_0, gopurs_runtime.Func3(func(i_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_2, i_5, s_7, a_6)
}), xs_4, s0_3)
})
}

func Call_scanlWithIndex(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
mapAccumLWithIndex1_1_0 := Call_mapAccumLWithIndex(dictTraversableWithIndex_0)
_ = mapAccumLWithIndex1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(mapAccumLWithIndex1_1_0, gopurs_runtime.Func3(func(i_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_8_1 := gopurs_runtime.Apply3(f_2, i_5, b_6, a_7)
_ = b_prime_8_1
return gopurs_runtime.RecordDict2("accum", "value", b_prime_8_1, b_prime_8_1)
}), b0_3, xs_4), "value")
})
}

func Call_forWithIndex(dictApplicative_0_loop gopurs_runtime.Value, dictTraversableWithIndex_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var dictTraversableWithIndex_1 gopurs_runtime.Value = dictTraversableWithIndex_1_loop
_ = dictTraversableWithIndex_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_1, "traverseWithIndex"), dictApplicative_0)
_ = __local_var_2_0
return gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_4, b_3)
})
}



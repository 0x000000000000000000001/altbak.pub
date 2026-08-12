package Data_TraversableWithIndex

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Functor_App "gopurs/output/Data.Functor.App"
	pkg_Data_Functor_Compose "gopurs/output/Data.Functor.Compose"
	pkg_Data_Functor_Product "gopurs/output/Data.Functor.Product"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Maybe_First "gopurs/output/Data.Maybe.First"
	pkg_Data_Maybe_Last "gopurs/output/Data.Maybe.Last"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Monoid_Additive "gopurs/output/Data.Monoid.Additive"
	pkg_Data_Monoid_Conj "gopurs/output/Data.Monoid.Conj"
	pkg_Data_Monoid_Disj "gopurs/output/Data.Monoid.Disj"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
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

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__1475749520 gopurs_runtime.Value
var once_pure__1475749520 sync.Once
func Get_pure__1475749520() gopurs_runtime.Value {
	once_pure__1475749520.Do(func() {
		cache_pure__1475749520 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1475749520(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__1475749520
}

var cache_pure__154576880 gopurs_runtime.Value
var once_pure__154576880 sync.Once
func Get_pure__154576880() gopurs_runtime.Value {
	once_pure__154576880.Do(func() {
		cache_pure__154576880 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__154576880(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__154576880
}

var cache_pure__1304937360 gopurs_runtime.Value
var once_pure__1304937360 sync.Once
func Get_pure__1304937360() gopurs_runtime.Value {
	once_pure__1304937360.Do(func() {
		cache_pure__1304937360 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1304937360(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__1304937360
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_lift2__2762258480 gopurs_runtime.Value
var once_lift2__2762258480 sync.Once
func Get_lift2__2762258480() gopurs_runtime.Value {
	once_lift2__2762258480.Do(func() {
		cache_lift2__2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2762258480(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2762258480
}

var cache_lift2__470376976 gopurs_runtime.Value
var once_lift2__470376976 sync.Once
func Get_lift2__470376976() gopurs_runtime.Value {
	once_lift2__470376976.Do(func() {
		cache_lift2__470376976 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__470376976(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__470376976
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_functorConst__1992455793 gopurs_runtime.Value
var once_functorConst__1992455793 sync.Once
func Get_functorConst__1992455793() gopurs_runtime.Value {
	once_functorConst__1992455793.Do(func() {
		cache_functorConst__1992455793 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return m_1
})
}))
	})
	return cache_functorConst__1992455793
}

var cache_functorEither__1771778897 gopurs_runtime.Value
var once_functorEither__1771778897 sync.Once
func Get_functorEither__1771778897() gopurs_runtime.Value {
	once_functorEither__1771778897.Do(func() {
		cache_functorEither__1771778897 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
}))
	})
	return cache_functorEither__1771778897
}

var cache_foldMap__4098395794 gopurs_runtime.Value
var once_foldMap__4098395794 sync.Once
func Get_foldMap__4098395794() gopurs_runtime.Value {
	once_foldMap__4098395794.Do(func() {
		cache_foldMap__4098395794 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__4098395794(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap__4098395794
}

var cache_foldMap__4073832436 gopurs_runtime.Value
var once_foldMap__4073832436 sync.Once
func Get_foldMap__4073832436() gopurs_runtime.Value {
	once_foldMap__4073832436.Do(func() {
		cache_foldMap__4073832436 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldMap")
	})
	return cache_foldMap__4073832436
}

var cache_foldableAdditive__1841171440 gopurs_runtime.Value
var once_foldableAdditive__1841171440 sync.Once
func Get_foldableAdditive__1841171440() gopurs_runtime.Value {
	once_foldableAdditive__1841171440.Do(func() {
		cache_foldableAdditive__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableAdditive__1841171440
}

var cache_foldableArray__2950015754 gopurs_runtime.Value
var once_foldableArray__2950015754 sync.Once
func Get_foldableArray__2950015754() gopurs_runtime.Value {
	once_foldableArray__2950015754.Do(func() {
		cache_foldableArray__2950015754 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), pkg_Data_Foldable.Get_foldlArray(), pkg_Data_Foldable.Get_foldrArray())
	})
	return cache_foldableArray__2950015754
}

var cache_foldableArray__3859409398 gopurs_runtime.Value
var once_foldableArray__3859409398 sync.Once
func Get_foldableArray__3859409398() gopurs_runtime.Value {
	once_foldableArray__3859409398.Do(func() {
		cache_foldableArray__3859409398 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), pkg_Data_Foldable.Get_foldlArray(), pkg_Data_Foldable.Get_foldrArray())
	})
	return cache_foldableArray__3859409398
}

var cache_foldableConj__1841171440 gopurs_runtime.Value
var once_foldableConj__1841171440 sync.Once
func Get_foldableConj__1841171440() gopurs_runtime.Value {
	once_foldableConj__1841171440.Do(func() {
		cache_foldableConj__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableConj__1841171440
}

var cache_foldableConst__943899702 gopurs_runtime.Value
var once_foldableConst__943899702 sync.Once
func Get_foldableConst__943899702() gopurs_runtime.Value {
	once_foldableConst__943899702.Do(func() {
		cache_foldableConst__943899702 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableConst__943899702
}

var cache_foldableDisj__1841171440 gopurs_runtime.Value
var once_foldableDisj__1841171440 sync.Once
func Get_foldableDisj__1841171440() gopurs_runtime.Value {
	once_foldableDisj__1841171440.Do(func() {
		cache_foldableDisj__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableDisj__1841171440
}

var cache_foldableDual__1841171440 gopurs_runtime.Value
var once_foldableDual__1841171440 sync.Once
func Get_foldableDual__1841171440() gopurs_runtime.Value {
	once_foldableDual__1841171440.Do(func() {
		cache_foldableDual__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableDual__1841171440
}

var cache_foldableEither__1622911640 gopurs_runtime.Value
var once_foldableEither__1622911640 sync.Once
func Get_foldableEither__1622911640() gopurs_runtime.Value {
	once_foldableEither__1622911640.Do(func() {
		cache_foldableEither__1622911640 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableEither__1622911640
}

var cache_foldableFirst__2831137713 gopurs_runtime.Value
var once_foldableFirst__2831137713 sync.Once
func Get_foldableFirst__2831137713() gopurs_runtime.Value {
	once_foldableFirst__2831137713.Do(func() {
		cache_foldableFirst__2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldl"), f_0, z_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldr"), f_0, z_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))})
})
})
}))
	})
	return cache_foldableFirst__2831137713
}

var cache_foldableIdentity__1841171440 gopurs_runtime.Value
var once_foldableIdentity__1841171440 sync.Once
func Get_foldableIdentity__1841171440() gopurs_runtime.Value {
	once_foldableIdentity__1841171440.Do(func() {
		cache_foldableIdentity__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableIdentity__1841171440
}

var cache_foldableLast__2831137713 gopurs_runtime.Value
var once_foldableLast__2831137713 sync.Once
func Get_foldableLast__2831137713() gopurs_runtime.Value {
	once_foldableLast__2831137713.Do(func() {
		cache_foldableLast__2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0))}, f_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldl"), f_0, z_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldr"), f_0, z_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))})
})
})
}))
	})
	return cache_foldableLast__2831137713
}

var cache_foldableMaybe__3653484922 gopurs_runtime.Value
var once_foldableMaybe__3653484922 sync.Once
func Get_foldableMaybe__3653484922() gopurs_runtime.Value {
	once_foldableMaybe__3653484922.Do(func() {
		cache_foldableMaybe__3653484922 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableMaybe__3653484922
}

var cache_foldableMaybe__2831137713 gopurs_runtime.Value
var once_foldableMaybe__2831137713 sync.Once
func Get_foldableMaybe__2831137713() gopurs_runtime.Value {
	once_foldableMaybe__2831137713.Do(func() {
		cache_foldableMaybe__2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableMaybe__2831137713
}

var cache_foldableMultiplicative__1841171440 gopurs_runtime.Value
var once_foldableMultiplicative__1841171440 sync.Once
func Get_foldableMultiplicative__1841171440() gopurs_runtime.Value {
	once_foldableMultiplicative__1841171440.Do(func() {
		cache_foldableMultiplicative__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableMultiplicative__1841171440
}

var cache_foldableTuple__1455669080 gopurs_runtime.Value
var once_foldableTuple__1455669080 sync.Once
func Get_foldableTuple__1455669080() gopurs_runtime.Value {
	once_foldableTuple__1455669080.Do(func() {
		cache_foldableTuple__1455669080 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableTuple__1455669080
}

var cache_foldl__2151204251 gopurs_runtime.Value
var once_foldl__2151204251 sync.Once
func Get_foldl__2151204251() gopurs_runtime.Value {
	once_foldl__2151204251.Do(func() {
		cache_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2151204251
}

var cache_foldl__3016550397 gopurs_runtime.Value
var once_foldl__3016550397 sync.Once
func Get_foldl__3016550397() gopurs_runtime.Value {
	once_foldl__3016550397.Do(func() {
		cache_foldl__3016550397 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldl")
	})
	return cache_foldl__3016550397
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldr__3016550397 gopurs_runtime.Value
var once_foldr__3016550397 sync.Once
func Get_foldr__3016550397() gopurs_runtime.Value {
	once_foldr__3016550397.Do(func() {
		cache_foldr__3016550397 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldr")
	})
	return cache_foldr__3016550397
}

var cache_foldableWithIndexAdditive__3548846699 gopurs_runtime.Value
var once_foldableWithIndexAdditive__3548846699 sync.Once
func Get_foldableWithIndexAdditive__3548846699() gopurs_runtime.Value {
	once_foldableWithIndexAdditive__3548846699.Do(func() {
		cache_foldableWithIndexAdditive__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableAdditive()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableAdditive(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableAdditive(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableAdditive(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexAdditive__3548846699
}

var cache_foldableWithIndexArray__740253118 gopurs_runtime.Value
var once_foldableWithIndexArray__740253118 sync.Once
func Get_foldableWithIndexArray__740253118() gopurs_runtime.Value {
	once_foldableWithIndexArray__740253118.Do(func() {
		cache_foldableWithIndexArray__740253118 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_FoldableWithIndex.Get_foldableWithIndexArray(), "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply2(f_3, i_4, x_5), acc_6)
})
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal), y_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
}), z_1)
_ = __local_var_2_2
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex"), pkg_Data_Tuple.Get_Tuple())
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(__local_var_3_3, x_4))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_3_5
__local_var_4_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_4_6
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_3_5.IntVal), __local_var_4_6, y_5)
})
}), z_1)
_ = __local_var_2_4
__local_var_3_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex"), pkg_Data_Tuple.Get_Tuple())
_ = __local_var_3_7
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(__local_var_3_7, x_4))
})
})
}))
	})
	return cache_foldableWithIndexArray__740253118
}

var cache_foldableWithIndexConj__3548846699 gopurs_runtime.Value
var once_foldableWithIndexConj__3548846699 sync.Once
func Get_foldableWithIndexConj__3548846699() gopurs_runtime.Value {
	once_foldableWithIndexConj__3548846699.Do(func() {
		cache_foldableWithIndexConj__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConj()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableConj(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableConj(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableConj(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexConj__3548846699
}

var cache_foldableWithIndexConst__1906290739 gopurs_runtime.Value
var once_foldableWithIndexConst__1906290739 sync.Once
func Get_foldableWithIndexConst__1906290739() gopurs_runtime.Value {
	once_foldableWithIndexConst__1906290739.Do(func() {
		cache_foldableWithIndexConst__1906290739 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConst()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_foldableWithIndexConst__1906290739
}

var cache_foldableWithIndexDisj__3548846699 gopurs_runtime.Value
var once_foldableWithIndexDisj__3548846699 sync.Once
func Get_foldableWithIndexDisj__3548846699() gopurs_runtime.Value {
	once_foldableWithIndexDisj__3548846699.Do(func() {
		cache_foldableWithIndexDisj__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDisj()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDisj(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDisj(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDisj(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexDisj__3548846699
}

var cache_foldableWithIndexDual__3548846699 gopurs_runtime.Value
var once_foldableWithIndexDual__3548846699 sync.Once
func Get_foldableWithIndexDual__3548846699() gopurs_runtime.Value {
	once_foldableWithIndexDual__3548846699.Do(func() {
		cache_foldableWithIndexDual__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDual(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDual(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDual(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexDual__3548846699
}

var cache_foldableWithIndexEither__162584107 gopurs_runtime.Value
var once_foldableWithIndexEither__162584107 sync.Once
func Get_foldableWithIndexEither__162584107() gopurs_runtime.Value {
	once_foldableWithIndexEither__162584107.Do(func() {
		cache_foldableWithIndexEither__162584107 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableEither()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply2(v_2, pkg_Data_Unit.Get_unit(), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
__t2 = gopurs_runtime.Apply3(v_0, pkg_Data_Unit.Get_unit(), v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
__t3 = gopurs_runtime.Apply3(v_0, pkg_Data_Unit.Get_unit(), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
	return cache_foldableWithIndexEither__162584107
}

var cache_foldableWithIndexFirst__1642852683 gopurs_runtime.Value
var once_foldableWithIndexFirst__1642852683 sync.Once
func Get_foldableWithIndexFirst__1642852683() gopurs_runtime.Value {
	once_foldableWithIndexFirst__1642852683.Do(func() {
		cache_foldableWithIndexFirst__1642852683 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableFirst()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableFirst(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableFirst(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableFirst(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexFirst__1642852683
}

var cache_foldableWithIndexIdentity__3548846699 gopurs_runtime.Value
var once_foldableWithIndexIdentity__3548846699 sync.Once
func Get_foldableWithIndexIdentity__3548846699() gopurs_runtime.Value {
	once_foldableWithIndexIdentity__3548846699.Do(func() {
		cache_foldableWithIndexIdentity__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, pkg_Data_Unit.Get_unit(), v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, pkg_Data_Unit.Get_unit(), z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, pkg_Data_Unit.Get_unit(), v_2, z_1)
})
})
}))
	})
	return cache_foldableWithIndexIdentity__3548846699
}

var cache_foldableWithIndexLast__1642852683 gopurs_runtime.Value
var once_foldableWithIndexLast__1642852683 sync.Once
func Get_foldableWithIndexLast__1642852683() gopurs_runtime.Value {
	once_foldableWithIndexLast__1642852683.Do(func() {
		cache_foldableWithIndexLast__1642852683 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableLast()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableLast(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableLast(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableLast(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexLast__1642852683
}

var cache_foldableWithIndexMaybe__1642852683 gopurs_runtime.Value
var once_foldableWithIndexMaybe__1642852683 sync.Once
func Get_foldableWithIndexMaybe__1642852683() gopurs_runtime.Value {
	once_foldableWithIndexMaybe__1642852683.Do(func() {
		cache_foldableWithIndexMaybe__1642852683 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMaybe()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexMaybe__1642852683
}

var cache_foldableWithIndexMultiplicative__3548846699 gopurs_runtime.Value
var once_foldableWithIndexMultiplicative__3548846699 sync.Once
func Get_foldableWithIndexMultiplicative__3548846699() gopurs_runtime.Value {
	once_foldableWithIndexMultiplicative__3548846699.Do(func() {
		cache_foldableWithIndexMultiplicative__3548846699 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMultiplicative(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMultiplicative(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMultiplicative(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexMultiplicative__3548846699
}

var cache_foldableWithIndexTuple__4170851755 gopurs_runtime.Value
var once_foldableWithIndexTuple__4170851755 sync.Once
func Get_foldableWithIndexTuple__4170851755() gopurs_runtime.Value {
	once_foldableWithIndexTuple__4170851755.Do(func() {
		cache_foldableWithIndexTuple__4170851755 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableTuple()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, pkg_Data_Unit.Get_unit(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, pkg_Data_Unit.Get_unit(), z_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, pkg_Data_Unit.Get_unit(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
})
})
}))
	})
	return cache_foldableWithIndexTuple__4170851755
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_const__3858583060 gopurs_runtime.Value
var once_const__3858583060 sync.Once
func Get_const__3858583060() gopurs_runtime.Value {
	once_const__3858583060.Do(func() {
		cache_const__3858583060 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__3858583060(a_0_box, v_1_box)
})
	})
	return cache_const__3858583060
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__1931071680 gopurs_runtime.Value
var once_flip__1931071680 sync.Once
func Get_flip__1931071680() gopurs_runtime.Value {
	once_flip__1931071680.Do(func() {
		cache_flip__1931071680 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1931071680(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1931071680
}

var cache_coproduct__413515331 gopurs_runtime.Value
var once_coproduct__413515331 sync.Once
func Get_coproduct__413515331() gopurs_runtime.Value {
	once_coproduct__413515331.Do(func() {
		cache_coproduct__413515331 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct__413515331(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_coproduct__413515331
}

var cache_coproduct__1706612365 gopurs_runtime.Value
var once_coproduct__1706612365 sync.Once
func Get_coproduct__1706612365() gopurs_runtime.Value {
	once_coproduct__1706612365.Do(func() {
		cache_coproduct__1706612365 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct__1706612365(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_coproduct__1706612365
}

var cache_product__2764631669 gopurs_runtime.Value
var once_product__2764631669 sync.Once
func Get_product__2764631669() gopurs_runtime.Value {
	once_product__2764631669.Do(func() {
		cache_product__2764631669 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_product__2764631669(fa_0_box, ga_1_box))}
})
	})
	return cache_product__2764631669
}

var cache_functorArray__361387505 gopurs_runtime.Value
var once_functorArray__361387505 sync.Once
func Get_functorArray__361387505() gopurs_runtime.Value {
	once_functorArray__361387505.Do(func() {
		cache_functorArray__361387505 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__361387505
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__328307316 gopurs_runtime.Value
var once_map__328307316 sync.Once
func Get_map__328307316() gopurs_runtime.Value {
	once_map__328307316.Do(func() {
		cache_map__328307316 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__328307316(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__328307316
}

var cache_map__2701008148 gopurs_runtime.Value
var once_map__2701008148 sync.Once
func Get_map__2701008148() gopurs_runtime.Value {
	once_map__2701008148.Do(func() {
		cache_map__2701008148 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2701008148(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2701008148
}

var cache_map__1762802164 gopurs_runtime.Value
var once_map__1762802164 sync.Once
func Get_map__1762802164() gopurs_runtime.Value {
	once_map__1762802164.Do(func() {
		cache_map__1762802164 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1762802164(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1762802164
}

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_map__3659954292 gopurs_runtime.Value
var once_map__3659954292 sync.Once
func Get_map__3659954292() gopurs_runtime.Value {
	once_map__3659954292.Do(func() {
		cache_map__3659954292 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3659954292(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3659954292
}

var cache_map__2251722612 gopurs_runtime.Value
var once_map__2251722612 sync.Once
func Get_map__2251722612() gopurs_runtime.Value {
	once_map__2251722612.Do(func() {
		cache_map__2251722612 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2251722612(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2251722612
}

var cache_map__1337616244 gopurs_runtime.Value
var once_map__1337616244 sync.Once
func Get_map__1337616244() gopurs_runtime.Value {
	once_map__1337616244.Do(func() {
		cache_map__1337616244 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1337616244(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1337616244
}

var cache_functorWithIndexAdditive__339744599 gopurs_runtime.Value
var once_functorWithIndexAdditive__339744599 sync.Once
func Get_functorWithIndexAdditive__339744599() gopurs_runtime.Value {
	once_functorWithIndexAdditive__339744599.Do(func() {
		cache_functorWithIndexAdditive__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Additive.Get_functorAdditive()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid_Additive.Get_functorAdditive(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexAdditive__339744599
}

var cache_functorWithIndexArray__3811533158 gopurs_runtime.Value
var once_functorWithIndexArray__3811533158 sync.Once
func Get_functorWithIndexArray__3811533158() gopurs_runtime.Value {
	once_functorWithIndexArray__3811533158.Do(func() {
		cache_functorWithIndexArray__3811533158 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), pkg_Data_FunctorWithIndex.Get_mapWithIndexArray())
	})
	return cache_functorWithIndexArray__3811533158
}

var cache_functorWithIndexArray__490015842 gopurs_runtime.Value
var once_functorWithIndexArray__490015842 sync.Once
func Get_functorWithIndexArray__490015842() gopurs_runtime.Value {
	once_functorWithIndexArray__490015842.Do(func() {
		cache_functorWithIndexArray__490015842 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), pkg_Data_FunctorWithIndex.Get_mapWithIndexArray())
	})
	return cache_functorWithIndexArray__490015842
}

var cache_functorWithIndexConj__339744599 gopurs_runtime.Value
var once_functorWithIndexConj__339744599 sync.Once
func Get_functorWithIndexConj__339744599() gopurs_runtime.Value {
	once_functorWithIndexConj__339744599.Do(func() {
		cache_functorWithIndexConj__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Conj.Get_functorConj()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid_Conj.Get_functorConj(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexConj__339744599
}

var cache_functorWithIndexConst__3232336655 gopurs_runtime.Value
var once_functorWithIndexConst__3232336655 sync.Once
func Get_functorWithIndexConst__3232336655() gopurs_runtime.Value {
	once_functorWithIndexConst__3232336655.Do(func() {
		cache_functorWithIndexConst__3232336655 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Const.Get_functorConst()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_functorWithIndexConst__3232336655
}

var cache_functorWithIndexDisj__339744599 gopurs_runtime.Value
var once_functorWithIndexDisj__339744599 sync.Once
func Get_functorWithIndexDisj__339744599() gopurs_runtime.Value {
	once_functorWithIndexDisj__339744599.Do(func() {
		cache_functorWithIndexDisj__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Disj.Get_functorDisj()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid_Disj.Get_functorDisj(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexDisj__339744599
}

var cache_functorWithIndexDual__339744599 gopurs_runtime.Value
var once_functorWithIndexDual__339744599 sync.Once
func Get_functorWithIndexDual__339744599() gopurs_runtime.Value {
	once_functorWithIndexDual__339744599.Do(func() {
		cache_functorWithIndexDual__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Dual.Get_functorDual()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid_Dual.Get_functorDual(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexDual__339744599
}

var cache_functorWithIndexEither__1698960599 gopurs_runtime.Value
var once_functorWithIndexEither__1698960599 sync.Once
func Get_functorWithIndexEither__1698960599() gopurs_runtime.Value {
	once_functorWithIndexEither__1698960599.Do(func() {
		cache_functorWithIndexEither__1698960599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_functorEither()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Either.Get_functorEither(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexEither__1698960599
}

var cache_functorWithIndexFirst__2982254679 gopurs_runtime.Value
var once_functorWithIndexFirst__2982254679 sync.Once
func Get_functorWithIndexFirst__2982254679() gopurs_runtime.Value {
	once_functorWithIndexFirst__2982254679.Do(func() {
		cache_functorWithIndexFirst__2982254679 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexFirst__2982254679
}

var cache_functorWithIndexIdentity__339744599 gopurs_runtime.Value
var once_functorWithIndexIdentity__339744599 sync.Once
func Get_functorWithIndexIdentity__339744599() gopurs_runtime.Value {
	once_functorWithIndexIdentity__339744599.Do(func() {
		cache_functorWithIndexIdentity__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, pkg_Data_Unit.Get_unit(), v_1)
})
}))
	})
	return cache_functorWithIndexIdentity__339744599
}

var cache_functorWithIndexLast__2982254679 gopurs_runtime.Value
var once_functorWithIndexLast__2982254679 sync.Once
func Get_functorWithIndexLast__2982254679() gopurs_runtime.Value {
	once_functorWithIndexLast__2982254679.Do(func() {
		cache_functorWithIndexLast__2982254679 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexLast__2982254679
}

var cache_functorWithIndexMaybe__2982254679 gopurs_runtime.Value
var once_functorWithIndexMaybe__2982254679 sync.Once
func Get_functorWithIndexMaybe__2982254679() gopurs_runtime.Value {
	once_functorWithIndexMaybe__2982254679.Do(func() {
		cache_functorWithIndexMaybe__2982254679 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexMaybe__2982254679
}

var cache_functorWithIndexMultiplicative__339744599 gopurs_runtime.Value
var once_functorWithIndexMultiplicative__339744599 sync.Once
func Get_functorWithIndexMultiplicative__339744599() gopurs_runtime.Value {
	once_functorWithIndexMultiplicative__339744599.Do(func() {
		cache_functorWithIndexMultiplicative__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Multiplicative.Get_functorMultiplicative()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid_Multiplicative.Get_functorMultiplicative(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexMultiplicative__339744599
}

var cache_functorWithIndexTuple__1273126103 gopurs_runtime.Value
var once_functorWithIndexTuple__1273126103 sync.Once
func Get_functorWithIndexTuple__1273126103() gopurs_runtime.Value {
	once_functorWithIndexTuple__1273126103.Do(func() {
		cache_functorWithIndexTuple__1273126103 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexTuple__1273126103
}

var cache_mapWithIndex__835054498 gopurs_runtime.Value
var once_mapWithIndex__835054498 sync.Once
func Get_mapWithIndex__835054498() gopurs_runtime.Value {
	once_mapWithIndex__835054498.Do(func() {
		cache_mapWithIndex__835054498 = gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex")
	})
	return cache_mapWithIndex__835054498
}

var cache_mapWithIndex__55256674 gopurs_runtime.Value
var once_mapWithIndex__55256674 sync.Once
func Get_mapWithIndex__55256674() gopurs_runtime.Value {
	once_mapWithIndex__55256674.Do(func() {
		cache_mapWithIndex__55256674 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex__55256674(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mapWithIndex__55256674
}

var cache_mapWithIndex__2239747170 gopurs_runtime.Value
var once_mapWithIndex__2239747170 sync.Once
func Get_mapWithIndex__2239747170() gopurs_runtime.Value {
	once_mapWithIndex__2239747170.Do(func() {
		cache_mapWithIndex__2239747170 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex__2239747170(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mapWithIndex__2239747170
}

var cache_functorIdentity__943655089 gopurs_runtime.Value
var once_functorIdentity__943655089 sync.Once
func Get_functorIdentity__943655089() gopurs_runtime.Value {
	once_functorIdentity__943655089.Do(func() {
		cache_functorIdentity__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorIdentity__943655089
}

var cache_functorFirst__2097654001 gopurs_runtime.Value
var once_functorFirst__2097654001 sync.Once
func Get_functorFirst__2097654001() gopurs_runtime.Value {
	once_functorFirst__2097654001.Do(func() {
		cache_functorFirst__2097654001 = pkg_Data_Maybe.Get_functorMaybe()
	})
	return cache_functorFirst__2097654001
}

var cache_functorLast__2097654001 gopurs_runtime.Value
var once_functorLast__2097654001 sync.Once
func Get_functorLast__2097654001() gopurs_runtime.Value {
	once_functorLast__2097654001.Do(func() {
		cache_functorLast__2097654001 = pkg_Data_Maybe.Get_functorMaybe()
	})
	return cache_functorLast__2097654001
}

var cache_functorMaybe__2097654001 gopurs_runtime.Value
var once_functorMaybe__2097654001 sync.Once
func Get_functorMaybe__2097654001() gopurs_runtime.Value {
	once_functorMaybe__2097654001.Do(func() {
		cache_functorMaybe__2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2097654001
}

var cache_functorAdditive__943655089 gopurs_runtime.Value
var once_functorAdditive__943655089 sync.Once
func Get_functorAdditive__943655089() gopurs_runtime.Value {
	once_functorAdditive__943655089.Do(func() {
		cache_functorAdditive__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorAdditive__943655089
}

var cache_functorConj__943655089 gopurs_runtime.Value
var once_functorConj__943655089 sync.Once
func Get_functorConj__943655089() gopurs_runtime.Value {
	once_functorConj__943655089.Do(func() {
		cache_functorConj__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorConj__943655089
}

var cache_functorDisj__943655089 gopurs_runtime.Value
var once_functorDisj__943655089 sync.Once
func Get_functorDisj__943655089() gopurs_runtime.Value {
	once_functorDisj__943655089.Do(func() {
		cache_functorDisj__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorDisj__943655089
}

var cache_functorDual__943655089 gopurs_runtime.Value
var once_functorDual__943655089 sync.Once
func Get_functorDual__943655089() gopurs_runtime.Value {
	once_functorDual__943655089.Do(func() {
		cache_functorDual__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorDual__943655089
}

var cache_functorMultiplicative__943655089 gopurs_runtime.Value
var once_functorMultiplicative__943655089 sync.Once
func Get_functorMultiplicative__943655089() gopurs_runtime.Value {
	once_functorMultiplicative__943655089.Do(func() {
		cache_functorMultiplicative__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorMultiplicative__943655089
}

var cache_applicativeStateL__2039640491 gopurs_runtime.Value
var once_applicativeStateL__2039640491 sync.Once
func Get_applicativeStateL__2039640491() gopurs_runtime.Value {
	once_applicativeStateL__2039640491.Do(func() {
		cache_applicativeStateL__2039640491 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable_Accum_Internal.Get_applyStateL()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
}))
	})
	return cache_applicativeStateL__2039640491
}

var cache_applicativeStateR__2039640491 gopurs_runtime.Value
var once_applicativeStateR__2039640491 sync.Once
func Get_applicativeStateR__2039640491() gopurs_runtime.Value {
	once_applicativeStateR__2039640491.Do(func() {
		cache_applicativeStateR__2039640491 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable_Accum_Internal.Get_applyStateR()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
}))
	})
	return cache_applicativeStateR__2039640491
}

var cache_applyStateL__1243455060 gopurs_runtime.Value
var once_applyStateL__1243455060 sync.Once
func Get_applyStateL__1243455060() gopurs_runtime.Value {
	once_applyStateL__1243455060.Do(func() {
		cache_applyStateL__1243455060 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable_Accum_Internal.Get_functorStateL()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(f_0, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(x_1, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v_3_0, "value"), gopurs_runtime.RecordGet(v1_4_1, "value")))
})
})
}))
	})
	return cache_applyStateL__1243455060
}

var cache_applyStateR__1243455060 gopurs_runtime.Value
var once_applyStateR__1243455060 sync.Once
func Get_applyStateR__1243455060() gopurs_runtime.Value {
	once_applyStateR__1243455060.Do(func() {
		cache_applyStateR__1243455060 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable_Accum_Internal.Get_functorStateR()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(x_1, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v1_4_1, "value"), gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_applyStateR__1243455060
}

var cache_functorStateL__830241200 gopurs_runtime.Value
var once_functorStateL__830241200 sync.Once
func Get_functorStateL__830241200() gopurs_runtime.Value {
	once_functorStateL__830241200.Do(func() {
		cache_functorStateL__830241200 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_functorStateL__830241200
}

var cache_functorStateR__830241200 gopurs_runtime.Value
var once_functorStateR__830241200 sync.Once
func Get_functorStateR__830241200() gopurs_runtime.Value {
	once_functorStateR__830241200.Do(func() {
		cache_functorStateR__830241200 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_functorStateR__830241200
}

var cache_stateL__1334064830 gopurs_runtime.Value
var once_stateL__1334064830 sync.Once
func Get_stateL__1334064830() gopurs_runtime.Value {
	once_stateL__1334064830.Do(func() {
		cache_stateL__1334064830 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateL__1334064830(v_0_box)
})
	})
	return cache_stateL__1334064830
}

var cache_stateL__1412771550 gopurs_runtime.Value
var once_stateL__1412771550 sync.Once
func Get_stateL__1412771550() gopurs_runtime.Value {
	once_stateL__1412771550.Do(func() {
		cache_stateL__1412771550 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateL__1412771550(v_0_box)
})
	})
	return cache_stateL__1412771550
}

var cache_stateR__1334064830 gopurs_runtime.Value
var once_stateR__1334064830 sync.Once
func Get_stateR__1334064830() gopurs_runtime.Value {
	once_stateR__1334064830.Do(func() {
		cache_stateR__1334064830 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateR__1334064830(v_0_box)
})
	})
	return cache_stateR__1334064830
}

var cache_stateR__1412771550 gopurs_runtime.Value
var once_stateR__1412771550 sync.Once
func Get_stateR__1412771550() gopurs_runtime.Value {
	once_stateR__1412771550.Do(func() {
		cache_stateR__1412771550 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateR__1412771550(v_0_box)
})
	})
	return cache_stateR__1412771550
}

var cache_sequence__1886310617 gopurs_runtime.Value
var once_sequence__1886310617 sync.Once
func Get_sequence__1886310617() gopurs_runtime.Value {
	once_sequence__1886310617.Do(func() {
		cache_sequence__1886310617 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence__1886310617(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequence__1886310617
}

var cache_sequence__2322703601 gopurs_runtime.Value
var once_sequence__2322703601 sync.Once
func Get_sequence__2322703601() gopurs_runtime.Value {
	once_sequence__2322703601.Do(func() {
		cache_sequence__2322703601 = gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableMaybe(), "sequence")
	})
	return cache_sequence__2322703601
}

var cache_traversableAdditive__3840848827 gopurs_runtime.Value
var once_traversableAdditive__3840848827 sync.Once
func Get_traversableAdditive__3840848827() gopurs_runtime.Value {
	once_traversableAdditive__3840848827.Do(func() {
		cache_traversableAdditive__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Additive.Get_functorAdditive()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Additive.Get_Additive(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Monoid_Additive.Get_Additive(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableAdditive__3840848827
}

var cache_traversableArray__2643873085 gopurs_runtime.Value
var once_traversableArray__2643873085 sync.Once
func Get_traversableArray__2643873085() gopurs_runtime.Value {
	once_traversableArray__2643873085.Do(func() {
		cache_traversableArray__2643873085 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableArray(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_Traversable.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply4(pkg_Data_Traversable.Get_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"))
}))
	})
	return cache_traversableArray__2643873085
}

var cache_traversableConj__3840848827 gopurs_runtime.Value
var once_traversableConj__3840848827 sync.Once
func Get_traversableConj__3840848827() gopurs_runtime.Value {
	once_traversableConj__3840848827.Do(func() {
		cache_traversableConj__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Conj.Get_functorConj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Conj.Get_Conj(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Monoid_Conj.Get_Conj(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableConj__3840848827
}

var cache_traversableConst__3861086397 gopurs_runtime.Value
var once_traversableConst__3861086397 sync.Once
func Get_traversableConst__3861086397() gopurs_runtime.Value {
	once_traversableConst__3861086397.Do(func() {
		cache_traversableConst__3861086397 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Const.Get_functorConst()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v_1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v1_2)
})
})
}))
	})
	return cache_traversableConst__3861086397
}

var cache_traversableDisj__3840848827 gopurs_runtime.Value
var once_traversableDisj__3840848827 sync.Once
func Get_traversableDisj__3840848827() gopurs_runtime.Value {
	once_traversableDisj__3840848827.Do(func() {
		cache_traversableDisj__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Disj.Get_functorDisj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Disj.Get_Disj(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Monoid_Disj.Get_Disj(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableDisj__3840848827
}

var cache_traversableDual__3840848827 gopurs_runtime.Value
var once_traversableDual__3840848827 sync.Once
func Get_traversableDual__3840848827() gopurs_runtime.Value {
	once_traversableDual__3840848827.Do(func() {
		cache_traversableDual__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Dual.Get_functorDual()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Dual.Get_Dual(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Monoid_Dual.Get_Dual(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableDual__3840848827
}

var cache_traversableEither__4232556979 gopurs_runtime.Value
var once_traversableEither__4232556979 sync.Once
func Get_traversableEither__4232556979() gopurs_runtime.Value {
	once_traversableEither__4232556979.Do(func() {
		cache_traversableEither__4232556979 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_functorEither()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0})})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Either.Get_Right(), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0})})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(Functor0_1_2.V0, pkg_Data_Either.Get_Right(), gopurs_runtime.Apply(v_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
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
	return cache_traversableEither__4232556979
}

var cache_traversableFirst__822285914 gopurs_runtime.Value
var once_traversableFirst__822285914 sync.Once
func Get_traversableFirst__822285914() gopurs_runtime.Value {
	once_traversableFirst__822285914.Do(func() {
		cache_traversableFirst__822285914 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableFirst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Maybe_First.Get_First(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableMaybe(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Maybe_First.Get_First(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableMaybe(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_3))}))
})
})
}))
	})
	return cache_traversableFirst__822285914
}

var cache_traversableIdentity__3840848827 gopurs_runtime.Value
var once_traversableIdentity__3840848827 sync.Once
func Get_traversableIdentity__3840848827() gopurs_runtime.Value {
	once_traversableIdentity__3840848827.Do(func() {
		cache_traversableIdentity__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Identity.Get_Identity(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableIdentity__3840848827
}

var cache_traversableLast__822285914 gopurs_runtime.Value
var once_traversableLast__822285914 sync.Once
func Get_traversableLast__822285914() gopurs_runtime.Value {
	once_traversableLast__822285914.Do(func() {
		cache_traversableLast__822285914 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableLast()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Maybe_Last.Get_Last(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableMaybe(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_2))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Maybe_Last.Get_Last(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableMaybe(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_3))}))
})
})
}))
	})
	return cache_traversableLast__822285914
}

var cache_traversableMaybe__98548986 gopurs_runtime.Value
var once_traversableMaybe__98548986 sync.Once
func Get_traversableMaybe__98548986() gopurs_runtime.Value {
	once_traversableMaybe__98548986.Do(func() {
		cache_traversableMaybe__98548986 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Maybe.Get_Just(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(Functor0_1_2.V0, pkg_Data_Maybe.Get_Just(), gopurs_runtime.Apply(v_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
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
	return cache_traversableMaybe__98548986
}

var cache_traversableMaybe__822285914 gopurs_runtime.Value
var once_traversableMaybe__822285914 sync.Once
func Get_traversableMaybe__822285914() gopurs_runtime.Value {
	once_traversableMaybe__822285914.Do(func() {
		cache_traversableMaybe__822285914 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Maybe.Get_Just(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(Functor0_1_2.V0, pkg_Data_Maybe.Get_Just(), gopurs_runtime.Apply(v_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
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
	return cache_traversableMaybe__822285914
}

var cache_traversableMultiplicative__3840848827 gopurs_runtime.Value
var once_traversableMultiplicative__3840848827 sync.Once
func Get_traversableMultiplicative__3840848827() gopurs_runtime.Value {
	once_traversableMultiplicative__3840848827.Do(func() {
		cache_traversableMultiplicative__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Multiplicative.Get_functorMultiplicative()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableMultiplicative__3840848827
}

var cache_traversableTuple__3228991731 gopurs_runtime.Value
var once_traversableTuple__3228991731 sync.Once
func Get_traversableTuple__3228991731() gopurs_runtime.Value {
	once_traversableTuple__3228991731.Do(func() {
		cache_traversableTuple__3228991731 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
})
}))
	})
	return cache_traversableTuple__3228991731
}

var cache_traverse__314957093 gopurs_runtime.Value
var once_traverse__314957093 sync.Once
func Get_traverse__314957093() gopurs_runtime.Value {
	once_traverse__314957093.Do(func() {
		cache_traverse__314957093 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__314957093(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__314957093
}

var cache_traverse__667327821 gopurs_runtime.Value
var once_traverse__667327821 sync.Once
func Get_traverse__667327821() gopurs_runtime.Value {
	once_traverse__667327821.Do(func() {
		cache_traverse__667327821 = gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableMaybe(), "traverse")
	})
	return cache_traverse__667327821
}

var cache_mapAccumLWithIndex__142050190 gopurs_runtime.Value
var once_mapAccumLWithIndex__142050190 sync.Once
func Get_mapAccumLWithIndex__142050190() gopurs_runtime.Value {
	once_mapAccumLWithIndex__142050190.Do(func() {
		cache_mapAccumLWithIndex__142050190 = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumLWithIndex__142050190(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumLWithIndex__142050190
}

var cache_mapAccumLWithIndex__1384596302 gopurs_runtime.Value
var once_mapAccumLWithIndex__1384596302 sync.Once
func Get_mapAccumLWithIndex__1384596302() gopurs_runtime.Value {
	once_mapAccumLWithIndex__1384596302.Do(func() {
		cache_mapAccumLWithIndex__1384596302 = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumLWithIndex__1384596302(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumLWithIndex__1384596302
}

var cache_mapAccumRWithIndex__142050190 gopurs_runtime.Value
var once_mapAccumRWithIndex__142050190 sync.Once
func Get_mapAccumRWithIndex__142050190() gopurs_runtime.Value {
	once_mapAccumRWithIndex__142050190.Do(func() {
		cache_mapAccumRWithIndex__142050190 = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumRWithIndex__142050190(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumRWithIndex__142050190
}

var cache_mapAccumRWithIndex__1384596302 gopurs_runtime.Value
var once_mapAccumRWithIndex__1384596302 sync.Once
func Get_mapAccumRWithIndex__1384596302() gopurs_runtime.Value {
	once_mapAccumRWithIndex__1384596302.Do(func() {
		cache_mapAccumRWithIndex__1384596302 = gopurs_runtime.Func4(func(dictTraversableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumRWithIndex__1384596302(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dictTraversableWithIndex_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumRWithIndex__1384596302
}

var cache_traverseWithIndex__2726076659 gopurs_runtime.Value
var once_traverseWithIndex__2726076659 sync.Once
func Get_traverseWithIndex__2726076659() gopurs_runtime.Value {
	once_traverseWithIndex__2726076659.Do(func() {
		cache_traverseWithIndex__2726076659 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex__2726076659(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverseWithIndex__2726076659
}

var cache_traverseWithIndex__2979276403 gopurs_runtime.Value
var once_traverseWithIndex__2979276403 sync.Once
func Get_traverseWithIndex__2979276403() gopurs_runtime.Value {
	once_traverseWithIndex__2979276403.Do(func() {
		cache_traverseWithIndex__2979276403 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex__2979276403(gopurs_runtime.CoerceToStruct[Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverseWithIndex__2979276403
}

var cache_curry__2567276513 gopurs_runtime.Value
var once_curry__2567276513 sync.Once
func Get_curry__2567276513() gopurs_runtime.Value {
	once_curry__2567276513.Do(func() {
		cache_curry__2567276513 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_curry__2567276513(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_curry__2567276513
}

var cache_curry__4081152289 gopurs_runtime.Value
var once_curry__4081152289 sync.Once
func Get_curry__4081152289() gopurs_runtime.Value {
	once_curry__4081152289.Do(func() {
		cache_curry__4081152289 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, b_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_curry__4081152289(f_0_box, a_1_box, b_2_box)
})
	})
	return cache_curry__4081152289
}

var cache_functorTuple__2249620049 gopurs_runtime.Value
var once_functorTuple__2249620049 sync.Once
func Get_functorTuple__2249620049() gopurs_runtime.Value {
	once_functorTuple__2249620049.Do(func() {
		cache_functorTuple__2249620049 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorTuple__2249620049
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
return gopurs_runtime.Apply2(Apply0_9_6.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Apply0_9_6.V0, gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product.Get_product(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_8))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_12})})
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_8))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
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
__local_var_11_8 := gopurs_runtime.Apply(Functor0_9_6.V0, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_11})}
}))
_ = __local_var_11_8
__local_var_12_9 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_8))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_12})})
}))
_ = __local_var_12_9
__local_var_11_7 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_8, gopurs_runtime.Apply(__local_var_12_9, x_13))
})
_ = __local_var_11_7
__local_var_12_11 := gopurs_runtime.Apply(Functor0_9_6.V0, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_12})}
}))
_ = __local_var_12_11
__local_var_13_12 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversableWithIndex1_4, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_8))}, gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_13})})
}))
_ = __local_var_13_12
__local_var_12_10 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_11, gopurs_runtime.Apply(__local_var_13_12, x_14))
})
_ = __local_var_12_10
return gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(__local_var_11_7, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t13 = gopurs_runtime.Apply(__local_var_12_10, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0)
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
return gopurs_runtime.Apply2(Functor0_10_10.V0, pkg_Data_Functor_Compose.Get_Compose(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_9))}, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Apply2(Functor0_5_5.V0, pkg_Data_Functor_App.Get_App(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_4))}, f_6, v_7))
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
return gopurs_runtime.Apply4(dictTraversableWithIndex_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Apply4(dictTraversableWithIndex_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1475749520(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__154576880(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1304937360(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lift2__2762258480(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__470376976(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldMap__4098395794(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__3858583060(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1931071680(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_coproduct__413515331(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_coproduct__1706612365(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_product__2764631669(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, fa_0, ga_1})})
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__328307316(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2701008148(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1762802164(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3659954292(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2251722612(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1337616244(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mapWithIndex__55256674(dict_0_loop *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_mapWithIndex__2239747170(dict_0_loop *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_stateL__1334064830(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_stateL__1412771550(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_stateR__1334064830(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_stateR__1412771550(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_sequence__1886310617(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_mapAccumLWithIndex__142050190(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversableWithIndex_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_mapAccumLWithIndex__1384596302(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversableWithIndex_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_mapAccumRWithIndex__142050190(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversableWithIndex_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_mapAccumRWithIndex__1384596302(dictTraversableWithIndex_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversableWithIndex_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR()))}, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_1, i_4, s_6, a_5)
})
})
}), xs_3, s0_2)
}

func Call_traverseWithIndex__2726076659(dict_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverseWithIndex__2979276403(dict_0_loop *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_curry__2567276513(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_1, b_2})})
}

func Call_curry__4081152289(f_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value, b_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_1, b_2})})
}



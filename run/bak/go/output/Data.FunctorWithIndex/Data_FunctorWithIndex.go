package Data_FunctorWithIndex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Disj "gopurs/output/Data.Monoid.Disj"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Monoid_Conj "gopurs/output/Data.Monoid.Conj"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Monoid_Additive "gopurs/output/Data.Monoid.Additive"
	pkg_Data_Functor_Product "gopurs/output/Data.Functor.Product"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Functor_Coproduct "gopurs/output/Data.Functor.Coproduct"
	unsafe "unsafe"
)

var cache_mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		cache_mapWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex(dict_0_box)
})
	})
	return cache_mapWithIndex
}

var cache_mapDefault gopurs_runtime.Value
var once_mapDefault sync.Once
func Get_mapDefault() gopurs_runtime.Value {
	once_mapDefault.Do(func() {
		cache_mapDefault = gopurs_runtime.Func2(func(dictFunctorWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapDefault(dictFunctorWithIndex_0_box, f_1_box)
})
	})
	return cache_mapDefault
}

var cache_functorWithIndexTuple gopurs_runtime.Value
var once_functorWithIndexTuple sync.Once
func Get_functorWithIndexTuple() gopurs_runtime.Value {
	once_functorWithIndexTuple.Do(func() {
		cache_functorWithIndexTuple = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexTuple
}

var cache_functorWithIndexProduct gopurs_runtime.Value
var once_functorWithIndexProduct sync.Once
func Get_functorWithIndexProduct() gopurs_runtime.Value {
	once_functorWithIndexProduct.Do(func() {
		cache_functorWithIndexProduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorWithIndexProduct(dictFunctorWithIndex_0_box)
})
	})
	return cache_functorWithIndexProduct
}

var cache_functorWithIndexMultiplicative gopurs_runtime.Value
var once_functorWithIndexMultiplicative sync.Once
func Get_functorWithIndexMultiplicative() gopurs_runtime.Value {
	once_functorWithIndexMultiplicative.Do(func() {
		cache_functorWithIndexMultiplicative = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Multiplicative.Get_functorMultiplicative()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid_Multiplicative.Get_functorMultiplicative(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexMultiplicative
}

var cache_functorWithIndexMaybe gopurs_runtime.Value
var once_functorWithIndexMaybe sync.Once
func Get_functorWithIndexMaybe() gopurs_runtime.Value {
	once_functorWithIndexMaybe.Do(func() {
		cache_functorWithIndexMaybe = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexMaybe
}

var cache_functorWithIndexLast gopurs_runtime.Value
var once_functorWithIndexLast sync.Once
func Get_functorWithIndexLast() gopurs_runtime.Value {
	once_functorWithIndexLast.Do(func() {
		cache_functorWithIndexLast = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexLast
}

var cache_functorWithIndexIdentity gopurs_runtime.Value
var once_functorWithIndexIdentity sync.Once
func Get_functorWithIndexIdentity() gopurs_runtime.Value {
	once_functorWithIndexIdentity.Do(func() {
		cache_functorWithIndexIdentity = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, pkg_Data_Unit.Get_unit(), v_1)
}))
	})
	return cache_functorWithIndexIdentity
}

var cache_functorWithIndexFirst gopurs_runtime.Value
var once_functorWithIndexFirst sync.Once
func Get_functorWithIndexFirst() gopurs_runtime.Value {
	once_functorWithIndexFirst.Do(func() {
		cache_functorWithIndexFirst = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexFirst
}

var cache_functorWithIndexEither gopurs_runtime.Value
var once_functorWithIndexEither sync.Once
func Get_functorWithIndexEither() gopurs_runtime.Value {
	once_functorWithIndexEither.Do(func() {
		cache_functorWithIndexEither = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_functorEither()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Either.Get_functorEither(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexEither
}

var cache_functorWithIndexDual gopurs_runtime.Value
var once_functorWithIndexDual sync.Once
func Get_functorWithIndexDual() gopurs_runtime.Value {
	once_functorWithIndexDual.Do(func() {
		cache_functorWithIndexDual = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Dual.Get_functorDual()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid_Dual.Get_functorDual(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexDual
}

var cache_functorWithIndexDisj gopurs_runtime.Value
var once_functorWithIndexDisj sync.Once
func Get_functorWithIndexDisj() gopurs_runtime.Value {
	once_functorWithIndexDisj.Do(func() {
		cache_functorWithIndexDisj = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Disj.Get_functorDisj()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid_Disj.Get_functorDisj(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexDisj
}

var cache_functorWithIndexCoproduct gopurs_runtime.Value
var once_functorWithIndexCoproduct sync.Once
func Get_functorWithIndexCoproduct() gopurs_runtime.Value {
	once_functorWithIndexCoproduct.Do(func() {
		cache_functorWithIndexCoproduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorWithIndexCoproduct(dictFunctorWithIndex_0_box)
})
	})
	return cache_functorWithIndexCoproduct
}

var cache_functorWithIndexConst gopurs_runtime.Value
var once_functorWithIndexConst sync.Once
func Get_functorWithIndexConst() gopurs_runtime.Value {
	once_functorWithIndexConst.Do(func() {
		cache_functorWithIndexConst = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Const.Get_functorConst()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}))
	})
	return cache_functorWithIndexConst
}

var cache_functorWithIndexConj gopurs_runtime.Value
var once_functorWithIndexConj sync.Once
func Get_functorWithIndexConj() gopurs_runtime.Value {
	once_functorWithIndexConj.Do(func() {
		cache_functorWithIndexConj = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Conj.Get_functorConj()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid_Conj.Get_functorConj(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexConj
}

var cache_functorWithIndexCompose gopurs_runtime.Value
var once_functorWithIndexCompose sync.Once
func Get_functorWithIndexCompose() gopurs_runtime.Value {
	once_functorWithIndexCompose.Do(func() {
		cache_functorWithIndexCompose = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorWithIndexCompose(dictFunctorWithIndex_0_box)
})
	})
	return cache_functorWithIndexCompose
}

var cache_functorWithIndexArray gopurs_runtime.Value
var once_functorWithIndexArray sync.Once
func Get_functorWithIndexArray() gopurs_runtime.Value {
	once_functorWithIndexArray.Do(func() {
		cache_functorWithIndexArray = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), Get_mapWithIndexArray())
	})
	return cache_functorWithIndexArray
}

var cache_functorWithIndexApp gopurs_runtime.Value
var once_functorWithIndexApp sync.Once
func Get_functorWithIndexApp() gopurs_runtime.Value {
	once_functorWithIndexApp.Do(func() {
		cache_functorWithIndexApp = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorWithIndexApp(dictFunctorWithIndex_0_box)
})
	})
	return cache_functorWithIndexApp
}

var cache_functorWithIndexAdditive gopurs_runtime.Value
var once_functorWithIndexAdditive sync.Once
func Get_functorWithIndexAdditive() gopurs_runtime.Value {
	once_functorWithIndexAdditive.Do(func() {
		cache_functorWithIndexAdditive = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Additive.Get_functorAdditive()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid_Additive.Get_functorAdditive(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexAdditive
}

func Call_mapWithIndex(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData1)(dict_0.UnsafePtr)).V0
}

func Call_mapDefault(dictFunctorWithIndex_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictFunctorWithIndex_0.UnsafePtr)).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
}

func Call_functorWithIndexProduct(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
functorProduct_1_0 := gopurs_runtime.Apply(pkg_Data_Functor_Product.Get_functorProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = functorProduct_1_0
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
functorProduct1_3_1 := gopurs_runtime.Apply(functorProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{}))
_ = functorProduct1_3_1
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct1_3_1
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Bifunctor.Get_bifunctorTuple(), "bimap"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictFunctorWithIndex_0.UnsafePtr)).V0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{x_6})})
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{x_6})})
})), v_5)
}))
})
}

func Call_functorWithIndexCoproduct(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
functorCoproduct_1_0 := gopurs_runtime.Apply(pkg_Data_Functor_Coproduct.Get_functorCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = functorCoproduct_1_0
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
functorCoproduct1_3_1 := gopurs_runtime.Apply(functorCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct1_3_1
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct1_3_1
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Bifunctor.Get_bifunctorEither(), "bimap"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictFunctorWithIndex_0.UnsafePtr)).V0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{x_6})})
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{x_6})})
})), v_5)
}))
})
}

func Call_functorWithIndexCompose(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
functorCompose1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "map"), f_4), v_5)
}))
_ = functorCompose1_4_2
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_4_2
}), gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctorWithIndex_0.UnsafePtr)).V0, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{x_7, b_8})})
}))
}), v_6)
}))
})
}

func Call_functorWithIndexApp(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctorWithIndex_0.UnsafePtr)).V0, f_2, v_3)
}))
}

func Get_mapWithIndexArray() gopurs_runtime.Value {
	return _Gopurs_MapWithIndexArray
}

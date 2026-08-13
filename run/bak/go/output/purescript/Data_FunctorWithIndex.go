package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_FunctorWithIndex_go__map gopurs_runtime.Value
var once_Data_FunctorWithIndex_go__map sync.Once
func Get_Data_FunctorWithIndex_go__map() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_go__map.Do(func() {
		cache_Data_FunctorWithIndex_go__map = gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map")
	})
	return cache_Data_FunctorWithIndex_go__map
}

var cache_Data_FunctorWithIndex_map1 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map1 sync.Once
func Get_Data_FunctorWithIndex_map1() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map1.Do(func() {
		cache_Data_FunctorWithIndex_map1 = gopurs_runtime.RecordGet(Get_Data_Monoid_Multiplicative_functorMultiplicative(), "map")
	})
	return cache_Data_FunctorWithIndex_map1
}

var cache_Data_FunctorWithIndex_map2 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map2 sync.Once
func Get_Data_FunctorWithIndex_map2() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map2.Do(func() {
		cache_Data_FunctorWithIndex_map2 = gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map")
	})
	return cache_Data_FunctorWithIndex_map2
}

var cache_Data_FunctorWithIndex_map3 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map3 sync.Once
func Get_Data_FunctorWithIndex_map3() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map3.Do(func() {
		cache_Data_FunctorWithIndex_map3 = gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map")
	})
	return cache_Data_FunctorWithIndex_map3
}

var cache_Data_FunctorWithIndex_map4 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map4 sync.Once
func Get_Data_FunctorWithIndex_map4() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map4.Do(func() {
		cache_Data_FunctorWithIndex_map4 = gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map")
	})
	return cache_Data_FunctorWithIndex_map4
}

var cache_Data_FunctorWithIndex_map5 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map5 sync.Once
func Get_Data_FunctorWithIndex_map5() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map5.Do(func() {
		cache_Data_FunctorWithIndex_map5 = gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map")
	})
	return cache_Data_FunctorWithIndex_map5
}

var cache_Data_FunctorWithIndex_map6 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map6 sync.Once
func Get_Data_FunctorWithIndex_map6() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map6.Do(func() {
		cache_Data_FunctorWithIndex_map6 = gopurs_runtime.RecordGet(Get_Data_Monoid_Dual_functorDual(), "map")
	})
	return cache_Data_FunctorWithIndex_map6
}

var cache_Data_FunctorWithIndex_map7 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map7 sync.Once
func Get_Data_FunctorWithIndex_map7() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map7.Do(func() {
		cache_Data_FunctorWithIndex_map7 = gopurs_runtime.RecordGet(Get_Data_Monoid_Disj_functorDisj(), "map")
	})
	return cache_Data_FunctorWithIndex_map7
}

var cache_Data_FunctorWithIndex_map8 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map8 sync.Once
func Get_Data_FunctorWithIndex_map8() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map8.Do(func() {
		cache_Data_FunctorWithIndex_map8 = gopurs_runtime.RecordGet(Get_Data_Monoid_Conj_functorConj(), "map")
	})
	return cache_Data_FunctorWithIndex_map8
}

var cache_Data_FunctorWithIndex_map9 gopurs_runtime.Value
var once_Data_FunctorWithIndex_map9 sync.Once
func Get_Data_FunctorWithIndex_map9() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_map9.Do(func() {
		cache_Data_FunctorWithIndex_map9 = gopurs_runtime.RecordGet(Get_Data_Monoid_Additive_functorAdditive(), "map")
	})
	return cache_Data_FunctorWithIndex_map9
}

var cache_Data_FunctorWithIndex_FunctorWithIndex_dollarDict gopurs_runtime.Value
var once_Data_FunctorWithIndex_FunctorWithIndex_dollarDict sync.Once
func Get_Data_FunctorWithIndex_FunctorWithIndex_dollarDict() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_FunctorWithIndex_dollarDict.Do(func() {
		cache_Data_FunctorWithIndex_FunctorWithIndex_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_FunctorWithIndex_dollarDict(x_0_box)
})
	})
	return cache_Data_FunctorWithIndex_FunctorWithIndex_dollarDict
}

var cache_Data_FunctorWithIndex_mapWithIndex gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapWithIndex sync.Once
func Get_Data_FunctorWithIndex_mapWithIndex() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapWithIndex.Do(func() {
		cache_Data_FunctorWithIndex_mapWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_mapWithIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](dict_0_box))
})
	})
	return cache_Data_FunctorWithIndex_mapWithIndex
}

var cache_Data_FunctorWithIndex_mapDefault gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapDefault sync.Once
func Get_Data_FunctorWithIndex_mapDefault() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapDefault.Do(func() {
		cache_Data_FunctorWithIndex_mapDefault = gopurs_runtime.Func2(func(dictFunctorWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_mapDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](dictFunctorWithIndex_0_box), f_1_box)
})
	})
	return cache_Data_FunctorWithIndex_mapDefault
}

var cache_Data_FunctorWithIndex_functorWithIndexTuple gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexTuple sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexTuple() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexTuple.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexTuple = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_functorTuple()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexTuple
}

var cache_Data_FunctorWithIndex_functorWithIndexProduct gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexProduct sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexProduct() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexProduct.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexProduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_functorWithIndexProduct(dictFunctorWithIndex_0_box)
})
	})
	return cache_Data_FunctorWithIndex_functorWithIndexProduct
}

var cache_Data_FunctorWithIndex_functorWithIndexMultiplicative gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexMultiplicative sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexMultiplicative() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexMultiplicative.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexMultiplicative = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Multiplicative_functorMultiplicative()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_Multiplicative_functorMultiplicative(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexMultiplicative
}

var cache_Data_FunctorWithIndex_functorWithIndexMaybe gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexMaybe sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexMaybe() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexMaybe.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexMaybe = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexMaybe
}

var cache_Data_FunctorWithIndex_functorWithIndexLast gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexLast sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexLast() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexLast.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexLast = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexLast
}

var cache_Data_FunctorWithIndex_functorWithIndexIdentity gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexIdentity sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexIdentity() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexIdentity.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexIdentity = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Identity_functorIdentity()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, Get_Data_Unit_unit(), v_1)
})
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexIdentity
}

var cache_Data_FunctorWithIndex_functorWithIndexFirst gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexFirst sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexFirst() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexFirst.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexFirst = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexFirst
}

var cache_Data_FunctorWithIndex_functorWithIndexEither gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexEither sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexEither() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexEither.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexEither = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_functorEither()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexEither
}

var cache_Data_FunctorWithIndex_functorWithIndexDual gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexDual sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexDual() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexDual.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexDual = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Dual_functorDual()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_Dual_functorDual(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexDual
}

var cache_Data_FunctorWithIndex_functorWithIndexDisj gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexDisj sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexDisj() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexDisj.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexDisj = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_functorDisj()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_Disj_functorDisj(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexDisj
}

var cache_Data_FunctorWithIndex_functorWithIndexCoproduct gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexCoproduct sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexCoproduct() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexCoproduct.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexCoproduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_functorWithIndexCoproduct(dictFunctorWithIndex_0_box)
})
	})
	return cache_Data_FunctorWithIndex_functorWithIndexCoproduct
}

var cache_Data_FunctorWithIndex_functorWithIndexConst gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexConst sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexConst() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexConst.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexConst = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Const_functorConst()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexConst
}

var cache_Data_FunctorWithIndex_functorWithIndexConj gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexConj sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexConj() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexConj.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexConj = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Conj_functorConj()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_Conj_functorConj(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexConj
}

var cache_Data_FunctorWithIndex_functorWithIndexCompose gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexCompose sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexCompose() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexCompose.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexCompose = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_functorWithIndexCompose(dictFunctorWithIndex_0_box)
})
	})
	return cache_Data_FunctorWithIndex_functorWithIndexCompose
}

var cache_Data_FunctorWithIndex_functorWithIndexArray gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexArray sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexArray() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexArray.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexArray = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorArray()
}), Get_Data_FunctorWithIndex_mapWithIndexArray())
	})
	return cache_Data_FunctorWithIndex_functorWithIndexArray
}

var cache_Data_FunctorWithIndex_functorWithIndexApp gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexApp sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexApp() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexApp.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexApp = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_functorWithIndexApp(dictFunctorWithIndex_0_box)
})
	})
	return cache_Data_FunctorWithIndex_functorWithIndexApp
}

var cache_Data_FunctorWithIndex_functorWithIndexAdditive gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexAdditive sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexAdditive() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexAdditive.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexAdditive = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Additive_functorAdditive()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_Additive_functorAdditive(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexAdditive
}

var cache_Data_FunctorWithIndex_functorWithIndexAdditive__339744599 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexAdditive__339744599 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexAdditive__339744599() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexAdditive__339744599.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexAdditive__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Additive_functorAdditive()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_Additive_functorAdditive(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexAdditive__339744599
}

var cache_Data_FunctorWithIndex_functorWithIndexArray__3811533158 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexArray__3811533158 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexArray__3811533158() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexArray__3811533158.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexArray__3811533158 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorArray()
}), Get_Data_FunctorWithIndex_mapWithIndexArray())
	})
	return cache_Data_FunctorWithIndex_functorWithIndexArray__3811533158
}

var cache_Data_FunctorWithIndex_functorWithIndexArray__490015842 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexArray__490015842 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexArray__490015842() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexArray__490015842.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexArray__490015842 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorArray()
}), Get_Data_FunctorWithIndex_mapWithIndexArray())
	})
	return cache_Data_FunctorWithIndex_functorWithIndexArray__490015842
}

var cache_Data_FunctorWithIndex_functorWithIndexConj__339744599 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexConj__339744599 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexConj__339744599() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexConj__339744599.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexConj__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Conj_functorConj()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_Conj_functorConj(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexConj__339744599
}

var cache_Data_FunctorWithIndex_functorWithIndexConst__3232336655 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexConst__3232336655 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexConst__3232336655() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexConst__3232336655.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexConst__3232336655 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Const_functorConst()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexConst__3232336655
}

var cache_Data_FunctorWithIndex_functorWithIndexDisj__339744599 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexDisj__339744599 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexDisj__339744599() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexDisj__339744599.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexDisj__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_functorDisj()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_Disj_functorDisj(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexDisj__339744599
}

var cache_Data_FunctorWithIndex_functorWithIndexDual__339744599 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexDual__339744599 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexDual__339744599() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexDual__339744599.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexDual__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Dual_functorDual()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_Dual_functorDual(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexDual__339744599
}

var cache_Data_FunctorWithIndex_functorWithIndexEither__1698960599 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexEither__1698960599 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexEither__1698960599() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexEither__1698960599.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexEither__1698960599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_functorEither()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexEither__1698960599
}

var cache_Data_FunctorWithIndex_functorWithIndexFirst__2982254679 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexFirst__2982254679 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexFirst__2982254679() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexFirst__2982254679.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexFirst__2982254679 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexFirst__2982254679
}

var cache_Data_FunctorWithIndex_functorWithIndexIdentity__339744599 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexIdentity__339744599 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexIdentity__339744599() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexIdentity__339744599.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexIdentity__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Identity_functorIdentity()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, Get_Data_Unit_unit(), v_1)
})
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexIdentity__339744599
}

var cache_Data_FunctorWithIndex_functorWithIndexLast__2982254679 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexLast__2982254679 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexLast__2982254679() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexLast__2982254679.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexLast__2982254679 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexLast__2982254679
}

var cache_Data_FunctorWithIndex_functorWithIndexMaybe__2982254679 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexMaybe__2982254679 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexMaybe__2982254679() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexMaybe__2982254679.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexMaybe__2982254679 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexMaybe__2982254679
}

var cache_Data_FunctorWithIndex_functorWithIndexMultiplicative__339744599 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexMultiplicative__339744599 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexMultiplicative__339744599() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexMultiplicative__339744599.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexMultiplicative__339744599 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Multiplicative_functorMultiplicative()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_Multiplicative_functorMultiplicative(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexMultiplicative__339744599
}

var cache_Data_FunctorWithIndex_functorWithIndexTuple__1273126103 gopurs_runtime.Value
var once_Data_FunctorWithIndex_functorWithIndexTuple__1273126103 sync.Once
func Get_Data_FunctorWithIndex_functorWithIndexTuple__1273126103() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_functorWithIndexTuple__1273126103.Do(func() {
		cache_Data_FunctorWithIndex_functorWithIndexTuple__1273126103 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_functorTuple()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_FunctorWithIndex_functorWithIndexTuple__1273126103
}

var cache_Data_FunctorWithIndex_mapWithIndex__835054498 gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapWithIndex__835054498 sync.Once
func Get_Data_FunctorWithIndex_mapWithIndex__835054498() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapWithIndex__835054498.Do(func() {
		cache_Data_FunctorWithIndex_mapWithIndex__835054498 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_mapWithIndex__835054498(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_FunctorWithIndex_mapWithIndex__835054498
}

var cache_Data_FunctorWithIndex_mapWithIndex__1847943938 gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapWithIndex__1847943938 sync.Once
func Get_Data_FunctorWithIndex_mapWithIndex__1847943938() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapWithIndex__1847943938.Do(func() {
		cache_Data_FunctorWithIndex_mapWithIndex__1847943938 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_mapWithIndex__1847943938(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_FunctorWithIndex_mapWithIndex__1847943938
}

var cache_Data_FunctorWithIndex_mapWithIndex__55256674 gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapWithIndex__55256674 sync.Once
func Get_Data_FunctorWithIndex_mapWithIndex__55256674() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapWithIndex__55256674.Do(func() {
		cache_Data_FunctorWithIndex_mapWithIndex__55256674 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_mapWithIndex__55256674(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](dict_0_box))
})
	})
	return cache_Data_FunctorWithIndex_mapWithIndex__55256674
}

var cache_Data_FunctorWithIndex_mapWithIndex__2239747170 gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapWithIndex__2239747170 sync.Once
func Get_Data_FunctorWithIndex_mapWithIndex__2239747170() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapWithIndex__2239747170.Do(func() {
		cache_Data_FunctorWithIndex_mapWithIndex__2239747170 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_mapWithIndex__2239747170(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](dict_0_box))
})
	})
	return cache_Data_FunctorWithIndex_mapWithIndex__2239747170
}

var cache_Data_FunctorWithIndex_mapWithIndex__3104159586 gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapWithIndex__3104159586 sync.Once
func Get_Data_FunctorWithIndex_mapWithIndex__3104159586() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapWithIndex__3104159586.Do(func() {
		cache_Data_FunctorWithIndex_mapWithIndex__3104159586 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_mapWithIndex__3104159586(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](dict_0_box))
})
	})
	return cache_Data_FunctorWithIndex_mapWithIndex__3104159586
}

var cache_Data_FunctorWithIndex_mapWithIndex__574674314 gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapWithIndex__574674314 sync.Once
func Get_Data_FunctorWithIndex_mapWithIndex__574674314() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapWithIndex__574674314.Do(func() {
		cache_Data_FunctorWithIndex_mapWithIndex__574674314 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_mapWithIndex__574674314(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_FunctorWithIndex_mapWithIndex__574674314
}

var cache_Data_FunctorWithIndex_mapWithIndex__3380890378 gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapWithIndex__3380890378 sync.Once
func Get_Data_FunctorWithIndex_mapWithIndex__3380890378() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapWithIndex__3380890378.Do(func() {
		cache_Data_FunctorWithIndex_mapWithIndex__3380890378 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_mapWithIndex__3380890378(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_FunctorWithIndex_mapWithIndex__3380890378
}

var cache_Data_FunctorWithIndex_mapWithIndex__598554346 gopurs_runtime.Value
var once_Data_FunctorWithIndex_mapWithIndex__598554346 sync.Once
func Get_Data_FunctorWithIndex_mapWithIndex__598554346() gopurs_runtime.Value {
	once_Data_FunctorWithIndex_mapWithIndex__598554346.Do(func() {
		cache_Data_FunctorWithIndex_mapWithIndex__598554346 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_FunctorWithIndex_mapWithIndex__598554346(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_FunctorWithIndex_mapWithIndex__598554346
}

type Constructor_Data_FunctorWithIndex_FunctorWithIndex struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4077743418] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_FunctorWithIndex_FunctorWithIndex)(ptr)
		_ = c
		switch key {
		case "Functor0": return gopurs_runtime.Box(c.V0)
		case "mapWithIndex": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_FunctorWithIndex_FunctorWithIndex: " + key)
		}
	}
}


func Call_Data_FunctorWithIndex_FunctorWithIndex_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_FunctorWithIndex_mapWithIndex(dict_0_loop *Constructor_Data_FunctorWithIndex_FunctorWithIndex) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FunctorWithIndex_FunctorWithIndex = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_FunctorWithIndex_mapDefault(dictFunctorWithIndex_0_loop *Constructor_Data_FunctorWithIndex_FunctorWithIndex, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 *Constructor_Data_FunctorWithIndex_FunctorWithIndex = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctorWithIndex_0.V1), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
}

func Call_Data_FunctorWithIndex_functorWithIndexProduct(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
// TAST (Let): functorProduct_1_0 -> gopurs_runtime.Value
functorProduct_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Product_functorProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorProduct_1_0
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorProduct1_3_1 -> gopurs_runtime.Value
functorProduct1_3_1 := gopurs_runtime.Apply(functorProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{}))
_ = functorProduct1_3_1
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct1_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Bifunctor_bifunctorTuple(), "bimap"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_6})})
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_6})})
})), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_5))})))}
})
}))
})
}

func Call_Data_FunctorWithIndex_functorWithIndexCoproduct(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
// TAST (Let): functorCoproduct_1_0 -> gopurs_runtime.Value
functorCoproduct_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Coproduct_functorCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct_1_0
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCoproduct1_3_1 -> gopurs_runtime.Value
functorCoproduct1_3_1 := gopurs_runtime.Apply(functorCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct1_3_1
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct1_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Bifunctor_bifunctorEither(), "bimap"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_6})})
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_6})})
})), v_5)
})
}))
})
}

func Call_Data_FunctorWithIndex_functorWithIndexCompose(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorCompose1_3_1 -> gopurs_runtime.Value
functorCompose1_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "map"), f_4), v_5)
})
}))
_ = functorCompose1_3_1
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, x_6, b_7})})
}))
}), v_5)
})
}))
})
}

func Call_Data_FunctorWithIndex_functorWithIndexApp(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
// TAST (Let): functorApp_1_0 -> gopurs_runtime.Value
functorApp_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = functorApp_1_0
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorApp_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), f_2, v_3)
})
}))
}

func Call_Data_FunctorWithIndex_mapWithIndex__835054498(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Data_FunctorWithIndex_mapWithIndexArray(), __eta0_0, __eta1_1)
}

func Call_Data_FunctorWithIndex_mapWithIndex__1847943938(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Data_FunctorWithIndex_mapWithIndexArray(), __eta0_0, __eta1_1)
}

func Call_Data_FunctorWithIndex_mapWithIndex__55256674(dict_0_loop *Constructor_Data_FunctorWithIndex_FunctorWithIndex) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FunctorWithIndex_FunctorWithIndex = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_FunctorWithIndex_mapWithIndex__2239747170(dict_0_loop *Constructor_Data_FunctorWithIndex_FunctorWithIndex) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FunctorWithIndex_FunctorWithIndex = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_FunctorWithIndex_mapWithIndex__3104159586(dict_0_loop *Constructor_Data_FunctorWithIndex_FunctorWithIndex) gopurs_runtime.Value {
var dict_0 *Constructor_Data_FunctorWithIndex_FunctorWithIndex = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_FunctorWithIndex_mapWithIndex__574674314(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](Get_Data_List_Lazy_Types_functorWithIndex()).V1), __eta0_0, __eta1_1)
}

func Call_Data_FunctorWithIndex_mapWithIndex__3380890378(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](Get_Data_List_Lazy_Types_functorWithIndex()).V1), __eta0_0, __eta1_1)
}

func Call_Data_FunctorWithIndex_mapWithIndex__598554346(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](Get_Data_List_Types_functorWithIndex()).V1), __eta0_0, __eta1_1)
}

func Get_Data_FunctorWithIndex_mapWithIndexArray() gopurs_runtime.Value {
	return _Gopurs_Data_FunctorWithIndex_MapWithIndexArray
}

package Data_FunctorWithIndex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Disj "gopurs/output/Data.Monoid.Disj"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Monoid_Conj "gopurs/output/Data.Monoid.Conj"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Monoid_Additive "gopurs/output/Data.Monoid.Additive"
	unsafe "unsafe"
)

var cache_mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		cache_mapWithIndex = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mapWithIndex")
}()
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
		cache_functorWithIndexTuple = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(m_2.UnsafePtr).V0, gopurs_runtime.Apply(__local_var_1_0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(m_2.UnsafePtr).V1)})}
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
}))
	})
	return cache_functorWithIndexTuple
}

var cache_functorWithIndexProduct gopurs_runtime.Value
var once_functorWithIndexProduct sync.Once
func Get_functorWithIndexProduct() gopurs_runtime.Value {
	once_functorWithIndexProduct.Do(func() {
		cache_functorWithIndexProduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
functorProduct1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_4, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "map"), f_4, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_5.UnsafePtr).V1)})}
}))
_ = functorProduct1_4_2
return gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{x_7})})
}), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{x_7})})
}), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_6.UnsafePtr).V1)})}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct1_4_2
}))
})
}()
})
	})
	return cache_functorWithIndexProduct
}

var cache_functorWithIndexMultiplicative gopurs_runtime.Value
var once_functorWithIndexMultiplicative sync.Once
func Get_functorWithIndexMultiplicative() gopurs_runtime.Value {
	once_functorWithIndexMultiplicative.Do(func() {
		cache_functorWithIndexMultiplicative = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Multiplicative.Get_functorMultiplicative()
}))
	})
	return cache_functorWithIndexMultiplicative
}

var cache_functorWithIndexMaybe gopurs_runtime.Value
var once_functorWithIndexMaybe sync.Once
func Get_functorWithIndexMaybe() gopurs_runtime.Value {
	once_functorWithIndexMaybe.Do(func() {
		cache_functorWithIndexMaybe = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(__local_var_1_0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_2.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}))
	})
	return cache_functorWithIndexMaybe
}

var cache_functorWithIndexLast gopurs_runtime.Value
var once_functorWithIndexLast sync.Once
func Get_functorWithIndexLast() gopurs_runtime.Value {
	once_functorWithIndexLast.Do(func() {
		cache_functorWithIndexLast = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(__local_var_1_0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_2.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}))
	})
	return cache_functorWithIndexLast
}

var cache_functorWithIndexIdentity gopurs_runtime.Value
var once_functorWithIndexIdentity sync.Once
func Get_functorWithIndexIdentity() gopurs_runtime.Value {
	once_functorWithIndexIdentity.Do(func() {
		cache_functorWithIndexIdentity = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, pkg_Data_Unit.Get_unit(), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
}))
	})
	return cache_functorWithIndexIdentity
}

var cache_functorWithIndexFirst gopurs_runtime.Value
var once_functorWithIndexFirst sync.Once
func Get_functorWithIndexFirst() gopurs_runtime.Value {
	once_functorWithIndexFirst.Do(func() {
		cache_functorWithIndexFirst = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{gopurs_runtime.Apply(__local_var_1_0, (*pkg_Data_Maybe.Data_Data_Maybe_Just)(v1_2.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
}
end_branch_1:
return __t1
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}))
	})
	return cache_functorWithIndexFirst
}

var cache_functorWithIndexEither gopurs_runtime.Value
var once_functorWithIndexEither sync.Once
func Get_functorWithIndexEither() gopurs_runtime.Value {
	once_functorWithIndexEither.Do(func() {
		cache_functorWithIndexEither = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (m_2.Type == 9 && m_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{(*pkg_Data_Either.Data_Data_Either_Left)(m_2.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
if (m_2.Type == 9 && m_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Apply(__local_var_1_0, (*pkg_Data_Either.Data_Data_Either_Right)(m_2.UnsafePtr).V0)})}
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
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_functorEither()
}))
	})
	return cache_functorWithIndexEither
}

var cache_functorWithIndexDual gopurs_runtime.Value
var once_functorWithIndexDual sync.Once
func Get_functorWithIndexDual() gopurs_runtime.Value {
	once_functorWithIndexDual.Do(func() {
		cache_functorWithIndexDual = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Dual.Get_functorDual()
}))
	})
	return cache_functorWithIndexDual
}

var cache_functorWithIndexDisj gopurs_runtime.Value
var once_functorWithIndexDisj sync.Once
func Get_functorWithIndexDisj() gopurs_runtime.Value {
	once_functorWithIndexDisj.Do(func() {
		cache_functorWithIndexDisj = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Disj.Get_functorDisj()
}))
	})
	return cache_functorWithIndexDisj
}

var cache_functorWithIndexCoproduct gopurs_runtime.Value
var once_functorWithIndexCoproduct sync.Once
func Get_functorWithIndexCoproduct() gopurs_runtime.Value {
	once_functorWithIndexCoproduct.Do(func() {
		cache_functorWithIndexCoproduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
functorCoproduct1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_4)
_ = __local_var_6_3
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "map"), f_4)
_ = __local_var_7_4
var __t5 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 3711209382) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{gopurs_runtime.Apply(__local_var_6_3, (*pkg_Data_Either.Data_Data_Either_Left)(v_5.UnsafePtr).V0)})}
goto end_branch_5
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 2465973597) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Apply(__local_var_7_4, (*pkg_Data_Either.Data_Data_Either_Right)(v_5.UnsafePtr).V0)})}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}))
_ = functorCoproduct1_4_2
return gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{x_7})})
}))
_ = __local_var_7_6
__local_var_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{x_8})})
}))
_ = __local_var_8_7
var __t8 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 3711209382) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Left{gopurs_runtime.Apply(__local_var_7_6, (*pkg_Data_Either.Data_Data_Either_Left)(v_6.UnsafePtr).V0)})}
goto end_branch_8
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 2465973597) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Data_Data_Either_Right{gopurs_runtime.Apply(__local_var_8_7, (*pkg_Data_Either.Data_Data_Either_Right)(v_6.UnsafePtr).V0)})}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct1_4_2
}))
})
}()
})
	})
	return cache_functorWithIndexCoproduct
}

var cache_functorWithIndexConst gopurs_runtime.Value
var once_functorWithIndexConst sync.Once
func Get_functorWithIndexConst() gopurs_runtime.Value {
	once_functorWithIndexConst.Do(func() {
		cache_functorWithIndexConst = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Const.Get_functorConst()
}))
	})
	return cache_functorWithIndexConst
}

var cache_functorWithIndexConj gopurs_runtime.Value
var once_functorWithIndexConj sync.Once
func Get_functorWithIndexConj() gopurs_runtime.Value {
	once_functorWithIndexConj.Do(func() {
		cache_functorWithIndexConj = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Conj.Get_functorConj()
}))
	})
	return cache_functorWithIndexConj
}

var cache_functorWithIndexCompose gopurs_runtime.Value
var once_functorWithIndexCompose sync.Once
func Get_functorWithIndexCompose() gopurs_runtime.Value {
	once_functorWithIndexCompose.Do(func() {
		cache_functorWithIndexCompose = gopurs_runtime.Func(func(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
functorCompose1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "map"), f_4), v_5)
}))
_ = functorCompose1_4_2
return gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{x_7, b_8})})
}))
}), v_6)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_4_2
}))
})
}()
})
	})
	return cache_functorWithIndexCompose
}

var cache_functorWithIndexArray gopurs_runtime.Value
var once_functorWithIndexArray sync.Once
func Get_functorWithIndexArray() gopurs_runtime.Value {
	once_functorWithIndexArray.Do(func() {
		cache_functorWithIndexArray = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", Get_mapWithIndexArray(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}))
	})
	return cache_functorWithIndexArray
}

var cache_functorWithIndexApp gopurs_runtime.Value
var once_functorWithIndexApp sync.Once
func Get_functorWithIndexApp() gopurs_runtime.Value {
	once_functorWithIndexApp.Do(func() {
		cache_functorWithIndexApp = gopurs_runtime.Func(func(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), f_2, v_3)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}))
}()
})
	})
	return cache_functorWithIndexApp
}

var cache_functorWithIndexAdditive gopurs_runtime.Value
var once_functorWithIndexAdditive sync.Once
func Get_functorWithIndexAdditive() gopurs_runtime.Value {
	once_functorWithIndexAdditive.Do(func() {
		cache_functorWithIndexAdditive = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Additive.Get_functorAdditive()
}))
	})
	return cache_functorWithIndexAdditive
}

func Call_mapDefault(dictFunctorWithIndex_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
}

func Get_mapWithIndexArray() gopurs_runtime.Value {
	return _Gopurs_MapWithIndexArray
}

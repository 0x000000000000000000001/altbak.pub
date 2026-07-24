package Data_FunctorWithIndex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
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
)

var mapWithIndex gopurs_runtime.Value
var once_mapWithIndex sync.Once
func Get_mapWithIndex() gopurs_runtime.Value {
	once_mapWithIndex.Do(func() {
		mapWithIndex = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mapWithIndex")
}()
})
	})
	return mapWithIndex
}

var mapDefault gopurs_runtime.Value
var once_mapDefault sync.Once
func Get_mapDefault() gopurs_runtime.Value {
	once_mapDefault.Do(func() {
		mapDefault = gopurs_runtime.Func2(func(dictFunctorWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapDefault(dictFunctorWithIndex_0_box, f_1_box)
})
	})
	return mapDefault
}

var functorWithIndexTuple gopurs_runtime.Value
var once_functorWithIndexTuple sync.Once
func Get_functorWithIndexTuple() gopurs_runtime.Value {
	once_functorWithIndexTuple.Do(func() {
		functorWithIndexTuple = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[0], gopurs_runtime.Apply(__local_var_1_0, (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[1]))
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
}))
	})
	return functorWithIndexTuple
}

var functorWithIndexProduct gopurs_runtime.Value
var once_functorWithIndexProduct sync.Once
func Get_functorWithIndexProduct() gopurs_runtime.Value {
	once_functorWithIndexProduct.Do(func() {
		functorWithIndexProduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFunctorWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_1
functorProduct1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "map"), f_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[1]))
}))
_ = functorProduct1_4_2
return gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Constructor1("Left", x_7))
}), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Constructor1("Right", x_7))
}), (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[1]))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct1_4_2
}))
})
}()
})
	})
	return functorWithIndexProduct
}

var functorWithIndexMultiplicative gopurs_runtime.Value
var once_functorWithIndexMultiplicative sync.Once
func Get_functorWithIndexMultiplicative() gopurs_runtime.Value {
	once_functorWithIndexMultiplicative.Do(func() {
		functorWithIndexMultiplicative = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Multiplicative.Get_functorMultiplicative()
}))
	})
	return functorWithIndexMultiplicative
}

var functorWithIndexMaybe gopurs_runtime.Value
var once_functorWithIndexMaybe sync.Once
func Get_functorWithIndexMaybe() gopurs_runtime.Value {
	once_functorWithIndexMaybe.Do(func() {
		functorWithIndexMaybe = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_2.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(__local_var_1_0, (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}))
	})
	return functorWithIndexMaybe
}

var functorWithIndexLast gopurs_runtime.Value
var once_functorWithIndexLast sync.Once
func Get_functorWithIndexLast() gopurs_runtime.Value {
	once_functorWithIndexLast.Do(func() {
		functorWithIndexLast = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_2.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(__local_var_1_0, (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}))
	})
	return functorWithIndexLast
}

var functorWithIndexIdentity gopurs_runtime.Value
var once_functorWithIndexIdentity sync.Once
func Get_functorWithIndexIdentity() gopurs_runtime.Value {
	once_functorWithIndexIdentity.Do(func() {
		functorWithIndexIdentity = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, pkg_Data_Unit.Get_unit(), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
}))
	})
	return functorWithIndexIdentity
}

var functorWithIndexFirst gopurs_runtime.Value
var once_functorWithIndexFirst sync.Once
func Get_functorWithIndexFirst() gopurs_runtime.Value {
	once_functorWithIndexFirst.Do(func() {
		functorWithIndexFirst = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v1_2.StrVal == "Just").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(__local_var_1_0, (*[1024]gopurs_runtime.Value)(v1_2.UnsafePtr)[0]))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}))
	})
	return functorWithIndexFirst
}

var functorWithIndexEither gopurs_runtime.Value
var once_functorWithIndexEither sync.Once
func Get_functorWithIndexEither() gopurs_runtime.Value {
	once_functorWithIndexEither.Do(func() {
		functorWithIndexEither = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
_ = __local_var_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(m_2.StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Left", (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(m_2.StrVal == "Right").IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(__local_var_1_0, (*[1024]gopurs_runtime.Value)(m_2.UnsafePtr)[0]))
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
	return functorWithIndexEither
}

var functorWithIndexDual gopurs_runtime.Value
var once_functorWithIndexDual sync.Once
func Get_functorWithIndexDual() gopurs_runtime.Value {
	once_functorWithIndexDual.Do(func() {
		functorWithIndexDual = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Dual.Get_functorDual()
}))
	})
	return functorWithIndexDual
}

var functorWithIndexDisj gopurs_runtime.Value
var once_functorWithIndexDisj sync.Once
func Get_functorWithIndexDisj() gopurs_runtime.Value {
	once_functorWithIndexDisj.Do(func() {
		functorWithIndexDisj = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Disj.Get_functorDisj()
}))
	})
	return functorWithIndexDisj
}

var functorWithIndexCoproduct gopurs_runtime.Value
var once_functorWithIndexCoproduct sync.Once
func Get_functorWithIndexCoproduct() gopurs_runtime.Value {
	once_functorWithIndexCoproduct.Do(func() {
		functorWithIndexCoproduct = gopurs_runtime.Func(func(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
if gopurs_runtime.Bool(v_5.StrVal == "Left").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply(__local_var_6_3, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]))
goto end_branch_5
} else {

}
}
{
if gopurs_runtime.Bool(v_5.StrVal == "Right").IntVal != 0 {
__t5 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(__local_var_7_4, (*[1024]gopurs_runtime.Value)(v_5.UnsafePtr)[0]))
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
return gopurs_runtime.Apply(f_5, gopurs_runtime.Constructor1("Left", x_7))
}))
_ = __local_var_7_6
__local_var_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex1_2, "mapWithIndex"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Constructor1("Right", x_8))
}))
_ = __local_var_8_7
var __t8 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v_6.StrVal == "Left").IntVal != 0 {
__t8 = gopurs_runtime.Constructor1("Left", gopurs_runtime.Apply(__local_var_7_6, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]))
goto end_branch_8
} else {

}
}
{
if gopurs_runtime.Bool(v_6.StrVal == "Right").IntVal != 0 {
__t8 = gopurs_runtime.Constructor1("Right", gopurs_runtime.Apply(__local_var_8_7, (*[1024]gopurs_runtime.Value)(v_6.UnsafePtr)[0]))
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
	return functorWithIndexCoproduct
}

var functorWithIndexConst gopurs_runtime.Value
var once_functorWithIndexConst sync.Once
func Get_functorWithIndexConst() gopurs_runtime.Value {
	once_functorWithIndexConst.Do(func() {
		functorWithIndexConst = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Const.Get_functorConst()
}))
	})
	return functorWithIndexConst
}

var functorWithIndexConj gopurs_runtime.Value
var once_functorWithIndexConj sync.Once
func Get_functorWithIndexConj() gopurs_runtime.Value {
	once_functorWithIndexConj.Do(func() {
		functorWithIndexConj = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Conj.Get_functorConj()
}))
	})
	return functorWithIndexConj
}

var functorWithIndexCompose gopurs_runtime.Value
var once_functorWithIndexCompose sync.Once
func Get_functorWithIndexCompose() gopurs_runtime.Value {
	once_functorWithIndexCompose.Do(func() {
		functorWithIndexCompose = gopurs_runtime.Func(func(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.Apply(f_5, gopurs_runtime.Constructor2("Tuple", x_7, b_8))
}))
}), v_6)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_4_2
}))
})
}()
})
	})
	return functorWithIndexCompose
}

var functorWithIndexArray gopurs_runtime.Value
var once_functorWithIndexArray sync.Once
func Get_functorWithIndexArray() gopurs_runtime.Value {
	once_functorWithIndexArray.Do(func() {
		functorWithIndexArray = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", Get_mapWithIndexArray(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}))
	})
	return functorWithIndexArray
}

var functorWithIndexApp gopurs_runtime.Value
var once_functorWithIndexApp sync.Once
func Get_functorWithIndexApp() gopurs_runtime.Value {
	once_functorWithIndexApp.Do(func() {
		functorWithIndexApp = gopurs_runtime.Func(func(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
	return functorWithIndexApp
}

var functorWithIndexAdditive gopurs_runtime.Value
var once_functorWithIndexAdditive sync.Once
func Get_functorWithIndexAdditive() gopurs_runtime.Value {
	once_functorWithIndexAdditive.Do(func() {
		functorWithIndexAdditive = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Additive.Get_functorAdditive()
}))
	})
	return functorWithIndexAdditive
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

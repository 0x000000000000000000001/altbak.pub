package Data_Traversable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Disj "gopurs/output/Data.Monoid.Disj"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Monoid_Conj "gopurs/output/Data.Monoid.Conj"
	pkg_Data_Functor_Compose "gopurs/output/Data.Functor.Compose"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Monoid_Additive "gopurs/output/Data.Monoid.Additive"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Functor_App "gopurs/output/Data.Functor.App"
	pkg_Data_Maybe_First "gopurs/output/Data.Maybe.First"
	pkg_Data_Maybe_Last "gopurs/output/Data.Maybe.Last"
	pkg_Data_Functor_Product "gopurs/output/Data.Functor.Product"
	pkg_Data_Traversable_Accum_Internal "gopurs/output/Data.Traversable.Accum.Internal"
)

var traverse gopurs_runtime.Value
var once_traverse sync.Once
func Get_traverse() gopurs_runtime.Value {
	once_traverse.Do(func() {
		traverse = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "traverse")
})
	})
	return traverse
}

var traversableTuple gopurs_runtime.Value
var once_traversableTuple sync.Once
func Get_traversableTuple() gopurs_runtime.Value {
	once_traversableTuple.Do(func() {
		traversableTuple = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.RecordGet(v_2, "value0")), gopurs_runtime.Apply(f_1, gopurs_runtime.RecordGet(v_2, "value1")))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.RecordGet(v_1, "value0")), gopurs_runtime.RecordGet(v_1, "value1"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableTuple()
}))
	})
	return traversableTuple
}

var traversableMultiplicative gopurs_runtime.Value
var once_traversableMultiplicative sync.Once
func Get_traversableMultiplicative() gopurs_runtime.Value {
	once_traversableMultiplicative.Do(func() {
		traversableMultiplicative = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), gopurs_runtime.Apply(f_1, v_2))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Multiplicative.Get_functorMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative()
}))
	})
	return traversableMultiplicative
}

var traversableMaybe gopurs_runtime.Value
var once_traversableMaybe sync.Once
func Get_traversableMaybe() gopurs_runtime.Value {
	once_traversableMaybe.Do(func() {
		traversableMaybe = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_2, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_2, "_tag").StrVal == "Just")).IntVal != 0 {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Apply(v_1, gopurs_runtime.RecordGet(v1_2, "value0")))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe.Get_Just(), gopurs_runtime.RecordGet(v_1, "value0"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMaybe()
}))
	})
	return traversableMaybe
}

var traversableIdentity gopurs_runtime.Value
var once_traversableIdentity sync.Once
func Get_traversableIdentity() gopurs_runtime.Value {
	once_traversableIdentity.Do(func() {
		traversableIdentity = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply(f_1, v_2))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Identity.Get_Identity(), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity()
}))
	})
	return traversableIdentity
}

var traversableEither gopurs_runtime.Value
var once_traversableEither sync.Once
func Get_traversableEither() gopurs_runtime.Value {
	once_traversableEither.Do(func() {
		traversableEither = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_2, "_tag").StrVal == "Left")).IntVal != 0 {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Left"), gopurs_runtime.RecordGet(v1_2, "value0")))
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v1_2, "_tag").StrVal == "Right")).IntVal != 0 {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Either.Get_Right(), gopurs_runtime.Apply(v_1, gopurs_runtime.RecordGet(v1_2, "value0")))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Left"), gopurs_runtime.RecordGet(v_1, "value0")))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_1, "_tag").StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Either.Get_Right(), gopurs_runtime.RecordGet(v_1, "value0"))
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_functorEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableEither()
}))
	})
	return traversableEither
}

var traversableDual gopurs_runtime.Value
var once_traversableDual sync.Once
func Get_traversableDual() gopurs_runtime.Value {
	once_traversableDual.Do(func() {
		traversableDual = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Dual.Get_Dual(), gopurs_runtime.Apply(f_1, v_2))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Dual.Get_Dual(), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Dual.Get_functorDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual()
}))
	})
	return traversableDual
}

var traversableDisj gopurs_runtime.Value
var once_traversableDisj sync.Once
func Get_traversableDisj() gopurs_runtime.Value {
	once_traversableDisj.Do(func() {
		traversableDisj = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Disj.Get_Disj(), gopurs_runtime.Apply(f_1, v_2))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Disj.Get_Disj(), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Disj.Get_functorDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDisj()
}))
	})
	return traversableDisj
}

var traversableConst gopurs_runtime.Value
var once_traversableConst sync.Once
func Get_traversableConst() gopurs_runtime.Value {
	once_traversableConst.Do(func() {
		traversableConst = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v1_2)
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Const.Get_functorConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConst()
}))
	})
	return traversableConst
}

var traversableConj gopurs_runtime.Value
var once_traversableConj sync.Once
func Get_traversableConj() gopurs_runtime.Value {
	once_traversableConj.Do(func() {
		traversableConj = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Conj.Get_Conj(), gopurs_runtime.Apply(f_1, v_2))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Conj.Get_Conj(), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Conj.Get_functorConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConj()
}))
	})
	return traversableConj
}

var traversableCompose gopurs_runtime.Value
var once_traversableCompose sync.Once
func Get_traversableCompose() gopurs_runtime.Value {
	once_traversableCompose.Do(func() {
		traversableCompose = gopurs_runtime.Func(func(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
traversableCompose:
for {
if false { continue traversableCompose }
var dictTraversable_0 = dictTraversable_0_loop
_ = dictTraversable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_2
functorCompose1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "map"), f_5), v_6)
}))
_ = functorCompose1_5_3
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_6_4
foldableCompose1_7_5 := gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_7 gopurs_runtime.Value, i_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_10_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_4, "foldr"), f_7)
_ = __local_var_10_6
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "foldr"), gopurs_runtime.Func2(func(b_11 gopurs_runtime.Value, a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_10_6, a_12, b_11)
}), i_8, v_9)
}), gopurs_runtime.Func3(func(f_7 gopurs_runtime.Value, i_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_4, "foldl"), f_7), i_8, v_9)
}), gopurs_runtime.Func(func(dictMonoid_7 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "foldMap"), dictMonoid_7)
_ = foldMap4_8_7
foldMap5_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_4, "foldMap"), dictMonoid_7)
_ = foldMap5_9_8
return gopurs_runtime.Func2(func(f_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap4_8_7, gopurs_runtime.Apply(foldMap5_9_8, f_10), v_11)
})
}))
_ = foldableCompose1_7_5
return gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
traverse4_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_8)
_ = traverse4_9_9
traverse5_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), dictApplicative_8)
_ = traverse5_10_10
return gopurs_runtime.Func2(func(f_11 gopurs_runtime.Value, v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_8, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Compose.Get_Compose(), gopurs_runtime.Apply2(traverse4_9_9, gopurs_runtime.Apply(traverse5_10_10, f_11), v_12))
})
}), gopurs_runtime.Func(func(dictApplicative_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_traversableCompose(), dictTraversable_0, dictTraversable1_3), "traverse"), dictApplicative_8, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCompose1_7_5
}))
})
}
}()
})
	})
	return traversableCompose
}

var traversableAdditive gopurs_runtime.Value
var once_traversableAdditive sync.Once
func Get_traversableAdditive() gopurs_runtime.Value {
	once_traversableAdditive.Do(func() {
		traversableAdditive = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Additive.Get_Additive(), gopurs_runtime.Apply(f_1, v_2))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Additive.Get_Additive(), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Additive.Get_functorAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableAdditive()
}))
	})
	return traversableAdditive
}

var sequenceDefault gopurs_runtime.Value
var once_sequenceDefault sync.Once
func Get_sequenceDefault() gopurs_runtime.Value {
	once_sequenceDefault.Do(func() {
		sequenceDefault = gopurs_runtime.Func2(func(dictTraversable_0 gopurs_runtime.Value, dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_1, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return sequenceDefault
}

var traversableArray gopurs_runtime.Value
var once_traversableArray sync.Once
func Get_traversableArray() gopurs_runtime.Value {
	once_traversableArray.Do(func() {
		traversableArray = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply3(Get_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_0, "pure"))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableArray(), "traverse"), dictApplicative_0, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
}))
	})
	return traversableArray
}

var sequence gopurs_runtime.Value
var once_sequence sync.Once
func Get_sequence() gopurs_runtime.Value {
	once_sequence.Do(func() {
		sequence = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "sequence")
})
	})
	return sequence
}

var traversableApp gopurs_runtime.Value
var once_traversableApp sync.Once
func Get_traversableApp() gopurs_runtime.Value {
	once_traversableApp.Do(func() {
		traversableApp = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_1
foldableApp_3_2 := gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, i_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "foldr"), f_3, i_4, v_5)
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, i_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "foldl"), f_3, i_4, v_5)
}), gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "foldMap"), dictMonoid_3)
}))
_ = foldableApp_3_2
return gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
traverse3_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_4)
_ = traverse3_5_3
return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_App.Get_App(), gopurs_runtime.Apply2(traverse3_5_3, f_6, v_7))
})
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
sequence3_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_4)
_ = sequence3_5_4
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_App.Get_App(), gopurs_runtime.Apply(sequence3_5_4, v_6))
})
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableApp_3_2
}))
})
	})
	return traversableApp
}

var traversableCoproduct gopurs_runtime.Value
var once_traversableCoproduct sync.Once
func Get_traversableCoproduct() gopurs_runtime.Value {
	once_traversableCoproduct.Do(func() {
		traversableCoproduct = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
foldableCoproduct_2_1 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableCoproduct_2_1
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_2
functorCoproduct1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_5)
_ = __local_var_7_4
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_2, "map"), f_5)
_ = __local_var_8_5
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6, "_tag").StrVal == "Left")).IntVal != 0 {
__t6 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Left"), gopurs_runtime.Apply(__local_var_7_4, gopurs_runtime.RecordGet(v_6, "value0")))
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_6, "_tag").StrVal == "Right")).IntVal != 0 {
__t6 = gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Right"), gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.RecordGet(v_6, "value0")))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
_ = functorCoproduct1_5_3
foldableCoproduct1_6_7 := gopurs_runtime.Apply(foldableCoproduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableCoproduct1_6_7
return gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func(func(dictApplicative_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_7, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_8
traverse4_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_7)
_ = traverse4_9_9
traverse5_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), dictApplicative_7)
_ = traverse5_10_10
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_12_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "map"), gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Left"), x_12)
}))
_ = __local_var_12_11
__local_var_13_12 := gopurs_runtime.Apply(traverse4_9_9, f_11)
_ = __local_var_13_12
__local_var_14_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "map"), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Right"), x_14)
}))
_ = __local_var_14_13
__local_var_15_14 := gopurs_runtime.Apply(traverse5_10_10, f_11)
_ = __local_var_15_14
return gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_16, "_tag").StrVal == "Left")).IntVal != 0 {
__t15 = gopurs_runtime.Apply(__local_var_12_11, gopurs_runtime.Apply(__local_var_13_12, gopurs_runtime.RecordGet(v2_16, "value0")))
goto end_branch_15
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_16, "_tag").StrVal == "Right")).IntVal != 0 {
__t15 = gopurs_runtime.Apply(__local_var_14_13, gopurs_runtime.Apply(__local_var_15_14, gopurs_runtime.RecordGet(v2_16, "value0")))
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
})
})
}), gopurs_runtime.Func(func(dictApplicative_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_7, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_16
__local_var_9_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_16, "map"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Left"), x_9)
}))
_ = __local_var_9_17
__local_var_10_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_7)
_ = __local_var_10_18
__local_var_11_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_16, "map"), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("_tag", "value0", gopurs_runtime.Str("Right"), x_11)
}))
_ = __local_var_11_19
__local_var_12_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "sequence"), dictApplicative_7)
_ = __local_var_12_20
return gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_13, "_tag").StrVal == "Left")).IntVal != 0 {
__t21 = gopurs_runtime.Apply(__local_var_9_17, gopurs_runtime.Apply(__local_var_10_18, gopurs_runtime.RecordGet(v2_13, "value0")))
goto end_branch_21
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_13, "_tag").StrVal == "Right")).IntVal != 0 {
__t21 = gopurs_runtime.Apply(__local_var_11_19, gopurs_runtime.Apply(__local_var_12_20, gopurs_runtime.RecordGet(v2_13, "value0")))
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
})
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCoproduct1_6_7
}))
})
})
	})
	return traversableCoproduct
}

var traversableFirst gopurs_runtime.Value
var once_traversableFirst sync.Once
func Get_traversableFirst() gopurs_runtime.Value {
	once_traversableFirst.Do(func() {
		traversableFirst = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe_First.Get_First(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_traversableMaybe(), "traverse"), dictApplicative_0, f_1, v_2))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe_First.Get_First(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMaybe(), "sequence"), dictApplicative_0, v_1))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableFirst()
}))
	})
	return traversableFirst
}

var traversableLast gopurs_runtime.Value
var once_traversableLast sync.Once
func Get_traversableLast() gopurs_runtime.Value {
	once_traversableLast.Do(func() {
		traversableLast = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe_Last.Get_Last(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_traversableMaybe(), "traverse"), dictApplicative_0, f_1, v_2))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe_Last.Get_Last(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMaybe(), "sequence"), dictApplicative_0, v_1))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableLast()
}))
	})
	return traversableLast
}

var traversableProduct gopurs_runtime.Value
var once_traversableProduct sync.Once
func Get_traversableProduct() gopurs_runtime.Value {
	once_traversableProduct.Do(func() {
		traversableProduct = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
foldableProduct_2_1 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableProduct_2_1
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_2
functorProduct1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Tuple"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_5, gopurs_runtime.RecordGet(v_6, "value0")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "map"), f_5, gopurs_runtime.RecordGet(v_6, "value1")))
}))
_ = functorProduct1_5_3
foldableProduct1_6_4 := gopurs_runtime.Apply(foldableProduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableProduct1_6_4
return gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func(func(dictApplicative_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_7, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_8_5
traverse4_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_7)
_ = traverse4_9_6
traverse5_10_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), dictApplicative_7)
_ = traverse5_10_7
return gopurs_runtime.Func2(func(f_11 gopurs_runtime.Value, v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_5, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_5, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product.Get_product(), gopurs_runtime.Apply2(traverse4_9_6, f_11, gopurs_runtime.RecordGet(v_12, "value0"))), gopurs_runtime.Apply2(traverse5_10_7, f_11, gopurs_runtime.RecordGet(v_12, "value1")))
})
}), gopurs_runtime.Func(func(dictApplicative_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_7, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_8_8
sequence4_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_7)
_ = sequence4_9_9
sequence5_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "sequence"), dictApplicative_7)
_ = sequence5_10_10
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_8, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product.Get_product(), gopurs_runtime.Apply(sequence4_9_9, gopurs_runtime.RecordGet(v_11, "value0"))), gopurs_runtime.Apply(sequence5_10_10, gopurs_runtime.RecordGet(v_11, "value1")))
})
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableProduct1_6_4
}))
})
})
	})
	return traversableProduct
}

var traverseDefault gopurs_runtime.Value
var once_traverseDefault sync.Once
func Get_traverseDefault() gopurs_runtime.Value {
	once_traverseDefault.Do(func() {
		traverseDefault = gopurs_runtime.Func2(func(dictTraversable_0 gopurs_runtime.Value, dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
sequence3_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_1)
_ = sequence3_2_0
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence3_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}), "map"), f_3, ta_4))
})
})
	})
	return traverseDefault
}

var mapAccumR gopurs_runtime.Value
var once_mapAccumR sync.Once
func Get_mapAccumR() gopurs_runtime.Value {
	once_mapAccumR.Do(func() {
		mapAccumR = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR())
_ = traverse2_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, s0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(traverse2_1_0, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, s_6, a_5)
}), xs_4, s0_3)
})
})
	})
	return mapAccumR
}

var scanr gopurs_runtime.Value
var once_scanr sync.Once
func Get_scanr() gopurs_runtime.Value {
	once_scanr.Do(func() {
		scanr = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
mapAccumR1_1_0 := gopurs_runtime.Apply(Get_mapAccumR(), dictTraversable_0)
_ = mapAccumR1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(mapAccumR1_1_0, gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_7_1 := gopurs_runtime.Apply2(f_2, a_6, b_5)
_ = b_prime_7_1
return gopurs_runtime.RecordDict2("accum", "value", b_prime_7_1, b_prime_7_1)
}), b0_3, xs_4), "value")
})
})
	})
	return scanr
}

var mapAccumL gopurs_runtime.Value
var once_mapAccumL sync.Once
func Get_mapAccumL() gopurs_runtime.Value {
	once_mapAccumL.Do(func() {
		mapAccumL = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL())
_ = traverse2_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, s0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(traverse2_1_0, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, s_6, a_5)
}), xs_4, s0_3)
})
})
	})
	return mapAccumL
}

var scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		scanl = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
mapAccumL1_1_0 := gopurs_runtime.Apply(Get_mapAccumL(), dictTraversable_0)
_ = mapAccumL1_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b0_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(mapAccumL1_1_0, gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_7_1 := gopurs_runtime.Apply2(f_2, b_5, a_6)
_ = b_prime_7_1
return gopurs_runtime.RecordDict2("accum", "value", b_prime_7_1, b_prime_7_1)
}), b0_3, xs_4), "value")
})
})
	})
	return scanl
}

var for_ gopurs_runtime.Value
var once_for_ sync.Once
func Get_for_() gopurs_runtime.Value {
	once_for_.Do(func() {
		for_ = gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, dictTraversable_1 gopurs_runtime.Value) gopurs_runtime.Value {
traverse2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_1, "traverse"), dictApplicative_0)
_ = traverse2_2_0
return gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(traverse2_2_0, f_4, x_3)
})
})
	})
	return for_
}

func Get_traverseArrayImpl() gopurs_runtime.Value {
	return TraverseArrayImpl
}

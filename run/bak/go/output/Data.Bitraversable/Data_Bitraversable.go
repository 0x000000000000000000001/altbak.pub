package Data_Bitraversable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Bifoldable "gopurs/output/Data.Bifoldable"
	pkg_Data_Functor_Joker "gopurs/output/Data.Functor.Joker"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Functor_Clown "gopurs/output/Data.Functor.Clown"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Functor_Flip "gopurs/output/Data.Functor.Flip"
	pkg_Data_Functor_Product2 "gopurs/output/Data.Functor.Product2"
)

var bitraverse gopurs_runtime.Value
var once_bitraverse sync.Once
func Get_bitraverse() gopurs_runtime.Value {
	once_bitraverse.Do(func() {
		bitraverse = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "bitraverse")
})
	})
	return bitraverse
}

var lfor gopurs_runtime.Value
var once_lfor sync.Once
func Get_lfor() gopurs_runtime.Value {
	once_lfor.Do(func() {
		lfor = gopurs_runtime.Func2(func(dictBitraversable_0 gopurs_runtime.Value, dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1)
_ = bitraverse2_2_0
pure_3_1 := gopurs_runtime.RecordGet(dictApplicative_1, "pure")
_ = pure_3_1
return gopurs_runtime.Func2(func(t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse2_2_0, f_5, pure_3_1, t_4)
})
})
	})
	return lfor
}

var ltraverse gopurs_runtime.Value
var once_ltraverse sync.Once
func Get_ltraverse() gopurs_runtime.Value {
	once_ltraverse.Do(func() {
		ltraverse = gopurs_runtime.Func2(func(dictBitraversable_0 gopurs_runtime.Value, dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1)
_ = bitraverse2_2_0
pure_3_1 := gopurs_runtime.RecordGet(dictApplicative_1, "pure")
_ = pure_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(bitraverse2_2_0, f_4, pure_3_1)
})
})
	})
	return ltraverse
}

var rfor gopurs_runtime.Value
var once_rfor sync.Once
func Get_rfor() gopurs_runtime.Value {
	once_rfor.Do(func() {
		rfor = gopurs_runtime.Func2(func(dictBitraversable_0 gopurs_runtime.Value, dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1)
_ = bitraverse2_2_0
pure_3_1 := gopurs_runtime.RecordGet(dictApplicative_1, "pure")
_ = pure_3_1
return gopurs_runtime.Func2(func(t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse2_2_0, pure_3_1, f_5, t_4)
})
})
	})
	return rfor
}

var rtraverse gopurs_runtime.Value
var once_rtraverse sync.Once
func Get_rtraverse() gopurs_runtime.Value {
	once_rtraverse.Do(func() {
		rtraverse = gopurs_runtime.Func2(func(dictBitraversable_0 gopurs_runtime.Value, dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1, gopurs_runtime.RecordGet(dictApplicative_1, "pure"))
})
	})
	return rtraverse
}

var bitraversableTuple gopurs_runtime.Value
var once_bitraversableTuple sync.Once
func Get_bitraversableTuple() gopurs_runtime.Value {
	once_bitraversableTuple.Do(func() {
		bitraversableTuple = gopurs_runtime.RecordDict4("bitraverse", "bisequence", "Bifunctor0", "Bifoldable1", gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Apply(f_2, gopurs_runtime.RecordGet(v_4, "value0"))), gopurs_runtime.Apply(g_3, gopurs_runtime.RecordGet(v_4, "value1")))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_1
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_1, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.RecordGet(v_2, "value0")), gopurs_runtime.RecordGet(v_2, "value1"))
})
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableTuple()
}))
	})
	return bitraversableTuple
}

var bitraversableJoker gopurs_runtime.Value
var once_bitraversableJoker sync.Once
func Get_bitraversableJoker() gopurs_runtime.Value {
	once_bitraversableJoker.Do(func() {
		bitraversableJoker = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifunctorJoker_2_1 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), g_3, v1_4)
}))
_ = bifunctorJoker_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_3_2
bifoldableJoker_4_3 := gopurs_runtime.RecordDict3("bifoldr", "bifoldl", "bifoldMap", gopurs_runtime.Func4(func(v_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldr"), r_5, u_6, v1_7)
}), gopurs_runtime.Func4(func(v_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldl"), r_5, u_6, v1_7)
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "foldMap"), dictMonoid_4)
_ = foldMap1_5_4
return gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_5_4, r_7, v1_8)
})
}))
_ = bifoldableJoker_4_3
return gopurs_runtime.RecordDict4("bitraverse", "bisequence", "Bifunctor0", "Bifoldable1", gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_5)
_ = traverse1_6_5
return gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Joker.Get_Joker(), gopurs_runtime.Apply2(traverse1_6_5, r_8, v1_9))
})
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_5)
_ = sequence1_6_6
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Joker.Get_Joker(), gopurs_runtime.Apply(sequence1_6_6, v_7))
})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker_2_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableJoker_4_3
}))
})
	})
	return bitraversableJoker
}

var bitraversableEither gopurs_runtime.Value
var once_bitraversableEither sync.Once
func Get_bitraversableEither() gopurs_runtime.Value {
	once_bitraversableEither.Do(func() {
		bitraversableEither = gopurs_runtime.RecordDict4("bitraverse", "bisequence", "Bifunctor0", "Bifoldable1", gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_4, "_tag").StrVal == "Left")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), pkg_Data_Either.Get_Left(), gopurs_runtime.Apply(v_2, gopurs_runtime.RecordGet(v2_4, "value0")))
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v2_4, "_tag").StrVal == "Right")).IntVal != 0 {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), pkg_Data_Either.Get_Right(), gopurs_runtime.Apply(v1_3, gopurs_runtime.RecordGet(v2_4, "value0")))
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
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_2
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Left")).IntVal != 0 {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "map"), pkg_Data_Either.Get_Left(), gopurs_runtime.RecordGet(v_2, "value0"))
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(v_2, "_tag").StrVal == "Right")).IntVal != 0 {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "map"), pkg_Data_Either.Get_Right(), gopurs_runtime.RecordGet(v_2, "value0"))
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
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableEither()
}))
	})
	return bitraversableEither
}

var bitraversableConst gopurs_runtime.Value
var once_bitraversableConst sync.Once
func Get_bitraversableConst() gopurs_runtime.Value {
	once_bitraversableConst.Do(func() {
		bitraversableConst = gopurs_runtime.RecordDict4("bitraverse", "bisequence", "Bifunctor0", "Bifoldable1", gopurs_runtime.Func4(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Const.Get_Const(), gopurs_runtime.Apply(f_1, v1_3))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Const.Get_Const(), v_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableConst()
}))
	})
	return bitraversableConst
}

var bitraversableClown gopurs_runtime.Value
var once_bitraversableClown sync.Once
func Get_bitraversableClown() gopurs_runtime.Value {
	once_bitraversableClown.Do(func() {
		bitraversableClown = gopurs_runtime.Func(func(dictTraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifunctorClown_2_1 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2, v1_4)
}))
_ = bifunctorClown_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_3_2
bifoldableClown_4_3 := gopurs_runtime.RecordDict3("bifoldr", "bifoldl", "bifoldMap", gopurs_runtime.Func4(func(l_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldr"), l_4, u_6, v1_7)
}), gopurs_runtime.Func4(func(l_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldl"), l_4, u_6, v1_7)
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "foldMap"), dictMonoid_4)
_ = foldMap1_5_4
return gopurs_runtime.Func3(func(l_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_5_4, l_6, v1_8)
})
}))
_ = bifoldableClown_4_3
return gopurs_runtime.RecordDict4("bitraverse", "bisequence", "Bifunctor0", "Bifoldable1", gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_5)
_ = traverse1_6_5
return gopurs_runtime.Func3(func(l_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Clown.Get_Clown(), gopurs_runtime.Apply2(traverse1_6_5, l_7, v1_9))
})
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_5)
_ = sequence1_6_6
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Clown.Get_Clown(), gopurs_runtime.Apply(sequence1_6_6, v_7))
})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorClown_2_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableClown_4_3
}))
})
	})
	return bitraversableClown
}

var bisequenceDefault gopurs_runtime.Value
var once_bisequenceDefault sync.Once
func Get_bisequenceDefault() gopurs_runtime.Value {
	once_bisequenceDefault.Do(func() {
		bisequenceDefault = gopurs_runtime.Func2(func(dictBitraversable_0 gopurs_runtime.Value, dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
})
	})
	return bisequenceDefault
}

var bisequence gopurs_runtime.Value
var once_bisequence sync.Once
func Get_bisequence() gopurs_runtime.Value {
	once_bisequence.Do(func() {
		bisequence = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "bisequence")
})
	})
	return bisequence
}

var bitraversableFlip gopurs_runtime.Value
var once_bitraversableFlip sync.Once
func Get_bitraversableFlip() gopurs_runtime.Value {
	once_bitraversableFlip.Do(func() {
		bitraversableFlip = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifunctorFlip_2_1 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), g_3, f_2, v_4)
}))
_ = bifunctorFlip_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{})
_ = __local_var_3_2
bifoldableFlip_4_3 := gopurs_runtime.RecordDict3("bifoldr", "bifoldl", "bifoldMap", gopurs_runtime.Func4(func(r_4 gopurs_runtime.Value, l_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_3_2, "bifoldr"), l_5, r_4, u_6, v_7)
}), gopurs_runtime.Func4(func(r_4 gopurs_runtime.Value, l_5 gopurs_runtime.Value, u_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_3_2, "bifoldl"), l_5, r_4, u_6, v_7)
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap2_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "bifoldMap"), dictMonoid_4)
_ = bifoldMap2_5_4
return gopurs_runtime.Func3(func(r_6 gopurs_runtime.Value, l_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bifoldMap2_5_4, l_7, r_6, v_8)
})
}))
_ = bifoldableFlip_4_3
return gopurs_runtime.RecordDict4("bitraverse", "bisequence", "Bifunctor0", "Bifoldable1", gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_5)
_ = bitraverse2_6_5
return gopurs_runtime.Func3(func(r_7 gopurs_runtime.Value, l_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Flip.Get_Flip(), gopurs_runtime.Apply3(bitraverse2_6_5, l_8, r_7, v_9))
})
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
bisequence2_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), dictApplicative_5)
_ = bisequence2_6_6
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Flip.Get_Flip(), gopurs_runtime.Apply(bisequence2_6_6, v_7))
})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorFlip_2_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableFlip_4_3
}))
})
	})
	return bitraversableFlip
}

var bitraversableProduct2 gopurs_runtime.Value
var once_bitraversableProduct2 sync.Once
func Get_bitraversableProduct2() gopurs_runtime.Value {
	once_bitraversableProduct2.Do(func() {
		bitraversableProduct2 = gopurs_runtime.Func(func(dictBitraversable_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifoldableProduct2_2_1 := gopurs_runtime.Apply(pkg_Data_Bifoldable.Get_bifoldableProduct2(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{}))
_ = bifoldableProduct2_2_1
return gopurs_runtime.Func(func(dictBitraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_4_2
bifunctorProduct21_5_3 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, g_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict3("_tag", "value0", "value1", gopurs_runtime.Str("Product2"), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_5, g_6, gopurs_runtime.RecordGet(v_7, "value0")), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_2, "bimap"), f_5, g_6, gopurs_runtime.RecordGet(v_7, "value1")))
}))
_ = bifunctorProduct21_5_3
bifoldableProduct21_6_4 := gopurs_runtime.Apply(bifoldableProduct2_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifoldable1"), gopurs_runtime.Value{}))
_ = bifoldableProduct21_6_4
return gopurs_runtime.RecordDict4("bitraverse", "bisequence", "Bifunctor0", "Bifoldable1", gopurs_runtime.Func(func(dictApplicative_7 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_7, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_8_5
bitraverse3_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_7)
_ = bitraverse3_9_6
bitraverse4_10_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "bitraverse"), dictApplicative_7)
_ = bitraverse4_10_7
return gopurs_runtime.Func3(func(l_11 gopurs_runtime.Value, r_12 gopurs_runtime.Value, v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_8_5, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_5, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product2.Get_Product2(), gopurs_runtime.Apply3(bitraverse3_9_6, l_11, r_12, gopurs_runtime.RecordGet(v_13, "value0"))), gopurs_runtime.Apply3(bitraverse4_10_7, l_11, r_12, gopurs_runtime.RecordGet(v_13, "value1")))
})
}), gopurs_runtime.Func(func(dictApplicative_7 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_7, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_8_8
bisequence3_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), dictApplicative_7)
_ = bisequence3_9_9
bisequence4_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "bisequence"), dictApplicative_7)
_ = bisequence4_10_10
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_8_8, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_8_8, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product2.Get_Product2(), gopurs_runtime.Apply(bisequence3_9_9, gopurs_runtime.RecordGet(v_11, "value0"))), gopurs_runtime.Apply(bisequence4_10_10, gopurs_runtime.RecordGet(v_11, "value1")))
})
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct21_5_3
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableProduct21_6_4
}))
})
})
	})
	return bitraversableProduct2
}

var bitraverseDefault gopurs_runtime.Value
var once_bitraverseDefault sync.Once
func Get_bitraverseDefault() gopurs_runtime.Value {
	once_bitraverseDefault.Do(func() {
		bitraverseDefault = gopurs_runtime.Func2(func(dictBitraversable_0 gopurs_runtime.Value, dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bisequence2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), dictApplicative_1)
_ = bisequence2_2_0
return gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(bisequence2_2_0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{}), "bimap"), f_3, g_4, t_5))
})
})
	})
	return bitraverseDefault
}

var bifor gopurs_runtime.Value
var once_bifor sync.Once
func Get_bifor() gopurs_runtime.Value {
	once_bifor.Do(func() {
		bifor = gopurs_runtime.Func2(func(dictBitraversable_0 gopurs_runtime.Value, dictApplicative_1 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1)
_ = bitraverse2_2_0
return gopurs_runtime.Func3(func(t_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse2_2_0, f_4, g_5, t_3)
})
})
	})
	return bifor
}



package Data_Bitraversable

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Bifoldable "gopurs/output/Data.Bifoldable"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Functor_Clown "gopurs/output/Data.Functor.Clown"
	pkg_Data_Functor_Flip "gopurs/output/Data.Functor.Flip"
	pkg_Data_Functor_Joker "gopurs/output/Data.Functor.Joker"
	pkg_Data_Functor_Product2 "gopurs/output/Data.Functor.Product2"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_identity1 gopurs_runtime.Value
var once_identity1 sync.Once
func Get_identity1() gopurs_runtime.Value {
	once_identity1.Do(func() {
		cache_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity1(x_0_box)
})
	})
	return cache_identity1
}

var cache_bitraverse gopurs_runtime.Value
var once_bitraverse sync.Once
func Get_bitraverse() gopurs_runtime.Value {
	once_bitraverse.Do(func() {
		cache_bitraverse = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bitraverse
}

var cache_lfor gopurs_runtime.Value
var once_lfor sync.Once
func Get_lfor() gopurs_runtime.Value {
	once_lfor.Do(func() {
		cache_lfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lfor(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_lfor
}

var cache_ltraverse gopurs_runtime.Value
var once_ltraverse sync.Once
func Get_ltraverse() gopurs_runtime.Value {
	once_ltraverse.Do(func() {
		cache_ltraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ltraverse(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_ltraverse
}

var cache_rfor gopurs_runtime.Value
var once_rfor sync.Once
func Get_rfor() gopurs_runtime.Value {
	once_rfor.Do(func() {
		cache_rfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rfor(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_rfor
}

var cache_rtraverse gopurs_runtime.Value
var once_rtraverse sync.Once
func Get_rtraverse() gopurs_runtime.Value {
	once_rtraverse.Do(func() {
		cache_rtraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rtraverse(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_rtraverse
}

var cache_bitraversableTuple gopurs_runtime.Value
var once_bitraversableTuple sync.Once
func Get_bitraversableTuple() gopurs_runtime.Value {
	once_bitraversableTuple.Do(func() {
		cache_bitraversableTuple = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorTuple()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_0.V1, gopurs_runtime.Apply2(Functor0_2_1.V0, pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_2 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_2
Functor0_2_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_3
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_1_2.V1, gopurs_runtime.Apply2(Functor0_2_3.V0, pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Apply(f_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)), gopurs_runtime.Apply(g_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
})
})
})
}))
	})
	return cache_bitraversableTuple
}

var cache_bitraversableJoker gopurs_runtime.Value
var once_bitraversableJoker sync.Once
func Get_bitraversableJoker() gopurs_runtime.Value {
	once_bitraversableJoker.Do(func() {
		cache_bitraversableJoker = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraversableJoker(dictTraversable_0_box)
})
	})
	return cache_bitraversableJoker
}

var cache_bitraversableEither gopurs_runtime.Value
var once_bitraversableEither sync.Once
func Get_bitraversableEither() gopurs_runtime.Value {
	once_bitraversableEither.Do(func() {
		cache_bitraversableEither = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorEither()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Either.Get_Left(), (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
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
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply2(Functor0_1_2.V0, pkg_Data_Either.Get_Left(), gopurs_runtime.Apply(v_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(Functor0_1_2.V0, pkg_Data_Either.Get_Right(), gopurs_runtime.Apply(v1_3, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0))
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
})
}))
	})
	return cache_bitraversableEither
}

var cache_bitraversableConst gopurs_runtime.Value
var once_bitraversableConst sync.Once
func Get_bitraversableConst() gopurs_runtime.Value {
	once_bitraversableConst.Do(func() {
		cache_bitraversableConst = gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifoldable.Get_bifoldableConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Bifunctor.Get_bifunctorConst()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Const.Get_Const(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Const.Get_Const(), gopurs_runtime.Apply(f_2, v1_4))
})
})
})
}))
	})
	return cache_bitraversableConst
}

var cache_bitraversableClown gopurs_runtime.Value
var once_bitraversableClown sync.Once
func Get_bitraversableClown() gopurs_runtime.Value {
	once_bitraversableClown.Do(func() {
		cache_bitraversableClown = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraversableClown(dictTraversable_0_box)
})
	})
	return cache_bitraversableClown
}

var cache_bisequenceDefault gopurs_runtime.Value
var once_bisequenceDefault sync.Once
func Get_bisequenceDefault() gopurs_runtime.Value {
	once_bisequenceDefault.Do(func() {
		cache_bisequenceDefault = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bisequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_bisequenceDefault
}

var cache_bisequence gopurs_runtime.Value
var once_bisequence sync.Once
func Get_bisequence() gopurs_runtime.Value {
	once_bisequence.Do(func() {
		cache_bisequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bisequence(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bisequence
}

var cache_bitraversableFlip gopurs_runtime.Value
var once_bitraversableFlip sync.Once
func Get_bitraversableFlip() gopurs_runtime.Value {
	once_bitraversableFlip.Do(func() {
		cache_bitraversableFlip = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraversableFlip(dictBitraversable_0_box)
})
	})
	return cache_bitraversableFlip
}

var cache_bitraversableProduct2 gopurs_runtime.Value
var once_bitraversableProduct2 sync.Once
func Get_bitraversableProduct2() gopurs_runtime.Value {
	once_bitraversableProduct2.Do(func() {
		cache_bitraversableProduct2 = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraversableProduct2(dictBitraversable_0_box)
})
	})
	return cache_bitraversableProduct2
}

var cache_bitraverseDefault gopurs_runtime.Value
var once_bitraverseDefault sync.Once
func Get_bitraverseDefault() gopurs_runtime.Value {
	once_bitraverseDefault.Do(func() {
		cache_bitraverseDefault = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverseDefault(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box))
})
	})
	return cache_bitraverseDefault
}

var cache_bifor gopurs_runtime.Value
var once_bifor sync.Once
func Get_bifor() gopurs_runtime.Value {
	once_bifor.Do(func() {
		cache_bifor = gopurs_runtime.Func5(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value, g_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifor(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), t_2_box, f_3_box, g_4_box)
})
	})
	return cache_bifor
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

var cache_apply__2623550860 gopurs_runtime.Value
var once_apply__2623550860 sync.Once
func Get_apply__2623550860() gopurs_runtime.Value {
	once_apply__2623550860.Do(func() {
		cache_apply__2623550860 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2623550860(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2623550860
}

var cache_apply__1941274444 gopurs_runtime.Value
var once_apply__1941274444 sync.Once
func Get_apply__1941274444() gopurs_runtime.Value {
	once_apply__1941274444.Do(func() {
		cache_apply__1941274444 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__1941274444(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__1941274444
}

var cache_bifoldableConst__1660068993 gopurs_runtime.Value
var once_bifoldableConst__1660068993 sync.Once
func Get_bifoldableConst__1660068993() gopurs_runtime.Value {
	once_bifoldableConst__1660068993.Do(func() {
		cache_bifoldableConst__1660068993 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v1_3)
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_2, v1_3)
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v1_3, z_2)
})
})
})
}))
	})
	return cache_bifoldableConst__1660068993
}

var cache_bifoldableEither__3757003471 gopurs_runtime.Value
var once_bifoldableEither__3757003471 sync.Once
func Get_bifoldableEither__3757003471() gopurs_runtime.Value {
	once_bifoldableEither__3757003471.Do(func() {
		cache_bifoldableEither__3757003471 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_3.Type == 9 && v2_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
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
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply2(v_0, v2_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(v1_1, v2_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
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
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 3711209382) {
__t2 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0, v2_2)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply2(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0, v2_2)
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
})
}))
	})
	return cache_bifoldableEither__3757003471
}

var cache_bifoldableTuple__2848462991 gopurs_runtime.Value
var once_bifoldableTuple__2848462991 sync.Once
func Get_bifoldableTuple__2848462991() gopurs_runtime.Value {
	once_bifoldableTuple__2848462991.Do(func() {
		cache_bifoldableTuple__2848462991 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply(g_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, z_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(g_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1, z_2))
})
})
})
}))
	})
	return cache_bifoldableTuple__2848462991
}

var cache_bifunctorConst__3156923452 gopurs_runtime.Value
var once_bifunctorConst__3156923452 sync.Once
func Get_bifunctorConst__3156923452() gopurs_runtime.Value {
	once_bifunctorConst__3156923452.Do(func() {
		cache_bifunctorConst__3156923452 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v1_2)
})
})
}))
	})
	return cache_bifunctorConst__3156923452
}

var cache_bifunctorEither__1585706332 gopurs_runtime.Value
var once_bifunctorEither__1585706332 sync.Once
func Get_bifunctorEither__1585706332() gopurs_runtime.Value {
	once_bifunctorEither__1585706332.Do(func() {
		cache_bifunctorEither__1585706332 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)})}
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
})
}))
	})
	return cache_bifunctorEither__1585706332
}

var cache_bifunctorTuple__553376860 gopurs_runtime.Value
var once_bifunctorTuple__553376860 sync.Once
func Get_bifunctorTuple__553376860() gopurs_runtime.Value {
	once_bifunctorTuple__553376860.Do(func() {
		cache_bifunctorTuple__553376860 = gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(g_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
})
})
}))
	})
	return cache_bifunctorTuple__553376860
}

var cache_bimap__132457202 gopurs_runtime.Value
var once_bimap__132457202 sync.Once
func Get_bimap__132457202() gopurs_runtime.Value {
	once_bimap__132457202.Do(func() {
		cache_bimap__132457202 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__132457202(gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bimap__132457202
}

var cache_bimap__4141710674 gopurs_runtime.Value
var once_bimap__4141710674 sync.Once
func Get_bimap__4141710674() gopurs_runtime.Value {
	once_bimap__4141710674.Do(func() {
		cache_bimap__4141710674 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bimap__4141710674(gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bimap__4141710674
}

var cache_bisequence__3517827200 gopurs_runtime.Value
var once_bisequence__3517827200 sync.Once
func Get_bisequence__3517827200() gopurs_runtime.Value {
	once_bisequence__3517827200.Do(func() {
		cache_bisequence__3517827200 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bisequence__3517827200(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bisequence__3517827200
}

var cache_bitraverse__3884078439 gopurs_runtime.Value
var once_bitraverse__3884078439 sync.Once
func Get_bitraverse__3884078439() gopurs_runtime.Value {
	once_bitraverse__3884078439.Do(func() {
		cache_bitraverse__3884078439 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse__3884078439(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bitraverse__3884078439
}

var cache_bitraverse__1091604167 gopurs_runtime.Value
var once_bitraverse__1091604167 sync.Once
func Get_bitraverse__1091604167() gopurs_runtime.Value {
	once_bitraverse__1091604167.Do(func() {
		cache_bitraverse__1091604167 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse__1091604167(gopurs_runtime.CoerceToStruct[Constructor_Bitraversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bitraverse__1091604167
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

var cache_map__2212490740 gopurs_runtime.Value
var once_map__2212490740 sync.Once
func Get_map__2212490740() gopurs_runtime.Value {
	once_map__2212490740.Do(func() {
		cache_map__2212490740 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2212490740(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2212490740
}

var cache_map__3061937460 gopurs_runtime.Value
var once_map__3061937460 sync.Once
func Get_map__3061937460() gopurs_runtime.Value {
	once_map__3061937460.Do(func() {
		cache_map__3061937460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3061937460(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3061937460
}

var cache_map__1974414836 gopurs_runtime.Value
var once_map__1974414836 sync.Once
func Get_map__1974414836() gopurs_runtime.Value {
	once_map__1974414836.Do(func() {
		cache_map__1974414836 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1974414836(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1974414836
}

var cache_map__3452061876 gopurs_runtime.Value
var once_map__3452061876 sync.Once
func Get_map__3452061876() gopurs_runtime.Value {
	once_map__3452061876.Do(func() {
		cache_map__3452061876 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3452061876(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3452061876
}

var cache_map__2511708020 gopurs_runtime.Value
var once_map__2511708020 sync.Once
func Get_map__2511708020() gopurs_runtime.Value {
	once_map__2511708020.Do(func() {
		cache_map__2511708020 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2511708020(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2511708020
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
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

type Constructor_Bitraversable[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3704227322] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Bitraversable[gopurs_runtime.Value])(ptr)
		switch key {
		case "Bifoldable1": return c.V0
		case "Bifunctor0": return c.V1
		case "bisequence": return c.V2
		case "bitraverse": return c.V3
		default: panic("Key not found in dictionary Constructor_Bitraversable: " + key)
		}
	}
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_bitraverse(dict_0_loop *Constructor_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_lfor(dictBitraversable_0_loop *Constructor_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
pure_2_0 := dictApplicative_1.V1
_ = pure_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, f_4, pure_2_0, t_3)
})
})
}

func Call_ltraverse(dictBitraversable_0_loop *Constructor_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
pure_2_0 := dictApplicative_1.V1
_ = pure_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, f_3, pure_2_0)
})
}

func Call_rfor(dictBitraversable_0_loop *Constructor_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
pure_2_0 := dictApplicative_1.V1
_ = pure_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, pure_2_0, f_4, t_3)
})
})
}

func Call_rtraverse(dictBitraversable_0_loop *Constructor_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, dictApplicative_1.V1)
}

func Call_bitraversableJoker(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
bifunctorJoker_1_0 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), g_3, v1_4)
})
})
}))
_ = bifunctorJoker_1_0
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
bifoldableJoker_2_2 := gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_3))}, r_5, v1_6)
})
})
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldl"), r_4, u_5, v1_6)
})
})
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldr"), r_4, u_5, v1_6)
})
})
})
}))
_ = bifoldableJoker_2_2
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableJoker_2_2
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_4.V0, pkg_Data_Functor_Joker.Get_Joker(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_5.V0, pkg_Data_Functor_Joker.Get_Joker(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, r_6, v1_7))
})
})
})
}))
}

func Call_bitraversableClown(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
bifunctorClown_1_0 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, v1_4)
})
})
}))
_ = bifunctorClown_1_0
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
bifoldableClown_2_2 := gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_3))}, l_4, v1_6)
})
})
})
}), gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldl"), l_3, u_5, v1_6)
})
})
})
}), gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldr"), l_3, u_5, v1_6)
})
})
})
}))
_ = bifoldableClown_2_2
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableClown_2_2
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorClown_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_4.V0, pkg_Data_Functor_Clown.Get_Clown(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func(func(l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_5.V0, pkg_Data_Functor_Clown.Get_Clown(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, l_5, v1_7))
})
})
})
}))
}

func Call_bisequenceDefault(dictBitraversable_0_loop *Constructor_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply3(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, Get_identity(), Get_identity1())
}

func Call_bisequence(dict_0_loop *Constructor_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_bitraversableFlip(dictBitraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
bifunctorFlip_1_0 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), g_3, f_2, v_4)
})
})
}))
_ = bifunctorFlip_1_0
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
bifoldableFlip_2_2 := gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_3, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_3))}, l_5, r_4, v_6)
})
})
})
}), gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_3, "bifoldl"), l_4, r_3, u_5, v_6)
})
})
})
}), gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_3, "bifoldr"), l_4, r_3, u_5, v_6)
})
})
})
}))
_ = bifoldableFlip_2_2
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableFlip_2_2
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorFlip_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_4.V0, pkg_Data_Functor_Flip.Get_Flip(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_5.V0, pkg_Data_Functor_Flip.Get_Flip(), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, l_6, r_5, v_7))
})
})
})
}))
}

func Call_bitraversableProduct2(dictBitraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifoldableProduct2_2_1 := gopurs_runtime.Apply(pkg_Data_Bifoldable.Get_bifoldableProduct2(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{}))
_ = bifoldableProduct2_2_1
return gopurs_runtime.Func(func(dictBitraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
bifunctorProduct21_4_2 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(&pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "bimap"), f_5, g_6, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_3, "bimap"), f_5, g_6, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1)})}
})
})
}))
_ = bifunctorProduct21_4_2
bifoldableProduct21_5_4 := gopurs_runtime.Apply(bifoldableProduct2_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifoldable1"), gopurs_runtime.Value{}))
_ = bifoldableProduct21_5_4
return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bifoldableProduct21_5_4
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorProduct21_4_2
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_7_5 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_5
Functor0_8_6 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_6
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_7_5.V1, gopurs_runtime.Apply2(Functor0_8_6.V0, pkg_Data_Functor_Product2.Get_Product2(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable1_3, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_7_7 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_7
Functor0_8_8 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_8
return gopurs_runtime.Func(func(l_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_7_7.V1, gopurs_runtime.Apply2(Functor0_8_8.V0, pkg_Data_Functor_Product2.Get_Product2(), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, l_9, r_10, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0)), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable1_3, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, l_9, r_10, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1))
})
})
})
}))
})
}

func Call_bitraverseDefault(dictBitraversable_0_loop *Constructor_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictBitraversable_0.V1, gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictBitraversable_0.V2, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_2))}, gopurs_runtime.Apply3(Bifunctor0_1_0.V0, f_3, g_4, t_5))
})
})
})
})
}

func Call_bifor(dictBitraversable_0_loop *Constructor_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], t_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value, g_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
var g_4 gopurs_runtime.Value = g_4_loop
_ = g_4
return gopurs_runtime.Apply4(dictBitraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, f_3, g_4, t_2)
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2623550860(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__1941274444(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bimap__132457202(dict_0_loop *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bimap__4141710674(dict_0_loop *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Bifunctor.Constructor_Bifunctor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bisequence__3517827200(dict_0_loop *Constructor_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_bitraverse__3884078439(dict_0_loop *Constructor_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_bitraverse__1091604167(dict_0_loop *Constructor_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bitraversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
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

func Call_map__2212490740(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3061937460(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1974414836(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3452061876(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2511708020(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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



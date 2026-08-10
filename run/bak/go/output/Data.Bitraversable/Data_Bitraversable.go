package Data_Bitraversable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Bifoldable "gopurs/output/Data.Bifoldable"
	pkg_Data_Bifunctor "gopurs/output/Data.Bifunctor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Functor_Joker "gopurs/output/Data.Functor.Joker"
	pkg_Data_Functor_Clown "gopurs/output/Data.Functor.Clown"
	pkg_Data_Functor_Flip "gopurs/output/Data.Functor.Flip"
	pkg_Data_Functor_Product2 "gopurs/output/Data.Functor.Product2"
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
return Call_bitraverse(dict_0_box)
})
	})
	return cache_bitraverse
}

var cache_bitraverse__gopurs_runtime_Value_3884078439 gopurs_runtime.Value
var once_bitraverse__gopurs_runtime_Value_3884078439 sync.Once
func Get_bitraverse__gopurs_runtime_Value_3884078439() gopurs_runtime.Value {
	once_bitraverse__gopurs_runtime_Value_3884078439.Do(func() {
		cache_bitraverse__gopurs_runtime_Value_3884078439 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse__gopurs_runtime_Value_3884078439(dict_0_box)
})
	})
	return cache_bitraverse__gopurs_runtime_Value_3884078439
}

var cache_lfor gopurs_runtime.Value
var once_lfor sync.Once
func Get_lfor() gopurs_runtime.Value {
	once_lfor.Do(func() {
		cache_lfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lfor(dictBitraversable_0_box, dictApplicative_1_box)
})
	})
	return cache_lfor
}

var cache_ltraverse gopurs_runtime.Value
var once_ltraverse sync.Once
func Get_ltraverse() gopurs_runtime.Value {
	once_ltraverse.Do(func() {
		cache_ltraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ltraverse(dictBitraversable_0_box, dictApplicative_1_box)
})
	})
	return cache_ltraverse
}

var cache_rfor gopurs_runtime.Value
var once_rfor sync.Once
func Get_rfor() gopurs_runtime.Value {
	once_rfor.Do(func() {
		cache_rfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rfor(dictBitraversable_0_box, dictApplicative_1_box)
})
	})
	return cache_rfor
}

var cache_rtraverse gopurs_runtime.Value
var once_rtraverse sync.Once
func Get_rtraverse() gopurs_runtime.Value {
	once_rtraverse.Do(func() {
		cache_rtraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_rtraverse(dictBitraversable_0_box, dictApplicative_1_box)
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
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_1, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)), gopurs_runtime.Apply(g_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
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
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), pkg_Data_Either.Get_Left(), (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), pkg_Data_Either.Get_Right(), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
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
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "map"), pkg_Data_Either.Get_Left(), gopurs_runtime.Apply(v_2, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "map"), pkg_Data_Either.Get_Right(), gopurs_runtime.Apply(v1_3, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0))
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
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Const.Get_Const(), v_1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Const.Get_Const(), gopurs_runtime.Apply(f_1, v1_3))
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
return Call_bisequenceDefault(dictBitraversable_0_box, dictApplicative_1_box)
})
	})
	return cache_bisequenceDefault
}

var cache_bisequence gopurs_runtime.Value
var once_bisequence sync.Once
func Get_bisequence() gopurs_runtime.Value {
	once_bisequence.Do(func() {
		cache_bisequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bisequence(dict_0_box)
})
	})
	return cache_bisequence
}

var cache_bisequence__gopurs_runtime_Value_3517827200 gopurs_runtime.Value
var once_bisequence__gopurs_runtime_Value_3517827200 sync.Once
func Get_bisequence__gopurs_runtime_Value_3517827200() gopurs_runtime.Value {
	once_bisequence__gopurs_runtime_Value_3517827200.Do(func() {
		cache_bisequence__gopurs_runtime_Value_3517827200 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bisequence__gopurs_runtime_Value_3517827200(dict_0_box)
})
	})
	return cache_bisequence__gopurs_runtime_Value_3517827200
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
		cache_bitraverseDefault = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverseDefault(dictBitraversable_0_box, dictApplicative_1_box)
})
	})
	return cache_bitraverseDefault
}

var cache_bifor gopurs_runtime.Value
var once_bifor sync.Once
func Get_bifor() gopurs_runtime.Value {
	once_bifor.Do(func() {
		cache_bifor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifor(dictBitraversable_0_box, dictApplicative_1_box)
})
	})
	return cache_bifor
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

func Call_bitraverse(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bitraverse")
}

func Call_bitraverse__gopurs_runtime_Value_3884078439(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bitraverse")
}

func Call_lfor(dictBitraversable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
bitraverse2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1)
_ = bitraverse2_2_0
pure_3_1 := gopurs_runtime.RecordGet(dictApplicative_1, "pure")
_ = pure_3_1
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse2_2_0, f_5, pure_3_1, t_4)
})
})
}

func Call_ltraverse(dictBitraversable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
bitraverse2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1)
_ = bitraverse2_2_0
pure_3_1 := gopurs_runtime.RecordGet(dictApplicative_1, "pure")
_ = pure_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(bitraverse2_2_0, f_4, pure_3_1)
})
}

func Call_rfor(dictBitraversable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
bitraverse2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1)
_ = bitraverse2_2_0
pure_3_1 := gopurs_runtime.RecordGet(dictApplicative_1, "pure")
_ = pure_3_1
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse2_2_0, pure_3_1, f_5, t_4)
})
})
}

func Call_rtraverse(dictBitraversable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1, gopurs_runtime.RecordGet(dictApplicative_1, "pure"))
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
foldMap1_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), dictMonoid_3)
_ = foldMap1_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_4_4, r_6, v1_7)
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
sequence1_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_3)
_ = sequence1_4_5
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Joker.Get_Joker(), gopurs_runtime.Apply(sequence1_4_5, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_3)
_ = traverse1_4_6
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Joker.Get_Joker(), gopurs_runtime.Apply2(traverse1_4_6, r_6, v1_7))
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
foldMap1_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), dictMonoid_3)
_ = foldMap1_4_4
return gopurs_runtime.Func(func(l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_4_4, l_5, v1_7)
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
sequence1_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_3)
_ = sequence1_4_5
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Clown.Get_Clown(), gopurs_runtime.Apply(sequence1_4_5, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_3)
_ = traverse1_4_6
return gopurs_runtime.Func(func(l_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Clown.Get_Clown(), gopurs_runtime.Apply2(traverse1_4_6, l_5, v1_7))
})
})
})
}))
}

func Call_bisequenceDefault(dictBitraversable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1, Get_identity(), Get_identity1())
}

func Call_bisequence(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bisequence")
}

func Call_bisequence__gopurs_runtime_Value_3517827200(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bisequence")
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
bifoldMap2_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "bifoldMap"), dictMonoid_3)
_ = bifoldMap2_4_4
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bifoldMap2_4_4, l_6, r_5, v_7)
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
bisequence2_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), dictApplicative_3)
_ = bisequence2_4_5
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Flip.Get_Flip(), gopurs_runtime.Apply(bisequence2_4_5, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
bitraverse2_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_3)
_ = bitraverse2_4_6
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Flip.Get_Flip(), gopurs_runtime.Apply3(bitraverse2_4_6, l_6, r_5, v_7))
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
Apply0_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_5
bisequence3_8_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), dictApplicative_6)
_ = bisequence3_8_6
bisequence4_9_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "bisequence"), dictApplicative_6)
_ = bisequence4_9_7
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_7_5, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_5, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product2.Get_Product2(), gopurs_runtime.Apply(bisequence3_8_6, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0)), gopurs_runtime.Apply(bisequence4_9_7, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_8
bitraverse3_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_6)
_ = bitraverse3_8_9
bitraverse4_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "bitraverse"), dictApplicative_6)
_ = bitraverse4_9_10
return gopurs_runtime.Func(func(l_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_7_8, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_8, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product2.Get_Product2(), gopurs_runtime.Apply3(bitraverse3_8_9, l_10, r_11, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V0)), gopurs_runtime.Apply3(bitraverse4_9_10, l_10, r_11, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_12.UnsafePtr).V1))
})
})
})
}))
})
}

func Call_bitraverseDefault(dictBitraversable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
bisequence2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), dictApplicative_1)
_ = bisequence2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(bisequence2_2_0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{}), "bimap"), f_3, g_4, t_5))
})
})
})
}

func Call_bifor(dictBitraversable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
bitraverse2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), dictApplicative_1)
_ = bitraverse2_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse2_2_0, f_4, g_5, t_3)
})
})
})
}



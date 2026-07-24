package Data_Bifoldable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		monoidEndo = func() gopurs_runtime.Value {
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, gopurs_runtime.Apply(v1_1, x_2))
}))
_ = semigroupEndo1_0_0
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}))
}()
	})
	return monoidEndo
}

var monoidDual gopurs_runtime.Value
var once_monoidDual sync.Once
func Get_monoidDual() gopurs_runtime.Value {
	once_monoidDual.Do(func() {
		monoidDual = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monoidEndo(), "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_0_0
semigroupDual1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "append"), v1_2, v_1)
}))
_ = semigroupDual1_1_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(Get_monoidEndo(), "mempty"), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_1_1
}))
}()
	})
	return monoidDual
}

var bifoldr gopurs_runtime.Value
var once_bifoldr sync.Once
func Get_bifoldr() gopurs_runtime.Value {
	once_bifoldr.Do(func() {
		bifoldr = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bifoldr")
}()
})
	})
	return bifoldr
}

var bitraverse_ gopurs_runtime.Value
var once_bitraverse_ sync.Once
func Get_bitraverse_() gopurs_runtime.Value {
	once_bitraverse_.Do(func() {
		bitraverse_ = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse_(dictBifoldable_0_box, dictApplicative_1_box)
})
	})
	return bitraverse_
}

var bifor_ gopurs_runtime.Value
var once_bifor_ sync.Once
func Get_bifor_() gopurs_runtime.Value {
	once_bifor_.Do(func() {
		bifor_ = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifor_(dictBifoldable_0_box, dictApplicative_1_box)
})
	})
	return bifor_
}

var bisequence_ gopurs_runtime.Value
var once_bisequence_ sync.Once
func Get_bisequence_() gopurs_runtime.Value {
	once_bisequence_.Do(func() {
		bisequence_ = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bisequence_(dictBifoldable_0_box, dictApplicative_1_box)
})
	})
	return bisequence_
}

var bifoldl gopurs_runtime.Value
var once_bifoldl sync.Once
func Get_bifoldl() gopurs_runtime.Value {
	once_bifoldl.Do(func() {
		bifoldl = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bifoldl")
}()
})
	})
	return bifoldl
}

var bifoldableTuple gopurs_runtime.Value
var once_bifoldableTuple sync.Once
func Get_bifoldableTuple() gopurs_runtime.Value {
	once_bifoldableTuple.Do(func() {
		bifoldableTuple = gopurs_runtime.RecordDict3("bifoldMap", "bifoldr", "bifoldl", gopurs_runtime.Func4(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]), gopurs_runtime.Apply(g_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1]))
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0], gopurs_runtime.Apply2(g_1, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1], z_2))
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, z_2, (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[0]), (*[1024]gopurs_runtime.Value)(v_3.UnsafePtr)[1])
}))
	})
	return bifoldableTuple
}

var bifoldableJoker gopurs_runtime.Value
var once_bifoldableJoker sync.Once
func Get_bifoldableJoker() gopurs_runtime.Value {
	once_bifoldableJoker.Do(func() {
		bifoldableJoker = gopurs_runtime.Func(func(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("bifoldr", "bifoldl", "bifoldMap", gopurs_runtime.Func4(func(v_1 gopurs_runtime.Value, r_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), r_2, u_3, v1_4)
}), gopurs_runtime.Func4(func(v_1 gopurs_runtime.Value, r_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), r_2, u_3, v1_4)
}), gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1)
_ = foldMap1_2_0
return gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_2_0, r_4, v1_5)
})
}))
}()
})
	})
	return bifoldableJoker
}

var bifoldableEither gopurs_runtime.Value
var once_bifoldableEither sync.Once
func Get_bifoldableEither() gopurs_runtime.Value {
	once_bifoldableEither.Do(func() {
		bifoldableEither = gopurs_runtime.RecordDict3("bifoldr", "bifoldl", "bifoldMap", gopurs_runtime.Func4(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value, v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v3_3.StrVal == "Left").IntVal != 0 {
__t0 = gopurs_runtime.Apply2(v_0, (*[1024]gopurs_runtime.Value)(v3_3.UnsafePtr)[0], v2_2)
goto end_branch_0
} else {

}
}
{
if gopurs_runtime.Bool(v3_3.StrVal == "Right").IntVal != 0 {
__t0 = gopurs_runtime.Apply2(v1_1, (*[1024]gopurs_runtime.Value)(v3_3.UnsafePtr)[0], v2_2)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Func4(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value, v3_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v3_3.StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply2(v_0, v2_2, (*[1024]gopurs_runtime.Value)(v3_3.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(v3_3.StrVal == "Right").IntVal != 0 {
__t1 = gopurs_runtime.Apply2(v1_1, v2_2, (*[1024]gopurs_runtime.Value)(v3_3.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}), gopurs_runtime.Func4(func(dictMonoid_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if gopurs_runtime.Bool(v2_3.StrVal == "Left").IntVal != 0 {
__t2 = gopurs_runtime.Apply(v_1, (*[1024]gopurs_runtime.Value)(v2_3.UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
if gopurs_runtime.Bool(v2_3.StrVal == "Right").IntVal != 0 {
__t2 = gopurs_runtime.Apply(v1_2, (*[1024]gopurs_runtime.Value)(v2_3.UnsafePtr)[0])
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}))
	})
	return bifoldableEither
}

var bifoldableConst gopurs_runtime.Value
var once_bifoldableConst sync.Once
func Get_bifoldableConst() gopurs_runtime.Value {
	once_bifoldableConst.Do(func() {
		bifoldableConst = gopurs_runtime.RecordDict3("bifoldr", "bifoldl", "bifoldMap", gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v1_3, z_2)
}), gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, z_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_2, v1_3)
}), gopurs_runtime.Func4(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v1_3)
}))
	})
	return bifoldableConst
}

var bifoldableClown gopurs_runtime.Value
var once_bifoldableClown sync.Once
func Get_bifoldableClown() gopurs_runtime.Value {
	once_bifoldableClown.Do(func() {
		bifoldableClown = gopurs_runtime.Func(func(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("bifoldr", "bifoldl", "bifoldMap", gopurs_runtime.Func4(func(l_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), l_1, u_3, v1_4)
}), gopurs_runtime.Func4(func(l_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), l_1, u_3, v1_4)
}), gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1)
_ = foldMap1_2_0
return gopurs_runtime.Func3(func(l_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_2_0, l_3, v1_5)
})
}))
}()
})
	})
	return bifoldableClown
}

var bifoldMapDefaultR gopurs_runtime.Value
var once_bifoldMapDefaultR sync.Once
func Get_bifoldMapDefaultR() gopurs_runtime.Value {
	once_bifoldMapDefaultR.Do(func() {
		bifoldMapDefaultR = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMapDefaultR(dictBifoldable_0_box, dictMonoid_1_box)
})
	})
	return bifoldMapDefaultR
}

var bifoldMapDefaultL gopurs_runtime.Value
var once_bifoldMapDefaultL sync.Once
func Get_bifoldMapDefaultL() gopurs_runtime.Value {
	once_bifoldMapDefaultL.Do(func() {
		bifoldMapDefaultL = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMapDefaultL(dictBifoldable_0_box, dictMonoid_1_box)
})
	})
	return bifoldMapDefaultL
}

var bifoldMap gopurs_runtime.Value
var once_bifoldMap sync.Once
func Get_bifoldMap() gopurs_runtime.Value {
	once_bifoldMap.Do(func() {
		bifoldMap = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bifoldMap")
}()
})
	})
	return bifoldMap
}

var bifoldableFlip gopurs_runtime.Value
var once_bifoldableFlip sync.Once
func Get_bifoldableFlip() gopurs_runtime.Value {
	once_bifoldableFlip.Do(func() {
		bifoldableFlip = gopurs_runtime.Func(func(dictBifoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
return gopurs_runtime.RecordDict3("bifoldr", "bifoldl", "bifoldMap", gopurs_runtime.Func4(func(r_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), l_2, r_1, u_3, v_4)
}), gopurs_runtime.Func4(func(r_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldl"), l_2, r_1, u_3, v_4)
}), gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), dictMonoid_1)
_ = bifoldMap2_2_0
return gopurs_runtime.Func3(func(r_3 gopurs_runtime.Value, l_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bifoldMap2_2_0, l_4, r_3, v_5)
})
}))
}()
})
	})
	return bifoldableFlip
}

var bifoldlDefault gopurs_runtime.Value
var once_bifoldlDefault sync.Once
func Get_bifoldlDefault() gopurs_runtime.Value {
	once_bifoldlDefault.Do(func() {
		bifoldlDefault = gopurs_runtime.Func(func(dictBifoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
bifoldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), Get_monoidDual())
_ = bifoldMap1_1_0
return gopurs_runtime.Func4(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, z_4 gopurs_runtime.Value, p_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(bifoldMap1_1_0, gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, a_7, x_6)
}), gopurs_runtime.Func2(func(x_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_3, a_7, x_6)
}), p_5, z_4)
})
}()
})
	})
	return bifoldlDefault
}

var bifoldrDefault gopurs_runtime.Value
var once_bifoldrDefault sync.Once
func Get_bifoldrDefault() gopurs_runtime.Value {
	once_bifoldrDefault.Do(func() {
		bifoldrDefault = gopurs_runtime.Func(func(dictBifoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
bifoldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), Get_monoidEndo())
_ = bifoldMap1_1_0
return gopurs_runtime.Func4(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, z_4 gopurs_runtime.Value, p_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(bifoldMap1_1_0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, x_6)
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_3, x_6)
}), p_5, z_4)
})
}()
})
	})
	return bifoldrDefault
}

var bifoldableProduct2 gopurs_runtime.Value
var once_bifoldableProduct2 sync.Once
func Get_bifoldableProduct2() gopurs_runtime.Value {
	once_bifoldableProduct2.Do(func() {
		bifoldableProduct2 = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBifoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldableProduct2(dictBifoldable_0_box, dictBifoldable1_1_box)
})
	})
	return bifoldableProduct2
}

var bifold gopurs_runtime.Value
var once_bifold sync.Once
func Get_bifold() gopurs_runtime.Value {
	once_bifold.Do(func() {
		bifold = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifold(dictBifoldable_0_box, dictMonoid_1_box)
})
	})
	return bifold
}

var biany gopurs_runtime.Value
var once_biany sync.Once
func Get_biany() gopurs_runtime.Value {
	once_biany.Do(func() {
		biany = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBooleanAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biany(dictBifoldable_0_box, dictBooleanAlgebra_1_box)
})
	})
	return biany
}

var biall gopurs_runtime.Value
var once_biall sync.Once
func Get_biall() gopurs_runtime.Value {
	once_biall.Do(func() {
		biall = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBooleanAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biall(dictBifoldable_0_box, dictBooleanAlgebra_1_box)
})
	})
	return biall
}

func Call_bitraverse_(dictBifoldable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
applySecond_2_0 := gopurs_runtime.Apply(pkg_Control_Apply.Get_applySecond(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_1, "Apply0"), gopurs_runtime.Value{}))
_ = applySecond_2_0
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_2_0, gopurs_runtime.Apply(f_3, x_5))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_2_0, gopurs_runtime.Apply(g_4, x_5))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_1, "pure"), pkg_Data_Unit.Get_unit()))
})
}

func Call_bifor_(dictBifoldable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
bitraverse_2_2_0 := Call_bitraverse_(dictBifoldable_0, dictApplicative_1)
_ = bitraverse_2_2_0
return gopurs_runtime.Func3(func(t_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(bitraverse_2_2_0, f_4, g_5, t_3)
})
}

func Call_bisequence_(dictBifoldable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(Call_bitraverse_(dictBifoldable_0, dictApplicative_1), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}

func Call_bifoldMapDefaultR(dictBifoldable_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_0
mempty_3_1 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_3_1
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "append"), gopurs_runtime.Apply(f_4, x_6))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "append"), gopurs_runtime.Apply(g_5, x_6))
}), mempty_3_1)
})
}

func Call_bifoldMapDefaultL(dictBifoldable_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_0
mempty_3_1 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_3_1
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldl"), gopurs_runtime.Func2(func(m_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), m_6, gopurs_runtime.Apply(f_4, a_7))
}), gopurs_runtime.Func2(func(m_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "append"), m_6, gopurs_runtime.Apply(g_5, b_7))
}), mempty_3_1)
})
}

func Call_bifoldableProduct2(dictBifoldable_0_loop gopurs_runtime.Value, dictBifoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bifoldableProduct2:
for {
if false { continue bifoldableProduct2 }
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBifoldable1_1 gopurs_runtime.Value = dictBifoldable1_1_loop
_ = dictBifoldable1_1
return gopurs_runtime.RecordDict3("bifoldr", "bifoldl", "bifoldMap", gopurs_runtime.Func4(func(l_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value, u_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply5(Get_bifoldrDefault(), Call_bifoldableProduct2(dictBifoldable_0, dictBifoldable1_1), l_2, r_3, u_4, m_5)
}), gopurs_runtime.Func4(func(l_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value, u_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply5(Get_bifoldlDefault(), Call_bifoldableProduct2(dictBifoldable_0, dictBifoldable1_1), l_2, r_3, u_4, m_5)
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
bifoldMap3_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), dictMonoid_2)
_ = bifoldMap3_3_0
bifoldMap4_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBifoldable1_1, "bifoldMap"), dictMonoid_2)
_ = bifoldMap4_4_1
return gopurs_runtime.Func3(func(l_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply3(bifoldMap3_3_0, l_5, r_6, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[0]), gopurs_runtime.Apply3(bifoldMap4_4_1, l_5, r_6, (*[1024]gopurs_runtime.Value)(v_7.UnsafePtr)[1]))
})
}))
}
}

func Call_bifold(dictBifoldable_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), dictMonoid_1, gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
}

func Call_biany(dictBifoldable_0_loop gopurs_runtime.Value, dictBooleanAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBooleanAlgebra_1 gopurs_runtime.Value = dictBooleanAlgebra_1_loop
_ = dictBooleanAlgebra_1
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_1, "HeytingAlgebra0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupDisj1_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "disj"), v_3, v1_4)
}))
_ = semigroupDisj1_3_2
bifoldMap2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(__local_var_2_1, "ff"), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_3_2
})))
_ = bifoldMap2_2_0
return gopurs_runtime.Func2(func(p_3 gopurs_runtime.Value, q_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(bifoldMap2_2_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(p_3, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(q_4, x_5)
}))
})
}

func Call_biall(dictBifoldable_0_loop gopurs_runtime.Value, dictBooleanAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBooleanAlgebra_1 gopurs_runtime.Value = dictBooleanAlgebra_1_loop
_ = dictBooleanAlgebra_1
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_1, "HeytingAlgebra0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupConj1_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "conj"), v_3, v1_4)
}))
_ = semigroupConj1_3_2
bifoldMap2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.RecordGet(__local_var_2_1, "tt"), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_3_2
})))
_ = bifoldMap2_2_0
return gopurs_runtime.Func2(func(p_3 gopurs_runtime.Value, q_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(bifoldMap2_2_0, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(p_3, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(q_4, x_5)
}))
})
}



package Data_Bifoldable

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_BooleanAlgebra "gopurs/output/Data.BooleanAlgebra"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Functor_Product2 "gopurs/output/Data.Functor.Product2"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
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

var cache_monoidDual gopurs_runtime.Value
var once_monoidDual sync.Once
func Get_monoidDual() gopurs_runtime.Value {
	once_monoidDual.Do(func() {
		cache_monoidDual = func() gopurs_runtime.Value {
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_1
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "compose"), v_1, v1_2)
})
}))
_ = semigroupEndo1_0_0
__local_var_1_2 := gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
_ = __local_var_1_2
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_4
semigroupDual1_2_3 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_4, "append"), v1_4, v_3)
})
}))
_ = semigroupDual1_2_3
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_2_3
}), gopurs_runtime.RecordGet(__local_var_1_2, "mempty"))))}
}()
	})
	return cache_monoidDual
}

var cache_monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		cache_monoidEndo = func() gopurs_runtime.Value {
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_1
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "compose"), v_1, v1_2)
})
}))
_ = semigroupEndo1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))))}
}()
	})
	return cache_monoidEndo
}

var cache_identity2 gopurs_runtime.Value
var once_identity2 sync.Once
func Get_identity2() gopurs_runtime.Value {
	once_identity2.Do(func() {
		cache_identity2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity2(x_0_box)
})
	})
	return cache_identity2
}

var cache_bifoldr gopurs_runtime.Value
var once_bifoldr sync.Once
func Get_bifoldr() gopurs_runtime.Value {
	once_bifoldr.Do(func() {
		cache_bifoldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldr(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldr
}

var cache_bitraverse_ gopurs_runtime.Value
var once_bitraverse_ sync.Once
func Get_bitraverse_() gopurs_runtime.Value {
	once_bitraverse_.Do(func() {
		cache_bitraverse_ = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse_(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_bitraverse_
}

var cache_bifor_ gopurs_runtime.Value
var once_bifor_ sync.Once
func Get_bifor_() gopurs_runtime.Value {
	once_bifor_.Do(func() {
		cache_bifor_ = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value, g_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifor_(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), t_2_box, f_3_box, g_4_box)
})
	})
	return cache_bifor_
}

var cache_bisequence_ gopurs_runtime.Value
var once_bisequence_ sync.Once
func Get_bisequence_() gopurs_runtime.Value {
	once_bisequence_.Do(func() {
		cache_bisequence_ = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bisequence_(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_bisequence_
}

var cache_bifoldl gopurs_runtime.Value
var once_bifoldl sync.Once
func Get_bifoldl() gopurs_runtime.Value {
	once_bifoldl.Do(func() {
		cache_bifoldl = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldl(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldl
}

var cache_bifoldableTuple gopurs_runtime.Value
var once_bifoldableTuple sync.Once
func Get_bifoldableTuple() gopurs_runtime.Value {
	once_bifoldableTuple.Do(func() {
		cache_bifoldableTuple = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_bifoldableTuple
}

var cache_bifoldableJoker gopurs_runtime.Value
var once_bifoldableJoker sync.Once
func Get_bifoldableJoker() gopurs_runtime.Value {
	once_bifoldableJoker.Do(func() {
		cache_bifoldableJoker = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldableJoker(dictFoldable_0_box)
})
	})
	return cache_bifoldableJoker
}

var cache_bifoldableEither gopurs_runtime.Value
var once_bifoldableEither sync.Once
func Get_bifoldableEither() gopurs_runtime.Value {
	once_bifoldableEither.Do(func() {
		cache_bifoldableEither = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_bifoldableEither
}

var cache_bifoldableConst gopurs_runtime.Value
var once_bifoldableConst sync.Once
func Get_bifoldableConst() gopurs_runtime.Value {
	once_bifoldableConst.Do(func() {
		cache_bifoldableConst = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_bifoldableConst
}

var cache_bifoldableClown gopurs_runtime.Value
var once_bifoldableClown sync.Once
func Get_bifoldableClown() gopurs_runtime.Value {
	once_bifoldableClown.Do(func() {
		cache_bifoldableClown = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldableClown(dictFoldable_0_box)
})
	})
	return cache_bifoldableClown
}

var cache_bifoldMapDefaultR gopurs_runtime.Value
var once_bifoldMapDefaultR sync.Once
func Get_bifoldMapDefaultR() gopurs_runtime.Value {
	once_bifoldMapDefaultR.Do(func() {
		cache_bifoldMapDefaultR = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_bifoldMapDefaultR
}

var cache_bifoldMapDefaultL gopurs_runtime.Value
var once_bifoldMapDefaultL sync.Once
func Get_bifoldMapDefaultL() gopurs_runtime.Value {
	once_bifoldMapDefaultL.Do(func() {
		cache_bifoldMapDefaultL = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMapDefaultL(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_bifoldMapDefaultL
}

var cache_bifoldMap gopurs_runtime.Value
var once_bifoldMap sync.Once
func Get_bifoldMap() gopurs_runtime.Value {
	once_bifoldMap.Do(func() {
		cache_bifoldMap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMap(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldMap
}

var cache_bifoldableFlip gopurs_runtime.Value
var once_bifoldableFlip sync.Once
func Get_bifoldableFlip() gopurs_runtime.Value {
	once_bifoldableFlip.Do(func() {
		cache_bifoldableFlip = gopurs_runtime.Func(func(dictBifoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldableFlip(dictBifoldable_0_box)
})
	})
	return cache_bifoldableFlip
}

var cache_bifoldlDefault gopurs_runtime.Value
var once_bifoldlDefault sync.Once
func Get_bifoldlDefault() gopurs_runtime.Value {
	once_bifoldlDefault.Do(func() {
		cache_bifoldlDefault = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldlDefault(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_bifoldlDefault
}

var cache_bifoldrDefault gopurs_runtime.Value
var once_bifoldrDefault sync.Once
func Get_bifoldrDefault() gopurs_runtime.Value {
	once_bifoldrDefault.Do(func() {
		cache_bifoldrDefault = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_bifoldrDefault
}

var cache_bifoldableProduct2 gopurs_runtime.Value
var once_bifoldableProduct2 sync.Once
func Get_bifoldableProduct2() gopurs_runtime.Value {
	once_bifoldableProduct2.Do(func() {
		cache_bifoldableProduct2 = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBifoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldableProduct2(dictBifoldable_0_box, dictBifoldable1_1_box)
})
	})
	return cache_bifoldableProduct2
}

var cache_bifold gopurs_runtime.Value
var once_bifold sync.Once
func Get_bifold() gopurs_runtime.Value {
	once_bifold.Do(func() {
		cache_bifold = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifold(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_bifold
}

var cache_biany gopurs_runtime.Value
var once_biany sync.Once
func Get_biany() gopurs_runtime.Value {
	once_biany.Do(func() {
		cache_biany = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBooleanAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biany(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_BooleanAlgebra.Constructor_BooleanAlgebra[gopurs_runtime.Value]](dictBooleanAlgebra_1_box))
})
	})
	return cache_biany
}

var cache_biall gopurs_runtime.Value
var once_biall sync.Once
func Get_biall() gopurs_runtime.Value {
	once_biall.Do(func() {
		cache_biall = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBooleanAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biall(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_BooleanAlgebra.Constructor_BooleanAlgebra[gopurs_runtime.Value]](dictBooleanAlgebra_1_box))
})
	})
	return cache_biall
}

var cache_pure__2935994064 gopurs_runtime.Value
var once_pure__2935994064 sync.Once
func Get_pure__2935994064() gopurs_runtime.Value {
	once_pure__2935994064.Do(func() {
		cache_pure__2935994064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__2935994064(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__2935994064
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

var cache_bifoldMap__4006734481 gopurs_runtime.Value
var once_bifoldMap__4006734481 sync.Once
func Get_bifoldMap__4006734481() gopurs_runtime.Value {
	once_bifoldMap__4006734481.Do(func() {
		cache_bifoldMap__4006734481 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMap__4006734481(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldMap__4006734481
}

var cache_bifoldMap__1302573585 gopurs_runtime.Value
var once_bifoldMap__1302573585 sync.Once
func Get_bifoldMap__1302573585() gopurs_runtime.Value {
	once_bifoldMap__1302573585.Do(func() {
		cache_bifoldMap__1302573585 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldMap__1302573585(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldMap__1302573585
}

var cache_bifoldl__2116322576 gopurs_runtime.Value
var once_bifoldl__2116322576 sync.Once
func Get_bifoldl__2116322576() gopurs_runtime.Value {
	once_bifoldl__2116322576.Do(func() {
		cache_bifoldl__2116322576 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldl__2116322576(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldl__2116322576
}

var cache_bifoldlDefault__2116322576 gopurs_runtime.Value
var once_bifoldlDefault__2116322576 sync.Once
func Get_bifoldlDefault__2116322576() gopurs_runtime.Value {
	once_bifoldlDefault__2116322576.Do(func() {
		cache_bifoldlDefault__2116322576 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldlDefault__2116322576(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_bifoldlDefault__2116322576
}

var cache_bifoldlDefault__2022005814 gopurs_runtime.Value
var once_bifoldlDefault__2022005814 sync.Once
func Get_bifoldlDefault__2022005814() gopurs_runtime.Value {
	once_bifoldlDefault__2022005814.Do(func() {
		cache_bifoldlDefault__2022005814 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldlDefault__2022005814(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_bifoldlDefault__2022005814
}

var cache_bifoldr__2116322576 gopurs_runtime.Value
var once_bifoldr__2116322576 sync.Once
func Get_bifoldr__2116322576() gopurs_runtime.Value {
	once_bifoldr__2116322576.Do(func() {
		cache_bifoldr__2116322576 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldr__2116322576(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldr__2116322576
}

var cache_bifoldr__656256240 gopurs_runtime.Value
var once_bifoldr__656256240 sync.Once
func Get_bifoldr__656256240() gopurs_runtime.Value {
	once_bifoldr__656256240.Do(func() {
		cache_bifoldr__656256240 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldr__656256240(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bifoldr__656256240
}

var cache_bifoldrDefault__2116322576 gopurs_runtime.Value
var once_bifoldrDefault__2116322576 sync.Once
func Get_bifoldrDefault__2116322576() gopurs_runtime.Value {
	once_bifoldrDefault__2116322576.Do(func() {
		cache_bifoldrDefault__2116322576 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldrDefault__2116322576(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_bifoldrDefault__2116322576
}

var cache_bifoldrDefault__2022005814 gopurs_runtime.Value
var once_bifoldrDefault__2022005814 sync.Once
func Get_bifoldrDefault__2022005814() gopurs_runtime.Value {
	once_bifoldrDefault__2022005814.Do(func() {
		cache_bifoldrDefault__2022005814 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldrDefault__2022005814(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_bifoldrDefault__2022005814
}

var cache_bitraverse___1288679761 gopurs_runtime.Value
var once_bitraverse___1288679761 sync.Once
func Get_bitraverse___1288679761() gopurs_runtime.Value {
	once_bitraverse___1288679761.Do(func() {
		cache_bitraverse___1288679761 = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse___1288679761(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_bitraverse___1288679761
}

var cache_bitraverse___648290481 gopurs_runtime.Value
var once_bitraverse___648290481 sync.Once
func Get_bitraverse___648290481() gopurs_runtime.Value {
	once_bitraverse___648290481.Do(func() {
		cache_bitraverse___648290481 = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bitraverse___648290481(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_bitraverse___648290481
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

var cache_unwrap__3267718003 gopurs_runtime.Value
var once_unwrap__3267718003 sync.Once
func Get_unwrap__3267718003() gopurs_runtime.Value {
	once_unwrap__3267718003.Do(func() {
		cache_unwrap__3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3267718003(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3267718003
}

var cache_unwrap__2928868755 gopurs_runtime.Value
var once_unwrap__2928868755 sync.Once
func Get_unwrap__2928868755() gopurs_runtime.Value {
	once_unwrap__2928868755.Do(func() {
		cache_unwrap__2928868755 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__2928868755(__eta0_0_box)
})
	})
	return cache_unwrap__2928868755
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

type Constructor_Bifoldable[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4001671834] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Bifoldable[gopurs_runtime.Value])(ptr)
		switch key {
		case "bifoldMap": return c.V0
		case "bifoldl": return c.V1
		case "bifoldr": return c.V2
		default: panic("Key not found in dictionary Constructor_Bifoldable: " + key)
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

func Call_identity2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_bifoldr(dict_0_loop *Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_bitraverse_(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
Apply0_2_0 := gopurs_runtime.Apply(dictApplicative_1.V0, gopurs_runtime.Value{})
_ = Apply0_2_0
Functor0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
applySecond_3_1 := gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(Functor0_3_2.V0, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return x_7
})
}), a_4), b_5)
})
})
_ = applySecond_3_1
Functor0_4_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
applySecond1_4_3 := gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(Functor0_4_4.V0, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
})
}), a_5), b_6)
})
})
_ = applySecond1_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictBifoldable_0.V2, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_3_1, gopurs_runtime.Apply(f_5, x_7))
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond1_4_3, gopurs_runtime.Apply(g_6, x_7))
}), gopurs_runtime.Apply(dictApplicative_1.V1, pkg_Data_Unit.Get_unit()))
})
})
}

func Call_bifor_(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], t_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value, g_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
var g_4 gopurs_runtime.Value = g_4_loop
_ = g_4
return gopurs_runtime.Apply3(Call_bitraverse_(dictBifoldable_0, dictApplicative_1), f_3, g_4, t_2)
}

func Call_bisequence_(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(Call_bitraverse_(dictBifoldable_0, dictApplicative_1), Get_identity(), Get_identity1())
}

func Call_bifoldl(dict_0_loop *Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bifoldableJoker(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, r_3, v1_4)
})
})
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), r_2, u_3, v1_4)
})
})
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), r_2, u_3, v1_4)
})
})
})
}))
}

func Call_bifoldableClown(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, l_2, v1_4)
})
})
})
}), gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), l_1, u_3, v1_4)
})
})
})
}), gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), l_1, u_3, v1_4)
})
})
})
}))
}

func Call_bifoldMapDefaultR(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
append_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}), "append")
_ = append_2_0
mempty_3_1 := dictMonoid_1.V1
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictBifoldable_0.V2, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(append_2_0, gopurs_runtime.Apply(f_4, x_6))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(append_2_0, gopurs_runtime.Apply(g_5, x_6))
}), mempty_3_1)
})
})
}

func Call_bifoldMapDefaultL(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
mempty_3_1 := dictMonoid_1.V1
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictBifoldable_0.V1, gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_2_0.V0, m_6, gopurs_runtime.Apply(f_4, a_7))
})
}), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_2_0.V0, m_6, gopurs_runtime.Apply(g_5, b_7))
})
}), mempty_3_1)
})
})
}

func Call_bifoldMap(dict_0_loop *Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bifoldableFlip(dictBifoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, l_3, r_2, v_4)
})
})
})
}), gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldl"), l_2, r_1, u_3, v_4)
})
})
})
}), gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldr"), l_2, r_1, u_3, v_4)
})
})
})
}))
}

func Call_bifoldlDefault(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidDual()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, a_6, x_5)
})
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_2, a_6, x_5)
})
}), p_4, z_3)
}

func Call_bifoldrDefault(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidEndo()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, x_5)
}), p_4, z_3)
}

func Call_bifoldableProduct2(dictBifoldable_0_loop gopurs_runtime.Value, dictBifoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bifoldableProduct2:
for {
if false { continue bifoldableProduct2 }
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBifoldable1_1 gopurs_runtime.Value = dictBifoldable1_1_loop
_ = dictBifoldable1_1
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_0
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_3_0.V0, gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, l_4, r_5, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable1_1, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, l_4, r_5, (*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
})
})
})
}), gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldlDefault(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]](Call_bifoldableProduct2(dictBifoldable_0, dictBifoldable1_1)))}), l_2, r_3, u_4, gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](m_5))})
})
})
})
}), gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifoldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Bifoldable[*pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]](Call_bifoldableProduct2(dictBifoldable_0, dictBifoldable1_1)))}), l_2, r_3, u_4, gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Functor_Product2.Constructor_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](m_5))})
})
})
})
}))
}
}

func Call_bifold(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply3(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, Get_identity2(), Get_identity2())
}

func Call_biany(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], dictBooleanAlgebra_1_loop *pkg_Data_BooleanAlgebra.Constructor_BooleanAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBooleanAlgebra_1 *pkg_Data_BooleanAlgebra.Constructor_BooleanAlgebra[gopurs_runtime.Value] = dictBooleanAlgebra_1_loop
_ = dictBooleanAlgebra_1
__local_var_2_1 := gopurs_runtime.Apply(dictBooleanAlgebra_1.V0, gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupDisj1_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "disj"), v_3, v1_4)
})
}))
_ = semigroupDisj1_3_2
monoidDisj_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_3_2
}), gopurs_runtime.RecordGet(__local_var_2_1, "ff")))
_ = monoidDisj_2_0
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(q_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply3(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(monoidDisj_2_0)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(p_3, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(q_4, x_5)
}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, x_6)
})
})
})
}

func Call_biall(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], dictBooleanAlgebra_1_loop *pkg_Data_BooleanAlgebra.Constructor_BooleanAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBooleanAlgebra_1 *pkg_Data_BooleanAlgebra.Constructor_BooleanAlgebra[gopurs_runtime.Value] = dictBooleanAlgebra_1_loop
_ = dictBooleanAlgebra_1
__local_var_2_1 := gopurs_runtime.Apply(dictBooleanAlgebra_1.V0, gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupConj1_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "conj"), v_3, v1_4)
})
}))
_ = semigroupConj1_3_2
monoidConj_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_3_2
}), gopurs_runtime.RecordGet(__local_var_2_1, "tt")))
_ = monoidConj_2_0
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(q_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply3(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(monoidConj_2_0)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(p_3, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(q_4, x_5)
}))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, x_6)
})
})
})
}

func Call_pure__2935994064(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bifoldMap__4006734481(dict_0_loop *Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bifoldMap__1302573585(dict_0_loop *Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(dict_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidDual()))})
}

func Call_bifoldl__2116322576(dict_0_loop *Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bifoldlDefault__2116322576(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidDual()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, a_6, x_5)
})
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_2, a_6, x_5)
})
}), p_4, z_3)
}

func Call_bifoldlDefault__2022005814(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidDual()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, a_6, x_5)
})
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_2, a_6, x_5)
})
}), p_4, z_3)
}

func Call_bifoldr__2116322576(dict_0_loop *Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_bifoldr__656256240(dict_0_loop *Constructor_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_bifoldrDefault__2116322576(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidEndo()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, x_5)
}), p_4, z_3)
}

func Call_bifoldrDefault__2022005814(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(dictBifoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidEndo()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, x_5)
}), p_4, z_3)
}

func Call_bitraverse___1288679761(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
Apply0_2_0 := gopurs_runtime.Apply(dictApplicative_1.V0, gopurs_runtime.Value{})
_ = Apply0_2_0
Functor0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
applySecond_3_1 := gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(Functor0_3_2.V0, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return x_7
})
}), a_4), b_5)
})
})
_ = applySecond_3_1
Functor0_4_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
applySecond1_4_3 := gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(Functor0_4_4.V0, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
})
}), a_5), b_6)
})
})
_ = applySecond1_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictBifoldable_0.V2, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_3_1, gopurs_runtime.Apply(f_5, x_7))
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond1_4_3, gopurs_runtime.Apply(g_6, x_7))
}), gopurs_runtime.Apply(dictApplicative_1.V1, pkg_Data_Unit.Get_unit()))
})
})
}

func Call_bitraverse___648290481(dictBifoldable_0_loop *Constructor_Bifoldable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
Apply0_2_0 := gopurs_runtime.Apply(dictApplicative_1.V0, gopurs_runtime.Value{})
_ = Apply0_2_0
Functor0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
applySecond_3_1 := gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(Functor0_3_2.V0, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return x_7
})
}), a_4), b_5)
})
})
_ = applySecond_3_1
Functor0_4_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
applySecond1_4_3 := gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(Functor0_4_4.V0, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
})
}), a_5), b_6)
})
})
_ = applySecond1_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(dictBifoldable_0.V2, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_3_1, gopurs_runtime.Apply(f_5, x_7))
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond1_4_3, gopurs_runtime.Apply(g_6, x_7))
}), gopurs_runtime.Apply(dictApplicative_1.V1, pkg_Data_Unit.Get_unit()))
})
})
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

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_unwrap__2928868755(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return __eta0_0
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}



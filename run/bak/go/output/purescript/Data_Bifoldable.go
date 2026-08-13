package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Bifoldable_identity gopurs_runtime.Value
var once_Data_Bifoldable_identity sync.Once
func Get_Data_Bifoldable_identity() gopurs_runtime.Value {
	once_Data_Bifoldable_identity.Do(func() {
		cache_Data_Bifoldable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_identity(x_0_box)
})
	})
	return cache_Data_Bifoldable_identity
}

var cache_Data_Bifoldable_identity1 gopurs_runtime.Value
var once_Data_Bifoldable_identity1 sync.Once
func Get_Data_Bifoldable_identity1() gopurs_runtime.Value {
	once_Data_Bifoldable_identity1.Do(func() {
		cache_Data_Bifoldable_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_identity1(x_0_box)
})
	})
	return cache_Data_Bifoldable_identity1
}

var cache_Data_Bifoldable_monoidDual gopurs_runtime.Value
var once_Data_Bifoldable_monoidDual sync.Once
func Get_Data_Bifoldable_monoidDual() gopurs_runtime.Value {
	once_Data_Bifoldable_monoidDual.Do(func() {
		cache_Data_Bifoldable_monoidDual = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Category_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_1
// TAST (Let): semigroupEndo1_0_0 -> gopurs_runtime.Value
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "compose"), v_1, v1_2)
})
}))
_ = semigroupEndo1_0_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
_ = __local_var_1_2
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_4
// TAST (Let): semigroupDual1_2_3 -> gopurs_runtime.Value
semigroupDual1_2_3 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_4, "append"), v1_4, v_3)
})
}))
_ = semigroupDual1_2_3
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_2_3
}), gopurs_runtime.RecordGet(__local_var_1_2, "mempty"))))}
}()
	})
	return cache_Data_Bifoldable_monoidDual
}

var cache_Data_Bifoldable_monoidEndo gopurs_runtime.Value
var once_Data_Bifoldable_monoidEndo sync.Once
func Get_Data_Bifoldable_monoidEndo() gopurs_runtime.Value {
	once_Data_Bifoldable_monoidEndo.Do(func() {
		cache_Data_Bifoldable_monoidEndo = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Category_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_1
// TAST (Let): semigroupEndo1_0_0 -> gopurs_runtime.Value
semigroupEndo1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "compose"), v_1, v1_2)
})
}))
_ = semigroupEndo1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_0_0
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))))}
}()
	})
	return cache_Data_Bifoldable_monoidEndo
}

var cache_Data_Bifoldable_identity2 gopurs_runtime.Value
var once_Data_Bifoldable_identity2 sync.Once
func Get_Data_Bifoldable_identity2() gopurs_runtime.Value {
	once_Data_Bifoldable_identity2.Do(func() {
		cache_Data_Bifoldable_identity2 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_identity2(x_0_box)
})
	})
	return cache_Data_Bifoldable_identity2
}

var cache_Data_Bifoldable_unwrap gopurs_runtime.Value
var once_Data_Bifoldable_unwrap sync.Once
func Get_Data_Bifoldable_unwrap() gopurs_runtime.Value {
	once_Data_Bifoldable_unwrap.Do(func() {
		cache_Data_Bifoldable_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Data_Bifoldable_unwrap
}

var cache_Data_Bifoldable_Bifoldable_dollarDict gopurs_runtime.Value
var once_Data_Bifoldable_Bifoldable_dollarDict sync.Once
func Get_Data_Bifoldable_Bifoldable_dollarDict() gopurs_runtime.Value {
	once_Data_Bifoldable_Bifoldable_dollarDict.Do(func() {
		cache_Data_Bifoldable_Bifoldable_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_Bifoldable_dollarDict(x_0_box)
})
	})
	return cache_Data_Bifoldable_Bifoldable_dollarDict
}

var cache_Data_Bifoldable_bifoldr gopurs_runtime.Value
var once_Data_Bifoldable_bifoldr sync.Once
func Get_Data_Bifoldable_bifoldr() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldr.Do(func() {
		cache_Data_Bifoldable_bifoldr = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldr
}

var cache_Data_Bifoldable_bitraverse_ gopurs_runtime.Value
var once_Data_Bifoldable_bitraverse_ sync.Once
func Get_Data_Bifoldable_bitraverse_() gopurs_runtime.Value {
	once_Data_Bifoldable_bitraverse_.Do(func() {
		cache_Data_Bifoldable_bitraverse_ = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bitraverse_(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_Data_Bifoldable_bitraverse_
}

var cache_Data_Bifoldable_bifor_ gopurs_runtime.Value
var once_Data_Bifoldable_bifor_ sync.Once
func Get_Data_Bifoldable_bifor_() gopurs_runtime.Value {
	once_Data_Bifoldable_bifor_.Do(func() {
		cache_Data_Bifoldable_bifor_ = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value, g_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifor_(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), t_2_box, f_3_box, g_4_box)
})
	})
	return cache_Data_Bifoldable_bifor_
}

var cache_Data_Bifoldable_bisequence_ gopurs_runtime.Value
var once_Data_Bifoldable_bisequence_ sync.Once
func Get_Data_Bifoldable_bisequence_() gopurs_runtime.Value {
	once_Data_Bifoldable_bisequence_.Do(func() {
		cache_Data_Bifoldable_bisequence_ = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bisequence_(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_Data_Bifoldable_bisequence_
}

var cache_Data_Bifoldable_bifoldl gopurs_runtime.Value
var once_Data_Bifoldable_bifoldl sync.Once
func Get_Data_Bifoldable_bifoldl() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldl.Do(func() {
		cache_Data_Bifoldable_bifoldl = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldl
}

var cache_Data_Bifoldable_bifoldableTuple gopurs_runtime.Value
var once_Data_Bifoldable_bifoldableTuple sync.Once
func Get_Data_Bifoldable_bifoldableTuple() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldableTuple.Do(func() {
		cache_Data_Bifoldable_bifoldableTuple = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply(g_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, z_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(g_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1, z_2))
})
})
})
}))
	})
	return cache_Data_Bifoldable_bifoldableTuple
}

var cache_Data_Bifoldable_bifoldableJoker gopurs_runtime.Value
var once_Data_Bifoldable_bifoldableJoker sync.Once
func Get_Data_Bifoldable_bifoldableJoker() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldableJoker.Do(func() {
		cache_Data_Bifoldable_bifoldableJoker = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldableJoker(dictFoldable_0_box)
})
	})
	return cache_Data_Bifoldable_bifoldableJoker
}

var cache_Data_Bifoldable_bifoldableEither gopurs_runtime.Value
var once_Data_Bifoldable_bifoldableEither sync.Once
func Get_Data_Bifoldable_bifoldableEither() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldableEither.Do(func() {
		cache_Data_Bifoldable_bifoldableEither = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_3.Type == 9 && v2_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
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
__t1 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
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
__t2 = gopurs_runtime.Apply2(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0, v2_2)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply2(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0, v2_2)
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
	return cache_Data_Bifoldable_bifoldableEither
}

var cache_Data_Bifoldable_bifoldableConst gopurs_runtime.Value
var once_Data_Bifoldable_bifoldableConst sync.Once
func Get_Data_Bifoldable_bifoldableConst() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldableConst.Do(func() {
		cache_Data_Bifoldable_bifoldableConst = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Bifoldable_bifoldableConst
}

var cache_Data_Bifoldable_bifoldableClown gopurs_runtime.Value
var once_Data_Bifoldable_bifoldableClown sync.Once
func Get_Data_Bifoldable_bifoldableClown() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldableClown.Do(func() {
		cache_Data_Bifoldable_bifoldableClown = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldableClown(dictFoldable_0_box)
})
	})
	return cache_Data_Bifoldable_bifoldableClown
}

var cache_Data_Bifoldable_bifoldMapDefaultR gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMapDefaultR sync.Once
func Get_Data_Bifoldable_bifoldMapDefaultR() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMapDefaultR.Do(func() {
		cache_Data_Bifoldable_bifoldMapDefaultR = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMapDefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_Data_Bifoldable_bifoldMapDefaultR
}

var cache_Data_Bifoldable_bifoldMapDefaultL gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMapDefaultL sync.Once
func Get_Data_Bifoldable_bifoldMapDefaultL() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMapDefaultL.Do(func() {
		cache_Data_Bifoldable_bifoldMapDefaultL = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMapDefaultL(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_Data_Bifoldable_bifoldMapDefaultL
}

var cache_Data_Bifoldable_bifoldMap gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMap sync.Once
func Get_Data_Bifoldable_bifoldMap() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMap.Do(func() {
		cache_Data_Bifoldable_bifoldMap = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldMap
}

var cache_Data_Bifoldable_bifoldableFlip gopurs_runtime.Value
var once_Data_Bifoldable_bifoldableFlip sync.Once
func Get_Data_Bifoldable_bifoldableFlip() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldableFlip.Do(func() {
		cache_Data_Bifoldable_bifoldableFlip = gopurs_runtime.Func(func(dictBifoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldableFlip(dictBifoldable_0_box)
})
	})
	return cache_Data_Bifoldable_bifoldableFlip
}

var cache_Data_Bifoldable_bifoldlDefault gopurs_runtime.Value
var once_Data_Bifoldable_bifoldlDefault sync.Once
func Get_Data_Bifoldable_bifoldlDefault() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldlDefault.Do(func() {
		cache_Data_Bifoldable_bifoldlDefault = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldlDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_Data_Bifoldable_bifoldlDefault
}

var cache_Data_Bifoldable_bifoldrDefault gopurs_runtime.Value
var once_Data_Bifoldable_bifoldrDefault sync.Once
func Get_Data_Bifoldable_bifoldrDefault() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldrDefault.Do(func() {
		cache_Data_Bifoldable_bifoldrDefault = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_Data_Bifoldable_bifoldrDefault
}

var cache_Data_Bifoldable_bifoldableProduct2 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldableProduct2 sync.Once
func Get_Data_Bifoldable_bifoldableProduct2() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldableProduct2.Do(func() {
		cache_Data_Bifoldable_bifoldableProduct2 = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBifoldable1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldableProduct2(dictBifoldable_0_box, dictBifoldable1_1_box)
})
	})
	return cache_Data_Bifoldable_bifoldableProduct2
}

var cache_Data_Bifoldable_bifold gopurs_runtime.Value
var once_Data_Bifoldable_bifold sync.Once
func Get_Data_Bifoldable_bifold() gopurs_runtime.Value {
	once_Data_Bifoldable_bifold.Do(func() {
		cache_Data_Bifoldable_bifold = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifold(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_Data_Bifoldable_bifold
}

var cache_Data_Bifoldable_biany gopurs_runtime.Value
var once_Data_Bifoldable_biany sync.Once
func Get_Data_Bifoldable_biany() gopurs_runtime.Value {
	once_Data_Bifoldable_biany.Do(func() {
		cache_Data_Bifoldable_biany = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBooleanAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_biany(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value]](dictBooleanAlgebra_1_box))
})
	})
	return cache_Data_Bifoldable_biany
}

var cache_Data_Bifoldable_biall gopurs_runtime.Value
var once_Data_Bifoldable_biall sync.Once
func Get_Data_Bifoldable_biall() gopurs_runtime.Value {
	once_Data_Bifoldable_biall.Do(func() {
		cache_Data_Bifoldable_biall = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictBooleanAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_biall(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value]](dictBooleanAlgebra_1_box))
})
	})
	return cache_Data_Bifoldable_biall
}

var cache_Data_Bifoldable_bifoldMap__4006734481 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMap__4006734481 sync.Once
func Get_Data_Bifoldable_bifoldMap__4006734481() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMap__4006734481.Do(func() {
		cache_Data_Bifoldable_bifoldMap__4006734481 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMap__4006734481(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldMap__4006734481
}

var cache_Data_Bifoldable_bifoldMap__1302573585 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMap__1302573585 sync.Once
func Get_Data_Bifoldable_bifoldMap__1302573585() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMap__1302573585.Do(func() {
		cache_Data_Bifoldable_bifoldMap__1302573585 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMap__1302573585(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldMap__1302573585
}

var cache_Data_Bifoldable_bifoldMap__581634711 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMap__581634711 sync.Once
func Get_Data_Bifoldable_bifoldMap__581634711() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMap__581634711.Do(func() {
		cache_Data_Bifoldable_bifoldMap__581634711 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMap__581634711(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldMap__581634711
}

var cache_Data_Bifoldable_bifoldMap__2444163767 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMap__2444163767 sync.Once
func Get_Data_Bifoldable_bifoldMap__2444163767() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMap__2444163767.Do(func() {
		cache_Data_Bifoldable_bifoldMap__2444163767 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMap__2444163767(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldMap__2444163767
}

var cache_Data_Bifoldable_bifoldMap__394664215 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMap__394664215 sync.Once
func Get_Data_Bifoldable_bifoldMap__394664215() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMap__394664215.Do(func() {
		cache_Data_Bifoldable_bifoldMap__394664215 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMap__394664215(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldMap__394664215
}

var cache_Data_Bifoldable_bifoldMap__3595123447 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMap__3595123447 sync.Once
func Get_Data_Bifoldable_bifoldMap__3595123447() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMap__3595123447.Do(func() {
		cache_Data_Bifoldable_bifoldMap__3595123447 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMap__3595123447(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldMap__3595123447
}

var cache_Data_Bifoldable_bifoldMap__237389305 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMap__237389305 sync.Once
func Get_Data_Bifoldable_bifoldMap__237389305() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMap__237389305.Do(func() {
		cache_Data_Bifoldable_bifoldMap__237389305 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMap__237389305(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldMap__237389305
}

var cache_Data_Bifoldable_bifoldMap__17102041 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldMap__17102041 sync.Once
func Get_Data_Bifoldable_bifoldMap__17102041() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldMap__17102041.Do(func() {
		cache_Data_Bifoldable_bifoldMap__17102041 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldMap__17102041(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldMap__17102041
}

var cache_Data_Bifoldable_bifoldableConst__1660068993 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldableConst__1660068993 sync.Once
func Get_Data_Bifoldable_bifoldableConst__1660068993() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldableConst__1660068993.Do(func() {
		cache_Data_Bifoldable_bifoldableConst__1660068993 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Bifoldable_bifoldableConst__1660068993
}

var cache_Data_Bifoldable_bifoldableEither__3757003471 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldableEither__3757003471 sync.Once
func Get_Data_Bifoldable_bifoldableEither__3757003471() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldableEither__3757003471.Do(func() {
		cache_Data_Bifoldable_bifoldableEither__3757003471 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v2_3.Type == 9 && v2_3.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_1, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_3.Type == 9 && v2_3.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_3.UnsafePtr).V0)
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
__t1 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
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
__t2 = gopurs_runtime.Apply2(v_0, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0, v2_2)
goto end_branch_2
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply2(v1_1, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0, v2_2)
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
	return cache_Data_Bifoldable_bifoldableEither__3757003471
}

var cache_Data_Bifoldable_bifoldableTuple__2848462991 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldableTuple__2848462991 sync.Once
func Get_Data_Bifoldable_bifoldableTuple__2848462991() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldableTuple__2848462991.Do(func() {
		cache_Data_Bifoldable_bifoldableTuple__2848462991 = gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply(g_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_1, gopurs_runtime.Apply2(f_0, z_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(g_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1, z_2))
})
})
})
}))
	})
	return cache_Data_Bifoldable_bifoldableTuple__2848462991
}

var cache_Data_Bifoldable_bifoldl__2116322576 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldl__2116322576 sync.Once
func Get_Data_Bifoldable_bifoldl__2116322576() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldl__2116322576.Do(func() {
		cache_Data_Bifoldable_bifoldl__2116322576 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldl__2116322576(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldl__2116322576
}

var cache_Data_Bifoldable_bifoldl__31590006 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldl__31590006 sync.Once
func Get_Data_Bifoldable_bifoldl__31590006() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldl__31590006.Do(func() {
		cache_Data_Bifoldable_bifoldl__31590006 = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldl__31590006(v_0_box, v1_1_box, v2_2_box, v3_3_box)
})
	})
	return cache_Data_Bifoldable_bifoldl__31590006
}

var cache_Data_Bifoldable_bifoldlDefault__2116322576 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldlDefault__2116322576 sync.Once
func Get_Data_Bifoldable_bifoldlDefault__2116322576() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldlDefault__2116322576.Do(func() {
		cache_Data_Bifoldable_bifoldlDefault__2116322576 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldlDefault__2116322576(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_Data_Bifoldable_bifoldlDefault__2116322576
}

var cache_Data_Bifoldable_bifoldlDefault__2022005814 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldlDefault__2022005814 sync.Once
func Get_Data_Bifoldable_bifoldlDefault__2022005814() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldlDefault__2022005814.Do(func() {
		cache_Data_Bifoldable_bifoldlDefault__2022005814 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldlDefault__2022005814(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](p_4_box))
})
	})
	return cache_Data_Bifoldable_bifoldlDefault__2022005814
}

var cache_Data_Bifoldable_bifoldr__2116322576 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldr__2116322576 sync.Once
func Get_Data_Bifoldable_bifoldr__2116322576() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldr__2116322576.Do(func() {
		cache_Data_Bifoldable_bifoldr__2116322576 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldr__2116322576(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldr__2116322576
}

var cache_Data_Bifoldable_bifoldr__656256240 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldr__656256240 sync.Once
func Get_Data_Bifoldable_bifoldr__656256240() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldr__656256240.Do(func() {
		cache_Data_Bifoldable_bifoldr__656256240 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldr__656256240(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldr__656256240
}

var cache_Data_Bifoldable_bifoldr__1631359728 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldr__1631359728 sync.Once
func Get_Data_Bifoldable_bifoldr__1631359728() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldr__1631359728.Do(func() {
		cache_Data_Bifoldable_bifoldr__1631359728 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldr__1631359728(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bifoldable_bifoldr__1631359728
}

var cache_Data_Bifoldable_bifoldr__31590006 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldr__31590006 sync.Once
func Get_Data_Bifoldable_bifoldr__31590006() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldr__31590006.Do(func() {
		cache_Data_Bifoldable_bifoldr__31590006 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldr__31590006(x_0_box)
})
	})
	return cache_Data_Bifoldable_bifoldr__31590006
}

var cache_Data_Bifoldable_bifoldrDefault__1989667951 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldrDefault__1989667951 sync.Once
func Get_Data_Bifoldable_bifoldrDefault__1989667951() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldrDefault__1989667951.Do(func() {
		cache_Data_Bifoldable_bifoldrDefault__1989667951 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldrDefault__1989667951(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_Data_Bifoldable_bifoldrDefault__1989667951
}

var cache_Data_Bifoldable_bifoldrDefault__2116322576 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldrDefault__2116322576 sync.Once
func Get_Data_Bifoldable_bifoldrDefault__2116322576() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldrDefault__2116322576.Do(func() {
		cache_Data_Bifoldable_bifoldrDefault__2116322576 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldrDefault__2116322576(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, p_4_box)
})
	})
	return cache_Data_Bifoldable_bifoldrDefault__2116322576
}

var cache_Data_Bifoldable_bifoldrDefault__2022005814 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldrDefault__2022005814 sync.Once
func Get_Data_Bifoldable_bifoldrDefault__2022005814() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldrDefault__2022005814.Do(func() {
		cache_Data_Bifoldable_bifoldrDefault__2022005814 = gopurs_runtime.Func5(func(dictBifoldable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, z_3_box gopurs_runtime.Value, p_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldrDefault__2022005814(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]](dictBifoldable_0_box), f_1_box, g_2_box, z_3_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](p_4_box))
})
	})
	return cache_Data_Bifoldable_bifoldrDefault__2022005814
}

var cache_Data_Bifoldable_bifoldrDefault__31590006 gopurs_runtime.Value
var once_Data_Bifoldable_bifoldrDefault__31590006 sync.Once
func Get_Data_Bifoldable_bifoldrDefault__31590006() gopurs_runtime.Value {
	once_Data_Bifoldable_bifoldrDefault__31590006.Do(func() {
		cache_Data_Bifoldable_bifoldrDefault__31590006 = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, z_2_box gopurs_runtime.Value, p_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldrDefault__31590006(f_0_box, g_1_box, z_2_box, p_3_box)
})
	})
	return cache_Data_Bifoldable_bifoldrDefault__31590006
}

var cache_Data_Bifoldable_bitraverse___1288679761 gopurs_runtime.Value
var once_Data_Bifoldable_bitraverse___1288679761 sync.Once
func Get_Data_Bifoldable_bitraverse___1288679761() gopurs_runtime.Value {
	once_Data_Bifoldable_bitraverse___1288679761.Do(func() {
		cache_Data_Bifoldable_bitraverse___1288679761 = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bitraverse___1288679761(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_Data_Bifoldable_bitraverse___1288679761
}

var cache_Data_Bifoldable_bitraverse___648290481 gopurs_runtime.Value
var once_Data_Bifoldable_bitraverse___648290481 sync.Once
func Get_Data_Bifoldable_bitraverse___648290481() gopurs_runtime.Value {
	once_Data_Bifoldable_bitraverse___648290481.Do(func() {
		cache_Data_Bifoldable_bitraverse___648290481 = gopurs_runtime.Func2(func(dictBifoldable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bitraverse___648290481(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](dictBifoldable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_Data_Bifoldable_bitraverse___648290481
}

type Constructor_Data_Bifoldable_Bifoldable[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4001671834] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "bifoldMap": return gopurs_runtime.Box(c.V0)
		case "bifoldl": return gopurs_runtime.Box(c.V1)
		case "bifoldr": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Bifoldable_Bifoldable: " + key)
		}
	}
}


func Call_Data_Bifoldable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bifoldable_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bifoldable_identity2(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bifoldable_Bifoldable_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bifoldable_bifoldr(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Bifoldable_bitraverse_(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
// TAST (Let): Apply0_2_0 -> gopurs_runtime.Value
Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_1.V0), gopurs_runtime.Value{})
_ = Apply0_2_0
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): applySecond_3_1 -> gopurs_runtime.Value
applySecond_3_1 := gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return x_7
})
}), a_4), b_5)
})
})
_ = applySecond_3_1
// TAST (Let): Functor0_4_4 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
// TAST (Let): applySecond1_4_3 -> gopurs_runtime.Value
applySecond1_4_3 := gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
})
}), a_5), b_6)
})
})
_ = applySecond1_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictBifoldable_0.V2), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_3_1, gopurs_runtime.Apply(f_5, x_7))
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond1_4_3, gopurs_runtime.Apply(g_6, x_7))
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_1.V1), Get_Data_Unit_unit()))
})
})
}

func Call_Data_Bifoldable_bifor_(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], t_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value, g_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
var g_4 gopurs_runtime.Value = g_4_loop
_ = g_4
return gopurs_runtime.Apply3(Call_Data_Bifoldable_bitraverse_(dictBifoldable_0, dictApplicative_1), f_3, g_4, t_2)
}

func Call_Data_Bifoldable_bisequence_(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(Call_Data_Bifoldable_bitraverse_(dictBifoldable_0, dictApplicative_1), Get_Data_Bifoldable_identity(), Get_Data_Bifoldable_identity1())
}

func Call_Data_Bifoldable_bifoldl(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Bifoldable_bifoldableJoker(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, r_3, v1_4)
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

func Call_Data_Bifoldable_bifoldableClown(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, l_2, v1_4)
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

func Call_Data_Bifoldable_bifoldMapDefaultR(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): append_2_0 -> gopurs_runtime.Value
append_2_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}), "append")
_ = append_2_0
// TAST (Let): mempty_3_1 -> gopurs_runtime.Value
mempty_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictBifoldable_0.V2), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(append_2_0, gopurs_runtime.Apply(f_4, x_6))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(append_2_0, gopurs_runtime.Apply(g_5, x_6))
}), mempty_3_1)
})
})
}

func Call_Data_Bifoldable_bifoldMapDefaultL(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): Semigroup0_2_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_1.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
// TAST (Let): mempty_3_1 -> gopurs_runtime.Value
mempty_3_1 := gopurs_runtime.Box(dictMonoid_1.V1)
_ = mempty_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictBifoldable_0.V1), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), m_6, gopurs_runtime.Apply(f_4, a_7))
})
}), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_0.V0), m_6, gopurs_runtime.Apply(g_5, b_7))
})
}), mempty_3_1)
})
})
}

func Call_Data_Bifoldable_bifoldMap(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifoldable_bifoldableFlip(dictBifoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, l_3, r_2, v_4)
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

func Call_Data_Bifoldable_bifoldlDefault(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(gopurs_runtime.Box(dictBifoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidDual()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, a_6, x_5)
})
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_2, a_6, x_5)
})
}), p_4, z_3)
}

func Call_Data_Bifoldable_bifoldrDefault(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(gopurs_runtime.Box(dictBifoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidEndo()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, x_5)
}), p_4, z_3)
}

func Call_Data_Bifoldable_bifoldableProduct2(dictBifoldable_0_loop gopurs_runtime.Value, dictBifoldable1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
bifoldableProduct2:
for {
if false { continue bifoldableProduct2 }
var dictBifoldable_0 gopurs_runtime.Value = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBifoldable1_1 gopurs_runtime.Value = dictBifoldable1_1_loop
_ = dictBifoldable1_1
return gopurs_runtime.RecordDict3("bifoldMap", "bifoldl", "bifoldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_0 -> *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]
Semigroup0_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_0
return gopurs_runtime.Func(func(l_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_0.V0), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable_0, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, l_4, r_5, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBifoldable1_1, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, l_4, r_5, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
})
})
})
}), gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldlDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]](Call_Data_Bifoldable_bifoldableProduct2(dictBifoldable_0, dictBifoldable1_1)))}), l_2, r_3, u_4, gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](m_5))})
})
})
})
}), gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(u_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bifoldable_bifoldrDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]](Call_Data_Bifoldable_bifoldableProduct2(dictBifoldable_0, dictBifoldable1_1)))}), l_2, r_3, u_4, gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](m_5))})
})
})
})
}))
}
}

func Call_Data_Bifoldable_bifold(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], dictMonoid_1_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictMonoid_1 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictBifoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, Get_Data_Bifoldable_identity2(), Get_Data_Bifoldable_identity2())
}

func Call_Data_Bifoldable_biany(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], dictBooleanAlgebra_1_loop *Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBooleanAlgebra_1 *Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value] = dictBooleanAlgebra_1_loop
_ = dictBooleanAlgebra_1
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictBooleanAlgebra_1.V0), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): semigroupDisj1_3_2 -> gopurs_runtime.Value
semigroupDisj1_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "disj"), v_3, v1_4)
})
}))
_ = semigroupDisj1_3_2
// TAST (Let): monoidDisj_2_0 -> *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]
monoidDisj_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_3_2
}), gopurs_runtime.RecordGet(__local_var_2_1, "ff")))
_ = monoidDisj_2_0
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(q_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply3(gopurs_runtime.Box(dictBifoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(monoidDisj_2_0)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Bifoldable_biall(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], dictBooleanAlgebra_1_loop *Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictBooleanAlgebra_1 *Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value] = dictBooleanAlgebra_1_loop
_ = dictBooleanAlgebra_1
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictBooleanAlgebra_1.V0), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): semigroupConj1_3_2 -> gopurs_runtime.Value
semigroupConj1_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "conj"), v_3, v1_4)
})
}))
_ = semigroupConj1_3_2
// TAST (Let): monoidConj_2_0 -> *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]
monoidConj_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_3_2
}), gopurs_runtime.RecordGet(__local_var_2_1, "tt")))
_ = monoidConj_2_0
return gopurs_runtime.Func(func(p_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(q_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply3(gopurs_runtime.Box(dictBifoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(monoidConj_2_0)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Bifoldable_bifoldMap__4006734481(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifoldable_bifoldMap__1302573585(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidDual()))})
}

func Call_Data_Bifoldable_bifoldMap__581634711(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifoldable_bifoldMap__2444163767(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidDual()))})
}

func Call_Data_Bifoldable_bifoldMap__394664215(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifoldable_bifoldMap__3595123447(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidEndo()))})
}

func Call_Data_Bifoldable_bifoldMap__237389305(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifoldable_bifoldMap__17102041(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Bifoldable_bifoldl__2116322576(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Bifoldable_bifoldl__31590006(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value, v3_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var v3_3 gopurs_runtime.Value = v3_3_loop
_ = v3_3
var __t0 gopurs_runtime.Value
{
if (v3_3.Type == 9 && v3_3.IntVal == 237113226) {
__t0 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v1_1, v2_2, (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_StartEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 1992629780) {
__t0 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0), (*Constructor_Data_Interval_DurationEnd[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2020675835) {
__t0 = gopurs_runtime.Apply2(v1_1, gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V1), (*Constructor_Data_Interval_StartDuration[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v3_3.Type == 9 && v3_3.IntVal == 2281256335) {
__t0 = gopurs_runtime.Apply2(v_0, v2_2, (*Constructor_Data_Interval_DurationOnly[gopurs_runtime.Value, gopurs_runtime.Value])(v3_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_Data_Bifoldable_bifoldlDefault__2116322576(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(gopurs_runtime.Box(dictBifoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidDual()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, a_6, x_5)
})
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_2, a_6, x_5)
})
}), p_4, z_3)
}

func Call_Data_Bifoldable_bifoldlDefault__2022005814(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop *Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 *Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(gopurs_runtime.Box(dictBifoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidDual()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, a_6, x_5)
})
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(g_2, a_6, x_5)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(p_4)}, z_3)
}

func Call_Data_Bifoldable_bifoldr__2116322576(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Bifoldable_bifoldr__656256240(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Bifoldable_bifoldr__1631359728(dict_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Bifoldable_bifoldr__31590006(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply2(Get_Data_Bifoldable_bifoldrDefault(), gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Interval_bifoldableInterval()))}, x_0)
}

func Call_Data_Bifoldable_bifoldrDefault__1989667951(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(gopurs_runtime.Box(dictBifoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidEndo()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, x_5)
}), p_4, z_3)
}

func Call_Data_Bifoldable_bifoldrDefault__2116322576(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 gopurs_runtime.Value = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(gopurs_runtime.Box(dictBifoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidEndo()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, x_5)
}), p_4, z_3)
}

func Call_Data_Bifoldable_bifoldrDefault__2022005814(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, z_3_loop gopurs_runtime.Value, p_4_loop *Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] = dictBifoldable_0_loop
_ = dictBifoldable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var z_3 gopurs_runtime.Value = z_3_loop
_ = z_3
var p_4 *Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = p_4_loop
_ = p_4
return gopurs_runtime.Apply5(gopurs_runtime.Box(dictBifoldable_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidEndo()))}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, x_5)
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_2, x_5)
}), gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer(p_4)}, z_3)
}

func Call_Data_Bifoldable_bifoldrDefault__31590006(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, z_2_loop gopurs_runtime.Value, p_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var z_2 gopurs_runtime.Value = z_2_loop
_ = z_2
var p_3 gopurs_runtime.Value = p_3_loop
_ = p_3
return gopurs_runtime.Apply5(gopurs_runtime.RecordGet(Get_Data_Interval_bifoldableInterval(), "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Bifoldable_monoidEndo()))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_4)
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, x_4)
}), p_3, z_2)
}

func Call_Data_Bifoldable_bitraverse___1288679761(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
// TAST (Let): Apply0_2_0 -> gopurs_runtime.Value
Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_1.V0), gopurs_runtime.Value{})
_ = Apply0_2_0
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): applySecond_3_1 -> gopurs_runtime.Value
applySecond_3_1 := gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return x_7
})
}), a_4), b_5)
})
})
_ = applySecond_3_1
// TAST (Let): Functor0_4_4 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
// TAST (Let): applySecond1_4_3 -> gopurs_runtime.Value
applySecond1_4_3 := gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
})
}), a_5), b_6)
})
})
_ = applySecond1_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictBifoldable_0.V2), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_3_1, gopurs_runtime.Apply(f_5, x_7))
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond1_4_3, gopurs_runtime.Apply(g_6, x_7))
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_1.V1), Get_Data_Unit_unit()))
})
})
}

func Call_Data_Bifoldable_bitraverse___648290481(dictBifoldable_0_loop *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBifoldable_0 *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] = dictBifoldable_0_loop
_ = dictBifoldable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
// TAST (Let): Apply0_2_0 -> gopurs_runtime.Value
Apply0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_1.V0), gopurs_runtime.Value{})
_ = Apply0_2_0
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): applySecond_3_1 -> gopurs_runtime.Value
applySecond_3_1 := gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return x_7
})
}), a_4), b_5)
})
})
_ = applySecond_3_1
// TAST (Let): Functor0_4_4 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_2_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
// TAST (Let): applySecond1_4_3 -> gopurs_runtime.Value
applySecond1_4_3 := gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_2_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
})
}), a_5), b_6)
})
})
_ = applySecond1_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictBifoldable_0.V2), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond_3_1, gopurs_runtime.Apply(f_5, x_7))
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(applySecond1_4_3, gopurs_runtime.Apply(g_6, x_7))
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_1.V1), Get_Data_Unit_unit()))
})
})
}



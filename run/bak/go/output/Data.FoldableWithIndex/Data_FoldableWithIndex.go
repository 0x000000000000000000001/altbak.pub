package Data_FoldableWithIndex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	unsafe "unsafe"
)

var cache_monoidEndo gopurs_runtime.Value
var once_monoidEndo sync.Once
func Get_monoidEndo() gopurs_runtime.Value {
	once_monoidEndo.Do(func() {
		cache_monoidEndo = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "Semigroupoid0"), gopurs_runtime.Value{})
_ = __local_var_0_0
semigroupEndo1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "compose"), v_1, v1_2)
}))
_ = semigroupEndo1_1_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupEndo1_1_1
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}()
	})
	return cache_monoidEndo
}

var cache_monoidDual gopurs_runtime.Value
var once_monoidDual sync.Once
func Get_monoidDual() gopurs_runtime.Value {
	once_monoidDual.Do(func() {
		cache_monoidDual = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monoidEndo(), "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_0_0
semigroupDual1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "append"), v1_2, v_1)
}))
_ = semigroupDual1_1_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDual1_1_1
}), gopurs_runtime.RecordGet(Get_monoidEndo(), "mempty"))
}()
	})
	return cache_monoidDual
}

var cache_foldrWithIndex gopurs_runtime.Value
var once_foldrWithIndex sync.Once
func Get_foldrWithIndex() gopurs_runtime.Value {
	once_foldrWithIndex.Do(func() {
		cache_foldrWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex(dict_0_box)
})
	})
	return cache_foldrWithIndex
}

var cache_traverseWithIndex_ gopurs_runtime.Value
var once_traverseWithIndex_ sync.Once
func Get_traverseWithIndex_() gopurs_runtime.Value {
	once_traverseWithIndex_.Do(func() {
		cache_traverseWithIndex_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex_(dictApplicative_0_box)
})
	})
	return cache_traverseWithIndex_
}

var cache_forWithIndex_ gopurs_runtime.Value
var once_forWithIndex_ sync.Once
func Get_forWithIndex_() gopurs_runtime.Value {
	once_forWithIndex_.Do(func() {
		cache_forWithIndex_ = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_forWithIndex_(dictApplicative_0_box)
})
	})
	return cache_forWithIndex_
}

var cache_foldrDefault gopurs_runtime.Value
var once_foldrDefault sync.Once
func Get_foldrDefault() gopurs_runtime.Value {
	once_foldrDefault.Do(func() {
		cache_foldrDefault = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrDefault(dictFoldableWithIndex_0_box, f_1_box)
})
	})
	return cache_foldrDefault
}

var cache_foldlWithIndex gopurs_runtime.Value
var once_foldlWithIndex sync.Once
func Get_foldlWithIndex() gopurs_runtime.Value {
	once_foldlWithIndex.Do(func() {
		cache_foldlWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex(dict_0_box)
})
	})
	return cache_foldlWithIndex
}

var cache_foldlDefault gopurs_runtime.Value
var once_foldlDefault sync.Once
func Get_foldlDefault() gopurs_runtime.Value {
	once_foldlDefault.Do(func() {
		cache_foldlDefault = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlDefault(dictFoldableWithIndex_0_box, f_1_box)
})
	})
	return cache_foldlDefault
}

var cache_foldableWithIndexTuple gopurs_runtime.Value
var once_foldableWithIndexTuple sync.Once
func Get_foldableWithIndexTuple() gopurs_runtime.Value {
	once_foldableWithIndexTuple.Do(func() {
		cache_foldableWithIndexTuple = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableTuple()
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, pkg_Data_Unit.Get_unit(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, pkg_Data_Unit.Get_unit(), z_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, pkg_Data_Unit.Get_unit(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
}))
	})
	return cache_foldableWithIndexTuple
}

var cache_foldableWithIndexMultiplicative gopurs_runtime.Value
var once_foldableWithIndexMultiplicative sync.Once
func Get_foldableWithIndexMultiplicative() gopurs_runtime.Value {
	once_foldableWithIndexMultiplicative.Do(func() {
		cache_foldableWithIndexMultiplicative = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMultiplicative(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMultiplicative(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMultiplicative(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexMultiplicative
}

var cache_foldableWithIndexMaybe gopurs_runtime.Value
var once_foldableWithIndexMaybe sync.Once
func Get_foldableWithIndexMaybe() gopurs_runtime.Value {
	once_foldableWithIndexMaybe.Do(func() {
		cache_foldableWithIndexMaybe = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMaybe()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexMaybe
}

var cache_foldableWithIndexLast gopurs_runtime.Value
var once_foldableWithIndexLast sync.Once
func Get_foldableWithIndexLast() gopurs_runtime.Value {
	once_foldableWithIndexLast.Do(func() {
		cache_foldableWithIndexLast = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableLast()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableLast(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableLast(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableLast(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexLast
}

var cache_foldableWithIndexIdentity gopurs_runtime.Value
var once_foldableWithIndexIdentity sync.Once
func Get_foldableWithIndexIdentity() gopurs_runtime.Value {
	once_foldableWithIndexIdentity.Do(func() {
		cache_foldableWithIndexIdentity = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity()
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, pkg_Data_Unit.Get_unit(), v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, pkg_Data_Unit.Get_unit(), z_1, v_2)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, pkg_Data_Unit.Get_unit(), v_2, z_1)
}))
	})
	return cache_foldableWithIndexIdentity
}

var cache_foldableWithIndexFirst gopurs_runtime.Value
var once_foldableWithIndexFirst sync.Once
func Get_foldableWithIndexFirst() gopurs_runtime.Value {
	once_foldableWithIndexFirst.Do(func() {
		cache_foldableWithIndexFirst = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableFirst()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableFirst(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableFirst(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableFirst(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexFirst
}

var cache_foldableWithIndexEither gopurs_runtime.Value
var once_foldableWithIndexEither sync.Once
func Get_foldableWithIndexEither() gopurs_runtime.Value {
	once_foldableWithIndexEither.Do(func() {
		cache_foldableWithIndexEither = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableEither()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t1 = mempty_1_0
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(v_2, pkg_Data_Unit.Get_unit(), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t2 = v1_1
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t2 = gopurs_runtime.Apply3(v_0, pkg_Data_Unit.Get_unit(), v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t3 = v1_1
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply3(v_0, pkg_Data_Unit.Get_unit(), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}))
	})
	return cache_foldableWithIndexEither
}

var cache_foldableWithIndexDual gopurs_runtime.Value
var once_foldableWithIndexDual sync.Once
func Get_foldableWithIndexDual() gopurs_runtime.Value {
	once_foldableWithIndexDual.Do(func() {
		cache_foldableWithIndexDual = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDual(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDual(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDual(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexDual
}

var cache_foldableWithIndexDisj gopurs_runtime.Value
var once_foldableWithIndexDisj sync.Once
func Get_foldableWithIndexDisj() gopurs_runtime.Value {
	once_foldableWithIndexDisj.Do(func() {
		cache_foldableWithIndexDisj = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDisj()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDisj(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDisj(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableDisj(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexDisj
}

var cache_foldableWithIndexConst gopurs_runtime.Value
var once_foldableWithIndexConst sync.Once
func Get_foldableWithIndexConst() gopurs_runtime.Value {
	once_foldableWithIndexConst.Do(func() {
		cache_foldableWithIndexConst = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConst()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
}))
	})
	return cache_foldableWithIndexConst
}

var cache_foldableWithIndexConj gopurs_runtime.Value
var once_foldableWithIndexConj sync.Once
func Get_foldableWithIndexConj() gopurs_runtime.Value {
	once_foldableWithIndexConj.Do(func() {
		cache_foldableWithIndexConj = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConj()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableConj(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableConj(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableConj(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexConj
}

var cache_foldableWithIndexAdditive gopurs_runtime.Value
var once_foldableWithIndexAdditive sync.Once
func Get_foldableWithIndexAdditive() gopurs_runtime.Value {
	once_foldableWithIndexAdditive.Do(func() {
		cache_foldableWithIndexAdditive = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableAdditive()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap8_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableAdditive(), "foldMap"), dictMonoid_0)
_ = foldMap8_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap8_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableAdditive(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableAdditive(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexAdditive
}

var cache_foldWithIndexM gopurs_runtime.Value
var once_foldWithIndexM sync.Once
func Get_foldWithIndexM() gopurs_runtime.Value {
	once_foldWithIndexM.Do(func() {
		cache_foldWithIndexM = gopurs_runtime.Func4(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonad_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, a0_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldWithIndexM(dictFoldableWithIndex_0_box, dictMonad_1_box, f_2_box, a0_3_box)
})
	})
	return cache_foldWithIndexM
}

var cache_foldMapWithIndexDefaultR gopurs_runtime.Value
var once_foldMapWithIndexDefaultR sync.Once
func Get_foldMapWithIndexDefaultR() gopurs_runtime.Value {
	once_foldMapWithIndexDefaultR.Do(func() {
		cache_foldMapWithIndexDefaultR = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapWithIndexDefaultR(dictFoldableWithIndex_0_box, dictMonoid_1_box)
})
	})
	return cache_foldMapWithIndexDefaultR
}

var cache_foldableWithIndexArray gopurs_runtime.Value
var once_foldableWithIndexArray sync.Once
func Get_foldableWithIndexArray() gopurs_runtime.Value {
	once_foldableWithIndexArray.Do(func() {
		cache_foldableWithIndexArray = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableWithIndexArray(), "foldrWithIndex"), gopurs_runtime.Func3(func(i_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(f_2, i_3, x_4), acc_5)
}), mempty_1_0)
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), gopurs_runtime.Func2(func(y_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, y_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
}), z_1)
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex"), pkg_Data_Tuple.Get_Tuple())
_ = __local_var_3_2
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1, gopurs_runtime.Apply(__local_var_3_2, x_4))
})
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_4 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_3_4
__local_var_4_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_4_5
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, __local_var_3_4, __local_var_4_5, y_5)
})
}), z_1)
_ = __local_var_2_3
__local_var_3_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex"), pkg_Data_Tuple.Get_Tuple())
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_3, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
}))
	})
	return cache_foldableWithIndexArray
}

var cache_foldMapWithIndexDefaultL gopurs_runtime.Value
var once_foldMapWithIndexDefaultL sync.Once
func Get_foldMapWithIndexDefaultL() gopurs_runtime.Value {
	once_foldMapWithIndexDefaultL.Do(func() {
		cache_foldMapWithIndexDefaultL = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapWithIndexDefaultL(dictFoldableWithIndex_0_box, dictMonoid_1_box)
})
	})
	return cache_foldMapWithIndexDefaultL
}

var cache_foldMapWithIndex gopurs_runtime.Value
var once_foldMapWithIndex sync.Once
func Get_foldMapWithIndex() gopurs_runtime.Value {
	once_foldMapWithIndex.Do(func() {
		cache_foldMapWithIndex = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapWithIndex(dict_0_box)
})
	})
	return cache_foldMapWithIndex
}

var cache_foldableWithIndexApp gopurs_runtime.Value
var once_foldableWithIndexApp sync.Once
func Get_foldableWithIndexApp() gopurs_runtime.Value {
	once_foldableWithIndexApp.Do(func() {
		cache_foldableWithIndexApp = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableWithIndexApp(dictFoldableWithIndex_0_box)
})
	})
	return cache_foldableWithIndexApp
}

var cache_foldableWithIndexCompose gopurs_runtime.Value
var once_foldableWithIndexCompose sync.Once
func Get_foldableWithIndexCompose() gopurs_runtime.Value {
	once_foldableWithIndexCompose.Do(func() {
		cache_foldableWithIndexCompose = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableWithIndexCompose(dictFoldableWithIndex_0_box)
})
	})
	return cache_foldableWithIndexCompose
}

var cache_foldableWithIndexCoproduct gopurs_runtime.Value
var once_foldableWithIndexCoproduct sync.Once
func Get_foldableWithIndexCoproduct() gopurs_runtime.Value {
	once_foldableWithIndexCoproduct.Do(func() {
		cache_foldableWithIndexCoproduct = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableWithIndexCoproduct(dictFoldableWithIndex_0_box)
})
	})
	return cache_foldableWithIndexCoproduct
}

var cache_foldableWithIndexProduct gopurs_runtime.Value
var once_foldableWithIndexProduct sync.Once
func Get_foldableWithIndexProduct() gopurs_runtime.Value {
	once_foldableWithIndexProduct.Do(func() {
		cache_foldableWithIndexProduct = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableWithIndexProduct(dictFoldableWithIndex_0_box)
})
	})
	return cache_foldableWithIndexProduct
}

var cache_foldlWithIndexDefault gopurs_runtime.Value
var once_foldlWithIndexDefault sync.Once
func Get_foldlWithIndexDefault() gopurs_runtime.Value {
	once_foldlWithIndexDefault.Do(func() {
		cache_foldlWithIndexDefault = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndexDefault(dictFoldableWithIndex_0_box)
})
	})
	return cache_foldlWithIndexDefault
}

var cache_foldrWithIndexDefault gopurs_runtime.Value
var once_foldrWithIndexDefault sync.Once
func Get_foldrWithIndexDefault() gopurs_runtime.Value {
	once_foldrWithIndexDefault.Do(func() {
		cache_foldrWithIndexDefault = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndexDefault(dictFoldableWithIndex_0_box)
})
	})
	return cache_foldrWithIndexDefault
}

var cache_surroundMapWithIndex gopurs_runtime.Value
var once_surroundMapWithIndex sync.Once
func Get_surroundMapWithIndex() gopurs_runtime.Value {
	once_surroundMapWithIndex.Do(func() {
		cache_surroundMapWithIndex = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_surroundMapWithIndex(dictFoldableWithIndex_0_box)
})
	})
	return cache_surroundMapWithIndex
}

var cache_foldMapDefault gopurs_runtime.Value
var once_foldMapDefault sync.Once
func Get_foldMapDefault() gopurs_runtime.Value {
	once_foldMapDefault.Do(func() {
		cache_foldMapDefault = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapDefault(dictFoldableWithIndex_0_box, dictMonoid_1_box)
})
	})
	return cache_foldMapDefault
}

var cache_findWithIndex gopurs_runtime.Value
var once_findWithIndex sync.Once
func Get_findWithIndex() gopurs_runtime.Value {
	once_findWithIndex.Do(func() {
		cache_findWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, p_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findWithIndex(dictFoldableWithIndex_0_box, p_1_box)
})
	})
	return cache_findWithIndex
}

var cache_findMapWithIndex gopurs_runtime.Value
var once_findMapWithIndex sync.Once
func Get_findMapWithIndex() gopurs_runtime.Value {
	once_findMapWithIndex.Do(func() {
		cache_findMapWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_findMapWithIndex(dictFoldableWithIndex_0_box, f_1_box)
})
	})
	return cache_findMapWithIndex
}

var cache_anyWithIndex gopurs_runtime.Value
var once_anyWithIndex sync.Once
func Get_anyWithIndex() gopurs_runtime.Value {
	once_anyWithIndex.Do(func() {
		cache_anyWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_anyWithIndex(dictFoldableWithIndex_0_box, dictHeytingAlgebra_1_box)
})
	})
	return cache_anyWithIndex
}

var cache_allWithIndex gopurs_runtime.Value
var once_allWithIndex sync.Once
func Get_allWithIndex() gopurs_runtime.Value {
	once_allWithIndex.Do(func() {
		cache_allWithIndex = gopurs_runtime.Func2(func(dictFoldableWithIndex_0_box gopurs_runtime.Value, dictHeytingAlgebra_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_allWithIndex(dictFoldableWithIndex_0_box, dictHeytingAlgebra_1_box)
})
	})
	return cache_allWithIndex
}

func Call_foldrWithIndex(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldrWithIndex")
}

func Call_traverseWithIndex_(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(dictFoldableWithIndex_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_2, "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(f_3, i_4)
_ = __local_var_5_1
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_2 := gopurs_runtime.Apply(__local_var_5_1, x_6)
_ = __local_var_7_2
return gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_identity()
}), __local_var_7_2), b_8)
})
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), pkg_Data_Unit.Get_unit()))
})
}

func Call_forWithIndex_(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
traverseWithIndex_1_1_0 := Call_traverseWithIndex_(dictApplicative_0)
_ = traverseWithIndex_1_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(traverseWithIndex_1_1_0, dictFoldableWithIndex_2)
_ = __local_var_3_1
return gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_3_1, a_5, b_4)
})
})
}

func Call_foldrDefault(dictFoldableWithIndex_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
}

func Call_foldlWithIndex(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldlWithIndex")
}

func Call_foldlDefault(dictFoldableWithIndex_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return f_1
}))
}

func Call_foldWithIndexM(dictFoldableWithIndex_0_loop gopurs_runtime.Value, dictMonad_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, a0_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonad_1 gopurs_runtime.Value = dictMonad_1_loop
_ = dictMonad_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var a0_3 gopurs_runtime.Value = a0_3_loop
_ = a0_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, ma_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_0 := gopurs_runtime.Apply(f_2, i_4)
_ = __local_var_7_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}), "bind"), ma_5, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_7_0, a_8, b_6)
}))
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Applicative0"), gopurs_runtime.Value{}), "pure"), a0_3))
}

func Call_foldMapWithIndexDefaultR(dictFoldableWithIndex_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
mempty_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value, acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(f_3, i_4, x_5), acc_6)
}), mempty_2_0)
})
}

func Call_foldMapWithIndexDefaultL(dictFoldableWithIndex_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
mempty_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func3(func(i_4 gopurs_runtime.Value, acc_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), acc_5, gopurs_runtime.Apply2(f_3, i_4, x_6))
}), mempty_2_0)
})
}

func Call_foldMapWithIndex(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "foldMapWithIndex")
}

func Call_foldableWithIndexApp(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_0
foldableApp_2_1 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "foldMap"), dictMonoid_2)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldl"), f_2, i_3, v_4)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, i_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldr"), f_2, i_3, v_4)
}))
_ = foldableApp_2_1
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableApp_2_1
}), gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), dictMonoid_3)
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, z_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), f_3, z_4, v_5)
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, z_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), f_3, z_4, v_5)
}))
}

func Call_foldableWithIndexCompose(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_3_1
foldableCompose1_4_2 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap4_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "foldMap"), dictMonoid_4)
_ = foldMap4_5_3
foldMap5_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "foldMap"), dictMonoid_4)
_ = foldMap5_6_4
return gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap4_5_3, gopurs_runtime.Apply(foldMap5_6_4, f_7), v_8)
})
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, i_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "foldl"), f_4), i_5, v_6)
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, i_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "foldr"), f_4)
_ = __local_var_7_5
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldr"), gopurs_runtime.Func2(func(b_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_7_5, a_9, b_8)
}), i_5, v_6)
}))
_ = foldableCompose1_4_2
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCompose1_4_2
}), gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex3_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), dictMonoid_5)
_ = foldMapWithIndex3_6_6
foldMapWithIndex4_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), dictMonoid_5)
_ = foldMapWithIndex4_7_7
return gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMapWithIndex3_6_6, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMapWithIndex4_7_7, gopurs_runtime.Func(func(b_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_8, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{x_10, b_11})})
}))
}), v_9)
})
}), gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, i_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{x_8, b_9})})
}))
}), i_6, v_7)
}), gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, i_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{a_8, b_9})})
}))
_ = __local_var_9_8
return gopurs_runtime.Func2(func(b_10 gopurs_runtime.Value, a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_9_8, a_11, b_10)
})
}), i_6, v_7)
}))
})
}

func Call_foldableWithIndexCoproduct(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
foldableCoproduct_1_0 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableCoproduct_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldableCoproduct1_3_1 := gopurs_runtime.Apply(foldableCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableCoproduct1_3_1
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCoproduct1_3_1
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex3_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), dictMonoid_4)
_ = foldMapWithIndex3_5_2
foldMapWithIndex4_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), dictMonoid_4)
_ = foldMapWithIndex4_6_3
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_4 := gopurs_runtime.Apply(foldMapWithIndex3_5_2, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{x_8})})
}))
_ = __local_var_8_4
__local_var_9_5 := gopurs_runtime.Apply(foldMapWithIndex4_6_3, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{x_9})})
}))
_ = __local_var_9_5
return gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t6 = gopurs_runtime.Apply(__local_var_8_4, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t6 = gopurs_runtime.Apply(__local_var_9_5, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_10.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
})
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_7 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{x_6})})
}), z_5)
_ = __local_var_6_7
__local_var_7_8 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{x_7})})
}), z_5)
_ = __local_var_7_8
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t9 = gopurs_runtime.Apply(__local_var_6_7, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_9
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t9 = gopurs_runtime.Apply(__local_var_7_8, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{x_6})})
}), z_5)
_ = __local_var_6_10
__local_var_7_11 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{x_7})})
}), z_5)
_ = __local_var_7_11
return gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t12 = gopurs_runtime.Apply(__local_var_6_10, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_12
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t12 = gopurs_runtime.Apply(__local_var_7_11, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_8.UnsafePtr).V0)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
}))
})
}

func Call_foldableWithIndexProduct(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
foldableProduct_1_0 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableProduct_1_0
return gopurs_runtime.Func(func(dictFoldableWithIndex1_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldableProduct1_3_1 := gopurs_runtime.Apply(foldableProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableProduct1_3_1
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableProduct1_3_1
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex3_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), dictMonoid_4)
_ = foldMapWithIndex3_5_2
foldMapWithIndex4_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldMapWithIndex"), dictMonoid_4)
_ = foldMapWithIndex4_6_3
return gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(foldMapWithIndex3_5_2, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{x_9})})
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply2(foldMapWithIndex4_6_3, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{x_9})})
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{x_7})})
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{x_7})})
}), z_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, z_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{x_7})})
}), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex1_2, "foldrWithIndex"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{x_7})})
}), z_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)
}))
})
}

func Call_foldlWithIndexDefault(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
foldMapWithIndex1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), Get_monoidDual())
_ = foldMapWithIndex1_1_0
return gopurs_runtime.Func3(func(c_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(foldMapWithIndex1_1_0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_1 := gopurs_runtime.Apply(c_2, i_5)
_ = __local_var_6_1
return gopurs_runtime.Func2(func(x_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_6_1, a_8, x_7)
})
}), xs_4, u_3)
})
}

func Call_foldrWithIndexDefault(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
foldMapWithIndex1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), Get_monoidEndo())
_ = foldMapWithIndex1_1_0
return gopurs_runtime.Func3(func(c_2 gopurs_runtime.Value, u_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(foldMapWithIndex1_1_0, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_2, i_5)
}), xs_4, u_3)
})
}

func Call_surroundMapWithIndex(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
foldMapWithIndex1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), Get_monoidEndo())
_ = foldMapWithIndex1_1_0
return gopurs_runtime.Func4(func(dictSemigroup_2 gopurs_runtime.Value, d_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(foldMapWithIndex1_1_0, gopurs_runtime.Func3(func(i_6 gopurs_runtime.Value, a_7 gopurs_runtime.Value, m_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), d_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), gopurs_runtime.Apply2(t_4, i_6, a_7), m_8))
}), f_5, d_3)
})
}

func Call_foldMapDefault(dictFoldableWithIndex_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
foldMapWithIndex2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), dictMonoid_1)
_ = foldMapWithIndex2_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMapWithIndex2_2_0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return f_3
}))
})
}

func Call_findWithIndex(dictFoldableWithIndex_0_loop gopurs_runtime.Value, p_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var p_1 gopurs_runtime.Value = p_1_loop
_ = p_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil)) && ((gopurs_runtime.Apply2(p_1, v_2, v2_4).IntVal) != (0)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.RecordDict2("index", "value", v_2, v2_4)})}
goto end_branch_0
} else {

}
}
{
__t0 = v1_3
}
end_branch_0:
return __t0
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_findMapWithIndex(dictFoldableWithIndex_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply2(f_1, v_2, v2_4)
goto end_branch_0
} else {

}
}
{
__t0 = v1_3
}
end_branch_0:
return __t0
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}

func Call_anyWithIndex(dictFoldableWithIndex_0_loop gopurs_runtime.Value, dictHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictHeytingAlgebra_1 gopurs_runtime.Value = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupDisj1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "disj"), v_2, v1_3)
}))
_ = semigroupDisj1_2_1
foldMapWithIndex2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupDisj1_2_1
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "ff")))
_ = foldMapWithIndex2_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMapWithIndex2_2_0, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(t_3, i_4)
}))
})
}

func Call_allWithIndex(dictFoldableWithIndex_0_loop gopurs_runtime.Value, dictHeytingAlgebra_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
var dictHeytingAlgebra_1 gopurs_runtime.Value = dictHeytingAlgebra_1_loop
_ = dictHeytingAlgebra_1
semigroupConj1_2_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "conj"), v_2, v1_3)
}))
_ = semigroupConj1_2_1
foldMapWithIndex2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupConj1_2_1
}), gopurs_runtime.RecordGet(dictHeytingAlgebra_1, "tt")))
_ = foldMapWithIndex2_2_0
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMapWithIndex2_2_0, gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(t_3, i_4)
}))
})
}



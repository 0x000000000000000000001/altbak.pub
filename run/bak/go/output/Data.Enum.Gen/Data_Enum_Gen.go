package Data_Enum_Gen

import (
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Control_Monad_Gen_Class "gopurs/output/Control.Monad.Gen.Class"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Semigroup_Last "gopurs/output/Data.Semigroup.Last"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		cache_foldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_Foldable.Get_foldableArray())))}
	})
	return cache_foldable1NonEmpty
}

var cache_genBoundedEnum gopurs_runtime.Value
var once_genBoundedEnum sync.Once
func Get_genBoundedEnum() gopurs_runtime.Value {
	once_genBoundedEnum.Do(func() {
		cache_genBoundedEnum = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genBoundedEnum(dictMonadGen_0_box)
})
	})
	return cache_genBoundedEnum
}

var cache_guard__2168855335 gopurs_runtime.Value
var once_guard__2168855335 sync.Once
func Get_guard__2168855335() gopurs_runtime.Value {
	once_guard__2168855335.Do(func() {
		cache_guard__2168855335 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_guard__2168855335(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_guard__2168855335
}

var cache_guard__666381313 gopurs_runtime.Value
var once_guard__666381313 sync.Once
func Get_guard__666381313() gopurs_runtime.Value {
	once_guard__666381313.Do(func() {
		cache_guard__666381313 = func() gopurs_runtime.Value {
Applicative0_0_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_alternativeMaybe(), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_0_0
empty_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_alternativeMaybe(), "Plus1"), gopurs_runtime.Value{}), "empty")
_ = empty_1_1
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_2.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(Applicative0_0_0.V1, pkg_Data_Unit.Get_unit())
goto end_branch_2
} else {

}
}
{
__t2 = empty_1_1
}
end_branch_2:
return __t2
})
}()
	})
	return cache_guard__666381313
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

var cache_bind__1858449959 gopurs_runtime.Value
var once_bind__1858449959 sync.Once
func Get_bind__1858449959() gopurs_runtime.Value {
	once_bind__1858449959.Do(func() {
		cache_bind__1858449959 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__1858449959(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__1858449959
}

var cache_bind__2601835655 gopurs_runtime.Value
var once_bind__2601835655 sync.Once
func Get_bind__2601835655() gopurs_runtime.Value {
	once_bind__2601835655.Do(func() {
		cache_bind__2601835655 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bind__2601835655(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_bind__2601835655
}

var cache_bind__1906657537 gopurs_runtime.Value
var once_bind__1906657537 sync.Once
func Get_bind__1906657537() gopurs_runtime.Value {
	once_bind__1906657537.Do(func() {
		cache_bind__1906657537 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind")
	})
	return cache_bind__1906657537
}

var cache_chooseInt__1063828903 gopurs_runtime.Value
var once_chooseInt__1063828903 sync.Once
func Get_chooseInt__1063828903() gopurs_runtime.Value {
	once_chooseInt__1063828903.Do(func() {
		cache_chooseInt__1063828903 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseInt__1063828903(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_chooseInt__1063828903
}

var cache_elements__1644372582 gopurs_runtime.Value
var once_elements__1644372582 sync.Once
func Get_elements__1644372582() gopurs_runtime.Value {
	once_elements__1644372582.Do(func() {
		cache_elements__1644372582 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elements__1644372582(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_elements__1644372582
}

var cache_elements__2000025632 gopurs_runtime.Value
var once_elements__2000025632 sync.Once
func Get_elements__2000025632() gopurs_runtime.Value {
	once_elements__2000025632.Do(func() {
		cache_elements__2000025632 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_elements__2000025632(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_0_box))
})
	})
	return cache_elements__2000025632
}

var cache_fromIndex__3111933544 gopurs_runtime.Value
var once_fromIndex__3111933544 sync.Once
func Get_fromIndex__3111933544() gopurs_runtime.Value {
	once_fromIndex__3111933544.Do(func() {
		cache_fromIndex__3111933544 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromIndex__3111933544(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_fromIndex__3111933544
}

var cache_bottom__338427193 gopurs_runtime.Value
var once_bottom__338427193 sync.Once
func Get_bottom__338427193() gopurs_runtime.Value {
	once_bottom__338427193.Do(func() {
		cache_bottom__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottom__338427193(dict_0_box)
})
	})
	return cache_bottom__338427193
}

var cache_top__338427193 gopurs_runtime.Value
var once_top__338427193 sync.Once
func Get_top__338427193() gopurs_runtime.Value {
	once_top__338427193.Do(func() {
		cache_top__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_top__338427193(dict_0_box)
})
	})
	return cache_top__338427193
}

var cache_enumFromTo__1480115131 gopurs_runtime.Value
var once_enumFromTo__1480115131 sync.Once
func Get_enumFromTo__1480115131() gopurs_runtime.Value {
	once_enumFromTo__1480115131.Do(func() {
		cache_enumFromTo__1480115131 = gopurs_runtime.Func(func(dictEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_enumFromTo__1480115131(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]](dictEnum_0_box))
})
	})
	return cache_enumFromTo__1480115131
}

var cache_succ__3199041328 gopurs_runtime.Value
var once_succ__3199041328 sync.Once
func Get_succ__3199041328() gopurs_runtime.Value {
	once_succ__3199041328.Do(func() {
		cache_succ__3199041328 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_succ__3199041328(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_succ__3199041328
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
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

var cache_foldr__3675782427 gopurs_runtime.Value
var once_foldr__3675782427 sync.Once
func Get_foldr__3675782427() gopurs_runtime.Value {
	once_foldr__3675782427.Do(func() {
		cache_foldr__3675782427 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3675782427(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__3675782427
}

var cache_length__854370588 gopurs_runtime.Value
var once_length__854370588 sync.Once
func Get_length__854370588() gopurs_runtime.Value {
	once_length__854370588.Do(func() {
		cache_length__854370588 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length__854370588(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_length__854370588
}

var cache_length__949294460 gopurs_runtime.Value
var once_length__949294460 sync.Once
func Get_length__949294460() gopurs_runtime.Value {
	once_length__949294460.Do(func() {
		cache_length__949294460 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictSemiring_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_length__949294460(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dictSemiring_1_box))
})
	})
	return cache_length__949294460
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
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

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__901270812
}

var cache_voidLeft__32301756 gopurs_runtime.Value
var once_voidLeft__32301756 sync.Once
func Get_voidLeft__32301756() gopurs_runtime.Value {
	once_voidLeft__32301756.Do(func() {
		cache_voidLeft__32301756 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidLeft__32301756(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), f_1_box, x_2_box)
})
	})
	return cache_voidLeft__32301756
}

var cache_voidLeft__3507060164 gopurs_runtime.Value
var once_voidLeft__3507060164 sync.Once
func Get_voidLeft__3507060164() gopurs_runtime.Value {
	once_voidLeft__3507060164.Do(func() {
		cache_voidLeft__3507060164 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidLeft__3507060164(f_0_box, x_1_box)
})
	})
	return cache_voidLeft__3507060164
}

var cache_altMaybe__4201091523 gopurs_runtime.Value
var once_altMaybe__4201091523 sync.Once
func Get_altMaybe__4201091523() gopurs_runtime.Value {
	once_altMaybe__4201091523.Do(func() {
		cache_altMaybe__4201091523 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
})
}))
	})
	return cache_altMaybe__4201091523
}

var cache_alternativeMaybe__196315533 gopurs_runtime.Value
var once_alternativeMaybe__196315533 sync.Once
func Get_alternativeMaybe__196315533() gopurs_runtime.Value {
	once_alternativeMaybe__196315533.Do(func() {
		cache_alternativeMaybe__196315533 = gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applicativeMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_plusMaybe()
}))
	})
	return cache_alternativeMaybe__196315533
}

var cache_applicativeMaybe__500933224 gopurs_runtime.Value
var once_applicativeMaybe__500933224 sync.Once
func Get_applicativeMaybe__500933224() gopurs_runtime.Value {
	once_applicativeMaybe__500933224.Do(func() {
		cache_applicativeMaybe__500933224 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), pkg_Data_Maybe.Get_Just())
	})
	return cache_applicativeMaybe__500933224
}

var cache_applyMaybe__3698865467 gopurs_runtime.Value
var once_applyMaybe__3698865467 sync.Once
func Get_applyMaybe__3698865467() gopurs_runtime.Value {
	once_applyMaybe__3698865467.Do(func() {
		cache_applyMaybe__3698865467 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, v1_1)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__3698865467
}

var cache_bindMaybe__1910292045 gopurs_runtime.Value
var once_bindMaybe__1910292045 sync.Once
func Get_bindMaybe__1910292045() gopurs_runtime.Value {
	once_bindMaybe__1910292045.Do(func() {
		cache_bindMaybe__1910292045 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindMaybe__1910292045
}

var cache_fromJust__1791383420 gopurs_runtime.Value
var once_fromJust__1791383420 sync.Once
func Get_fromJust__1791383420() gopurs_runtime.Value {
	once_fromJust__1791383420.Do(func() {
		cache_fromJust__1791383420 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1791383420(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust__1791383420
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_functorMaybe__2097654001 gopurs_runtime.Value
var once_functorMaybe__2097654001 sync.Once
func Get_functorMaybe__2097654001() gopurs_runtime.Value {
	once_functorMaybe__2097654001.Do(func() {
		cache_functorMaybe__2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2097654001
}

var cache_isNothing__2591355336 gopurs_runtime.Value
var once_isNothing__2591355336 sync.Once
func Get_isNothing__2591355336() gopurs_runtime.Value {
	once_isNothing__2591355336.Do(func() {
		cache_isNothing__2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__2591355336(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__2591355336
}

var cache_maybe__1594528518 gopurs_runtime.Value
var once_maybe__1594528518 sync.Once
func Get_maybe__1594528518() gopurs_runtime.Value {
	once_maybe__1594528518.Do(func() {
		cache_maybe__1594528518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1594528518(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1594528518
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_plusMaybe__400696082 gopurs_runtime.Value
var once_plusMaybe__400696082 sync.Once
func Get_plusMaybe__400696082() gopurs_runtime.Value {
	once_plusMaybe__400696082.Do(func() {
		cache_plusMaybe__400696082 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_altMaybe()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_plusMaybe__400696082
}

var cache_un__3470773042 gopurs_runtime.Value
var once_un__3470773042 sync.Once
func Get_un__3470773042() gopurs_runtime.Value {
	once_un__3470773042.Do(func() {
		cache_un__3470773042 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_un__3470773042(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box), v_1_box)
})
	})
	return cache_un__3470773042
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
}

var cache_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_lessThanOrEq__4087042607 sync.Once
func Get_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_lessThanOrEq__4087042607.Do(func() {
		cache_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__4087042607
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_foldMap1__3342855683 gopurs_runtime.Value
var once_foldMap1__3342855683 sync.Once
func Get_foldMap1__3342855683() gopurs_runtime.Value {
	once_foldMap1__3342855683.Do(func() {
		cache_foldMap1__3342855683 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__3342855683(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__3342855683
}

var cache_semigroupLast__2108226578 gopurs_runtime.Value
var once_semigroupLast__2108226578 sync.Once
func Get_semigroupLast__2108226578() gopurs_runtime.Value {
	once_semigroupLast__2108226578.Do(func() {
		cache_semigroupLast__2108226578 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
})
}))
	})
	return cache_semigroupLast__2108226578
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_one__1204848985 gopurs_runtime.Value
var once_one__1204848985 sync.Once
func Get_one__1204848985() gopurs_runtime.Value {
	once_one__1204848985.Do(func() {
		cache_one__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_one__1204848985(dict_0_box)
})
	})
	return cache_one__1204848985
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
}

var cache_fst__20422131 gopurs_runtime.Value
var once_fst__20422131 sync.Once
func Get_fst__20422131() gopurs_runtime.Value {
	once_fst__20422131.Do(func() {
		cache_fst__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fst__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fst__20422131
}

var cache_snd__20422131 gopurs_runtime.Value
var once_snd__20422131 sync.Once
func Get_snd__20422131() gopurs_runtime.Value {
	once_snd__20422131.Do(func() {
		cache_snd__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__20422131
}

var cache_replicate1__3169098027 gopurs_runtime.Value
var once_replicate1__3169098027 sync.Once
func Get_replicate1__3169098027() gopurs_runtime.Value {
	once_replicate1__3169098027.Do(func() {
		cache_replicate1__3169098027 = gopurs_runtime.Func3(func(dictUnfoldable1_0_box gopurs_runtime.Value, n_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replicate1__3169098027(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box), n_1_box.IntVal, v_2_box)
})
	})
	return cache_replicate1__3169098027
}

var cache_singleton__1620623815 gopurs_runtime.Value
var once_singleton__1620623815 sync.Once
func Get_singleton__1620623815() gopurs_runtime.Value {
	once_singleton__1620623815.Do(func() {
		cache_singleton__1620623815 = gopurs_runtime.Func(func(dictUnfoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton__1620623815(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](dictUnfoldable1_0_box))
})
	})
	return cache_singleton__1620623815
}

var cache_unfoldable1Array__2415700810 gopurs_runtime.Value
var once_unfoldable1Array__2415700810 sync.Once
func Get_unfoldable1Array__2415700810() gopurs_runtime.Value {
	once_unfoldable1Array__2415700810.Do(func() {
		cache_unfoldable1Array__2415700810 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(pkg_Data_Unfoldable1.Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldable1Array__2415700810
}

var cache_unfoldr1__169691813 gopurs_runtime.Value
var once_unfoldr1__169691813 sync.Once
func Get_unfoldr1__169691813() gopurs_runtime.Value {
	once_unfoldr1__169691813.Do(func() {
		cache_unfoldr1__169691813 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr1__169691813(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr1__169691813
}

var cache_unfoldr1__2402610528 gopurs_runtime.Value
var once_unfoldr1__2402610528 sync.Once
func Get_unfoldr1__2402610528() gopurs_runtime.Value {
	once_unfoldr1__2402610528.Do(func() {
		cache_unfoldr1__2402610528 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr1__2402610528(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr1__2402610528
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__1090765981 gopurs_runtime.Value
var once_unsafePartial__1090765981 sync.Once
func Get_unsafePartial__1090765981() gopurs_runtime.Value {
	once_unsafePartial__1090765981.Do(func() {
		cache_unsafePartial__1090765981 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1090765981
}

func Call_genBoundedEnum(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
Enum1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_3_1
Bounded0_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Bounded.Constructor_Bounded[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_4_2
v_5_3 := gopurs_runtime.Apply(Enum1_3_1.V2, Bounded0_4_2.V1)
_ = v_5_3
var __t4 gopurs_runtime.Value
{
if (v_5_3.Type == 9 && v_5_3.IntVal == 930809136 && v_5_3.UnsafePtr != nil) {
__t4 = gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_elements(), dictMonadGen_0, gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldable1NonEmpty()))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, Bounded0_4_2.V1, gopurs_runtime.Apply4(pkg_Data_Enum.Get_enumFromTo(), gopurs_runtime.Value{Type: 9, IntVal: 4075786298, UnsafePtr: unsafe.Pointer(Enum1_3_1)}, pkg_Data_Unfoldable1.Get_unfoldable1Array(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_3.UnsafePtr).V0, Bounded0_4_2.V2)})})
goto end_branch_4
} else {

}
}
{
if (v_5_3.Type == 9 && v_5_3.IntVal == 930809136 && v_5_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply(Applicative0_1_0.V1, Bounded0_4_2.V1)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})
}

func Call_guard__2168855335(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_1_0
empty_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "empty")
_ = empty_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v_3.IntVal) != (0) {
__t2 = gopurs_runtime.Apply(Applicative0_1_0.V1, pkg_Data_Unit.Get_unit())
goto end_branch_2
} else {

}
}
{
__t2 = empty_2_1
}
end_branch_2:
return __t2
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

func Call_bind__1858449959(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_bind__2601835655(dict_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_chooseInt__1063828903(dict_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_elements__1644372582(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(dictFoldable1_4 gopurs_runtime.Value) gopurs_runtime.Value {
Foldable0_5_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_4, "Foldable0"), gopurs_runtime.Value{}))
_ = Foldable0_5_3
return gopurs_runtime.Func(func(xs_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(dictMonadGen_0.V3, gopurs_runtime.Int(0), gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply3(Foldable0_5_3.V1, gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_7.IntVal))
})
}), gopurs_runtime.Int(0), xs_6), gopurs_runtime.Int(1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_fromIndex(), dictFoldable1_4, n_7, xs_6))
}))
})
})
}

func Call_elements__2000025632(dictMonadGen_0_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadGen_0 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.V0, gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
Foldable0_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldable1NonEmpty()).V0, gopurs_runtime.Value{}))
_ = Foldable0_4_3
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(dictMonadGen_0.V3, gopurs_runtime.Int(0), gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply3(Foldable0_4_3.V1, gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_6.IntVal))
})
}), gopurs_runtime.Int(0), xs_5), gopurs_runtime.Int(1))), gopurs_runtime.Func(func(n_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_fromIndex(), gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldable1NonEmpty()))}, n_6, xs_5))
}))
})
}

func Call_fromIndex__3111933544(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_1_0 gopurs_runtime.Value
go__go_4_1_0 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop gopurs_runtime.Value = v_5_loop_val
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_1_0:
for {
if false { continue go__go_4_1_0 }
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t4 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 759514854 && __t_tag_2.UnsafePtr == nil) {
__t3 = (*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(v_5, gopurs_runtime.Int(0))).IntVal) != (0) {
__t3 = (*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
v_5_loop = gopurs_runtime.Apply2(Get_sub__1043827704(), v_5, gopurs_runtime.Int(1))
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(v1_6.UnsafePtr).V1)}
continue go__go_4_1_0
__t3 = gopurs_runtime.Value{}
}
end_branch_3:
__t4 = __t3
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t4 = gopurs_runtime.Apply3(dictFoldable1_0.V1, pkg_Data_Semigroup_Last.Get_semigroupLast(), pkg_Data_Semigroup_Last.Get_Last(), xs_3)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_1_0, i_2, gopurs_runtime.Apply3(Foldable0_1_0.V2, pkg_Control_Monad_Gen.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*pkg_Control_Monad_Gen.Constructor_Cons[gopurs_runtime.Value])(nil))}, xs_3))
})
})
}

func Call_bottom__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_top__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_enumFromTo__1480115131(dictEnum_0_loop *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEnum_0 *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value] = dictEnum_0_loop
_ = dictEnum_0
Ord0_1_0 := gopurs_runtime.Apply(dictEnum_0.V0, gopurs_runtime.Value{})
_ = Ord0_1_0
Eq0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_1_0, "Eq0"), gopurs_runtime.Value{}))
_ = Eq0_2_1
Ord01_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(dictEnum_0.V0, gopurs_runtime.Value{}))
_ = Ord01_3_2
return gopurs_runtime.Func(func(dictUnfoldable1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Eq0_2_1.V0, v_5, v1_6).IntVal) != (0) {
__t9 = gopurs_runtime.Apply3(pkg_Data_Unfoldable1.Get_replicate1(), dictUnfoldable1_4, gopurs_runtime.Int(1), v_5)
goto end_branch_9
} else {

}
}
{
var __t6 bool
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.Apply2(Ord01_3_2.V1, v_5, v1_6)
if (uint32(__t_tag_5.IntVal) == 1527465420) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
if __t6 {
__t9 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_7, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(dictEnum_0.V2, a_7), gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 bool
{
var __t_tag_7 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_8, v1_6)
if (uint32(__t_tag_7.IntVal) == 380165415) {
__t8 = false
goto end_branch_8
} else {

}
}
{
__t8 = true
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_8
}), gopurs_runtime.Apply(Get_guard__666381313(), gopurs_runtime.Bool(__t8)))))}
}))})}
}), v_5)
goto end_branch_9
} else {

}
}
{
__t9 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictUnfoldable1_4, "unfoldr1"), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_7, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(dictEnum_0.V1, a_7), gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 bool
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_1_0, "compare"), a_prime_8, v1_6)
if (uint32(__t_tag_3.IntVal) == 1527465420) {
__t4 = false
goto end_branch_4
} else {

}
}
{
__t4 = true
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return a_prime_8
}), gopurs_runtime.Apply(Get_guard__666381313(), gopurs_runtime.Bool(__t4)))))}
}))})}
}), v_5)
}
end_branch_9:
return __t9
})
})
})
}

func Call_succ__3199041328(dict_0_loop *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_Enum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
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

func Call_foldr__3675782427(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_length__854370588(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(c_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_1.IntVal))
})
}), gopurs_runtime.Int(0))
}

func Call_length__949294460(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictSemiring_1_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictSemiring_1 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dictSemiring_1_loop
_ = dictSemiring_1
return gopurs_runtime.Apply2(dictFoldable_0.V1, gopurs_runtime.Func(func(c_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemiring_1.V0, dictSemiring_1.V2, c_2)
})
}), dictSemiring_1.V3)
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_voidLeft__32301756(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(dictFunctor_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), f_1)
}

func Call_voidLeft__3507060164(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), f_0)
}

func Call_fromJust__1791383420(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
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

func Call_isNothing__2591355336(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__1594528518(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
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

func Call_un__3470773042(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lessThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_lessThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMap1__3342855683(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_one__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_fst__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_snd__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_replicate1__3169098027(dictUnfoldable1_0_loop *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value], n_1_loop int64, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable1_0 *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
var n_1 int64 = n_1_loop
_ = n_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(dictUnfoldable1_0.V0, gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]
{
var __t0 bool
{
if (i_3.IntVal) > (0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = &pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}}
goto end_branch_1
} else {

}
}
{
__t1 = &pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, v_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int((i_3.IntVal) - (1))})}}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Int((n_1) - (1)))
}

func Call_singleton__1620623815(dictUnfoldable1_0_loop *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable1_0 *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value] = dictUnfoldable1_0_loop
_ = dictUnfoldable1_0
return gopurs_runtime.Apply2(pkg_Data_Unfoldable1.Get_replicate1(), gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(dictUnfoldable1_0)}, gopurs_runtime.Int(1))
}

func Call_unfoldr1__169691813(dict_0_loop *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_unfoldr1__2402610528(dict_0_loop *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable1.Constructor_Unfoldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}



package Data_Traversable

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Functor_App "gopurs/output/Data.Functor.App"
	pkg_Data_Functor_Compose "gopurs/output/Data.Functor.Compose"
	pkg_Data_Functor_Coproduct "gopurs/output/Data.Functor.Coproduct"
	pkg_Data_Functor_Product "gopurs/output/Data.Functor.Product"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Maybe_First "gopurs/output/Data.Maybe.First"
	pkg_Data_Maybe_Last "gopurs/output/Data.Maybe.Last"
	pkg_Data_Monoid_Additive "gopurs/output/Data.Monoid.Additive"
	pkg_Data_Monoid_Conj "gopurs/output/Data.Monoid.Conj"
	pkg_Data_Monoid_Disj "gopurs/output/Data.Monoid.Disj"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Traversable_Accum_Internal "gopurs/output/Data.Traversable.Accum.Internal"
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

var cache_traverse gopurs_runtime.Value
var once_traverse sync.Once
func Get_traverse() gopurs_runtime.Value {
	once_traverse.Do(func() {
		cache_traverse = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse
}

var cache_traversableTuple gopurs_runtime.Value
var once_traversableTuple sync.Once
func Get_traversableTuple() gopurs_runtime.Value {
	once_traversableTuple.Do(func() {
		cache_traversableTuple = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_functorTuple()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
})
}))
	})
	return cache_traversableTuple
}

var cache_traversableMultiplicative gopurs_runtime.Value
var once_traversableMultiplicative sync.Once
func Get_traversableMultiplicative() gopurs_runtime.Value {
	once_traversableMultiplicative.Do(func() {
		cache_traversableMultiplicative = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Multiplicative.Get_functorMultiplicative()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableMultiplicative
}

var cache_traversableMaybe gopurs_runtime.Value
var once_traversableMaybe sync.Once
func Get_traversableMaybe() gopurs_runtime.Value {
	once_traversableMaybe.Do(func() {
		cache_traversableMaybe = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Maybe.Get_Just(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
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
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(Functor0_1_2.V0, pkg_Data_Maybe.Get_Just(), gopurs_runtime.Apply(v_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
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
}))
	})
	return cache_traversableMaybe
}

var cache_traversableIdentity gopurs_runtime.Value
var once_traversableIdentity sync.Once
func Get_traversableIdentity() gopurs_runtime.Value {
	once_traversableIdentity.Do(func() {
		cache_traversableIdentity = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Identity.Get_functorIdentity()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Identity.Get_Identity(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableIdentity
}

var cache_traversableEither gopurs_runtime.Value
var once_traversableEither sync.Once
func Get_traversableEither() gopurs_runtime.Value {
	once_traversableEither.Do(func() {
		cache_traversableEither = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Either.Get_functorEither()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0})})
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
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0})})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(Functor0_1_2.V0, pkg_Data_Either.Get_Right(), gopurs_runtime.Apply(v_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
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
}))
	})
	return cache_traversableEither
}

var cache_traversableDual gopurs_runtime.Value
var once_traversableDual sync.Once
func Get_traversableDual() gopurs_runtime.Value {
	once_traversableDual.Do(func() {
		cache_traversableDual = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Dual.Get_functorDual()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Dual.Get_Dual(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Monoid_Dual.Get_Dual(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableDual
}

var cache_traversableDisj gopurs_runtime.Value
var once_traversableDisj sync.Once
func Get_traversableDisj() gopurs_runtime.Value {
	once_traversableDisj.Do(func() {
		cache_traversableDisj = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Disj.Get_functorDisj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Disj.Get_Disj(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Monoid_Disj.Get_Disj(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableDisj
}

var cache_traversableConst gopurs_runtime.Value
var once_traversableConst sync.Once
func Get_traversableConst() gopurs_runtime.Value {
	once_traversableConst.Do(func() {
		cache_traversableConst = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Const.Get_functorConst()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v_1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v1_2)
})
})
}))
	})
	return cache_traversableConst
}

var cache_traversableConj gopurs_runtime.Value
var once_traversableConj sync.Once
func Get_traversableConj() gopurs_runtime.Value {
	once_traversableConj.Do(func() {
		cache_traversableConj = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Conj.Get_functorConj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Conj.Get_Conj(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Monoid_Conj.Get_Conj(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableConj
}

var cache_traversableCompose gopurs_runtime.Value
var once_traversableCompose sync.Once
func Get_traversableCompose() gopurs_runtime.Value {
	once_traversableCompose.Do(func() {
		cache_traversableCompose = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableCompose(dictTraversable_0_box)
})
	})
	return cache_traversableCompose
}

var cache_traversableAdditive gopurs_runtime.Value
var once_traversableAdditive sync.Once
func Get_traversableAdditive() gopurs_runtime.Value {
	once_traversableAdditive.Do(func() {
		cache_traversableAdditive = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Monoid_Additive.Get_functorAdditive()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Monoid_Additive.Get_Additive(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Monoid_Additive.Get_Additive(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_traversableAdditive
}

var cache_sequenceDefault gopurs_runtime.Value
var once_sequenceDefault sync.Once
func Get_sequenceDefault() gopurs_runtime.Value {
	once_sequenceDefault.Do(func() {
		cache_sequenceDefault = gopurs_runtime.Func2(func(dictTraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_sequenceDefault
}

var cache_traversableArray gopurs_runtime.Value
var once_traversableArray sync.Once
func Get_traversableArray() gopurs_runtime.Value {
	once_traversableArray.Do(func() {
		cache_traversableArray = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableArray(), "traverse"), dictApplicative_0, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply4(Get_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"))
}))
	})
	return cache_traversableArray
}

var cache_sequence gopurs_runtime.Value
var once_sequence sync.Once
func Get_sequence() gopurs_runtime.Value {
	once_sequence.Do(func() {
		cache_sequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequence
}

var cache_traversableApp gopurs_runtime.Value
var once_traversableApp sync.Once
func Get_traversableApp() gopurs_runtime.Value {
	once_traversableApp.Do(func() {
		cache_traversableApp = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableApp(dictTraversable_0_box)
})
	})
	return cache_traversableApp
}

var cache_traversableCoproduct gopurs_runtime.Value
var once_traversableCoproduct sync.Once
func Get_traversableCoproduct() gopurs_runtime.Value {
	once_traversableCoproduct.Do(func() {
		cache_traversableCoproduct = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableCoproduct(dictTraversable_0_box)
})
	})
	return cache_traversableCoproduct
}

var cache_traversableFirst gopurs_runtime.Value
var once_traversableFirst sync.Once
func Get_traversableFirst() gopurs_runtime.Value {
	once_traversableFirst.Do(func() {
		cache_traversableFirst = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableFirst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Maybe_First.Get_First(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMaybe(), "sequence"), dictApplicative_0, v_2))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Maybe_First.Get_First(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_traversableMaybe(), "traverse"), dictApplicative_0, f_2, v_3))
})
})
}))
	})
	return cache_traversableFirst
}

var cache_traversableLast gopurs_runtime.Value
var once_traversableLast sync.Once
func Get_traversableLast() gopurs_runtime.Value {
	once_traversableLast.Do(func() {
		cache_traversableLast = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableLast()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Maybe_Last.Get_Last(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableMaybe(), "sequence"), dictApplicative_0, v_2))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, pkg_Data_Maybe_Last.Get_Last(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_traversableMaybe(), "traverse"), dictApplicative_0, f_2, v_3))
})
})
}))
	})
	return cache_traversableLast
}

var cache_traversableProduct gopurs_runtime.Value
var once_traversableProduct sync.Once
func Get_traversableProduct() gopurs_runtime.Value {
	once_traversableProduct.Do(func() {
		cache_traversableProduct = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableProduct(dictTraversable_0_box)
})
	})
	return cache_traversableProduct
}

var cache_traverseDefault gopurs_runtime.Value
var once_traverseDefault sync.Once
func Get_traverseDefault() gopurs_runtime.Value {
	once_traverseDefault.Do(func() {
		cache_traverseDefault = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseDefault(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_0_box))
})
	})
	return cache_traverseDefault
}

var cache_mapAccumR gopurs_runtime.Value
var once_mapAccumR sync.Once
func Get_mapAccumR() gopurs_runtime.Value {
	once_mapAccumR.Do(func() {
		cache_mapAccumR = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumR(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumR
}

var cache_scanr gopurs_runtime.Value
var once_scanr sync.Once
func Get_scanr() gopurs_runtime.Value {
	once_scanr.Do(func() {
		cache_scanr = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanr(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_scanr
}

var cache_mapAccumL gopurs_runtime.Value
var once_mapAccumL sync.Once
func Get_mapAccumL() gopurs_runtime.Value {
	once_mapAccumL.Do(func() {
		cache_mapAccumL = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumL(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumL
}

var cache_scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		cache_scanl = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanl(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_scanl
}

var cache_go__for gopurs_runtime.Value
var once_go__for sync.Once
func Get_go__for() gopurs_runtime.Value {
	once_go__for.Do(func() {
		cache_go__for = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, dictTraversable_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_go__for(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_1_box), x_2_box, f_3_box)
})
	})
	return cache_go__for
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

var cache_pure__1475749520 gopurs_runtime.Value
var once_pure__1475749520 sync.Once
func Get_pure__1475749520() gopurs_runtime.Value {
	once_pure__1475749520.Do(func() {
		cache_pure__1475749520 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1475749520(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__1475749520
}

var cache_pure__154576880 gopurs_runtime.Value
var once_pure__154576880 sync.Once
func Get_pure__154576880() gopurs_runtime.Value {
	once_pure__154576880.Do(func() {
		cache_pure__154576880 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__154576880(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__154576880
}

var cache_pure__1304937360 gopurs_runtime.Value
var once_pure__1304937360 sync.Once
func Get_pure__1304937360() gopurs_runtime.Value {
	once_pure__1304937360.Do(func() {
		cache_pure__1304937360 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__1304937360(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__1304937360
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

var cache_lift2__2762258480 gopurs_runtime.Value
var once_lift2__2762258480 sync.Once
func Get_lift2__2762258480() gopurs_runtime.Value {
	once_lift2__2762258480.Do(func() {
		cache_lift2__2762258480 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__2762258480(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__2762258480
}

var cache_lift2__470376976 gopurs_runtime.Value
var once_lift2__470376976 sync.Once
func Get_lift2__470376976() gopurs_runtime.Value {
	once_lift2__470376976.Do(func() {
		cache_lift2__470376976 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift2__470376976(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_lift2__470376976
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

var cache_functorConst__1992455793 gopurs_runtime.Value
var once_functorConst__1992455793 sync.Once
func Get_functorConst__1992455793() gopurs_runtime.Value {
	once_functorConst__1992455793.Do(func() {
		cache_functorConst__1992455793 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return m_1
})
}))
	})
	return cache_functorConst__1992455793
}

var cache_functorEither__1771778897 gopurs_runtime.Value
var once_functorEither__1771778897 sync.Once
func Get_functorEither__1771778897() gopurs_runtime.Value {
	once_functorEither__1771778897.Do(func() {
		cache_functorEither__1771778897 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (m_1.Type == 9 && m_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0})}
goto end_branch_0
} else {

}
}
{
if (m_1.Type == 9 && m_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0)})}
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
}))
	})
	return cache_functorEither__1771778897
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

var cache_foldMap__4073832436 gopurs_runtime.Value
var once_foldMap__4073832436 sync.Once
func Get_foldMap__4073832436() gopurs_runtime.Value {
	once_foldMap__4073832436.Do(func() {
		cache_foldMap__4073832436 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldMap")
	})
	return cache_foldMap__4073832436
}

var cache_foldableAdditive__1841171440 gopurs_runtime.Value
var once_foldableAdditive__1841171440 sync.Once
func Get_foldableAdditive__1841171440() gopurs_runtime.Value {
	once_foldableAdditive__1841171440.Do(func() {
		cache_foldableAdditive__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_foldableAdditive__1841171440
}

var cache_foldableArray__3859409398 gopurs_runtime.Value
var once_foldableArray__3859409398 sync.Once
func Get_foldableArray__3859409398() gopurs_runtime.Value {
	once_foldableArray__3859409398.Do(func() {
		cache_foldableArray__3859409398 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), pkg_Data_Foldable.Get_foldlArray(), pkg_Data_Foldable.Get_foldrArray())
	})
	return cache_foldableArray__3859409398
}

var cache_foldableConj__1841171440 gopurs_runtime.Value
var once_foldableConj__1841171440 sync.Once
func Get_foldableConj__1841171440() gopurs_runtime.Value {
	once_foldableConj__1841171440.Do(func() {
		cache_foldableConj__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_foldableConj__1841171440
}

var cache_foldableConst__943899702 gopurs_runtime.Value
var once_foldableConst__943899702 sync.Once
func Get_foldableConst__943899702() gopurs_runtime.Value {
	once_foldableConst__943899702.Do(func() {
		cache_foldableConst__943899702 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return z_1
})
})
}))
	})
	return cache_foldableConst__943899702
}

var cache_foldableDisj__1841171440 gopurs_runtime.Value
var once_foldableDisj__1841171440 sync.Once
func Get_foldableDisj__1841171440() gopurs_runtime.Value {
	once_foldableDisj__1841171440.Do(func() {
		cache_foldableDisj__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_foldableDisj__1841171440
}

var cache_foldableDual__1841171440 gopurs_runtime.Value
var once_foldableDual__1841171440 sync.Once
func Get_foldableDual__1841171440() gopurs_runtime.Value {
	once_foldableDual__1841171440.Do(func() {
		cache_foldableDual__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_foldableDual__1841171440
}

var cache_foldableEither__1622911640 gopurs_runtime.Value
var once_foldableEither__1622911640 sync.Once
func Get_foldableEither__1622911640() gopurs_runtime.Value {
	once_foldableEither__1622911640.Do(func() {
		cache_foldableEither__1622911640 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t1 = gopurs_runtime.Apply(v_2, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t3 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
}))
	})
	return cache_foldableEither__1622911640
}

var cache_foldableFirst__2831137713 gopurs_runtime.Value
var once_foldableFirst__2831137713 sync.Once
func Get_foldableFirst__2831137713() gopurs_runtime.Value {
	once_foldableFirst__2831137713.Do(func() {
		cache_foldableFirst__2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldMap"), dictMonoid_0, f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldl"), f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldr"), f_0, z_1, v_2)
})
})
}))
	})
	return cache_foldableFirst__2831137713
}

var cache_foldableIdentity__1841171440 gopurs_runtime.Value
var once_foldableIdentity__1841171440 sync.Once
func Get_foldableIdentity__1841171440() gopurs_runtime.Value {
	once_foldableIdentity__1841171440.Do(func() {
		cache_foldableIdentity__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_foldableIdentity__1841171440
}

var cache_foldableLast__2831137713 gopurs_runtime.Value
var once_foldableLast__2831137713 sync.Once
func Get_foldableLast__2831137713() gopurs_runtime.Value {
	once_foldableLast__2831137713.Do(func() {
		cache_foldableLast__2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldMap"), dictMonoid_0, f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldl"), f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldr"), f_0, z_1, v_2)
})
})
}))
	})
	return cache_foldableLast__2831137713
}

var cache_foldableMaybe__3653484922 gopurs_runtime.Value
var once_foldableMaybe__3653484922 sync.Once
func Get_foldableMaybe__3653484922() gopurs_runtime.Value {
	once_foldableMaybe__3653484922.Do(func() {
		cache_foldableMaybe__3653484922 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t1 = mempty_1_0
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply(v_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t2 = v1_1
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t3 = v1_1
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
}))
	})
	return cache_foldableMaybe__3653484922
}

var cache_foldableMaybe__2831137713 gopurs_runtime.Value
var once_foldableMaybe__2831137713 sync.Once
func Get_foldableMaybe__2831137713() gopurs_runtime.Value {
	once_foldableMaybe__2831137713.Do(func() {
		cache_foldableMaybe__2831137713 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t1 = mempty_1_0
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply(v_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t2 = v1_1
goto end_branch_2
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply2(v_0, v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr == nil) {
__t3 = v1_1
goto end_branch_3
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 930809136 && v2_2.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2.UnsafePtr).V0, v1_1)
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
}))
	})
	return cache_foldableMaybe__2831137713
}

var cache_foldableMultiplicative__1841171440 gopurs_runtime.Value
var once_foldableMultiplicative__1841171440 sync.Once
func Get_foldableMultiplicative__1841171440() gopurs_runtime.Value {
	once_foldableMultiplicative__1841171440.Do(func() {
		cache_foldableMultiplicative__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_foldableMultiplicative__1841171440
}

var cache_foldableTuple__1455669080 gopurs_runtime.Value
var once_foldableTuple__1455669080 sync.Once
func Get_foldableTuple__1455669080() gopurs_runtime.Value {
	once_foldableTuple__1455669080.Do(func() {
		cache_foldableTuple__1455669080 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
})
})
}))
	})
	return cache_foldableTuple__1455669080
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

var cache_foldl__3016550397 gopurs_runtime.Value
var once_foldl__3016550397 sync.Once
func Get_foldl__3016550397() gopurs_runtime.Value {
	once_foldl__3016550397.Do(func() {
		cache_foldl__3016550397 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldl")
	})
	return cache_foldl__3016550397
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

var cache_foldr__3016550397 gopurs_runtime.Value
var once_foldr__3016550397 sync.Once
func Get_foldr__3016550397() gopurs_runtime.Value {
	once_foldr__3016550397.Do(func() {
		cache_foldr__3016550397 = gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableMaybe(), "foldr")
	})
	return cache_foldr__3016550397
}

var cache_coproduct__413515331 gopurs_runtime.Value
var once_coproduct__413515331 sync.Once
func Get_coproduct__413515331() gopurs_runtime.Value {
	once_coproduct__413515331.Do(func() {
		cache_coproduct__413515331 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct__413515331(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_coproduct__413515331
}

var cache_coproduct__1706612365 gopurs_runtime.Value
var once_coproduct__1706612365 sync.Once
func Get_coproduct__1706612365() gopurs_runtime.Value {
	once_coproduct__1706612365.Do(func() {
		cache_coproduct__1706612365 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct__1706612365(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_coproduct__1706612365
}

var cache_coproduct__1642299426 gopurs_runtime.Value
var once_coproduct__1642299426 sync.Once
func Get_coproduct__1642299426() gopurs_runtime.Value {
	once_coproduct__1642299426.Do(func() {
		cache_coproduct__1642299426 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coproduct__1642299426(v_0_box, v1_1_box, v2_2_box)
})
	})
	return cache_coproduct__1642299426
}

var cache_product__2764631669 gopurs_runtime.Value
var once_product__2764631669 sync.Once
func Get_product__2764631669() gopurs_runtime.Value {
	once_product__2764631669.Do(func() {
		cache_product__2764631669 = gopurs_runtime.Func2(func(fa_0_box gopurs_runtime.Value, ga_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_product__2764631669(fa_0_box, ga_1_box))}
})
	})
	return cache_product__2764631669
}

var cache_functorArray__361387505 gopurs_runtime.Value
var once_functorArray__361387505 sync.Once
func Get_functorArray__361387505() gopurs_runtime.Value {
	once_functorArray__361387505.Do(func() {
		cache_functorArray__361387505 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__361387505
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

var cache_map__2701008148 gopurs_runtime.Value
var once_map__2701008148 sync.Once
func Get_map__2701008148() gopurs_runtime.Value {
	once_map__2701008148.Do(func() {
		cache_map__2701008148 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2701008148(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2701008148
}

var cache_map__1762802164 gopurs_runtime.Value
var once_map__1762802164 sync.Once
func Get_map__1762802164() gopurs_runtime.Value {
	once_map__1762802164.Do(func() {
		cache_map__1762802164 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1762802164(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1762802164
}

var cache_map__2562444020 gopurs_runtime.Value
var once_map__2562444020 sync.Once
func Get_map__2562444020() gopurs_runtime.Value {
	once_map__2562444020.Do(func() {
		cache_map__2562444020 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2562444020(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2562444020
}

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_map__1162593300 gopurs_runtime.Value
var once_map__1162593300 sync.Once
func Get_map__1162593300() gopurs_runtime.Value {
	once_map__1162593300.Do(func() {
		cache_map__1162593300 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1162593300(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1162593300
}

var cache_map__3659954292 gopurs_runtime.Value
var once_map__3659954292 sync.Once
func Get_map__3659954292() gopurs_runtime.Value {
	once_map__3659954292.Do(func() {
		cache_map__3659954292 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3659954292(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3659954292
}

var cache_map__2251722612 gopurs_runtime.Value
var once_map__2251722612 sync.Once
func Get_map__2251722612() gopurs_runtime.Value {
	once_map__2251722612.Do(func() {
		cache_map__2251722612 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2251722612(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2251722612
}

var cache_map__1337616244 gopurs_runtime.Value
var once_map__1337616244 sync.Once
func Get_map__1337616244() gopurs_runtime.Value {
	once_map__1337616244.Do(func() {
		cache_map__1337616244 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1337616244(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1337616244
}

var cache_functorIdentity__943655089 gopurs_runtime.Value
var once_functorIdentity__943655089 sync.Once
func Get_functorIdentity__943655089() gopurs_runtime.Value {
	once_functorIdentity__943655089.Do(func() {
		cache_functorIdentity__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorIdentity__943655089
}

var cache_functorFirst__2097654001 gopurs_runtime.Value
var once_functorFirst__2097654001 sync.Once
func Get_functorFirst__2097654001() gopurs_runtime.Value {
	once_functorFirst__2097654001.Do(func() {
		cache_functorFirst__2097654001 = pkg_Data_Maybe.Get_functorMaybe()
	})
	return cache_functorFirst__2097654001
}

var cache_functorLast__2097654001 gopurs_runtime.Value
var once_functorLast__2097654001 sync.Once
func Get_functorLast__2097654001() gopurs_runtime.Value {
	once_functorLast__2097654001.Do(func() {
		cache_functorLast__2097654001 = pkg_Data_Maybe.Get_functorMaybe()
	})
	return cache_functorLast__2097654001
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

var cache_functorAdditive__943655089 gopurs_runtime.Value
var once_functorAdditive__943655089 sync.Once
func Get_functorAdditive__943655089() gopurs_runtime.Value {
	once_functorAdditive__943655089.Do(func() {
		cache_functorAdditive__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorAdditive__943655089
}

var cache_functorConj__943655089 gopurs_runtime.Value
var once_functorConj__943655089 sync.Once
func Get_functorConj__943655089() gopurs_runtime.Value {
	once_functorConj__943655089.Do(func() {
		cache_functorConj__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorConj__943655089
}

var cache_functorDisj__943655089 gopurs_runtime.Value
var once_functorDisj__943655089 sync.Once
func Get_functorDisj__943655089() gopurs_runtime.Value {
	once_functorDisj__943655089.Do(func() {
		cache_functorDisj__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorDisj__943655089
}

var cache_functorDual__943655089 gopurs_runtime.Value
var once_functorDual__943655089 sync.Once
func Get_functorDual__943655089() gopurs_runtime.Value {
	once_functorDual__943655089.Do(func() {
		cache_functorDual__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorDual__943655089
}

var cache_functorMultiplicative__943655089 gopurs_runtime.Value
var once_functorMultiplicative__943655089 sync.Once
func Get_functorMultiplicative__943655089() gopurs_runtime.Value {
	once_functorMultiplicative__943655089.Do(func() {
		cache_functorMultiplicative__943655089 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, m_1)
})
}))
	})
	return cache_functorMultiplicative__943655089
}

var cache_applicativeStateL__2039640491 gopurs_runtime.Value
var once_applicativeStateL__2039640491 sync.Once
func Get_applicativeStateL__2039640491() gopurs_runtime.Value {
	once_applicativeStateL__2039640491.Do(func() {
		cache_applicativeStateL__2039640491 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable_Accum_Internal.Get_applyStateL()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
}))
	})
	return cache_applicativeStateL__2039640491
}

var cache_applicativeStateR__2039640491 gopurs_runtime.Value
var once_applicativeStateR__2039640491 sync.Once
func Get_applicativeStateR__2039640491() gopurs_runtime.Value {
	once_applicativeStateR__2039640491.Do(func() {
		cache_applicativeStateR__2039640491 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable_Accum_Internal.Get_applyStateR()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("accum", "value", s_1, a_0)
})
}))
	})
	return cache_applicativeStateR__2039640491
}

var cache_applyStateL__1243455060 gopurs_runtime.Value
var once_applyStateL__1243455060 sync.Once
func Get_applyStateL__1243455060() gopurs_runtime.Value {
	once_applyStateL__1243455060.Do(func() {
		cache_applyStateL__1243455060 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable_Accum_Internal.Get_functorStateL()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(f_0, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(x_1, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v_3_0, "value"), gopurs_runtime.RecordGet(v1_4_1, "value")))
})
})
}))
	})
	return cache_applyStateL__1243455060
}

var cache_applyStateR__1243455060 gopurs_runtime.Value
var once_applyStateR__1243455060 sync.Once
func Get_applyStateR__1243455060() gopurs_runtime.Value {
	once_applyStateR__1243455060.Do(func() {
		cache_applyStateR__1243455060 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable_Accum_Internal.Get_functorStateR()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(x_1, s_2)
_ = v_3_0
v1_4_1 := gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "accum"))
_ = v1_4_1
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v1_4_1, "accum"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(v1_4_1, "value"), gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_applyStateR__1243455060
}

var cache_functorStateL__830241200 gopurs_runtime.Value
var once_functorStateL__830241200 sync.Once
func Get_functorStateL__830241200() gopurs_runtime.Value {
	once_functorStateL__830241200.Do(func() {
		cache_functorStateL__830241200 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_functorStateL__830241200
}

var cache_functorStateR__830241200 gopurs_runtime.Value
var once_functorStateR__830241200 sync.Once
func Get_functorStateR__830241200() gopurs_runtime.Value {
	once_functorStateR__830241200.Do(func() {
		cache_functorStateR__830241200 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(k_1, s_2)
_ = v_3_0
return gopurs_runtime.RecordDict2("accum", "value", gopurs_runtime.RecordGet(v_3_0, "accum"), gopurs_runtime.Apply(f_0, gopurs_runtime.RecordGet(v_3_0, "value")))
})
})
}))
	})
	return cache_functorStateR__830241200
}

var cache_stateL__1334064830 gopurs_runtime.Value
var once_stateL__1334064830 sync.Once
func Get_stateL__1334064830() gopurs_runtime.Value {
	once_stateL__1334064830.Do(func() {
		cache_stateL__1334064830 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateL__1334064830(v_0_box)
})
	})
	return cache_stateL__1334064830
}

var cache_stateL__1412771550 gopurs_runtime.Value
var once_stateL__1412771550 sync.Once
func Get_stateL__1412771550() gopurs_runtime.Value {
	once_stateL__1412771550.Do(func() {
		cache_stateL__1412771550 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateL__1412771550(v_0_box)
})
	})
	return cache_stateL__1412771550
}

var cache_stateR__1334064830 gopurs_runtime.Value
var once_stateR__1334064830 sync.Once
func Get_stateR__1334064830() gopurs_runtime.Value {
	once_stateR__1334064830.Do(func() {
		cache_stateR__1334064830 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateR__1334064830(v_0_box)
})
	})
	return cache_stateR__1334064830
}

var cache_stateR__1412771550 gopurs_runtime.Value
var once_stateR__1412771550 sync.Once
func Get_stateR__1412771550() gopurs_runtime.Value {
	once_stateR__1412771550.Do(func() {
		cache_stateR__1412771550 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stateR__1412771550(v_0_box)
})
	})
	return cache_stateR__1412771550
}

var cache_mapAccumL__2528537048 gopurs_runtime.Value
var once_mapAccumL__2528537048 sync.Once
func Get_mapAccumL__2528537048() gopurs_runtime.Value {
	once_mapAccumL__2528537048.Do(func() {
		cache_mapAccumL__2528537048 = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumL__2528537048(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumL__2528537048
}

var cache_mapAccumL__1189480088 gopurs_runtime.Value
var once_mapAccumL__1189480088 sync.Once
func Get_mapAccumL__1189480088() gopurs_runtime.Value {
	once_mapAccumL__1189480088.Do(func() {
		cache_mapAccumL__1189480088 = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumL__1189480088(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumL__1189480088
}

var cache_mapAccumR__2528537048 gopurs_runtime.Value
var once_mapAccumR__2528537048 sync.Once
func Get_mapAccumR__2528537048() gopurs_runtime.Value {
	once_mapAccumR__2528537048.Do(func() {
		cache_mapAccumR__2528537048 = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumR__2528537048(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumR__2528537048
}

var cache_mapAccumR__1189480088 gopurs_runtime.Value
var once_mapAccumR__1189480088 sync.Once
func Get_mapAccumR__1189480088() gopurs_runtime.Value {
	once_mapAccumR__1189480088.Do(func() {
		cache_mapAccumR__1189480088 = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumR__1189480088(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_mapAccumR__1189480088
}

var cache_sequence__1886310617 gopurs_runtime.Value
var once_sequence__1886310617 sync.Once
func Get_sequence__1886310617() gopurs_runtime.Value {
	once_sequence__1886310617.Do(func() {
		cache_sequence__1886310617 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence__1886310617(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequence__1886310617
}

var cache_sequence__2322703601 gopurs_runtime.Value
var once_sequence__2322703601 sync.Once
func Get_sequence__2322703601() gopurs_runtime.Value {
	once_sequence__2322703601.Do(func() {
		cache_sequence__2322703601 = gopurs_runtime.RecordGet(Get_traversableMaybe(), "sequence")
	})
	return cache_sequence__2322703601
}

var cache_traversableMaybe__98548986 gopurs_runtime.Value
var once_traversableMaybe__98548986 sync.Once
func Get_traversableMaybe__98548986() gopurs_runtime.Value {
	once_traversableMaybe__98548986.Do(func() {
		cache_traversableMaybe__98548986 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply2(Functor0_1_0.V0, pkg_Data_Maybe.Get_Just(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
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
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(Functor0_1_2.V0, pkg_Data_Maybe.Get_Just(), gopurs_runtime.Apply(v_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
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
}))
	})
	return cache_traversableMaybe__98548986
}

var cache_traverse__314957093 gopurs_runtime.Value
var once_traverse__314957093 sync.Once
func Get_traverse__314957093() gopurs_runtime.Value {
	once_traverse__314957093.Do(func() {
		cache_traverse__314957093 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__314957093(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__314957093
}

var cache_traverse__3459678245 gopurs_runtime.Value
var once_traverse__3459678245 sync.Once
func Get_traverse__3459678245() gopurs_runtime.Value {
	once_traverse__3459678245.Do(func() {
		cache_traverse__3459678245 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__3459678245(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__3459678245
}

var cache_traverse__23995045 gopurs_runtime.Value
var once_traverse__23995045 sync.Once
func Get_traverse__23995045() gopurs_runtime.Value {
	once_traverse__23995045.Do(func() {
		cache_traverse__23995045 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__23995045(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__23995045
}

var cache_traverse__667327821 gopurs_runtime.Value
var once_traverse__667327821 sync.Once
func Get_traverse__667327821() gopurs_runtime.Value {
	once_traverse__667327821.Do(func() {
		cache_traverse__667327821 = gopurs_runtime.RecordGet(Get_traversableMaybe(), "traverse")
	})
	return cache_traverse__667327821
}

var cache_traverse__4214075557 gopurs_runtime.Value
var once_traverse__4214075557 sync.Once
func Get_traverse__4214075557() gopurs_runtime.Value {
	once_traverse__4214075557.Do(func() {
		cache_traverse__4214075557 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__4214075557(gopurs_runtime.CoerceToStruct[Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__4214075557
}

var cache_functorTuple__2249620049 gopurs_runtime.Value
var once_functorTuple__2249620049 sync.Once
func Get_functorTuple__2249620049() gopurs_runtime.Value {
	once_functorTuple__2249620049.Do(func() {
		cache_functorTuple__2249620049 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorTuple__2249620049
}

type Constructor_Traversable[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3941073978] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Traversable[gopurs_runtime.Value])(ptr)
		switch key {
		case "Foldable1": return c.V0
		case "Functor0": return c.V1
		case "sequence": return c.V2
		case "traverse": return c.V3
		default: panic("Key not found in dictionary Constructor_Traversable: " + key)
		}
	}
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_traverse(dict_0_loop *Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traversableCompose(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
traversableCompose:
for {
if false { continue traversableCompose }
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
functorCompose1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), f_5), v_6)
})
}))
_ = functorCompose1_4_2
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_5_5
foldableCompose1_5_4 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "foldMap"), dictMonoid_6, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "foldMap"), dictMonoid_6, f_7), v_8)
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "foldl"), f_6), i_7, v_8)
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "foldr"), f_6)
_ = __local_var_9_6
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "foldr"), gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_9_6, a_11, b_10)
})
}), i_7, v_8)
})
})
}))
_ = foldableCompose1_5_4
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCompose1_5_4
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCompose1_4_2
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Call_traversableCompose(dictTraversable_0), dictTraversable1_3), "traverse"), dictApplicative_6, Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_7_7 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_7
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_7_7.V0, pkg_Data_Functor_Compose.Get_Compose(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_6, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), dictApplicative_6, f_8), v_9))
})
})
}))
})
}
}

func Call_sequenceDefault(dictTraversable_0_loop *Constructor_Traversable[gopurs_runtime.Value], dictApplicative_1_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var dictApplicative_1 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(dictTraversable_0.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, Get_identity())
}

func Call_sequence(dict_0_loop *Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_traversableApp(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
functorApp_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = functorApp_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_2
foldableApp_2_1 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_2, "foldMap"), dictMonoid_3, f_4, v_5)
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_2, "foldl"), f_3, i_4, v_5)
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_2, "foldr"), f_3, i_4, v_5)
})
})
}))
_ = foldableApp_2_1
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableApp_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorApp_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_3.V0, pkg_Data_Functor_App.Get_App(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_3, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_4.V0, pkg_Data_Functor_App.Get_App(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_3, f_5, v_6))
})
})
}))
}

func Call_traversableCoproduct(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
functorCoproduct_1_0 := gopurs_runtime.Apply(pkg_Data_Functor_Coproduct.Get_functorCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct_1_0
foldableCoproduct_2_1 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableCoproduct_2_1
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
functorCoproduct1_4_2 := gopurs_runtime.Apply(functorCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct1_4_2
foldableCoproduct1_5_3 := gopurs_runtime.Apply(foldableCoproduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableCoproduct1_5_3
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCoproduct1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct1_4_2
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_7_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_4
__local_var_8_5 := gopurs_runtime.Apply(Functor0_7_4.V0, gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_8})}
}))
_ = __local_var_8_5
__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_6)
_ = __local_var_9_6
__local_var_10_7 := gopurs_runtime.Apply(Functor0_7_4.V0, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_10})}
}))
_ = __local_var_10_7
__local_var_11_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "sequence"), dictApplicative_6)
_ = __local_var_11_8
return gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t9 = gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.Apply(__local_var_9_6, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0))
goto end_branch_9
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t9 = gopurs_runtime.Apply(__local_var_10_7, gopurs_runtime.Apply(__local_var_11_8, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0))
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
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_7_10 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_10
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_11 := gopurs_runtime.Apply(Functor0_7_10.V0, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_9})}
}))
_ = __local_var_9_11
__local_var_10_12 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_6, f_8)
_ = __local_var_10_12
__local_var_11_13 := gopurs_runtime.Apply(Functor0_7_10.V0, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_11})}
}))
_ = __local_var_11_13
__local_var_12_14 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), dictApplicative_6, f_8)
_ = __local_var_12_14
return gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(__local_var_9_11, gopurs_runtime.Apply(__local_var_10_12, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0))
goto end_branch_15
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t15 = gopurs_runtime.Apply(__local_var_11_13, gopurs_runtime.Apply(__local_var_12_14, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_13.UnsafePtr).V0))
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
}))
})
}

func Call_traversableProduct(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
functorProduct_1_0 := gopurs_runtime.Apply(pkg_Data_Functor_Product.Get_functorProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorProduct_1_0
foldableProduct_2_1 := gopurs_runtime.Apply(pkg_Data_Foldable.Get_foldableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableProduct_2_1
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
functorProduct1_4_2 := gopurs_runtime.Apply(functorProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{}))
_ = functorProduct1_4_2
foldableProduct1_5_3 := gopurs_runtime.Apply(foldableProduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableProduct1_5_3
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableProduct1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct1_4_2
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_7_4 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_4
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_7_4.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Apply0_7_4.V0, gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product.Get_product(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_6, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_3, "sequence"), dictApplicative_6, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_7_5 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_5
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_7_5.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Apply0_7_5.V0, gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product.Get_product(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_6, f_8, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), dictApplicative_6, f_8, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1))
})
})
}))
})
}

func Call_traverseDefault(dictTraversable_0_loop *Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictTraversable_0.V1, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictTraversable_0.V2, dictApplicative_2, gopurs_runtime.Apply2(Functor0_1_0.V0, f_3, ta_4))
})
})
})
}

func Call_mapAccumR(dictTraversable_0_loop *Constructor_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversable_0.V3, pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR(), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_scanr(dictTraversable_0_loop *Constructor_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.RecordGet(Call_mapAccumR(dictTraversable_0, gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_6_0 := gopurs_runtime.Apply2(f_1, a_5, b_4)
_ = b_prime_6_0
return gopurs_runtime.RecordDict2("accum", "value", b_prime_6_0, b_prime_6_0)
})
}), b0_2, xs_3), "value")
}

func Call_mapAccumL(dictTraversable_0_loop *Constructor_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversable_0.V3, pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL(), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_scanl(dictTraversable_0_loop *Constructor_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.RecordGet(Call_mapAccumL(dictTraversable_0, gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_6_0 := gopurs_runtime.Apply2(f_1, b_4, a_5)
_ = b_prime_6_0
return gopurs_runtime.RecordDict2("accum", "value", b_prime_6_0, b_prime_6_0)
})
}), b0_2, xs_3), "value")
}

func Call_go__for(dictApplicative_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value], dictTraversable_1_loop *Constructor_Traversable[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var dictTraversable_1 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_1_loop
_ = dictTraversable_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
return gopurs_runtime.Apply3(dictTraversable_1.V3, gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, f_3, x_2)
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1475749520(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__154576880(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__1304937360(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_lift2__2762258480(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_lift2__470376976(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, f_2, a_3), b_4)
})
})
})
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_coproduct__413515331(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_coproduct__1706612365(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_coproduct__1642299426(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (v2_2.Type == 9 && v2_2.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(v_0, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
if (v2_2.Type == 9 && v2_2.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_2.UnsafePtr).V0)
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

func Call_product__2764631669(fa_0_loop gopurs_runtime.Value, ga_1_loop gopurs_runtime.Value) *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var fa_0 gopurs_runtime.Value = fa_0_loop
_ = fa_0
var ga_1 gopurs_runtime.Value = ga_1_loop
_ = ga_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, fa_0, ga_1})})
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

func Call_map__2701008148(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1762802164(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2562444020(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1162593300(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3659954292(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2251722612(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1337616244(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_stateL__1334064830(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_stateL__1412771550(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_stateR__1334064830(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_stateR__1412771550(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_mapAccumL__2528537048(dictTraversable_0_loop *Constructor_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversable_0.V3, pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL(), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_mapAccumL__1189480088(dictTraversable_0_loop *Constructor_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversable_0.V3, pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL(), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_mapAccumR__2528537048(dictTraversable_0_loop *Constructor_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversable_0.V3, pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR(), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_mapAccumR__1189480088(dictTraversable_0_loop *Constructor_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(dictTraversable_0.V3, pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR(), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_sequence__1886310617(dict_0_loop *Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_traverse__314957093(dict_0_loop *Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse__3459678245(dict_0_loop *Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse__23995045(dict_0_loop *Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverse__4214075557(dict_0_loop *Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Get_traverseArrayImpl() gopurs_runtime.Value {
	return _Gopurs_TraverseArrayImpl
}

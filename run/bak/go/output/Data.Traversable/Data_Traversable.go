package Data_Traversable

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Monoid_Multiplicative "gopurs/output/Data.Monoid.Multiplicative"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Monoid_Dual "gopurs/output/Data.Monoid.Dual"
	pkg_Data_Monoid_Disj "gopurs/output/Data.Monoid.Disj"
	pkg_Data_Const "gopurs/output/Data.Const"
	pkg_Data_Monoid_Conj "gopurs/output/Data.Monoid.Conj"
	pkg_Data_Monoid_Additive "gopurs/output/Data.Monoid.Additive"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Maybe_First "gopurs/output/Data.Maybe.First"
	pkg_Data_Maybe_Last "gopurs/output/Data.Maybe.Last"
	pkg_Data_Functor_Compose "gopurs/output/Data.Functor.Compose"
	pkg_Data_Functor_App "gopurs/output/Data.Functor.App"
	pkg_Data_Functor_Coproduct "gopurs/output/Data.Functor.Coproduct"
	pkg_Data_Functor_Product "gopurs/output/Data.Functor.Product"
	pkg_Data_Traversable_Accum_Internal "gopurs/output/Data.Traversable.Accum.Internal"
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
return Call_traverse(dict_0_box)
})
	})
	return cache_traverse
}

var cache_traverse__gopurs_runtime_Value_2265355576 gopurs_runtime.Value
var once_traverse__gopurs_runtime_Value_2265355576 sync.Once
func Get_traverse__gopurs_runtime_Value_2265355576() gopurs_runtime.Value {
	once_traverse__gopurs_runtime_Value_2265355576.Do(func() {
		cache_traverse__gopurs_runtime_Value_2265355576 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__gopurs_runtime_Value_2265355576(dict_0_box)
})
	})
	return cache_traverse__gopurs_runtime_Value_2265355576
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
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1))
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
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), v_1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Multiplicative.Get_Multiplicative(), gopurs_runtime.Apply(f_1, v_2))
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
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe.Get_Just(), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 930809136 && v1_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Apply(v_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_2.UnsafePtr).V0))
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
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Identity.Get_Identity(), v_1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Identity.Get_Identity(), gopurs_runtime.Apply(f_1, v_2))
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
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 3711209382) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0})})
goto end_branch_0
} else {

}
}
{
if (v_1.Type == 9 && v_1.IntVal == 2465973597) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Either.Get_Right(), (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0})})
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Either.Get_Right(), gopurs_runtime.Apply(v_1, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_2.UnsafePtr).V0))
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
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Dual.Get_Dual(), v_1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Dual.Get_Dual(), gopurs_runtime.Apply(f_1, v_2))
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
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Disj.Get_Disj(), v_1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Disj.Get_Disj(), gopurs_runtime.Apply(f_1, v_2))
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
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Conj.Get_Conj(), v_1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Conj.Get_Conj(), gopurs_runtime.Apply(f_1, v_2))
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
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Additive.Get_Additive(), v_1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Monoid_Additive.Get_Additive(), gopurs_runtime.Apply(f_1, v_2))
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
return Call_sequenceDefault(dictTraversable_0_box, dictApplicative_1_box)
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
return Call_sequence(dict_0_box)
})
	})
	return cache_sequence
}

var cache_sequence__gopurs_runtime_Value_1188743833 gopurs_runtime.Value
var once_sequence__gopurs_runtime_Value_1188743833 sync.Once
func Get_sequence__gopurs_runtime_Value_1188743833() gopurs_runtime.Value {
	once_sequence__gopurs_runtime_Value_1188743833.Do(func() {
		cache_sequence__gopurs_runtime_Value_1188743833 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence__gopurs_runtime_Value_1188743833(dict_0_box)
})
	})
	return cache_sequence__gopurs_runtime_Value_1188743833
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
sequence2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableMaybe(), "sequence"), dictApplicative_0)
_ = sequence2_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe_First.Get_First(), gopurs_runtime.Apply(sequence2_1_0, v_2))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse2_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableMaybe(), "traverse"), dictApplicative_0)
_ = traverse2_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe_First.Get_First(), gopurs_runtime.Apply2(traverse2_1_1, f_2, v_3))
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
sequence2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableMaybe(), "sequence"), dictApplicative_0)
_ = sequence2_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe_Last.Get_Last(), gopurs_runtime.Apply(sequence2_1_0, v_2))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse2_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableMaybe(), "traverse"), dictApplicative_0)
_ = traverse2_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Maybe_Last.Get_Last(), gopurs_runtime.Apply2(traverse2_1_1, f_2, v_3))
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
		cache_traverseDefault = gopurs_runtime.Func2(func(dictTraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseDefault(dictTraversable_0_box, dictApplicative_1_box)
})
	})
	return cache_traverseDefault
}

var cache_mapAccumR gopurs_runtime.Value
var once_mapAccumR sync.Once
func Get_mapAccumR() gopurs_runtime.Value {
	once_mapAccumR.Do(func() {
		cache_mapAccumR = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumR(dictTraversable_0_box)
})
	})
	return cache_mapAccumR
}

var cache_mapAccumR__gopurs_runtime_Value_2780270502 gopurs_runtime.Value
var once_mapAccumR__gopurs_runtime_Value_2780270502 sync.Once
func Get_mapAccumR__gopurs_runtime_Value_2780270502() gopurs_runtime.Value {
	once_mapAccumR__gopurs_runtime_Value_2780270502.Do(func() {
		cache_mapAccumR__gopurs_runtime_Value_2780270502 = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumR__gopurs_runtime_Value_2780270502(dictTraversable_0_box)
})
	})
	return cache_mapAccumR__gopurs_runtime_Value_2780270502
}

var cache_scanr gopurs_runtime.Value
var once_scanr sync.Once
func Get_scanr() gopurs_runtime.Value {
	once_scanr.Do(func() {
		cache_scanr = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanr(dictTraversable_0_box)
})
	})
	return cache_scanr
}

var cache_mapAccumL gopurs_runtime.Value
var once_mapAccumL sync.Once
func Get_mapAccumL() gopurs_runtime.Value {
	once_mapAccumL.Do(func() {
		cache_mapAccumL = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumL(dictTraversable_0_box)
})
	})
	return cache_mapAccumL
}

var cache_mapAccumL__gopurs_runtime_Value_2780270502 gopurs_runtime.Value
var once_mapAccumL__gopurs_runtime_Value_2780270502 sync.Once
func Get_mapAccumL__gopurs_runtime_Value_2780270502() gopurs_runtime.Value {
	once_mapAccumL__gopurs_runtime_Value_2780270502.Do(func() {
		cache_mapAccumL__gopurs_runtime_Value_2780270502 = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapAccumL__gopurs_runtime_Value_2780270502(dictTraversable_0_box)
})
	})
	return cache_mapAccumL__gopurs_runtime_Value_2780270502
}

var cache_scanl gopurs_runtime.Value
var once_scanl sync.Once
func Get_scanl() gopurs_runtime.Value {
	once_scanl.Do(func() {
		cache_scanl = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_scanl(dictTraversable_0_box)
})
	})
	return cache_scanl
}

var cache_go__for gopurs_runtime.Value
var once_go__for sync.Once
func Get_go__for() gopurs_runtime.Value {
	once_go__for.Do(func() {
		cache_go__for = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictTraversable_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_go__for(dictApplicative_0_box, dictTraversable_1_box)
})
	})
	return cache_go__for
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_traverse(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "traverse")
}

func Call_traverse__gopurs_runtime_Value_2265355576(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "traverse")
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
foldMap4_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "foldMap"), dictMonoid_6)
_ = foldMap4_7_6
foldMap5_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "foldMap"), dictMonoid_6)
_ = foldMap5_8_7
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap4_7_6, gopurs_runtime.Apply(foldMap5_8_7, f_9), v_10)
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
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "foldr"), f_6)
_ = __local_var_9_8
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "foldr"), gopurs_runtime.Func(func(b_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_9_8, a_11, b_10)
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
traverse4_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_6)
_ = traverse4_7_9
traverse5_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), dictApplicative_6)
_ = traverse5_8_10
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Compose.Get_Compose(), gopurs_runtime.Apply2(traverse4_7_9, gopurs_runtime.Apply(traverse5_8_10, f_9), v_10))
})
})
}))
})
}
}

func Call_sequenceDefault(dictTraversable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_1, Get_identity())
}

func Call_sequence(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "sequence")
}

func Call_sequence__gopurs_runtime_Value_1188743833(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "sequence")
}

func Call_traversableApp(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
functorApp_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = functorApp_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_2
foldableApp_2_1 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap3_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "foldMap"), dictMonoid_3)
_ = foldMap3_4_3
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap3_4_3, f_5, v_6)
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
sequence3_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_3)
_ = sequence3_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_App.Get_App(), gopurs_runtime.Apply(sequence3_4_4, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
traverse3_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_3)
_ = traverse3_4_5
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_App.Get_App(), gopurs_runtime.Apply2(traverse3_4_5, f_5, v_6))
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
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_4
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_4, "map"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_8})}
}))
_ = __local_var_8_5
__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_6)
_ = __local_var_9_6
__local_var_10_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_4, "map"), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
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
__local_var_7_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_10
traverse4_8_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_6)
_ = traverse4_8_11
traverse5_9_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), dictApplicative_6)
_ = traverse5_9_12
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_11_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_10, "map"), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_11})}
}))
_ = __local_var_11_13
__local_var_12_14 := gopurs_runtime.Apply(traverse4_8_11, f_10)
_ = __local_var_12_14
__local_var_13_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_10, "map"), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, x_13})}
}))
_ = __local_var_13_15
__local_var_14_16 := gopurs_runtime.Apply(traverse5_9_12, f_10)
_ = __local_var_14_16
return gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(__local_var_11_13, gopurs_runtime.Apply(__local_var_12_14, (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0))
goto end_branch_17
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t17 = gopurs_runtime.Apply(__local_var_13_15, gopurs_runtime.Apply(__local_var_14_16, (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_15.UnsafePtr).V0))
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
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
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_7_4
sequence4_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_6)
_ = sequence4_8_5
sequence5_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "sequence"), dictApplicative_6)
_ = sequence5_9_6
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_4, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_4, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product.Get_product(), gopurs_runtime.Apply(sequence4_8_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0)), gopurs_runtime.Apply(sequence5_9_6, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_7_7
traverse4_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_6)
_ = traverse4_8_8
traverse5_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), dictApplicative_6)
_ = traverse5_9_9
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_7, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Functor_Product.Get_product(), gopurs_runtime.Apply2(traverse4_8_8, f_10, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0)), gopurs_runtime.Apply2(traverse5_9_9, f_10, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1))
})
})
}))
})
}

func Call_traverseDefault(dictTraversable_0_loop gopurs_runtime.Value, dictApplicative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
var dictApplicative_1 gopurs_runtime.Value = dictApplicative_1_loop
_ = dictApplicative_1
sequence3_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_1)
_ = sequence3_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence3_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}), "map"), f_3, ta_4))
})
})
}

func Call_mapAccumR(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
traverse2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR())
_ = traverse2_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(traverse2_1_0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, s_6, a_5)
})
}), xs_4, s0_3)
})
})
})
}

func Call_mapAccumR__gopurs_runtime_Value_2780270502(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
traverse2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), pkg_Data_Traversable_Accum_Internal.Get_applicativeStateR())
_ = traverse2_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(traverse2_1_0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, s_6, a_5)
})
}), xs_4, s0_3)
})
})
})
}

func Call_scanr(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
mapAccumR1_1_0 := Call_mapAccumR(dictTraversable_0)
_ = mapAccumR1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(mapAccumR1_1_0, gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_7_1 := gopurs_runtime.Apply2(f_2, a_6, b_5)
_ = b_prime_7_1
return gopurs_runtime.RecordDict2("accum", "value", b_prime_7_1, b_prime_7_1)
})
}), b0_3, xs_4), "value")
})
})
})
}

func Call_mapAccumL(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
traverse2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL())
_ = traverse2_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(traverse2_1_0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, s_6, a_5)
})
}), xs_4, s0_3)
})
})
})
}

func Call_mapAccumL__gopurs_runtime_Value_2780270502(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
traverse2_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), pkg_Data_Traversable_Accum_Internal.Get_applicativeStateL())
_ = traverse2_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(traverse2_1_0, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, s_6, a_5)
})
}), xs_4, s0_3)
})
})
})
}

func Call_scanl(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
mapAccumL1_1_0 := Call_mapAccumL(dictTraversable_0)
_ = mapAccumL1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b0_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(mapAccumL1_1_0, gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
b_prime_7_1 := gopurs_runtime.Apply2(f_2, b_5, a_6)
_ = b_prime_7_1
return gopurs_runtime.RecordDict2("accum", "value", b_prime_7_1, b_prime_7_1)
})
}), b0_3, xs_4), "value")
})
})
})
}

func Call_go__for(dictApplicative_0_loop gopurs_runtime.Value, dictTraversable_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var dictTraversable_1 gopurs_runtime.Value = dictTraversable_1_loop
_ = dictTraversable_1
traverse2_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_1, "traverse"), dictApplicative_0)
_ = traverse2_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(traverse2_2_0, f_4, x_3)
})
})
}

func Get_traverseArrayImpl() gopurs_runtime.Value {
	return _Gopurs_TraverseArrayImpl
}

package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Traversable_identity gopurs_runtime.Value
var once_Data_Traversable_identity sync.Once
func Get_Data_Traversable_identity() gopurs_runtime.Value {
	once_Data_Traversable_identity.Do(func() {
		cache_Data_Traversable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_identity(x_0_box)
})
	})
	return cache_Data_Traversable_identity
}

var cache_Data_Traversable_append gopurs_runtime.Value
var once_Data_Traversable_append sync.Once
func Get_Data_Traversable_append() gopurs_runtime.Value {
	once_Data_Traversable_append.Do(func() {
		cache_Data_Traversable_append = gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupArray(), "append")
	})
	return cache_Data_Traversable_append
}

var cache_Data_Traversable_Traversable_dollarDict gopurs_runtime.Value
var once_Data_Traversable_Traversable_dollarDict sync.Once
func Get_Data_Traversable_Traversable_dollarDict() gopurs_runtime.Value {
	once_Data_Traversable_Traversable_dollarDict.Do(func() {
		cache_Data_Traversable_Traversable_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_Traversable_dollarDict(x_0_box)
})
	})
	return cache_Data_Traversable_Traversable_dollarDict
}

var cache_Data_Traversable_traverse gopurs_runtime.Value
var once_Data_Traversable_traverse sync.Once
func Get_Data_Traversable_traverse() gopurs_runtime.Value {
	once_Data_Traversable_traverse.Do(func() {
		cache_Data_Traversable_traverse = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dict_0_box))
})
	})
	return cache_Data_Traversable_traverse
}

var cache_Data_Traversable_traversableTuple gopurs_runtime.Value
var once_Data_Traversable_traversableTuple sync.Once
func Get_Data_Traversable_traversableTuple() gopurs_runtime.Value {
	once_Data_Traversable_traversableTuple.Do(func() {
		cache_Data_Traversable_traversableTuple = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_functorTuple()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1))
})
})
}))
	})
	return cache_Data_Traversable_traversableTuple
}

var cache_Data_Traversable_traversableMultiplicative gopurs_runtime.Value
var once_Data_Traversable_traversableMultiplicative sync.Once
func Get_Data_Traversable_traversableMultiplicative() gopurs_runtime.Value {
	once_Data_Traversable_traversableMultiplicative.Do(func() {
		cache_Data_Traversable_traversableMultiplicative = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Multiplicative_functorMultiplicative()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Multiplicative_Multiplicative(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Multiplicative_Multiplicative(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableMultiplicative
}

var cache_Data_Traversable_traversableMaybe gopurs_runtime.Value
var once_Data_Traversable_traversableMaybe sync.Once
func Get_Data_Traversable_traversableMaybe() gopurs_runtime.Value {
	once_Data_Traversable_traversableMaybe.Do(func() {
		cache_Data_Traversable_traversableMaybe = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Just(), (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
// TAST (Let): Functor0_1_2 -> *Constructor_Data_Functor_Functor
Functor0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_2.V0), Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just)(v1_3.UnsafePtr).V0))
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
	return cache_Data_Traversable_traversableMaybe
}

var cache_Data_Traversable_traversableIdentity gopurs_runtime.Value
var once_Data_Traversable_traversableIdentity sync.Once
func Get_Data_Traversable_traversableIdentity() gopurs_runtime.Value {
	once_Data_Traversable_traversableIdentity.Do(func() {
		cache_Data_Traversable_traversableIdentity = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Identity_functorIdentity()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Identity_Identity(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Identity_Identity(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableIdentity
}

var cache_Data_Traversable_traversableEither gopurs_runtime.Value
var once_Data_Traversable_traversableEither sync.Once
func Get_Data_Traversable_traversableEither() gopurs_runtime.Value {
	once_Data_Traversable_traversableEither.Do(func() {
		cache_Data_Traversable_traversableEither = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_functorEither()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v_2.UnsafePtr).V0})})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Right(), (*Constructor_Data_Either_Right)(v_2.UnsafePtr).V0)
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
// TAST (Let): Functor0_1_2 -> *Constructor_Data_Functor_Functor
Functor0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v1_3.UnsafePtr).V0})})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_2.V0), Get_Data_Either_Right(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Either_Right)(v1_3.UnsafePtr).V0))
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
	return cache_Data_Traversable_traversableEither
}

var cache_Data_Traversable_traversableDual gopurs_runtime.Value
var once_Data_Traversable_traversableDual sync.Once
func Get_Data_Traversable_traversableDual() gopurs_runtime.Value {
	once_Data_Traversable_traversableDual.Do(func() {
		cache_Data_Traversable_traversableDual = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Dual_functorDual()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Dual_Dual(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Dual_Dual(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableDual
}

var cache_Data_Traversable_traversableDisj gopurs_runtime.Value
var once_Data_Traversable_traversableDisj sync.Once
func Get_Data_Traversable_traversableDisj() gopurs_runtime.Value {
	once_Data_Traversable_traversableDisj.Do(func() {
		cache_Data_Traversable_traversableDisj = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_functorDisj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Disj_Disj(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Disj_Disj(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableDisj
}

var cache_Data_Traversable_traversableConst gopurs_runtime.Value
var once_Data_Traversable_traversableConst sync.Once
func Get_Data_Traversable_traversableConst() gopurs_runtime.Value {
	once_Data_Traversable_traversableConst.Do(func() {
		cache_Data_Traversable_traversableConst = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Const_functorConst()
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
	return cache_Data_Traversable_traversableConst
}

var cache_Data_Traversable_traversableConj gopurs_runtime.Value
var once_Data_Traversable_traversableConj sync.Once
func Get_Data_Traversable_traversableConj() gopurs_runtime.Value {
	once_Data_Traversable_traversableConj.Do(func() {
		cache_Data_Traversable_traversableConj = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Conj_functorConj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Conj_Conj(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Conj_Conj(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableConj
}

var cache_Data_Traversable_traversableCompose gopurs_runtime.Value
var once_Data_Traversable_traversableCompose sync.Once
func Get_Data_Traversable_traversableCompose() gopurs_runtime.Value {
	once_Data_Traversable_traversableCompose.Do(func() {
		cache_Data_Traversable_traversableCompose = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traversableCompose(dictTraversable_0_box)
})
	})
	return cache_Data_Traversable_traversableCompose
}

var cache_Data_Traversable_traversableAdditive gopurs_runtime.Value
var once_Data_Traversable_traversableAdditive sync.Once
func Get_Data_Traversable_traversableAdditive() gopurs_runtime.Value {
	once_Data_Traversable_traversableAdditive.Do(func() {
		cache_Data_Traversable_traversableAdditive = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Additive_functorAdditive()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Additive_Additive(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Additive_Additive(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableAdditive
}

var cache_Data_Traversable_sequenceDefault gopurs_runtime.Value
var once_Data_Traversable_sequenceDefault sync.Once
func Get_Data_Traversable_sequenceDefault() gopurs_runtime.Value {
	once_Data_Traversable_sequenceDefault.Do(func() {
		cache_Data_Traversable_sequenceDefault = gopurs_runtime.Func2(func(dictTraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_sequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_1_box))
})
	})
	return cache_Data_Traversable_sequenceDefault
}

var cache_Data_Traversable_traversableArray gopurs_runtime.Value
var once_Data_Traversable_traversableArray sync.Once
func Get_Data_Traversable_traversableArray() gopurs_runtime.Value {
	once_Data_Traversable_traversableArray.Do(func() {
		cache_Data_Traversable_traversableArray = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorArray()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableArray(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, Get_Data_Traversable_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> gopurs_runtime.Value
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply4(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupArray(), "append"))
}))
	})
	return cache_Data_Traversable_traversableArray
}

var cache_Data_Traversable_sequence gopurs_runtime.Value
var once_Data_Traversable_sequence sync.Once
func Get_Data_Traversable_sequence() gopurs_runtime.Value {
	once_Data_Traversable_sequence.Do(func() {
		cache_Data_Traversable_sequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dict_0_box))
})
	})
	return cache_Data_Traversable_sequence
}

var cache_Data_Traversable_traversableApp gopurs_runtime.Value
var once_Data_Traversable_traversableApp sync.Once
func Get_Data_Traversable_traversableApp() gopurs_runtime.Value {
	once_Data_Traversable_traversableApp.Do(func() {
		cache_Data_Traversable_traversableApp = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traversableApp(dictTraversable_0_box)
})
	})
	return cache_Data_Traversable_traversableApp
}

var cache_Data_Traversable_traversableCoproduct gopurs_runtime.Value
var once_Data_Traversable_traversableCoproduct sync.Once
func Get_Data_Traversable_traversableCoproduct() gopurs_runtime.Value {
	once_Data_Traversable_traversableCoproduct.Do(func() {
		cache_Data_Traversable_traversableCoproduct = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traversableCoproduct(dictTraversable_0_box)
})
	})
	return cache_Data_Traversable_traversableCoproduct
}

var cache_Data_Traversable_traversableFirst gopurs_runtime.Value
var once_Data_Traversable_traversableFirst sync.Once
func Get_Data_Traversable_traversableFirst() gopurs_runtime.Value {
	once_Data_Traversable_traversableFirst.Do(func() {
		cache_Data_Traversable_traversableFirst = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableFirst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_First_First(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMaybe(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Maybe_First_First(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMaybe(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_3))}))
})
})
}))
	})
	return cache_Data_Traversable_traversableFirst
}

var cache_Data_Traversable_traversableLast gopurs_runtime.Value
var once_Data_Traversable_traversableLast sync.Once
func Get_Data_Traversable_traversableLast() gopurs_runtime.Value {
	once_Data_Traversable_traversableLast.Do(func() {
		cache_Data_Traversable_traversableLast = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableLast()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Last_Last(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMaybe(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Maybe_Last_Last(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMaybe(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_3))}))
})
})
}))
	})
	return cache_Data_Traversable_traversableLast
}

var cache_Data_Traversable_traversableProduct gopurs_runtime.Value
var once_Data_Traversable_traversableProduct sync.Once
func Get_Data_Traversable_traversableProduct() gopurs_runtime.Value {
	once_Data_Traversable_traversableProduct.Do(func() {
		cache_Data_Traversable_traversableProduct = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traversableProduct(dictTraversable_0_box)
})
	})
	return cache_Data_Traversable_traversableProduct
}

var cache_Data_Traversable_traverseDefault gopurs_runtime.Value
var once_Data_Traversable_traverseDefault sync.Once
func Get_Data_Traversable_traverseDefault() gopurs_runtime.Value {
	once_Data_Traversable_traverseDefault.Do(func() {
		cache_Data_Traversable_traverseDefault = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverseDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box))
})
	})
	return cache_Data_Traversable_traverseDefault
}

var cache_Data_Traversable_mapAccumR gopurs_runtime.Value
var once_Data_Traversable_mapAccumR sync.Once
func Get_Data_Traversable_mapAccumR() gopurs_runtime.Value {
	once_Data_Traversable_mapAccumR.Do(func() {
		cache_Data_Traversable_mapAccumR = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_mapAccumR(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_Traversable_mapAccumR
}

var cache_Data_Traversable_scanr gopurs_runtime.Value
var once_Data_Traversable_scanr sync.Once
func Get_Data_Traversable_scanr() gopurs_runtime.Value {
	once_Data_Traversable_scanr.Do(func() {
		cache_Data_Traversable_scanr = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_scanr(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_Data_Traversable_scanr
}

var cache_Data_Traversable_mapAccumL gopurs_runtime.Value
var once_Data_Traversable_mapAccumL sync.Once
func Get_Data_Traversable_mapAccumL() gopurs_runtime.Value {
	once_Data_Traversable_mapAccumL.Do(func() {
		cache_Data_Traversable_mapAccumL = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_mapAccumL(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_Traversable_mapAccumL
}

var cache_Data_Traversable_scanl gopurs_runtime.Value
var once_Data_Traversable_scanl sync.Once
func Get_Data_Traversable_scanl() gopurs_runtime.Value {
	once_Data_Traversable_scanl.Do(func() {
		cache_Data_Traversable_scanl = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_scanl(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_Data_Traversable_scanl
}

var cache_Data_Traversable_go__for gopurs_runtime.Value
var once_Data_Traversable_go__for sync.Once
func Get_Data_Traversable_go__for() gopurs_runtime.Value {
	once_Data_Traversable_go__for.Do(func() {
		cache_Data_Traversable_go__for = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, dictTraversable_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_go__for(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_1_box), x_2_box, f_3_box)
})
	})
	return cache_Data_Traversable_go__for
}

var cache_Data_Traversable_mapAccumL__2528537048 gopurs_runtime.Value
var once_Data_Traversable_mapAccumL__2528537048 sync.Once
func Get_Data_Traversable_mapAccumL__2528537048() gopurs_runtime.Value {
	once_Data_Traversable_mapAccumL__2528537048.Do(func() {
		cache_Data_Traversable_mapAccumL__2528537048 = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_mapAccumL__2528537048(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_Traversable_mapAccumL__2528537048
}

var cache_Data_Traversable_mapAccumL__1189480088 gopurs_runtime.Value
var once_Data_Traversable_mapAccumL__1189480088 sync.Once
func Get_Data_Traversable_mapAccumL__1189480088() gopurs_runtime.Value {
	once_Data_Traversable_mapAccumL__1189480088.Do(func() {
		cache_Data_Traversable_mapAccumL__1189480088 = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_mapAccumL__1189480088(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_Traversable_mapAccumL__1189480088
}

var cache_Data_Traversable_mapAccumR__2528537048 gopurs_runtime.Value
var once_Data_Traversable_mapAccumR__2528537048 sync.Once
func Get_Data_Traversable_mapAccumR__2528537048() gopurs_runtime.Value {
	once_Data_Traversable_mapAccumR__2528537048.Do(func() {
		cache_Data_Traversable_mapAccumR__2528537048 = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_mapAccumR__2528537048(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_Traversable_mapAccumR__2528537048
}

var cache_Data_Traversable_mapAccumR__1189480088 gopurs_runtime.Value
var once_Data_Traversable_mapAccumR__1189480088 sync.Once
func Get_Data_Traversable_mapAccumR__1189480088() gopurs_runtime.Value {
	once_Data_Traversable_mapAccumR__1189480088.Do(func() {
		cache_Data_Traversable_mapAccumR__1189480088 = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_mapAccumR__1189480088(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
})
	})
	return cache_Data_Traversable_mapAccumR__1189480088
}

var cache_Data_Traversable_sequence__1886310617 gopurs_runtime.Value
var once_Data_Traversable_sequence__1886310617 sync.Once
func Get_Data_Traversable_sequence__1886310617() gopurs_runtime.Value {
	once_Data_Traversable_sequence__1886310617.Do(func() {
		cache_Data_Traversable_sequence__1886310617 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_sequence__1886310617(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dict_0_box))
})
	})
	return cache_Data_Traversable_sequence__1886310617
}

var cache_Data_Traversable_sequence__2904194897 gopurs_runtime.Value
var once_Data_Traversable_sequence__2904194897 sync.Once
func Get_Data_Traversable_sequence__2904194897() gopurs_runtime.Value {
	once_Data_Traversable_sequence__2904194897.Do(func() {
		cache_Data_Traversable_sequence__2904194897 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_sequence__2904194897(dictApplicative_0_box)
})
	})
	return cache_Data_Traversable_sequence__2904194897
}

var cache_Data_Traversable_sequence__2616145509 gopurs_runtime.Value
var once_Data_Traversable_sequence__2616145509 sync.Once
func Get_Data_Traversable_sequence__2616145509() gopurs_runtime.Value {
	once_Data_Traversable_sequence__2616145509.Do(func() {
		cache_Data_Traversable_sequence__2616145509 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_sequence__2616145509(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](dictApplicative_0_box))
})
	})
	return cache_Data_Traversable_sequence__2616145509
}

var cache_Data_Traversable_sequence__2322703601 gopurs_runtime.Value
var once_Data_Traversable_sequence__2322703601 sync.Once
func Get_Data_Traversable_sequence__2322703601() gopurs_runtime.Value {
	once_Data_Traversable_sequence__2322703601.Do(func() {
		cache_Data_Traversable_sequence__2322703601 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_sequence__2322703601(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](dictApplicative_0_box))
})
	})
	return cache_Data_Traversable_sequence__2322703601
}

var cache_Data_Traversable_sequence__3720417425 gopurs_runtime.Value
var once_Data_Traversable_sequence__3720417425 sync.Once
func Get_Data_Traversable_sequence__3720417425() gopurs_runtime.Value {
	once_Data_Traversable_sequence__3720417425.Do(func() {
		cache_Data_Traversable_sequence__3720417425 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_sequence__3720417425(__eta0_0_box)
})
	})
	return cache_Data_Traversable_sequence__3720417425
}

var cache_Data_Traversable_sequence__3634000753 gopurs_runtime.Value
var once_Data_Traversable_sequence__3634000753 sync.Once
func Get_Data_Traversable_sequence__3634000753() gopurs_runtime.Value {
	once_Data_Traversable_sequence__3634000753.Do(func() {
		cache_Data_Traversable_sequence__3634000753 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_sequence__3634000753(__eta0_0_box)
})
	})
	return cache_Data_Traversable_sequence__3634000753
}

var cache_Data_Traversable_traversableAdditive__3840848827 gopurs_runtime.Value
var once_Data_Traversable_traversableAdditive__3840848827 sync.Once
func Get_Data_Traversable_traversableAdditive__3840848827() gopurs_runtime.Value {
	once_Data_Traversable_traversableAdditive__3840848827.Do(func() {
		cache_Data_Traversable_traversableAdditive__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableAdditive()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Additive_functorAdditive()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Additive_Additive(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Additive_Additive(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableAdditive__3840848827
}

var cache_Data_Traversable_traversableArray__2090378122 gopurs_runtime.Value
var once_Data_Traversable_traversableArray__2090378122 sync.Once
func Get_Data_Traversable_traversableArray__2090378122() gopurs_runtime.Value {
	once_Data_Traversable_traversableArray__2090378122.Do(func() {
		cache_Data_Traversable_traversableArray__2090378122 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorArray()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableArray(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, Get_Data_Traversable_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> gopurs_runtime.Value
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply4(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupArray(), "append"))
}))
	})
	return cache_Data_Traversable_traversableArray__2090378122
}

var cache_Data_Traversable_traversableArray__2643873085 gopurs_runtime.Value
var once_Data_Traversable_traversableArray__2643873085 sync.Once
func Get_Data_Traversable_traversableArray__2643873085() gopurs_runtime.Value {
	once_Data_Traversable_traversableArray__2643873085.Do(func() {
		cache_Data_Traversable_traversableArray__2643873085 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Functor_functorArray()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableArray(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, Get_Data_Traversable_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 -> gopurs_runtime.Value
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply4(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupArray(), "append"))
}))
	})
	return cache_Data_Traversable_traversableArray__2643873085
}

var cache_Data_Traversable_traversableConj__3840848827 gopurs_runtime.Value
var once_Data_Traversable_traversableConj__3840848827 sync.Once
func Get_Data_Traversable_traversableConj__3840848827() gopurs_runtime.Value {
	once_Data_Traversable_traversableConj__3840848827.Do(func() {
		cache_Data_Traversable_traversableConj__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableConj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Conj_functorConj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Conj_Conj(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Conj_Conj(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableConj__3840848827
}

var cache_Data_Traversable_traversableConst__3861086397 gopurs_runtime.Value
var once_Data_Traversable_traversableConst__3861086397 sync.Once
func Get_Data_Traversable_traversableConst__3861086397() gopurs_runtime.Value {
	once_Data_Traversable_traversableConst__3861086397.Do(func() {
		cache_Data_Traversable_traversableConst__3861086397 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableConst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Const_functorConst()
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
	return cache_Data_Traversable_traversableConst__3861086397
}

var cache_Data_Traversable_traversableDisj__3840848827 gopurs_runtime.Value
var once_Data_Traversable_traversableDisj__3840848827 sync.Once
func Get_Data_Traversable_traversableDisj__3840848827() gopurs_runtime.Value {
	once_Data_Traversable_traversableDisj__3840848827.Do(func() {
		cache_Data_Traversable_traversableDisj__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableDisj()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Disj_functorDisj()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Disj_Disj(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Disj_Disj(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableDisj__3840848827
}

var cache_Data_Traversable_traversableDual__3840848827 gopurs_runtime.Value
var once_Data_Traversable_traversableDual__3840848827 sync.Once
func Get_Data_Traversable_traversableDual__3840848827() gopurs_runtime.Value {
	once_Data_Traversable_traversableDual__3840848827.Do(func() {
		cache_Data_Traversable_traversableDual__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableDual()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Dual_functorDual()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Dual_Dual(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Dual_Dual(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableDual__3840848827
}

var cache_Data_Traversable_traversableEither__4232556979 gopurs_runtime.Value
var once_Data_Traversable_traversableEither__4232556979 sync.Once
func Get_Data_Traversable_traversableEither__4232556979() gopurs_runtime.Value {
	once_Data_Traversable_traversableEither__4232556979.Do(func() {
		cache_Data_Traversable_traversableEither__4232556979 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableEither()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Either_functorEither()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v_2.UnsafePtr).V0})})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Right(), (*Constructor_Data_Either_Right)(v_2.UnsafePtr).V0)
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
// TAST (Let): Functor0_1_2 -> *Constructor_Data_Functor_Functor
Functor0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v1_3.UnsafePtr).V0})})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_2.V0), Get_Data_Either_Right(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Either_Right)(v1_3.UnsafePtr).V0))
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
	return cache_Data_Traversable_traversableEither__4232556979
}

var cache_Data_Traversable_traversableFirst__822285914 gopurs_runtime.Value
var once_Data_Traversable_traversableFirst__822285914 sync.Once
func Get_Data_Traversable_traversableFirst__822285914() gopurs_runtime.Value {
	once_Data_Traversable_traversableFirst__822285914.Do(func() {
		cache_Data_Traversable_traversableFirst__822285914 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableFirst()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_First_First(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMaybe(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Maybe_First_First(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMaybe(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_3))}))
})
})
}))
	})
	return cache_Data_Traversable_traversableFirst__822285914
}

var cache_Data_Traversable_traversableIdentity__3840848827 gopurs_runtime.Value
var once_Data_Traversable_traversableIdentity__3840848827 sync.Once
func Get_Data_Traversable_traversableIdentity__3840848827() gopurs_runtime.Value {
	once_Data_Traversable_traversableIdentity__3840848827.Do(func() {
		cache_Data_Traversable_traversableIdentity__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableIdentity()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Identity_functorIdentity()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Identity_Identity(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Identity_Identity(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableIdentity__3840848827
}

var cache_Data_Traversable_traversableLast__822285914 gopurs_runtime.Value
var once_Data_Traversable_traversableLast__822285914 sync.Once
func Get_Data_Traversable_traversableLast__822285914() gopurs_runtime.Value {
	once_Data_Traversable_traversableLast__822285914.Do(func() {
		cache_Data_Traversable_traversableLast__822285914 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableLast()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Last_Last(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMaybe(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_2))}))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Maybe_Last_Last(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableMaybe(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_3))}))
})
})
}))
	})
	return cache_Data_Traversable_traversableLast__822285914
}

var cache_Data_Traversable_traversableMaybe__98548986 gopurs_runtime.Value
var once_Data_Traversable_traversableMaybe__98548986 sync.Once
func Get_Data_Traversable_traversableMaybe__98548986() gopurs_runtime.Value {
	once_Data_Traversable_traversableMaybe__98548986.Do(func() {
		cache_Data_Traversable_traversableMaybe__98548986 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Just(), (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
// TAST (Let): Functor0_1_2 -> *Constructor_Data_Functor_Functor
Functor0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_2.V0), Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just)(v1_3.UnsafePtr).V0))
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
	return cache_Data_Traversable_traversableMaybe__98548986
}

var cache_Data_Traversable_traversableMaybe__822285914 gopurs_runtime.Value
var once_Data_Traversable_traversableMaybe__822285914 sync.Once
func Get_Data_Traversable_traversableMaybe__822285914() gopurs_runtime.Value {
	once_Data_Traversable_traversableMaybe__822285914.Do(func() {
		cache_Data_Traversable_traversableMaybe__822285914 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMaybe()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_functorMaybe()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Just(), (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
// TAST (Let): Functor0_1_2 -> *Constructor_Data_Functor_Functor
Functor0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_2.V0), Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just)(v1_3.UnsafePtr).V0))
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
	return cache_Data_Traversable_traversableMaybe__822285914
}

var cache_Data_Traversable_traversableMultiplicative__3840848827 gopurs_runtime.Value
var once_Data_Traversable_traversableMultiplicative__3840848827 sync.Once
func Get_Data_Traversable_traversableMultiplicative__3840848827() gopurs_runtime.Value {
	once_Data_Traversable_traversableMultiplicative__3840848827.Do(func() {
		cache_Data_Traversable_traversableMultiplicative__3840848827 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMultiplicative()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Monoid_Multiplicative_functorMultiplicative()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Multiplicative_Multiplicative(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Multiplicative_Multiplicative(), gopurs_runtime.Apply(f_2, v_3))
})
})
}))
	})
	return cache_Data_Traversable_traversableMultiplicative__3840848827
}

var cache_Data_Traversable_traversableTuple__3228991731 gopurs_runtime.Value
var once_Data_Traversable_traversableTuple__3228991731 sync.Once
func Get_Data_Traversable_traversableTuple__3228991731() gopurs_runtime.Value {
	once_Data_Traversable_traversableTuple__3228991731.Do(func() {
		cache_Data_Traversable_traversableTuple__3228991731 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableTuple()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_functorTuple()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1))
})
})
}))
	})
	return cache_Data_Traversable_traversableTuple__3228991731
}

var cache_Data_Traversable_traverse__3956862083 gopurs_runtime.Value
var once_Data_Traversable_traverse__3956862083 sync.Once
func Get_Data_Traversable_traverse__3956862083() gopurs_runtime.Value {
	once_Data_Traversable_traverse__3956862083.Do(func() {
		cache_Data_Traversable_traverse__3956862083 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__3956862083(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dict_0_box))
})
	})
	return cache_Data_Traversable_traverse__3956862083
}

var cache_Data_Traversable_traverse__1933302275 gopurs_runtime.Value
var once_Data_Traversable_traverse__1933302275 sync.Once
func Get_Data_Traversable_traverse__1933302275() gopurs_runtime.Value {
	once_Data_Traversable_traverse__1933302275.Do(func() {
		cache_Data_Traversable_traverse__1933302275 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__1933302275(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dict_0_box))
})
	})
	return cache_Data_Traversable_traverse__1933302275
}

var cache_Data_Traversable_traverse__314957093 gopurs_runtime.Value
var once_Data_Traversable_traverse__314957093 sync.Once
func Get_Data_Traversable_traverse__314957093() gopurs_runtime.Value {
	once_Data_Traversable_traverse__314957093.Do(func() {
		cache_Data_Traversable_traverse__314957093 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__314957093(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dict_0_box))
})
	})
	return cache_Data_Traversable_traverse__314957093
}

var cache_Data_Traversable_traverse__303957893 gopurs_runtime.Value
var once_Data_Traversable_traverse__303957893 sync.Once
func Get_Data_Traversable_traverse__303957893() gopurs_runtime.Value {
	once_Data_Traversable_traverse__303957893.Do(func() {
		cache_Data_Traversable_traverse__303957893 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__303957893(dictApplicative_0_box)
})
	})
	return cache_Data_Traversable_traverse__303957893
}

var cache_Data_Traversable_traverse__3459678245 gopurs_runtime.Value
var once_Data_Traversable_traverse__3459678245 sync.Once
func Get_Data_Traversable_traverse__3459678245() gopurs_runtime.Value {
	once_Data_Traversable_traverse__3459678245.Do(func() {
		cache_Data_Traversable_traverse__3459678245 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__3459678245(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dict_0_box))
})
	})
	return cache_Data_Traversable_traverse__3459678245
}

var cache_Data_Traversable_traverse__23995045 gopurs_runtime.Value
var once_Data_Traversable_traverse__23995045 sync.Once
func Get_Data_Traversable_traverse__23995045() gopurs_runtime.Value {
	once_Data_Traversable_traverse__23995045.Do(func() {
		cache_Data_Traversable_traverse__23995045 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__23995045(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dict_0_box))
})
	})
	return cache_Data_Traversable_traverse__23995045
}

var cache_Data_Traversable_traverse__4126651533 gopurs_runtime.Value
var once_Data_Traversable_traverse__4126651533 sync.Once
func Get_Data_Traversable_traverse__4126651533() gopurs_runtime.Value {
	once_Data_Traversable_traverse__4126651533.Do(func() {
		cache_Data_Traversable_traverse__4126651533 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__4126651533(dictApplicative_0_box)
})
	})
	return cache_Data_Traversable_traverse__4126651533
}

var cache_Data_Traversable_traverse__894989549 gopurs_runtime.Value
var once_Data_Traversable_traverse__894989549 sync.Once
func Get_Data_Traversable_traverse__894989549() gopurs_runtime.Value {
	once_Data_Traversable_traverse__894989549.Do(func() {
		cache_Data_Traversable_traverse__894989549 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__894989549(dictApplicative_0_box)
})
	})
	return cache_Data_Traversable_traverse__894989549
}

var cache_Data_Traversable_traverse__1157172365 gopurs_runtime.Value
var once_Data_Traversable_traverse__1157172365 sync.Once
func Get_Data_Traversable_traverse__1157172365() gopurs_runtime.Value {
	once_Data_Traversable_traverse__1157172365.Do(func() {
		cache_Data_Traversable_traverse__1157172365 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__1157172365(dictApplicative_0_box)
})
	})
	return cache_Data_Traversable_traverse__1157172365
}

var cache_Data_Traversable_traverse__878259545 gopurs_runtime.Value
var once_Data_Traversable_traverse__878259545 sync.Once
func Get_Data_Traversable_traverse__878259545() gopurs_runtime.Value {
	once_Data_Traversable_traverse__878259545.Do(func() {
		cache_Data_Traversable_traverse__878259545 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__878259545(dictApplicative_0_box)
})
	})
	return cache_Data_Traversable_traverse__878259545
}

var cache_Data_Traversable_traverse__2839486329 gopurs_runtime.Value
var once_Data_Traversable_traverse__2839486329 sync.Once
func Get_Data_Traversable_traverse__2839486329() gopurs_runtime.Value {
	once_Data_Traversable_traverse__2839486329.Do(func() {
		cache_Data_Traversable_traverse__2839486329 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__2839486329(dictApplicative_0_box)
})
	})
	return cache_Data_Traversable_traverse__2839486329
}

var cache_Data_Traversable_traverse__3640625269 gopurs_runtime.Value
var once_Data_Traversable_traverse__3640625269 sync.Once
func Get_Data_Traversable_traverse__3640625269() gopurs_runtime.Value {
	once_Data_Traversable_traverse__3640625269.Do(func() {
		cache_Data_Traversable_traverse__3640625269 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__3640625269(dictApplicative_0_box)
})
	})
	return cache_Data_Traversable_traverse__3640625269
}

var cache_Data_Traversable_traverse__667327821 gopurs_runtime.Value
var once_Data_Traversable_traverse__667327821 sync.Once
func Get_Data_Traversable_traverse__667327821() gopurs_runtime.Value {
	once_Data_Traversable_traverse__667327821.Do(func() {
		cache_Data_Traversable_traverse__667327821 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__667327821(dictApplicative_0_box)
})
	})
	return cache_Data_Traversable_traverse__667327821
}

var cache_Data_Traversable_traverse__3246764013 gopurs_runtime.Value
var once_Data_Traversable_traverse__3246764013 sync.Once
func Get_Data_Traversable_traverse__3246764013() gopurs_runtime.Value {
	once_Data_Traversable_traverse__3246764013.Do(func() {
		cache_Data_Traversable_traverse__3246764013 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__3246764013(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Traversable_traverse__3246764013
}

var cache_Data_Traversable_traverse__694301997 gopurs_runtime.Value
var once_Data_Traversable_traverse__694301997 sync.Once
func Get_Data_Traversable_traverse__694301997() gopurs_runtime.Value {
	once_Data_Traversable_traverse__694301997.Do(func() {
		cache_Data_Traversable_traverse__694301997 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__694301997(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Traversable_traverse__694301997
}

var cache_Data_Traversable_traverse__4214075557 gopurs_runtime.Value
var once_Data_Traversable_traverse__4214075557 sync.Once
func Get_Data_Traversable_traverse__4214075557() gopurs_runtime.Value {
	once_Data_Traversable_traverse__4214075557.Do(func() {
		cache_Data_Traversable_traverse__4214075557 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse__4214075557(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](dict_0_box))
})
	})
	return cache_Data_Traversable_traverse__4214075557
}

type Constructor_Data_Traversable_Traversable struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3941073978] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Traversable_Traversable)(ptr)
		_ = c
		switch key {
		case "Foldable1": return gopurs_runtime.Box(c.V0)
		case "Functor0": return gopurs_runtime.Box(c.V1)
		case "sequence": return gopurs_runtime.Box(c.V2)
		case "traverse": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_Traversable_Traversable: " + key)
		}
	}
}


func Call_Data_Traversable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Traversable_Traversable_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Traversable_traverse(dict_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Traversable_traversableCompose(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
traversableCompose:
for {
if false { continue traversableCompose }
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorCompose1_4_2 -> gopurs_runtime.Value
functorCompose1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), f_5), v_6)
})
}))
_ = functorCompose1_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): foldableCompose1_5_4 -> gopurs_runtime.Value
foldableCompose1_5_4 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_6))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_6))}, f_7), v_8)
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
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Call_Data_Traversable_traversableCompose(dictTraversable_0), dictTraversable1_3), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, Get_Data_Traversable_identity())
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_7 -> *Constructor_Data_Functor_Functor
Functor0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_7
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_7.V0), Get_Data_Functor_Compose_Compose(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, f_8), v_9))
})
})
}))
})
}
}

func Call_Data_Traversable_sequenceDefault(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable, dictApplicative_1_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
_ = dictTraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, Get_Data_Traversable_identity())
}

func Call_Data_Traversable_sequence(dict_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Traversable_traversableApp(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): functorApp_1_0 -> gopurs_runtime.Value
functorApp_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = functorApp_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): foldableApp_2_1 -> gopurs_runtime.Value
foldableApp_2_1 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_2, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, f_4, v_5)
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
// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), Get_Data_Functor_App_App(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_4 -> *Constructor_Data_Functor_Functor
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), Get_Data_Functor_App_App(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_5, v_6))
})
})
}))
}

func Call_Data_Traversable_traversableCoproduct(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): functorCoproduct_1_0 -> gopurs_runtime.Value
functorCoproduct_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Coproduct_functorCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct_1_0
// TAST (Let): foldableCoproduct_2_1 -> gopurs_runtime.Value
foldableCoproduct_2_1 := gopurs_runtime.Apply(Get_Data_Foldable_foldableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableCoproduct_2_1
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCoproduct1_4_2 -> gopurs_runtime.Value
functorCoproduct1_4_2 := gopurs_runtime.Apply(functorCoproduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{}))
_ = functorCoproduct1_4_2
// TAST (Let): foldableCoproduct1_5_3 -> gopurs_runtime.Value
foldableCoproduct1_5_3 := gopurs_runtime.Apply(foldableCoproduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableCoproduct1_5_3
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableCoproduct1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorCoproduct1_4_2
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_4 -> *Constructor_Data_Functor_Functor
Functor0_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_4
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_7_4.V0), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_8})}
}))
_ = __local_var_8_6
// TAST (Let): __local_var_9_7 -> gopurs_runtime.Value
__local_var_9_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_6)
_ = __local_var_9_7
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_6, gopurs_runtime.Apply(__local_var_9_7, x_10))
})
_ = __local_var_8_5
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_7_4.V0), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})}
}))
_ = __local_var_9_9
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "sequence"), dictApplicative_6)
_ = __local_var_10_10
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_9, gopurs_runtime.Apply(__local_var_10_10, x_11))
})
_ = __local_var_9_8
return gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t11 = gopurs_runtime.Apply(__local_var_8_5, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0)
goto end_branch_11
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t11 = gopurs_runtime.Apply(__local_var_9_8, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_12 -> *Constructor_Data_Functor_Functor
Functor0_7_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_12
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_14 -> gopurs_runtime.Value
__local_var_9_14 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_7_12.V0), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_9})}
}))
_ = __local_var_9_14
// TAST (Let): __local_var_10_15 -> gopurs_runtime.Value
__local_var_10_15 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, f_8)
_ = __local_var_10_15
// TAST (Let): __local_var_9_13 -> gopurs_runtime.Value
__local_var_9_13 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_14, gopurs_runtime.Apply(__local_var_10_15, x_11))
})
_ = __local_var_9_13
// TAST (Let): __local_var_10_17 -> gopurs_runtime.Value
__local_var_10_17 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_7_12.V0), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})}
}))
_ = __local_var_10_17
// TAST (Let): __local_var_11_18 -> gopurs_runtime.Value
__local_var_11_18 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, f_8)
_ = __local_var_11_18
// TAST (Let): __local_var_10_16 -> gopurs_runtime.Value
__local_var_10_16 := gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_17, gopurs_runtime.Apply(__local_var_11_18, x_12))
})
_ = __local_var_10_16
return gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(__local_var_9_13, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t19 = gopurs_runtime.Apply(__local_var_10_16, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
})
})
}))
})
}

func Call_Data_Traversable_traversableProduct(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): functorProduct_1_0 -> gopurs_runtime.Value
functorProduct_1_0 := gopurs_runtime.Apply(Get_Data_Functor_Product_functorProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorProduct_1_0
// TAST (Let): foldableProduct_2_1 -> gopurs_runtime.Value
foldableProduct_2_1 := gopurs_runtime.Apply(Get_Data_Foldable_foldableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableProduct_2_1
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorProduct1_4_2 -> gopurs_runtime.Value
functorProduct1_4_2 := gopurs_runtime.Apply(functorProduct_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{}))
_ = functorProduct1_4_2
// TAST (Let): foldableProduct1_5_3 -> gopurs_runtime.Value
foldableProduct1_5_3 := gopurs_runtime.Apply(foldableProduct_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableProduct1_5_3
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableProduct1_5_3
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorProduct1_4_2
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_4 -> *Constructor_Control_Apply_Apply
Apply0_7_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_4
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_4.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_7_4.V0), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Product_product(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_3, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_5 -> *Constructor_Control_Apply_Apply
Apply0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_5
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_5.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_7_5.V0), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Product_product(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, f_8, (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, f_8, (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1))
})
})
}))
})
}

func Call_Data_Traversable_traverseDefault(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversable_0.V1), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(dictApplicative_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable_0.V2), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_2))}, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_3, ta_4))
})
})
})
}

func Call_Data_Traversable_mapAccumR(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateR()))}, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_Data_Traversable_scanr(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable, f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.RecordGet(Call_Data_Traversable_mapAccumR(dictTraversable_0, gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): b_prime_6_0 -> gopurs_runtime.Value
b_prime_6_0 := gopurs_runtime.Apply2(f_1, a_5, b_4)
_ = b_prime_6_0
return gopurs_runtime.RecordDict2("accum", "value", b_prime_6_0, b_prime_6_0)
})
}), b0_2, xs_3), "value")
}

func Call_Data_Traversable_mapAccumL(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateL()))}, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_Data_Traversable_scanl(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable, f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.RecordGet(Call_Data_Traversable_mapAccumL(dictTraversable_0, gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): b_prime_6_0 -> gopurs_runtime.Value
b_prime_6_0 := gopurs_runtime.Apply2(f_1, b_4, a_5)
_ = b_prime_6_0
return gopurs_runtime.RecordDict2("accum", "value", b_prime_6_0, b_prime_6_0)
})
}), b0_2, xs_3), "value")
}

func Call_Data_Traversable_go__for(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, dictTraversable_1_loop *Constructor_Data_Traversable_Traversable, x_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var dictTraversable_1 *Constructor_Data_Traversable_Traversable = dictTraversable_1_loop
_ = dictTraversable_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictTraversable_1.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, f_3, x_2)
}

func Call_Data_Traversable_mapAccumL__2528537048(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateL()))}, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_Data_Traversable_mapAccumL__1189480088(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateL()))}, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_Data_Traversable_mapAccumR__2528537048(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateR()))}, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_Data_Traversable_mapAccumR__1189480088(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable, f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Data_Traversable_Accum_Internal_applicativeStateR()))}, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, s_5, a_4)
})
}), xs_3, s0_2)
}

func Call_Data_Traversable_sequence__1886310617(dict_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Traversable_sequence__2904194897(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0))}, Get_Data_List_Lazy_Types_identity())
}

func Call_Data_Traversable_sequence__2616145509(dictApplicative_0_loop *Constructor_Data_List_Types_Cons) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Data_List_Types_Cons = dictApplicative_0_loop
_ = dictApplicative_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_traversableList(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, Get_Data_List_Types_identity())
}

func Call_Data_Traversable_sequence__2322703601(dictApplicative_0_loop *Constructor_Data_Maybe_Just) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Data_Maybe_Just = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Just(), (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0)
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
}

func Call_Data_Traversable_sequence__3720417425(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_List_Lazy_Types_traversableNonEmpty()).V2), __eta0_0)
}

func Call_Data_Traversable_sequence__3634000753(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_List_Lazy_Types_traversableNonEmpty()).V2), __eta0_0)
}

func Call_Data_Traversable_traverse__3956862083(dict_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Traversable_traverse__1933302275(dict_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Traversable_traverse__314957093(dict_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Traversable_traverse__303957893(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 -> gopurs_runtime.Value
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply4(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupArray(), "append"))
}

func Call_Data_Traversable_traverse__3459678245(dict_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Traversable_traverse__23995045(dict_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Traversable_traverse__4126651533(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_4.Type == 9 && v1_4.IntVal == 237113226) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Interval_StartEnd(), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartEnd)(v1_4.UnsafePtr).V0)), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartEnd)(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 1992629780) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Apply(Get_Data_Interval_DurationEnd(), (*Constructor_Data_Interval_DurationEnd)(v1_4.UnsafePtr).V0), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_DurationEnd)(v1_4.UnsafePtr).V1))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2020675835) {
// TAST (Let): __local_var_5_2 -> gopurs_runtime.Value
__local_var_5_2 := (*Constructor_Data_Interval_StartDuration)(v1_4.UnsafePtr).V1
_ = __local_var_5_2
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Interval_StartDuration(), v2_6, __local_var_5_2)
}), gopurs_runtime.Apply(v_3, (*Constructor_Data_Interval_StartDuration)(v1_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v1_4.Type == 9 && v1_4.IntVal == 2281256335) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Apply(Get_Data_Interval_DurationOnly(), (*Constructor_Data_Interval_DurationOnly)(v1_4.UnsafePtr).V0))
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
}

func Call_Data_Traversable_traverse__894989549(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_3, a_4)), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()))
})
}

func Call_Data_Traversable_traverse__1157172365(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_foldableList(), "foldr"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_List_Lazy_Types_cons(), gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](a_4))})), b_5)
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Lazy_Types_nil()))
})
}

func Call_Data_Traversable_traverse__878259545(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 -> *Constructor_Control_Apply_Apply
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), a_5, b_4)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](Get_Data_List_Types_Nil()))}))
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_6_5 -> *Constructor_Data_Functor_Functor
Functor0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_2_1.V0), gopurs_runtime.Value{}))
_ = Functor0_6_5
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_5.V0), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), a_9, b_8)
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Types_Nil()))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}

func Call_Data_Traversable_traverse__2839486329(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): Apply0_2_1 -> *Constructor_Control_Apply_Apply
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), a_5, b_4)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](Get_Data_List_Types_Nil()))}))
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_6_5 -> *Constructor_Data_Functor_Functor
Functor0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_2_1.V0), gopurs_runtime.Value{}))
_ = Functor0_6_5
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_5.V0), gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_List_Types_Cons(), a_9, b_8)
})
}), acc_5), b_7)
})
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, gopurs_runtime.Apply(f_3, x_7))
})
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_List_Types_Nil()))
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_5_3, x_6))
})
})
}

func Call_Data_Traversable_traverse__3640625269(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_0 gopurs_runtime.Value
_ = go__go_4_2_0
go__go_4_2_0 = gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_Map_Internal_Leaf())
goto end_branch_6
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 324739070 && v_5.UnsafePtr != nil) {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
var __local_var_6_3 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V0)
// TAST (Let): __local_var_7_4 -> gopurs_runtime.Value
__local_var_7_4 := (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V2
_ = __local_var_7_4
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
var __local_var_8_5 gopurs_runtime.Value = gopurs_runtime.Int((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V1)
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(l_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply6(Get_Data_Map_Internal_Node(), gopurs_runtime.Int(__local_var_6_3.IntVal), gopurs_runtime.Int(__local_var_8_5.IntVal), __local_var_7_4, v_prime_10, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](l_prime_9))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](r_prime_11))})))}
})
})
}), gopurs_runtime.Apply(go__go_4_2_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V4)})), gopurs_runtime.Apply(f_3, (*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V3)), gopurs_runtime.Apply(go__go_4_2_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v_5.UnsafePtr).V5)}))
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
return go__go_4_2_0
})
}

func Call_Data_Traversable_traverse__667327821(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just)(v1_3.UnsafePtr).V0))
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
}

func Call_Data_Traversable_traverse__3246764013(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_List_Lazy_Types_traversableNonEmpty()).V3), __eta0_0, __eta1_1)
}

func Call_Data_Traversable_traverse__694301997(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_List_Lazy_Types_traversableNonEmpty()).V3), __eta0_0, __eta1_1)
}

func Call_Data_Traversable_traverse__4214075557(dict_0_loop *Constructor_Data_Traversable_Traversable) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Get_Data_Traversable_traverseArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Traversable_TraverseArrayImpl
}

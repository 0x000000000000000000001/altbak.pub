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
		cache_Data_Traversable_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Traversable_identity
}

var cache_Data_Traversable_go__append gopurs_runtime.Value
var once_Data_Traversable_go__append sync.Once
func Get_Data_Traversable_go__append() gopurs_runtime.Value {
	once_Data_Traversable_go__append.Do(func() {
		cache_Data_Traversable_go__append = Get_Data_Semigroup_concatArray()
	})
	return cache_Data_Traversable_go__append
}

var cache_Data_Traversable_Traversable_dollar_Dict gopurs_runtime.Value
var once_Data_Traversable_Traversable_dollar_Dict sync.Once
func Get_Data_Traversable_Traversable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Traversable_Traversable_dollar_Dict.Do(func() {
		cache_Data_Traversable_Traversable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Call_Data_Traversable_Traversable_dollar_Dict(func() struct{
	Foldable1 gopurs_runtime.Value
	Functor0 gopurs_runtime.Value
	sequence gopurs_runtime.Value
	traverse gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Foldable1 gopurs_runtime.Value
	Functor0 gopurs_runtime.Value
	sequence gopurs_runtime.Value
	traverse gopurs_runtime.Value
}{}
					clone.Foldable1 = gopurs_runtime.RecordGet(orig, "Foldable1")
					clone.Functor0 = gopurs_runtime.RecordGet(orig, "Functor0")
					clone.sequence = gopurs_runtime.RecordGet(orig, "sequence")
					clone.traverse = gopurs_runtime.RecordGet(orig, "traverse")
					return clone
				}()))}
})
	})
	return cache_Data_Traversable_Traversable_dollar_Dict
}

var cache_Data_Traversable_traverse gopurs_runtime.Value
var once_Data_Traversable_traverse sync.Once
func Get_Data_Traversable_traverse() gopurs_runtime.Value {
	once_Data_Traversable_traverse.Do(func() {
		cache_Data_Traversable_traverse = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Traversable_traverse
}

var cache_Data_Traversable_traversableTuple gopurs_runtime.Value
var once_Data_Traversable_traversableTuple sync.Once
func Get_Data_Traversable_traversableTuple() gopurs_runtime.Value {
	once_Data_Traversable_traversableTuple.Do(func() {
		cache_Data_Traversable_traversableTuple = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3543431075_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_4173511203_1680800814(Rebox_Data_Traversable_1680800814_4173511203(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableTuple()))))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_2363162019_2812149806(Rebox_Data_Traversable_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Tuple_functorTuple()))))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
})})))}
	})
	return cache_Data_Traversable_traversableTuple
}

var cache_Data_Traversable_traversableMultiplicative gopurs_runtime.Value
var once_Data_Traversable_traversableMultiplicative sync.Once
func Get_Data_Traversable_traversableMultiplicative() gopurs_runtime.Value {
	once_Data_Traversable_traversableMultiplicative.Do(func() {
		cache_Data_Traversable_traversableMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Multiplicative_functorMultiplicative()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Multiplicative_Multiplicative(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Multiplicative_Multiplicative(), gopurs_runtime.Apply(f_2, v_3))
})
})}))}
	})
	return cache_Data_Traversable_traversableMultiplicative
}

var cache_Data_Traversable_traversableMaybe gopurs_runtime.Value
var once_Data_Traversable_traversableMaybe sync.Once
func Get_Data_Traversable_traversableMaybe() gopurs_runtime.Value {
	once_Data_Traversable_traversableMaybe.Do(func() {
		cache_Data_Traversable_traversableMaybe = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3188237647_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_1146820559_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableMaybe())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3689823567_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_functorMaybe())))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_1 == nil) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
goto end_branch_3
} else {

}
}
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_2 != nil) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Just(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_4
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
var __t_tag_5 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
if (__t_tag_5 == nil) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
goto end_branch_7
} else {

}
}
{
var __t_tag_6 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v1_3)
if (__t_tag_6 != nil) {
__t7 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_4.V0), Get_Data_Maybe_Just(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
})})))}
	})
	return cache_Data_Traversable_traversableMaybe
}

var cache_Data_Traversable_traversableIdentity gopurs_runtime.Value
var once_Data_Traversable_traversableIdentity sync.Once
func Get_Data_Traversable_traversableIdentity() gopurs_runtime.Value {
	once_Data_Traversable_traversableIdentity.Do(func() {
		cache_Data_Traversable_traversableIdentity = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableIdentity()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Identity_functorIdentity()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Identity_Identity(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Identity_Identity(), gopurs_runtime.Apply(f_2, v_3))
})
})}))}
	})
	return cache_Data_Traversable_traversableIdentity
}

var cache_Data_Traversable_traversableEither gopurs_runtime.Value
var once_Data_Traversable_traversableEither sync.Once
func Get_Data_Traversable_traversableEither() gopurs_runtime.Value {
	once_Data_Traversable_traversableEither.Do(func() {
		cache_Data_Traversable_traversableEither = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableEither()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Either_functorEither()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Right(), (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
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
// TAST (Let): Functor0_1_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{(*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}())
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_2.V0), Get_Data_Either_Right(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0))
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
})}))}
	})
	return cache_Data_Traversable_traversableEither
}

var cache_Data_Traversable_traversableDual gopurs_runtime.Value
var once_Data_Traversable_traversableDual sync.Once
func Get_Data_Traversable_traversableDual() gopurs_runtime.Value {
	once_Data_Traversable_traversableDual.Do(func() {
		cache_Data_Traversable_traversableDual = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Dual_functorDual()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Dual_Dual(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Dual_Dual(), gopurs_runtime.Apply(f_2, v_3))
})
})}))}
	})
	return cache_Data_Traversable_traversableDual
}

var cache_Data_Traversable_traversableDisj gopurs_runtime.Value
var once_Data_Traversable_traversableDisj sync.Once
func Get_Data_Traversable_traversableDisj() gopurs_runtime.Value {
	once_Data_Traversable_traversableDisj.Do(func() {
		cache_Data_Traversable_traversableDisj = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDisj()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Disj_functorDisj()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Disj_Disj(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Disj_Disj(), gopurs_runtime.Apply(f_2, v_3))
})
})}))}
	})
	return cache_Data_Traversable_traversableDisj
}

var cache_Data_Traversable_traversableConst gopurs_runtime.Value
var once_Data_Traversable_traversableConst sync.Once
func Get_Data_Traversable_traversableConst() gopurs_runtime.Value {
	once_Data_Traversable_traversableConst.Do(func() {
		cache_Data_Traversable_traversableConst = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConst()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Const_functorConst()))}
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v_1)
}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), v1_2)
})}))}
	})
	return cache_Data_Traversable_traversableConst
}

var cache_Data_Traversable_traversableConj gopurs_runtime.Value
var once_Data_Traversable_traversableConj sync.Once
func Get_Data_Traversable_traversableConj() gopurs_runtime.Value {
	once_Data_Traversable_traversableConj.Do(func() {
		cache_Data_Traversable_traversableConj = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableConj()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Conj_functorConj()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Conj_Conj(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Conj_Conj(), gopurs_runtime.Apply(f_2, v_3))
})
})}))}
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
		cache_Data_Traversable_traversableAdditive = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableAdditive()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Monoid_Additive_functorAdditive()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Monoid_Additive_Additive(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Monoid_Additive_Additive(), gopurs_runtime.Apply(f_2, v_3))
})
})}))}
	})
	return cache_Data_Traversable_traversableAdditive
}

var cache_Data_Traversable_sequenceDefault gopurs_runtime.Value
var once_Data_Traversable_sequenceDefault sync.Once
func Get_Data_Traversable_sequenceDefault() gopurs_runtime.Value {
	once_Data_Traversable_sequenceDefault.Do(func() {
		cache_Data_Traversable_sequenceDefault = gopurs_runtime.Func2(func(dictTraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_sequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_Data_Traversable_sequenceDefault
}

var cache_Data_Traversable_traversableArray gopurs_runtime.Value
var once_Data_Traversable_traversableArray sync.Once
func Get_Data_Traversable_traversableArray() gopurs_runtime.Value {
	once_Data_Traversable_traversableArray.Do(func() {
		cache_Data_Traversable_traversableArray = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=Any
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply4(Get_Data_Traversable_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_0, "pure"), Get_Data_Semigroup_concatArray())
})}))}
	})
	return cache_Data_Traversable_traversableArray
}

var cache_Data_Traversable_sequence gopurs_runtime.Value
var once_Data_Traversable_sequence sync.Once
func Get_Data_Traversable_sequence() gopurs_runtime.Value {
	once_Data_Traversable_sequence.Do(func() {
		cache_Data_Traversable_sequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_sequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dict_0_box))
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
		cache_Data_Traversable_traversableFirst = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3188237647_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_1146820559_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableFirst())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3689823567_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_functorMaybe())))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_2 == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_3 != nil) {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), Get_Data_Maybe_Just(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_First_First(), __t4)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_5
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_6 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_6
var __t9 gopurs_runtime.Value
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_7 == nil) {
__t9 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
goto end_branch_9
} else {

}
}
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_8 != nil) {
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_6.V0), Get_Data_Maybe_Just(), gopurs_runtime.Apply(f_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_5.V0), Get_Data_Maybe_First_First(), __t9)
})
})})))}
	})
	return cache_Data_Traversable_traversableFirst
}

var cache_Data_Traversable_traversableLast gopurs_runtime.Value
var once_Data_Traversable_traversableLast sync.Once
func Get_Data_Traversable_traversableLast() gopurs_runtime.Value {
	once_Data_Traversable_traversableLast.Do(func() {
		cache_Data_Traversable_traversableLast = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3188237647_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_1146820559_1680800814(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Foldable_foldableLast())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3689823567_2812149806(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]](Get_Data_Maybe_functorMaybe())))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_1
var __t4 gopurs_runtime.Value
{
var __t_tag_2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_2 == nil) {
__t4 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_2)
if (__t_tag_3 != nil) {
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), Get_Data_Maybe_Just(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_2.UnsafePtr).V0)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Maybe_Last_Last(), __t4)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_5
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_6 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_6
var __t9 gopurs_runtime.Value
{
var __t_tag_7 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_7 == nil) {
__t9 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))})
goto end_branch_9
} else {

}
}
{
var __t_tag_8 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](v_3)
if (__t_tag_8 != nil) {
__t9 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_6.V0), Get_Data_Maybe_Just(), gopurs_runtime.Apply(f_2, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0))
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_5.V0), Get_Data_Maybe_Last_Last(), __t9)
})
})})))}
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
return Call_Data_Traversable_traverseDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box))
})
	})
	return cache_Data_Traversable_traverseDefault
}

var cache_Data_Traversable_mapAccumR gopurs_runtime.Value
var once_Data_Traversable_mapAccumR sync.Once
func Get_Data_Traversable_mapAccumR() gopurs_runtime.Value {
	once_Data_Traversable_mapAccumR.Do(func() {
		cache_Data_Traversable_mapAccumR = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_Traversable_mapAccumR(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
				_ = orig
				return gopurs_runtime.RecordDict2("accum", "value", orig.accum, orig.value)
				}()
})
	})
	return cache_Data_Traversable_mapAccumR
}

var cache_Data_Traversable_scanr gopurs_runtime.Value
var once_Data_Traversable_scanr sync.Once
func Get_Data_Traversable_scanr() gopurs_runtime.Value {
	once_Data_Traversable_scanr.Do(func() {
		cache_Data_Traversable_scanr = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_scanr(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_Data_Traversable_scanr
}

var cache_Data_Traversable_mapAccumL gopurs_runtime.Value
var once_Data_Traversable_mapAccumL sync.Once
func Get_Data_Traversable_mapAccumL() gopurs_runtime.Value {
	once_Data_Traversable_mapAccumL.Do(func() {
		cache_Data_Traversable_mapAccumL = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, s0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_Traversable_mapAccumL(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, s0_2_box, xs_3_box)
				_ = orig
				return gopurs_runtime.RecordDict2("accum", "value", orig.accum, orig.value)
				}()
})
	})
	return cache_Data_Traversable_mapAccumL
}

var cache_Data_Traversable_scanl gopurs_runtime.Value
var once_Data_Traversable_scanl sync.Once
func Get_Data_Traversable_scanl() gopurs_runtime.Value {
	once_Data_Traversable_scanl.Do(func() {
		cache_Data_Traversable_scanl = gopurs_runtime.Func4(func(dictTraversable_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, b0_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_scanl(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_0_box), f_1_box, b0_2_box, xs_3_box)
})
	})
	return cache_Data_Traversable_scanl
}

var cache_Data_Traversable_go__for gopurs_runtime.Value
var once_Data_Traversable_go__for sync.Once
func Get_Data_Traversable_go__for() gopurs_runtime.Value {
	once_Data_Traversable_go__for.Do(func() {
		cache_Data_Traversable_go__for = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, dictTraversable_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Traversable_go__for(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](dictTraversable_1_box), x_2_box, f_3_box)
})
	})
	return cache_Data_Traversable_go__for
}

type Constructor_Data_Traversable_Traversable[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3941073978] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Traversable_Traversable[any])(ptr)
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


func Call_Data_Traversable_Traversable_dollar_Dict(x_0_loop struct{
	Foldable1 gopurs_runtime.Value
	Functor0 gopurs_runtime.Value
	sequence gopurs_runtime.Value
	traverse gopurs_runtime.Value
}) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
var x_0 struct{
	Foldable1 gopurs_runtime.Value
	Functor0 gopurs_runtime.Value
	sequence gopurs_runtime.Value
	traverse gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", orig.Foldable1, orig.Functor0, orig.sequence, orig.traverse)
				}())
}

func Call_Data_Traversable_traverse(dict_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Traversable_traversableCompose(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
traversableCompose:
for {
if false { continue traversableCompose }
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorCompose__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
functorCompose__193435443_1_0 := gopurs_runtime.Func(func(dictFunctor1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_2, "map"), f_3), v_4)
})}))}
})
_ = functorCompose__193435443_1_0
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): foldableCompose__193435443_2_2 shape=Let(Abs(LitRecord)) bindingType=Any
foldableCompose__193435443_2_2 := gopurs_runtime.Func(func(dictFoldable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_4 gopurs_runtime.Value, f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_4))}, f_5), v_6)
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, i_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldl"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_3, "foldl"), f_4), i_5, v_6)
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, i_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_4 shape=App(Other) bindingType=Any
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_3, "foldr"), f_4)
_ = __local_var_7_4
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldr"), gopurs_runtime.Func2(func(b_8 gopurs_runtime.Value, a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_7_4, a_9, b_8)
}), i_5, v_6)
})}))}
})
_ = foldableCompose__193435443_2_2
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCompose1_4_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCompose1_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCompose__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{})))
_ = functorCompose1_4_5
// TAST (Let): foldableCompose1_5_6 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (TypeApp (TypeVar f) [(TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
foldableCompose1_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableCompose__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{})))
_ = foldableCompose1_5_6
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableCompose1_5_6)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCompose1_4_5)}
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Call_Data_Traversable_traversableCompose(dictTraversable_0), dictTraversable1_3), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_7 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_7
return gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_7.V0), Get_Data_Functor_Compose_Compose(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, f_8), v_9))
})
})}))}
})
}
}

func Call_Data_Traversable_sequenceDefault(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Data_Traversable_sequence(dict_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Traversable_traversableApp(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): functorApp_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f)])])
functorApp_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{}))
_ = functorApp_1_0
// TAST (Let): __local_var_2_2 shape=App(Other) bindingType=Any
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): foldableApp_2_1 shape=Let(LitRecord) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f)])])
foldableApp_2_1 := (&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_2, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_3))}, f_4, v_5)
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, i_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_2, "foldl"), f_3, i_4, v_5)
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, i_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_2, "foldr"), f_3, i_4, v_5)
})})
_ = foldableApp_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableApp_2_1)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorApp_1_0)}
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), Get_Data_Functor_App_App(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), Get_Data_Functor_App_App(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_5, v_6))
})
})}))}
}

func Call_Data_Traversable_traversableCoproduct(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorCoproduct__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
functorCoproduct__193435443_1_0 := gopurs_runtime.Func(func(dictFunctor1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_2 shape=App(Other) bindingType=Any
__local_var_5_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_3)
_ = __local_var_5_2
// TAST (Let): __local_var_6_3 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar g) [(TypeVar b)]))
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor1_2, "map"), f_3)
_ = __local_var_6_3
var __t4 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 3711209382) {
__t4 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Apply(__local_var_5_2, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_4
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2465973597) {
__t4 = func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, gopurs_runtime.Apply(__local_var_6_3, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
})}))}
})
_ = functorCoproduct__193435443_1_0
// TAST (Let): foldableCoproduct__193435443_2_5 shape=App(Var) bindingType=Any
foldableCoproduct__193435443_2_5 := gopurs_runtime.Apply(Get_Data_Foldable_foldableCoproduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableCoproduct__193435443_2_5
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorCoproduct1_4_6 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorCoproduct1_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorCoproduct__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{})))
_ = functorCoproduct1_4_6
// TAST (Let): foldableCoproduct1_5_7 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
foldableCoproduct1_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableCoproduct__193435443_2_5, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{})))
_ = foldableCoproduct1_5_7
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableCoproduct1_5_7)}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorCoproduct1_4_6)}
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_8 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_8
// TAST (Let): __local_var_8_9 shape=App(Other) bindingType=Any
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_7_8.V0), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_8, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
}))
_ = __local_var_8_9
// TAST (Let): __local_var_9_11 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeApp (TypeVar m) [(TypeVar a)])])] (TypeApp (TypeVar m) [(TypeApp (TypeVar f) [(TypeVar a)])]))
__local_var_9_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_6)
_ = __local_var_9_11
// TAST (Let): __local_var_9_10 shape=Let(Abs(App(Other))) bindingType=(TypeApp (TypeVar m) [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar a)])])
__local_var_9_10 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_9, gopurs_runtime.Apply(__local_var_9_11, x_10))
})
_ = __local_var_9_10
// TAST (Let): __local_var_10_13 shape=App(Other) bindingType=Any
__local_var_10_13 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_7_8.V0), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_10, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
}))
_ = __local_var_10_13
// TAST (Let): __local_var_11_14 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeApp (TypeVar m) [(TypeVar a)])])] (TypeApp (TypeVar m) [(TypeApp (TypeVar g) [(TypeVar a)])]))
__local_var_11_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "sequence"), dictApplicative_6)
_ = __local_var_11_14
// TAST (Let): __local_var_10_12 shape=Let(Let(Abs(App(Other)))) bindingType=(Func [(TypeApp (TypeVar g) [(TypeApp (TypeVar m) [(TypeVar a)])])] (TypeApp (TypeVar m) [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar a)])]))
__local_var_10_12 := gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_13, gopurs_runtime.Apply(__local_var_11_14, x_12))
})
_ = __local_var_10_12
return gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(__local_var_9_10, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0)
goto end_branch_15
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t15 = gopurs_runtime.Apply(__local_var_10_12, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_11.UnsafePtr).V0)
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
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_16 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_7_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_16
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_17 shape=App(Other) bindingType=Any
__local_var_9_17 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_7_16.V0), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{x_9, gopurs_runtime.Value{}, false}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
}))
_ = __local_var_9_17
// TAST (Let): __local_var_10_19 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(TypeVar a)])] (TypeApp (TypeVar m) [(TypeApp (TypeVar f) [(TypeVar b)])]))
__local_var_10_19 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, f_8)
_ = __local_var_10_19
// TAST (Let): __local_var_10_18 shape=Let(Abs(App(Other))) bindingType=(TypeApp (TypeVar m) [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar b)])])
__local_var_10_18 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_17, gopurs_runtime.Apply(__local_var_10_19, x_11))
})
_ = __local_var_10_18
// TAST (Let): __local_var_11_21 shape=App(Other) bindingType=Any
__local_var_11_21 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_7_16.V0), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{gopurs_runtime.Value{}, x_11, true}
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
}))
_ = __local_var_11_21
// TAST (Let): __local_var_12_22 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar m) [(TypeApp (TypeVar g) [(TypeVar b)])]))
__local_var_12_22 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, f_8)
_ = __local_var_12_22
// TAST (Let): __local_var_11_20 shape=Let(Let(Abs(App(Other)))) bindingType=(Func [(TypeApp (TypeVar g) [(TypeVar a)])] (TypeApp (TypeVar m) [(TypeApp (ADT ["Data","Either","Either"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g), (TypeVar b)])]))
__local_var_11_20 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_21, gopurs_runtime.Apply(__local_var_12_22, x_13))
})
_ = __local_var_11_20
return gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t23 = gopurs_runtime.Apply(__local_var_10_18, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0)
goto end_branch_23
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t23 = gopurs_runtime.Apply(__local_var_11_20, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_12.UnsafePtr).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
})
})
})}))}
})
}

func Call_Data_Traversable_traversableProduct(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorProduct__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
functorProduct__193435443_1_0 := gopurs_runtime.Func(func(dictFunctor1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_2363162019_2812149806((&Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor1_2, "map"), f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
})})))}
})
_ = functorProduct__193435443_1_0
// TAST (Let): foldableProduct__193435443_2_2 shape=App(Var) bindingType=Any
foldableProduct__193435443_2_2 := gopurs_runtime.Apply(Get_Data_Foldable_foldableProduct(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableProduct__193435443_2_2
return gopurs_runtime.Func(func(dictTraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): functorProduct1_4_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
functorProduct1_4_3 := Rebox_Data_Traversable_2812149806_2363162019(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(functorProduct__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Functor0"), gopurs_runtime.Value{}))))
_ = functorProduct1_4_3
// TAST (Let): foldableProduct1_5_4 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar f) [(TypeVar a)]), (TypeApp (TypeVar g) [(TypeVar a)])]) [(TypeVar f), (TypeVar g)])])
foldableProduct1_5_4 := Rebox_Data_Traversable_1680800814_4173511203(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(foldableProduct__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable1_3, "Foldable1"), gopurs_runtime.Value{}))))
_ = foldableProduct1_5_4
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_3543431075_3043886126((&Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_4173511203_1680800814(foldableProduct1_5_4))}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(Rebox_Data_Traversable_2363162019_2812149806(functorProduct1_4_3))}
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_5 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_5
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_5.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_7_5.V0), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Product_product(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable1_3, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_6 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar m)])
Apply0_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_6
return gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_6.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Apply0_7_6.V0), gopurs_runtime.Value{}), "map"), Get_Data_Functor_Product_product(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable1_3, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, f_8, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1))
})
})})))}
})
}

func Call_Data_Traversable_traverseDefault(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar t)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictTraversable_0.V1), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func3(func(dictApplicative_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, ta_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictTraversable_0.V2), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_2))}, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_3, ta_4))
})
}

func Call_Data_Traversable_mapAccumR(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return func() struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Traversable_Accum_Internal_applicativeStateR()))}, gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := func() struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply2(f_1, s_5, a_4)
					_ = orig
					clone := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.accum = gopurs_runtime.RecordGet(orig, "accum")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict2("accum", "value", orig.accum, orig.value)
				}()
}), xs_3, s0_2)
					_ = orig
					clone := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.accum = gopurs_runtime.RecordGet(orig, "accum")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
}

func Call_Data_Traversable_scanr(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return Call_Data_Traversable_mapAccumR(dictTraversable_0, gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): b_prime__6_0 shape=App(Other) bindingType=(TypeVar b)
b_prime__6_0 := gopurs_runtime.Apply2(f_1, a_5, b_4)
_ = b_prime__6_0
return func() gopurs_runtime.Value {
				orig := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{b_prime__6_0, b_prime__6_0}
				_ = orig
				return gopurs_runtime.RecordDict2("accum", "value", orig.accum, orig.value)
				}()
}), b0_2, xs_3).value
}

func Call_Data_Traversable_mapAccumL(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, s0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var s0_2 gopurs_runtime.Value = s0_2_loop
_ = s0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return func() struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply4(gopurs_runtime.Box(dictTraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Traversable_Accum_Internal_applicativeStateL()))}, gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := func() struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
} {
					orig := gopurs_runtime.Apply2(f_1, s_5, a_4)
					_ = orig
					clone := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.accum = gopurs_runtime.RecordGet(orig, "accum")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
				_ = orig
				return gopurs_runtime.RecordDict2("accum", "value", orig.accum, orig.value)
				}()
}), xs_3, s0_2)
					_ = orig
					clone := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{}
					clone.accum = gopurs_runtime.RecordGet(orig, "accum")
					clone.value = gopurs_runtime.RecordGet(orig, "value")
					return clone
				}()
}

func Call_Data_Traversable_scanl(dictTraversable_0_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, b0_2_loop gopurs_runtime.Value, xs_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_0_loop
_ = dictTraversable_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var b0_2 gopurs_runtime.Value = b0_2_loop
_ = b0_2
var xs_3 gopurs_runtime.Value = xs_3_loop
_ = xs_3
return Call_Data_Traversable_mapAccumL(dictTraversable_0, gopurs_runtime.Func2(func(b_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): b_prime__6_0 shape=App(Other) bindingType=(TypeVar b)
b_prime__6_0 := gopurs_runtime.Apply2(f_1, b_4, a_5)
_ = b_prime__6_0
return func() gopurs_runtime.Value {
				orig := struct{
	accum gopurs_runtime.Value
	value gopurs_runtime.Value
}{b_prime__6_0, b_prime__6_0}
				_ = orig
				return gopurs_runtime.RecordDict2("accum", "value", orig.accum, orig.value)
				}()
}), b0_2, xs_3).value
}

func Call_Data_Traversable_go__for(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], dictTraversable_1_loop *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var dictTraversable_1 *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] = dictTraversable_1_loop
_ = dictTraversable_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictTraversable_1.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, f_3, x_2)
}

func Rebox_Data_Traversable_1146820559_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Traversable_1680800814_4173511203(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Traversable_2363162019_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Traversable_2812149806_2363162019(in *Constructor_Data_Functor_Functor[gopurs_runtime.Value]) *Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Traversable_3188237647_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Traversable_3543431075_3043886126(in *Constructor_Data_Traversable_Traversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Traversable_Traversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Traversable_3689823567_2812149806(in *Constructor_Data_Functor_Functor[*Constructor_Data_Maybe_Just[gopurs_runtime.Value]]) *Constructor_Data_Functor_Functor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Functor_Functor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Traversable_4173511203_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Get_Data_Traversable_traverseArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Traversable_TraverseArrayImpl
}

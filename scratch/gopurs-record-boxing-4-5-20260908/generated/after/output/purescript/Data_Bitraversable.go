package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Bitraversable_identity gopurs_runtime.Value
var once_Data_Bitraversable_identity sync.Once
func Get_Data_Bitraversable_identity() gopurs_runtime.Value {
	once_Data_Bitraversable_identity.Do(func() {
		cache_Data_Bitraversable_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Bitraversable_identity
}

var cache_Data_Bitraversable_identity1 gopurs_runtime.Value
var once_Data_Bitraversable_identity1 sync.Once
func Get_Data_Bitraversable_identity1() gopurs_runtime.Value {
	once_Data_Bitraversable_identity1.Do(func() {
		cache_Data_Bitraversable_identity1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Bitraversable_identity1
}

var cache_Data_Bitraversable_Bitraversable_dollar_Dict gopurs_runtime.Value
var once_Data_Bitraversable_Bitraversable_dollar_Dict sync.Once
func Get_Data_Bitraversable_Bitraversable_dollar_Dict() gopurs_runtime.Value {
	once_Data_Bitraversable_Bitraversable_dollar_Dict.Do(func() {
		cache_Data_Bitraversable_Bitraversable_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer(Call_Data_Bitraversable_Bitraversable_dollar_Dict(func() struct{
	Bifoldable1 gopurs_runtime.Value
	Bifunctor0 gopurs_runtime.Value
	bisequence gopurs_runtime.Value
	bitraverse gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Bifoldable1 gopurs_runtime.Value
	Bifunctor0 gopurs_runtime.Value
	bisequence gopurs_runtime.Value
	bitraverse gopurs_runtime.Value
}{}
					clone.Bifoldable1 = gopurs_runtime.RecordGet(orig, "Bifoldable1")
					clone.Bifunctor0 = gopurs_runtime.RecordGet(orig, "Bifunctor0")
					clone.bisequence = gopurs_runtime.RecordGet(orig, "bisequence")
					clone.bitraverse = gopurs_runtime.RecordGet(orig, "bitraverse")
					return clone
				}()))}
})
	})
	return cache_Data_Bitraversable_Bitraversable_dollar_Dict
}

var cache_Data_Bitraversable_bitraverse gopurs_runtime.Value
var once_Data_Bitraversable_bitraverse sync.Once
func Get_Data_Bitraversable_bitraverse() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraverse.Do(func() {
		cache_Data_Bitraversable_bitraverse = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bitraversable_bitraverse
}

var cache_Data_Bitraversable_lfor gopurs_runtime.Value
var once_Data_Bitraversable_lfor sync.Once
func Get_Data_Bitraversable_lfor() gopurs_runtime.Value {
	once_Data_Bitraversable_lfor.Do(func() {
		cache_Data_Bitraversable_lfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_lfor(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_Data_Bitraversable_lfor
}

var cache_Data_Bitraversable_ltraverse gopurs_runtime.Value
var once_Data_Bitraversable_ltraverse sync.Once
func Get_Data_Bitraversable_ltraverse() gopurs_runtime.Value {
	once_Data_Bitraversable_ltraverse.Do(func() {
		cache_Data_Bitraversable_ltraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_ltraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_Data_Bitraversable_ltraverse
}

var cache_Data_Bitraversable_rfor gopurs_runtime.Value
var once_Data_Bitraversable_rfor sync.Once
func Get_Data_Bitraversable_rfor() gopurs_runtime.Value {
	once_Data_Bitraversable_rfor.Do(func() {
		cache_Data_Bitraversable_rfor = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_rfor(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_Data_Bitraversable_rfor
}

var cache_Data_Bitraversable_rtraverse gopurs_runtime.Value
var once_Data_Bitraversable_rtraverse sync.Once
func Get_Data_Bitraversable_rtraverse() gopurs_runtime.Value {
	once_Data_Bitraversable_rtraverse.Do(func() {
		cache_Data_Bitraversable_rtraverse = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_rtraverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_Data_Bitraversable_rtraverse
}

var cache_Data_Bitraversable_bitraversableTuple gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableTuple sync.Once
func Get_Data_Bitraversable_bitraversableTuple() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableTuple.Do(func() {
		cache_Data_Bitraversable_bitraversableTuple = gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_2243670499_3561684974((&Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_2812820739_3566843086(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Bifoldable_bifoldableTuple())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_1495429347_1688994542(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Data_Bifunctor_bifunctorTuple())))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_1_2 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_2
// TAST (Let): Functor0_2_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_3
return gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_3.V0), Get_Data_Tuple_Tuple(), gopurs_runtime.Apply(f_3, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)), gopurs_runtime.Apply(g_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
})
})})))}
	})
	return cache_Data_Bitraversable_bitraversableTuple
}

var cache_Data_Bitraversable_bitraversableJoker gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableJoker sync.Once
func Get_Data_Bitraversable_bitraversableJoker() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableJoker.Do(func() {
		cache_Data_Bitraversable_bitraversableJoker = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraversableJoker(dictTraversable_0_box)
})
	})
	return cache_Data_Bitraversable_bitraversableJoker
}

var cache_Data_Bitraversable_bitraversableEither gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableEither sync.Once
func Get_Data_Bitraversable_bitraversableEither() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableEither.Do(func() {
		cache_Data_Bitraversable_bitraversableEither = gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableEither()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Bifunctor_bifunctorEither()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Either_Left(), (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0)
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
// TAST (Let): Functor0_1_2 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_2
return gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, v2_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (v2_4.Type == 9 && v2_4.IntVal == 3711209382) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_2.V0), Get_Data_Either_Left(), gopurs_runtime.Apply(v_2, (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0))
goto end_branch_3
} else {

}
}
{
if (v2_4.Type == 9 && v2_4.IntVal == 2465973597) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_2.V0), Get_Data_Either_Right(), gopurs_runtime.Apply(v1_3, (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v2_4.UnsafePtr).V0))
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
	return cache_Data_Bitraversable_bitraversableEither
}

var cache_Data_Bitraversable_bitraversableConst gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableConst sync.Once
func Get_Data_Bitraversable_bitraversableConst() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableConst.Do(func() {
		cache_Data_Bitraversable_bitraversableConst = gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](Get_Data_Bifoldable_bifoldableConst()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](Get_Data_Bifunctor_bifunctorConst()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), Get_Data_Const_Const(), v_2)
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), Get_Data_Const_Const(), gopurs_runtime.Apply(f_2, v1_4))
})
})}))}
	})
	return cache_Data_Bitraversable_bitraversableConst
}

var cache_Data_Bitraversable_bitraversableClown gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableClown sync.Once
func Get_Data_Bitraversable_bitraversableClown() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableClown.Do(func() {
		cache_Data_Bitraversable_bitraversableClown = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraversableClown(dictTraversable_0_box)
})
	})
	return cache_Data_Bitraversable_bitraversableClown
}

var cache_Data_Bitraversable_bisequenceDefault gopurs_runtime.Value
var once_Data_Bitraversable_bisequenceDefault sync.Once
func Get_Data_Bitraversable_bisequenceDefault() gopurs_runtime.Value {
	once_Data_Bitraversable_bisequenceDefault.Do(func() {
		cache_Data_Bitraversable_bisequenceDefault = gopurs_runtime.Func2(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bisequenceDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box))
})
	})
	return cache_Data_Bitraversable_bisequenceDefault
}

var cache_Data_Bitraversable_bisequence gopurs_runtime.Value
var once_Data_Bitraversable_bisequence sync.Once
func Get_Data_Bitraversable_bisequence() gopurs_runtime.Value {
	once_Data_Bitraversable_bisequence.Do(func() {
		cache_Data_Bitraversable_bisequence = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bisequence(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bitraversable_bisequence
}

var cache_Data_Bitraversable_bitraversableFlip gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableFlip sync.Once
func Get_Data_Bitraversable_bitraversableFlip() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableFlip.Do(func() {
		cache_Data_Bitraversable_bitraversableFlip = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraversableFlip(dictBitraversable_0_box)
})
	})
	return cache_Data_Bitraversable_bitraversableFlip
}

var cache_Data_Bitraversable_bitraversableProduct2 gopurs_runtime.Value
var once_Data_Bitraversable_bitraversableProduct2 sync.Once
func Get_Data_Bitraversable_bitraversableProduct2() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraversableProduct2.Do(func() {
		cache_Data_Bitraversable_bitraversableProduct2 = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraversableProduct2(dictBitraversable_0_box)
})
	})
	return cache_Data_Bitraversable_bitraversableProduct2
}

var cache_Data_Bitraversable_bitraverseDefault gopurs_runtime.Value
var once_Data_Bitraversable_bitraverseDefault sync.Once
func Get_Data_Bitraversable_bitraverseDefault() gopurs_runtime.Value {
	once_Data_Bitraversable_bitraverseDefault.Do(func() {
		cache_Data_Bitraversable_bitraverseDefault = gopurs_runtime.Func(func(dictBitraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bitraverseDefault(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box))
})
	})
	return cache_Data_Bitraversable_bitraverseDefault
}

var cache_Data_Bitraversable_bifor gopurs_runtime.Value
var once_Data_Bitraversable_bifor sync.Once
func Get_Data_Bitraversable_bifor() gopurs_runtime.Value {
	once_Data_Bitraversable_bifor.Do(func() {
		cache_Data_Bitraversable_bifor = gopurs_runtime.Func5(func(dictBitraversable_0_box gopurs_runtime.Value, dictApplicative_1_box gopurs_runtime.Value, t_2_box gopurs_runtime.Value, f_3_box gopurs_runtime.Value, g_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bitraversable_bifor(gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](dictBitraversable_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_1_box), t_2_box, f_3_box, g_4_box)
})
	})
	return cache_Data_Bitraversable_bifor
}

type Constructor_Data_Bitraversable_Bitraversable[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3704227322] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bitraversable_Bitraversable[any])(ptr)
		_ = c
		switch key {
		case "Bifoldable1": return gopurs_runtime.Box(c.V0)
		case "Bifunctor0": return gopurs_runtime.Box(c.V1)
		case "bisequence": return gopurs_runtime.Box(c.V2)
		case "bitraverse": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_Bitraversable_Bitraversable: " + key)
		}
	}
}


func Call_Data_Bitraversable_Bitraversable_dollar_Dict(x_0_loop struct{
	Bifoldable1 gopurs_runtime.Value
	Bifunctor0 gopurs_runtime.Value
	bisequence gopurs_runtime.Value
	bitraverse gopurs_runtime.Value
}) *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] {
var x_0 struct{
	Bifoldable1 gopurs_runtime.Value
	Bifunctor0 gopurs_runtime.Value
	bisequence gopurs_runtime.Value
	bitraverse gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict4("Bifoldable1", "Bifunctor0", "bisequence", "bitraverse", orig.Bifoldable1, orig.Bifunctor0, orig.bisequence, orig.bitraverse)
				}())
}

func Call_Data_Bitraversable_bitraverse(dict_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Bitraversable_lfor(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
// TAST (Let): pure_2_0 shape=Other bindingType=(Func [(TypeVar b)] (TypeApp (TypeVar f) [(TypeVar b)]))
pure_2_0 := gopurs_runtime.Box(dictApplicative_1.V1)
_ = pure_2_0
return gopurs_runtime.Func2(func(t_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, f_4, pure_2_0, t_3)
})
}

func Call_Data_Bitraversable_ltraverse(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
// TAST (Let): pure_2_0 shape=Other bindingType=(Func [(TypeVar b)] (TypeApp (TypeVar f) [(TypeVar b)]))
pure_2_0 := gopurs_runtime.Box(dictApplicative_1.V1)
_ = pure_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, f_3, pure_2_0)
})
}

func Call_Data_Bitraversable_rfor(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
// TAST (Let): pure_2_0 shape=Other bindingType=(Func [(TypeVar a)] (TypeApp (TypeVar f) [(TypeVar a)]))
pure_2_0 := gopurs_runtime.Box(dictApplicative_1.V1)
_ = pure_2_0
return gopurs_runtime.Func2(func(t_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, pure_2_0, f_4, t_3)
})
}

func Call_Data_Bitraversable_rtraverse(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Box(dictApplicative_1.V1))
}

func Call_Data_Bitraversable_bitraversableJoker(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorJoker_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeApp (TypeVar f) [(TypeVar b)])])
bifunctorJoker_1_0 := (&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), g_3, v1_4)
})})
_ = bifunctorJoker_1_0
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifoldableJoker_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Bifoldable","Bifoldable"] [(TypeApp (TypeVar f) [(TypeVar b)])])
bifoldableJoker_2_2 := (&Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(dictMonoid_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_3))}, r_5, v1_6)
}), gopurs_runtime.Func4(func(v_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, u_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldl"), r_4, u_5, v1_6)
}), gopurs_runtime.Func4(func(v_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, u_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldr"), r_4, u_5, v1_6)
})})
_ = bifoldableJoker_2_2
return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(bifoldableJoker_2_2)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorJoker_1_0)}
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), Get_Data_Functor_Joker_Joker(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), Get_Data_Functor_Joker_Joker(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, r_6, v1_7))
})
})}))}
}

func Call_Data_Bitraversable_bitraversableClown(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorClown_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeApp (TypeVar f) [(TypeVar a)])])
bifunctorClown_1_0 := (&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, v1_4)
})})
_ = bifunctorClown_1_0
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifoldableClown_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Bifoldable","Bifoldable"] [(TypeApp (TypeVar f) [(TypeVar a)])])
bifoldableClown_2_2 := (&Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(dictMonoid_3 gopurs_runtime.Value, l_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_3))}, l_4, v1_6)
}), gopurs_runtime.Func4(func(l_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, u_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldl"), l_3, u_5, v1_6)
}), gopurs_runtime.Func4(func(l_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, u_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldr"), l_3, u_5, v1_6)
})})
_ = bifoldableClown_2_2
return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(bifoldableClown_2_2)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorClown_1_0)}
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), Get_Data_Functor_Clown_Clown(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func3(func(l_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), Get_Data_Functor_Clown_Clown(), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, l_5, v1_7))
})
})}))}
}

func Call_Data_Bitraversable_bisequenceDefault(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
return gopurs_runtime.Apply3(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Data_Bitraversable_bisequence(dict_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Bitraversable_bitraversableFlip(dictBitraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorFlip_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeApp (TypeApp (TypeVar p) [(TypeVar b), (TypeVar a)]) [(TypeVar p)])])
bifunctorFlip_1_0 := (&Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), g_3, f_2, v_4)
})})
_ = bifunctorFlip_1_0
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifoldableFlip_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Bifoldable","Bifoldable"] [(TypeApp (TypeApp (TypeVar p) [(TypeVar b), (TypeVar a)]) [(TypeVar p)])])
bifoldableFlip_2_2 := (&Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]{1, gopurs_runtime.Func4(func(dictMonoid_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, l_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_3, "bifoldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_3))}, l_5, r_4, v_6)
}), gopurs_runtime.Func4(func(r_3 gopurs_runtime.Value, l_4 gopurs_runtime.Value, u_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_3, "bifoldl"), l_4, r_3, u_5, v_6)
}), gopurs_runtime.Func4(func(r_3 gopurs_runtime.Value, l_4 gopurs_runtime.Value, u_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.RecordGet(__local_var_2_3, "bifoldr"), l_4, r_3, u_5, v_6)
})})
_ = bifoldableFlip_2_2
return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(bifoldableFlip_2_2)}
}), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(bifunctorFlip_1_0)}
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), Get_Data_Functor_Flip_Flip(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, v_5))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func3(func(r_5 gopurs_runtime.Value, l_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), Get_Data_Functor_Flip_Flip(), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, l_6, r_5, v_7))
})
})}))}
}

func Call_Data_Bitraversable_bitraversableProduct2(dictBitraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 gopurs_runtime.Value = dictBitraversable_0_loop
_ = dictBitraversable_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifunctor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorProduct2__193435443_1_0 shape=Let(Abs(LitRecord)) bindingType=Any
bifunctorProduct2__193435443_1_0 := gopurs_runtime.Func(func(dictBifunctor1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_2864580459_1688994542((&Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3559137202, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "bimap"), f_3, g_4, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictBifunctor1_2, "bimap"), f_3, g_4, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)}))}
})})))}
})
_ = bifunctorProduct2__193435443_1_0
// TAST (Let): bifoldableProduct2__193435443_2_2 shape=App(Var) bindingType=Any
bifoldableProduct2__193435443_2_2 := gopurs_runtime.Apply(Get_Data_Bifoldable_bifoldableProduct2(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable_0, "Bifoldable1"), gopurs_runtime.Value{}))
_ = bifoldableProduct2__193435443_2_2
return gopurs_runtime.Func(func(dictBitraversable1_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): bifunctorProduct21_4_3 shape=App(Other) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(ADT ["Data","Functor","Product2","Product2"] [(TypeVar f), (TypeVar g)])])
bifunctorProduct21_4_3 := Rebox_Data_Bitraversable_1688994542_2864580459(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(bifunctorProduct2__193435443_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifunctor0"), gopurs_runtime.Value{}))))
_ = bifunctorProduct21_4_3
// TAST (Let): bifoldableProduct21_5_4 shape=App(Other) bindingType=(ADT ["Data","Bifoldable","Bifoldable"] [(ADT ["Data","Functor","Product2","Product2"] [(TypeVar f), (TypeVar g)])])
bifoldableProduct21_5_4 := Rebox_Data_Bitraversable_3566843086_4280144779(gopurs_runtime.CoerceToStruct[Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]](gopurs_runtime.Apply(bifoldableProduct2__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBitraversable1_3, "Bifoldable1"), gopurs_runtime.Value{}))))
_ = bifoldableProduct21_5_4
return gopurs_runtime.Value{Type: 9, IntVal: 3704227322, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_4211574891_3561684974((&Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4001671834, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_4280144779_3566843086(bifoldableProduct21_5_4))}
}), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4141114362, UnsafePtr: unsafe.Pointer(Rebox_Data_Bitraversable_2864580459_1688994542(bifunctorProduct21_4_3))}
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_5 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_5
// TAST (Let): Functor0_8_6 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_8_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_6
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_5.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_6.V0), Get_Data_Functor_Product2_Product2(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable_0, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBitraversable1_3, "bisequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Apply0_7_7 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_7_7
// TAST (Let): Functor0_8_8 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_8
return gopurs_runtime.Func3(func(l_9 gopurs_runtime.Value, r_10 gopurs_runtime.Value, v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_7_7.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_8.V0), Get_Data_Functor_Product2_Product2(), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable_0, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, l_9, r_10, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V0)), gopurs_runtime.Apply4(gopurs_runtime.RecordGet(dictBitraversable1_3, "bitraverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_6))}, l_9, r_10, (*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_11.UnsafePtr).V1))
})
})})))}
})
}

func Call_Data_Bitraversable_bitraverseDefault(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
// TAST (Let): Bifunctor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Bifunctor","Bifunctor"] [(TypeVar t)])
Bifunctor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictBitraversable_0.V1), gopurs_runtime.Value{}))
_ = Bifunctor0_1_0
return gopurs_runtime.Func4(func(dictApplicative_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictBitraversable_0.V2), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_2))}, gopurs_runtime.Apply3(gopurs_runtime.Box(Bifunctor0_1_0.V0), f_3, g_4, t_5))
})
}

func Call_Data_Bitraversable_bifor(dictBitraversable_0_loop *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value], dictApplicative_1_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], t_2_loop gopurs_runtime.Value, f_3_loop gopurs_runtime.Value, g_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBitraversable_0 *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] = dictBitraversable_0_loop
_ = dictBitraversable_0
var dictApplicative_1 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_1_loop
_ = dictApplicative_1
var t_2 gopurs_runtime.Value = t_2_loop
_ = t_2
var f_3 gopurs_runtime.Value = f_3_loop
_ = f_3
var g_4 gopurs_runtime.Value = g_4_loop
_ = g_4
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictBitraversable_0.V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_1)}, f_3, g_4, t_2)
}

func Rebox_Data_Bitraversable_1495429347_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Bitraversable_1688994542_2864580459(in *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]) *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Bitraversable_2243670499_3561684974(in *Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Bitraversable_2812820739_3566843086(in *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Bitraversable_2864580459_1688994542(in *Constructor_Data_Bifunctor_Bifunctor[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifunctor_Bifunctor[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Bitraversable_3566843086_4280144779(in *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]) *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Bitraversable_4211574891_3561684974(in *Constructor_Data_Bitraversable_Bitraversable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bitraversable_Bitraversable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Bitraversable_4280144779_3566843086(in *Constructor_Data_Bifoldable_Bifoldable[*Constructor_Data_Functor_Product2_Product2[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bifoldable_Bifoldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}



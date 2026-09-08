package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semigroup_Foldable_identity gopurs_runtime.Value
var once_Data_Semigroup_Foldable_identity sync.Once
func Get_Data_Semigroup_Foldable_identity() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_identity.Do(func() {
		cache_Data_Semigroup_Foldable_identity = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Semigroup_Foldable_identity
}

var cache_Data_Semigroup_Foldable_identity1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_identity1 sync.Once
func Get_Data_Semigroup_Foldable_identity1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_identity1.Do(func() {
		cache_Data_Semigroup_Foldable_identity1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)
	})
	return cache_Data_Semigroup_Foldable_identity1
}

var cache_Data_Semigroup_Foldable_JoinWith gopurs_runtime.Value
var once_Data_Semigroup_Foldable_JoinWith sync.Once
func Get_Data_Semigroup_Foldable_JoinWith() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_JoinWith.Do(func() {
		cache_Data_Semigroup_Foldable_JoinWith = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_JoinWith(x_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_JoinWith
}

var cache_Data_Semigroup_Foldable_Foldable1_dollar_Dict gopurs_runtime.Value
var once_Data_Semigroup_Foldable_Foldable1_dollar_Dict sync.Once
func Get_Data_Semigroup_Foldable_Foldable1_dollar_Dict() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_Foldable1_dollar_Dict.Do(func() {
		cache_Data_Semigroup_Foldable_Foldable1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Call_Data_Semigroup_Foldable_Foldable1_dollar_Dict(func() struct{
	Foldable0 gopurs_runtime.Value
	foldMap1 gopurs_runtime.Value
	foldl1 gopurs_runtime.Value
	foldr1 gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Foldable0 gopurs_runtime.Value
	foldMap1 gopurs_runtime.Value
	foldl1 gopurs_runtime.Value
	foldr1 gopurs_runtime.Value
}{}
					clone.Foldable0 = gopurs_runtime.RecordGet(orig, "Foldable0")
					clone.foldMap1 = gopurs_runtime.RecordGet(orig, "foldMap1")
					clone.foldl1 = gopurs_runtime.RecordGet(orig, "foldl1")
					clone.foldr1 = gopurs_runtime.RecordGet(orig, "foldr1")
					return clone
				}()))}
})
	})
	return cache_Data_Semigroup_Foldable_Foldable1_dollar_Dict
}

var cache_Data_Semigroup_Foldable_FoldRight1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_FoldRight1 sync.Once
func Get_Data_Semigroup_Foldable_FoldRight1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_FoldRight1.Do(func() {
		cache_Data_Semigroup_Foldable_FoldRight1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value]{1, value0, value1}))}
})
})
	})
	return cache_Data_Semigroup_Foldable_FoldRight1
}

var cache_Data_Semigroup_Foldable_Act gopurs_runtime.Value
var once_Data_Semigroup_Foldable_Act sync.Once
func Get_Data_Semigroup_Foldable_Act() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_Act.Do(func() {
		cache_Data_Semigroup_Foldable_Act = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_Act(x_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_Act
}

var cache_Data_Semigroup_Foldable_semigroupJoinWith gopurs_runtime.Value
var once_Data_Semigroup_Foldable_semigroupJoinWith sync.Once
func Get_Data_Semigroup_Foldable_semigroupJoinWith() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_semigroupJoinWith.Do(func() {
		cache_Data_Semigroup_Foldable_semigroupJoinWith = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_semigroupJoinWith(dictSemigroup_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_semigroupJoinWith
}

var cache_Data_Semigroup_Foldable_semigroupAct gopurs_runtime.Value
var once_Data_Semigroup_Foldable_semigroupAct sync.Once
func Get_Data_Semigroup_Foldable_semigroupAct() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_semigroupAct.Do(func() {
		cache_Data_Semigroup_Foldable_semigroupAct = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_semigroupAct(dictApply_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_semigroupAct
}

var cache_Data_Semigroup_Foldable_runFoldRight1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_runFoldRight1 sync.Once
func Get_Data_Semigroup_Foldable_runFoldRight1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_runFoldRight1.Do(func() {
		cache_Data_Semigroup_Foldable_runFoldRight1 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_runFoldRight1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_runFoldRight1
}

var cache_Data_Semigroup_Foldable_mkFoldRight1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_mkFoldRight1 sync.Once
func Get_Data_Semigroup_Foldable_mkFoldRight1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_mkFoldRight1.Do(func() {
		cache_Data_Semigroup_Foldable_mkFoldRight1 = gopurs_runtime.Apply(Get_Data_Semigroup_Foldable_FoldRight1(), Get_Data_Function_go__const())
	})
	return cache_Data_Semigroup_Foldable_mkFoldRight1
}

var cache_Data_Semigroup_Foldable_joinee gopurs_runtime.Value
var once_Data_Semigroup_Foldable_joinee sync.Once
func Get_Data_Semigroup_Foldable_joinee() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_joinee.Do(func() {
		cache_Data_Semigroup_Foldable_joinee = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_joinee(v_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_joinee
}

var cache_Data_Semigroup_Foldable_getAct gopurs_runtime.Value
var once_Data_Semigroup_Foldable_getAct sync.Once
func Get_Data_Semigroup_Foldable_getAct() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_getAct.Do(func() {
		cache_Data_Semigroup_Foldable_getAct = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_getAct(v_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_getAct
}

var cache_Data_Semigroup_Foldable_foldr1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldr1 sync.Once
func Get_Data_Semigroup_Foldable_foldr1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldr1.Do(func() {
		cache_Data_Semigroup_Foldable_foldr1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldr1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldr1
}

var cache_Data_Semigroup_Foldable_foldl1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldl1 sync.Once
func Get_Data_Semigroup_Foldable_foldl1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldl1.Do(func() {
		cache_Data_Semigroup_Foldable_foldl1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldl1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldl1
}

var cache_Data_Semigroup_Foldable_maximumBy gopurs_runtime.Value
var once_Data_Semigroup_Foldable_maximumBy sync.Once
func Get_Data_Semigroup_Foldable_maximumBy() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_maximumBy.Do(func() {
		cache_Data_Semigroup_Foldable_maximumBy = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_maximumBy(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), cmp_1_box)
})
	})
	return cache_Data_Semigroup_Foldable_maximumBy
}

var cache_Data_Semigroup_Foldable_minimumBy gopurs_runtime.Value
var once_Data_Semigroup_Foldable_minimumBy sync.Once
func Get_Data_Semigroup_Foldable_minimumBy() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_minimumBy.Do(func() {
		cache_Data_Semigroup_Foldable_minimumBy = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_minimumBy(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), cmp_1_box)
})
	})
	return cache_Data_Semigroup_Foldable_minimumBy
}

var cache_Data_Semigroup_Foldable_foldableTuple gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableTuple sync.Once
func Get_Data_Semigroup_Foldable_foldableTuple() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableTuple.Do(func() {
		cache_Data_Semigroup_Foldable_foldableTuple = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_Foldable_1136416832_4151366573((&Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_Foldable_4173511203_1680800814(Rebox_Data_Semigroup_Foldable_1680800814_4173511203(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableTuple()))))}
}), gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1
})})))}
	})
	return cache_Data_Semigroup_Foldable_foldableTuple
}

var cache_Data_Semigroup_Foldable_foldableMultiplicative gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableMultiplicative sync.Once
func Get_Data_Semigroup_Foldable_foldableMultiplicative() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableMultiplicative.Do(func() {
		cache_Data_Semigroup_Foldable_foldableMultiplicative = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableMultiplicative()))}
}), gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})}))}
	})
	return cache_Data_Semigroup_Foldable_foldableMultiplicative
}

var cache_Data_Semigroup_Foldable_foldableIdentity gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableIdentity sync.Once
func Get_Data_Semigroup_Foldable_foldableIdentity() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableIdentity.Do(func() {
		cache_Data_Semigroup_Foldable_foldableIdentity = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableIdentity()))}
}), gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})}))}
	})
	return cache_Data_Semigroup_Foldable_foldableIdentity
}

var cache_Data_Semigroup_Foldable_foldableDual gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableDual sync.Once
func Get_Data_Semigroup_Foldable_foldableDual() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableDual.Do(func() {
		cache_Data_Semigroup_Foldable_foldableDual = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableDual()))}
}), gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})}))}
	})
	return cache_Data_Semigroup_Foldable_foldableDual
}

var cache_Data_Semigroup_Foldable_foldRight1Semigroup gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldRight1Semigroup sync.Once
func Get_Data_Semigroup_Foldable_foldRight1Semigroup() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldRight1Semigroup.Do(func() {
		cache_Data_Semigroup_Foldable_foldRight1Semigroup = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_Foldable_4053213292_4179793454((&Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value]]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
__local_var_2_0 := (*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Apply2(f_4, __local_var_2_0, gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, a_3, f_4)), f_4)
}), (*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V1}))}
})})))}
	})
	return cache_Data_Semigroup_Foldable_foldRight1Semigroup
}

var cache_Data_Semigroup_Foldable_semigroupDual gopurs_runtime.Value
var once_Data_Semigroup_Foldable_semigroupDual sync.Once
func Get_Data_Semigroup_Foldable_semigroupDual() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_semigroupDual.Do(func() {
		cache_Data_Semigroup_Foldable_semigroupDual = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=Other bindingType=Any
__local_var_2_0 := (*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, gopurs_runtime.Apply2(f_4, __local_var_2_0, gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V0, a_3, f_4)), f_4)
}), (*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V1}))}
})}))}
	})
	return cache_Data_Semigroup_Foldable_semigroupDual
}

var cache_Data_Semigroup_Foldable_foldMap1DefaultR gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1DefaultR sync.Once
func Get_Data_Semigroup_Foldable_foldMap1DefaultR() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1DefaultR.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1DefaultR = gopurs_runtime.Func3(func(dictFoldable1_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1DefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_2_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1DefaultR
}

var cache_Data_Semigroup_Foldable_foldMap1DefaultL gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1DefaultL sync.Once
func Get_Data_Semigroup_Foldable_foldMap1DefaultL() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1DefaultL.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1DefaultL = gopurs_runtime.Func3(func(dictFoldable1_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1DefaultL(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_2_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1DefaultL
}

var cache_Data_Semigroup_Foldable_foldMap1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1
}

var cache_Data_Semigroup_Foldable_foldMap1__3266958652 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__3266958652 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__3266958652() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__3266958652.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__3266958652 = gopurs_runtime.Func2(func(dict_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__3266958652(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](dict_unused_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](__eta_norm_0_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__3266958652
}

var cache_Data_Semigroup_Foldable_foldl1Default gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldl1Default sync.Once
func Get_Data_Semigroup_Foldable_foldl1Default() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldl1Default.Do(func() {
		cache_Data_Semigroup_Foldable_foldl1Default = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldl1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldl1Default
}

var cache_Data_Semigroup_Foldable_foldr1Default gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldr1Default sync.Once
func Get_Data_Semigroup_Foldable_foldr1Default() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldr1Default.Do(func() {
		cache_Data_Semigroup_Foldable_foldr1Default = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldr1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldr1Default
}

var cache_Data_Semigroup_Foldable_intercalateMap gopurs_runtime.Value
var once_Data_Semigroup_Foldable_intercalateMap sync.Once
func Get_Data_Semigroup_Foldable_intercalateMap() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_intercalateMap.Do(func() {
		cache_Data_Semigroup_Foldable_intercalateMap = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_intercalateMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_intercalateMap
}

var cache_Data_Semigroup_Foldable_intercalate gopurs_runtime.Value
var once_Data_Semigroup_Foldable_intercalate sync.Once
func Get_Data_Semigroup_Foldable_intercalate() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_intercalate.Do(func() {
		cache_Data_Semigroup_Foldable_intercalate = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_intercalate
}

var cache_Data_Semigroup_Foldable_maximum gopurs_runtime.Value
var once_Data_Semigroup_Foldable_maximum sync.Once
func Get_Data_Semigroup_Foldable_maximum() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_maximum.Do(func() {
		cache_Data_Semigroup_Foldable_maximum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_maximum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_maximum
}

var cache_Data_Semigroup_Foldable_minimum gopurs_runtime.Value
var once_Data_Semigroup_Foldable_minimum sync.Once
func Get_Data_Semigroup_Foldable_minimum() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_minimum.Do(func() {
		cache_Data_Semigroup_Foldable_minimum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_minimum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_minimum
}

var cache_Data_Semigroup_Foldable_traverse1_ gopurs_runtime.Value
var once_Data_Semigroup_Foldable_traverse1_ sync.Once
func Get_Data_Semigroup_Foldable_traverse1_() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_traverse1_.Do(func() {
		cache_Data_Semigroup_Foldable_traverse1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_traverse1_(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_traverse1_
}

var cache_Data_Semigroup_Foldable_for1_ gopurs_runtime.Value
var once_Data_Semigroup_Foldable_for1_ sync.Once
func Get_Data_Semigroup_Foldable_for1_() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_for1_.Do(func() {
		cache_Data_Semigroup_Foldable_for1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_for1_(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_for1_
}

var cache_Data_Semigroup_Foldable_sequence1_ gopurs_runtime.Value
var once_Data_Semigroup_Foldable_sequence1_ sync.Once
func Get_Data_Semigroup_Foldable_sequence1_() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_sequence1_.Do(func() {
		cache_Data_Semigroup_Foldable_sequence1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_sequence1_(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_sequence1_
}

var cache_Data_Semigroup_Foldable_fold1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_fold1 sync.Once
func Get_Data_Semigroup_Foldable_fold1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_fold1.Do(func() {
		cache_Data_Semigroup_Foldable_fold1 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_fold1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_fold1
}

type Constructor_Data_Semigroup_Foldable_FoldRight1[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
}


type Constructor_Data_Semigroup_Foldable_Foldable1[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2465059545] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semigroup_Foldable_Foldable1[any])(ptr)
		_ = c
		switch key {
		case "Foldable0": return gopurs_runtime.Box(c.V0)
		case "foldMap1": return gopurs_runtime.Box(c.V1)
		case "foldl1": return gopurs_runtime.Box(c.V2)
		case "foldr1": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_Semigroup_Foldable_Foldable1: " + key)
		}
	}
}


func Call_Data_Semigroup_Foldable_JoinWith(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Foldable_Foldable1_dollar_Dict(x_0_loop struct{
	Foldable0 gopurs_runtime.Value
	foldMap1 gopurs_runtime.Value
	foldl1 gopurs_runtime.Value
	foldr1 gopurs_runtime.Value
}) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
var x_0 struct{
	Foldable0 gopurs_runtime.Value
	foldMap1 gopurs_runtime.Value
	foldl1 gopurs_runtime.Value
	foldr1 gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", orig.Foldable0, orig.foldMap1, orig.foldl1, orig.foldr1)
				}())
}

func Call_Data_Semigroup_Foldable_Act(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Foldable_semigroupJoinWith(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value, j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(v_1, j_3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), j_3, gopurs_runtime.Apply(v1_2, j_3)))
})}))}
}

func Call_Data_Semigroup_Foldable_semigroupAct(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Control_Apply_applySecond(), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}, v_1, v1_2)
})}))}
}

func Call_Data_Semigroup_Foldable_runFoldRight1(v_0_loop *Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.Apply((v_0).V0, (v_0).V1)
}

func Call_Data_Semigroup_Foldable_joinee(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Semigroup_Foldable_getAct(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Semigroup_Foldable_foldr1(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Semigroup_Foldable_foldl1(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Semigroup_Foldable_maximumBy(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V2), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}

func Call_Data_Semigroup_Foldable_minimumBy(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V2), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}

func Call_Data_Semigroup_Foldable_foldMap1DefaultR(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], dictFunctor_1_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], dictSemigroup_2_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictFunctor_1 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var dictSemigroup_2 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_2_loop
_ = dictSemigroup_2
// TAST (Let): append_3_0 shape=Other bindingType=(Func [(TypeVar m), (TypeVar m)] (TypeVar m))
append_3_0 := gopurs_runtime.Box(dictSemigroup_2.V0)
_ = append_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (TypeApp (TypeVar t) [(TypeVar m)]))
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_1.V0), f_4)
_ = __local_var_5_1
// TAST (Let): __local_var_6_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar m)])] (TypeVar m))
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V3), append_3_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
}

func Call_Data_Semigroup_Foldable_foldMap1DefaultL(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], dictFunctor_1_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], dictSemigroup_2_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictFunctor_1 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var dictSemigroup_2 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_2_loop
_ = dictSemigroup_2
// TAST (Let): append_3_0 shape=Other bindingType=(Func [(TypeVar m), (TypeVar m)] (TypeVar m))
append_3_0 := gopurs_runtime.Box(dictSemigroup_2.V0)
_ = append_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (TypeApp (TypeVar t) [(TypeVar m)]))
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_1.V0), f_4)
_ = __local_var_5_1
// TAST (Let): __local_var_6_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar m)])] (TypeVar m))
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V2), append_3_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
}

func Call_Data_Semigroup_Foldable_foldMap1(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semigroup_Foldable_foldMap1__3266958652(dict_unused_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]], __eta_norm_0_1_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
foldMap1__3266958652:
for {
if false { continue foldMap1__3266958652 }
var dict_unused_0 *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] = dict_unused_0_loop
_ = dict_unused_0
var __eta_norm_0_1 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_Last_Last(), "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(__eta_norm_0_1)})
}
}

func Call_Data_Semigroup_Foldable_foldl1Default(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (ADT ["Data","Semigroup","Foldable","FoldRight1"] [(TypeVar a)]))
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), Get_Data_Semigroup_Foldable_semigroupDual(), Get_Data_Semigroup_Foldable_mkFoldRight1()))
_ = __local_var_1_0
return gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(__local_var_1_0)}, a_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, (*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, gopurs_runtime.Func2(func(b_5 gopurs_runtime.Value, a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(x_2, a_6, b_5)
}))
})
}

func Call_Data_Semigroup_Foldable_foldr1Default(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): __local_var_1_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (ADT ["Data","Semigroup","Foldable","FoldRight1"] [(TypeVar a)]))
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Semigroup_Foldable_4053213292_4179793454(Rebox_Data_Semigroup_Foldable_4179793454_4053213292(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_Foldable_foldRight1Semigroup()))))}, Get_Data_Semigroup_Foldable_mkFoldRight1())
_ = __local_var_1_0
return gopurs_runtime.Func2(func(b_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=Any
__local_var_4_1 := gopurs_runtime.Apply(__local_var_1_0, a_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V0, (*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value])(__local_var_4_1.UnsafePtr).V1, b_2)
})
}

func Call_Data_Semigroup_Foldable_intercalateMap(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): semigroupJoinWith1_2_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar m)] (TypeVar m))])
semigroupJoinWith1_2_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), gopurs_runtime.Apply(v_2, j_4), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), j_4, gopurs_runtime.Apply(v1_3, j_4)))
})})
_ = semigroupJoinWith1_2_0
return gopurs_runtime.Func3(func(j_3 gopurs_runtime.Value, f_4 gopurs_runtime.Value, foldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_2_0)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_1 shape=App(Other) bindingType=(TypeVar m)
__local_var_7_1 := gopurs_runtime.Apply(f_4, x_6)
_ = __local_var_7_1
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_1
})
}), foldable_5, j_3)
})
}

func Call_Data_Semigroup_Foldable_intercalate(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): semigroupJoinWith1_2_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar m)] (TypeVar m))])
semigroupJoinWith1_2_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), gopurs_runtime.Apply(v_2, j_4), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), j_4, gopurs_runtime.Apply(v1_3, j_4)))
})})
_ = semigroupJoinWith1_2_0
return gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, foldable_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_2_0)}, gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
}), foldable_4, a_3)
})
}

func Call_Data_Semigroup_Foldable_maximum(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): semigroupMax_1_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupMax_1_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_3_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), v_1, v1_2).IntVal)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1 == 1527465420) {
__t2 = v1_2
goto end_branch_2
} else {

}
}
{
if (v_3_1 == 902936544) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (v_3_1 == 380165415) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})})
_ = semigroupMax_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupMax_1_0)}, Get_Unsafe_Coerce_unsafeCoerce())
})
}

func Call_Data_Semigroup_Foldable_minimum(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): semigroupMin_1_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupMin_1_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_1 shape=App(Other) bindingType=(ADT ["Data","Ordering","Ordering"] [])
v_3_1 := uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), v_1, v1_2).IntVal)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1 == 1527465420) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (v_3_1 == 902936544) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (v_3_1 == 380165415) {
__t2 = v1_2
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})})
_ = semigroupMin_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupMin_1_0)}, Get_Unsafe_Coerce_unsafeCoerce())
})
}

func Call_Data_Semigroup_Foldable_traverse1_(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], dictApply_1_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
// TAST (Let): Functor0_2_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}))
_ = Functor0_2_0
// TAST (Let): semigroupAct1_3_1 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeApp (TypeApp (TypeVar f) [(TypeVar a)]) [(TypeVar f), (TypeVar b)])])
semigroupAct1_3_1 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Control_Apply_applySecond(), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(dictApply_1)}, v_3, v1_4)
})})
_ = semigroupAct1_3_1
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
}

func Call_Data_Semigroup_Foldable_for1_(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], dictApply_1_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := Call_Data_Semigroup_Foldable_traverse1_(dictFoldable1_0, dictApply_1)
_ = __local_var_2_0
return gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_4, b_3)
})
}

func Call_Data_Semigroup_Foldable_sequence1_(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], dictApply_1_loop *Constructor_Control_Apply_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *Constructor_Control_Apply_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
return gopurs_runtime.Apply(Call_Data_Semigroup_Foldable_traverse1_(dictFoldable1_0, dictApply_1), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Data_Semigroup_Foldable_fold1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value], dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_1)}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Rebox_Data_Semigroup_Foldable_1136416832_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Semigroup_Foldable_1680800814_4173511203(in *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Semigroup_Foldable_4053213292_4179793454(in *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Semigroup_Foldable_4173511203_1680800814(in *Constructor_Data_Foldable_Foldable[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Data_Semigroup_Foldable_4179793454_4053213292(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[*Constructor_Data_Semigroup_Foldable_FoldRight1[gopurs_runtime.Value]]{}
		out.V0 = in.V0
	return out
}



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
		cache_Data_Semigroup_Foldable_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_identity(x_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_identity
}

var cache_Data_Semigroup_Foldable_identity1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_identity1 sync.Once
func Get_Data_Semigroup_Foldable_identity1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_identity1.Do(func() {
		cache_Data_Semigroup_Foldable_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_identity1(x_0_box)
})
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

var cache_Data_Semigroup_Foldable_Foldable1_dollarDict gopurs_runtime.Value
var once_Data_Semigroup_Foldable_Foldable1_dollarDict sync.Once
func Get_Data_Semigroup_Foldable_Foldable1_dollarDict() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_Foldable1_dollarDict.Do(func() {
		cache_Data_Semigroup_Foldable_Foldable1_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_Foldable1_dollarDict(x_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_Foldable1_dollarDict
}

var cache_Data_Semigroup_Foldable_FoldRight1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_FoldRight1 sync.Once
func Get_Data_Semigroup_Foldable_FoldRight1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_FoldRight1.Do(func() {
		cache_Data_Semigroup_Foldable_FoldRight1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_FoldRight1{1, value0, value1})}
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
return Call_Data_Semigroup_Foldable_runFoldRight1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_FoldRight1](v_0_box))
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
return Call_Data_Semigroup_Foldable_foldr1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldr1
}

var cache_Data_Semigroup_Foldable_foldl1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldl1 sync.Once
func Get_Data_Semigroup_Foldable_foldl1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldl1.Do(func() {
		cache_Data_Semigroup_Foldable_foldl1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldl1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldl1
}

var cache_Data_Semigroup_Foldable_maximumBy gopurs_runtime.Value
var once_Data_Semigroup_Foldable_maximumBy sync.Once
func Get_Data_Semigroup_Foldable_maximumBy() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_maximumBy.Do(func() {
		cache_Data_Semigroup_Foldable_maximumBy = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_maximumBy(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), cmp_1_box)
})
	})
	return cache_Data_Semigroup_Foldable_maximumBy
}

var cache_Data_Semigroup_Foldable_minimumBy gopurs_runtime.Value
var once_Data_Semigroup_Foldable_minimumBy sync.Once
func Get_Data_Semigroup_Foldable_minimumBy() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_minimumBy.Do(func() {
		cache_Data_Semigroup_Foldable_minimumBy = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_minimumBy(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), cmp_1_box)
})
	})
	return cache_Data_Semigroup_Foldable_minimumBy
}

var cache_Data_Semigroup_Foldable_foldableTuple gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableTuple sync.Once
func Get_Data_Semigroup_Foldable_foldableTuple() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableTuple.Do(func() {
		cache_Data_Semigroup_Foldable_foldableTuple = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableTuple()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V1
})
}))
	})
	return cache_Data_Semigroup_Foldable_foldableTuple
}

var cache_Data_Semigroup_Foldable_foldableMultiplicative gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableMultiplicative sync.Once
func Get_Data_Semigroup_Foldable_foldableMultiplicative() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableMultiplicative.Do(func() {
		cache_Data_Semigroup_Foldable_foldableMultiplicative = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMultiplicative()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_Data_Semigroup_Foldable_foldableMultiplicative
}

var cache_Data_Semigroup_Foldable_foldableIdentity gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableIdentity sync.Once
func Get_Data_Semigroup_Foldable_foldableIdentity() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableIdentity.Do(func() {
		cache_Data_Semigroup_Foldable_foldableIdentity = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableIdentity()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_Data_Semigroup_Foldable_foldableIdentity
}

var cache_Data_Semigroup_Foldable_foldableDual gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableDual sync.Once
func Get_Data_Semigroup_Foldable_foldableDual() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableDual.Do(func() {
		cache_Data_Semigroup_Foldable_foldableDual = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableDual()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_Data_Semigroup_Foldable_foldableDual
}

var cache_Data_Semigroup_Foldable_foldRight1Semigroup gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldRight1Semigroup sync.Once
func Get_Data_Semigroup_Foldable_foldRight1Semigroup() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldRight1Semigroup.Do(func() {
		cache_Data_Semigroup_Foldable_foldRight1Semigroup = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (*Constructor_Data_Semigroup_Foldable_FoldRight1)(v_0.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_FoldRight1{1, gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1)(v_0.UnsafePtr).V0, gopurs_runtime.Apply2(f_4, __local_var_2_0, gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1)(v1_1.UnsafePtr).V0, a_3, f_4)), f_4)
})
}), (*Constructor_Data_Semigroup_Foldable_FoldRight1)(v1_1.UnsafePtr).V1})}
})
}))
	})
	return cache_Data_Semigroup_Foldable_foldRight1Semigroup
}

var cache_Data_Semigroup_Foldable_semigroupDual gopurs_runtime.Value
var once_Data_Semigroup_Foldable_semigroupDual sync.Once
func Get_Data_Semigroup_Foldable_semigroupDual() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_semigroupDual.Do(func() {
		cache_Data_Semigroup_Foldable_semigroupDual = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (*Constructor_Data_Semigroup_Foldable_FoldRight1)(v1_1.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_FoldRight1{1, gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1)(v1_1.UnsafePtr).V0, gopurs_runtime.Apply2(f_4, __local_var_2_0, gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1)(v_0.UnsafePtr).V0, a_3, f_4)), f_4)
})
}), (*Constructor_Data_Semigroup_Foldable_FoldRight1)(v_0.UnsafePtr).V1})}
})
}))
	})
	return cache_Data_Semigroup_Foldable_semigroupDual
}

var cache_Data_Semigroup_Foldable_foldMap1DefaultR gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1DefaultR sync.Once
func Get_Data_Semigroup_Foldable_foldMap1DefaultR() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1DefaultR.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1DefaultR = gopurs_runtime.Func3(func(dictFoldable1_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1DefaultR(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_2_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1DefaultR
}

var cache_Data_Semigroup_Foldable_foldMap1DefaultL gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1DefaultL sync.Once
func Get_Data_Semigroup_Foldable_foldMap1DefaultL() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1DefaultL.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1DefaultL = gopurs_runtime.Func3(func(dictFoldable1_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1DefaultL(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_2_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1DefaultL
}

var cache_Data_Semigroup_Foldable_foldMap1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1
}

var cache_Data_Semigroup_Foldable_foldl1Default gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldl1Default sync.Once
func Get_Data_Semigroup_Foldable_foldl1Default() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldl1Default.Do(func() {
		cache_Data_Semigroup_Foldable_foldl1Default = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldl1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldl1Default
}

var cache_Data_Semigroup_Foldable_foldr1Default gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldr1Default sync.Once
func Get_Data_Semigroup_Foldable_foldr1Default() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldr1Default.Do(func() {
		cache_Data_Semigroup_Foldable_foldr1Default = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldr1Default(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldr1Default
}

var cache_Data_Semigroup_Foldable_intercalateMap gopurs_runtime.Value
var once_Data_Semigroup_Foldable_intercalateMap sync.Once
func Get_Data_Semigroup_Foldable_intercalateMap() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_intercalateMap.Do(func() {
		cache_Data_Semigroup_Foldable_intercalateMap = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_intercalateMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_intercalateMap
}

var cache_Data_Semigroup_Foldable_intercalate gopurs_runtime.Value
var once_Data_Semigroup_Foldable_intercalate sync.Once
func Get_Data_Semigroup_Foldable_intercalate() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_intercalate.Do(func() {
		cache_Data_Semigroup_Foldable_intercalate = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_intercalate
}

var cache_Data_Semigroup_Foldable_maximum gopurs_runtime.Value
var once_Data_Semigroup_Foldable_maximum sync.Once
func Get_Data_Semigroup_Foldable_maximum() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_maximum.Do(func() {
		cache_Data_Semigroup_Foldable_maximum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_maximum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_maximum
}

var cache_Data_Semigroup_Foldable_minimum gopurs_runtime.Value
var once_Data_Semigroup_Foldable_minimum sync.Once
func Get_Data_Semigroup_Foldable_minimum() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_minimum.Do(func() {
		cache_Data_Semigroup_Foldable_minimum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_minimum(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_minimum
}

var cache_Data_Semigroup_Foldable_traverse1_ gopurs_runtime.Value
var once_Data_Semigroup_Foldable_traverse1_ sync.Once
func Get_Data_Semigroup_Foldable_traverse1_() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_traverse1_.Do(func() {
		cache_Data_Semigroup_Foldable_traverse1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_traverse1_(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_traverse1_
}

var cache_Data_Semigroup_Foldable_for1_ gopurs_runtime.Value
var once_Data_Semigroup_Foldable_for1_ sync.Once
func Get_Data_Semigroup_Foldable_for1_() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_for1_.Do(func() {
		cache_Data_Semigroup_Foldable_for1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_for1_(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_for1_
}

var cache_Data_Semigroup_Foldable_sequence1_ gopurs_runtime.Value
var once_Data_Semigroup_Foldable_sequence1_ sync.Once
func Get_Data_Semigroup_Foldable_sequence1_() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_sequence1_.Do(func() {
		cache_Data_Semigroup_Foldable_sequence1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_sequence1_(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_sequence1_
}

var cache_Data_Semigroup_Foldable_fold1 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_fold1 sync.Once
func Get_Data_Semigroup_Foldable_fold1() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_fold1.Do(func() {
		cache_Data_Semigroup_Foldable_fold1 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_fold1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dictSemigroup_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_fold1
}

var cache_Data_Semigroup_Foldable_foldMap1__1646952192 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__1646952192 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__1646952192() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__1646952192.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__1646952192 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__1646952192(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__1646952192
}

var cache_Data_Semigroup_Foldable_foldMap1__3675913824 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__3675913824 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__3675913824() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__3675913824.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__3675913824 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__3675913824(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__3675913824
}

var cache_Data_Semigroup_Foldable_foldMap1__4118857699 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__4118857699 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__4118857699() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__4118857699.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__4118857699 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__4118857699(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__4118857699
}

var cache_Data_Semigroup_Foldable_foldMap1__3342855683 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__3342855683 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__3342855683() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__3342855683.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__3342855683 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__3342855683(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__3342855683
}

var cache_Data_Semigroup_Foldable_foldMap1__2273408387 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__2273408387 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__2273408387() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__2273408387.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__2273408387 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__2273408387(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__2273408387
}

var cache_Data_Semigroup_Foldable_foldMap1__3962279491 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__3962279491 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__3962279491() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__3962279491.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__3962279491 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__3962279491(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__3962279491
}

var cache_Data_Semigroup_Foldable_foldMap1__3881804011 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__3881804011 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__3881804011() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__3881804011.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__3881804011 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__3881804011(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__3881804011
}

var cache_Data_Semigroup_Foldable_foldMap1__3273114988 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__3273114988 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__3273114988() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__3273114988.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__3273114988 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__3273114988(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__3273114988
}

var cache_Data_Semigroup_Foldable_foldMap1__1749181674 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__1749181674 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__1749181674() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__1749181674.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__1749181674 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__1749181674(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__1749181674
}

var cache_Data_Semigroup_Foldable_foldMap1__912539845 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__912539845 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__912539845() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__912539845.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__912539845 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__912539845(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__912539845
}

var cache_Data_Semigroup_Foldable_foldMap1__4160988333 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__4160988333 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__4160988333() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__4160988333.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__4160988333 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__4160988333(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__4160988333
}

var cache_Data_Semigroup_Foldable_foldMap1__3028551755 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__3028551755 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__3028551755() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__3028551755.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__3028551755 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__3028551755(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__3028551755
}

var cache_Data_Semigroup_Foldable_foldMap1__2836832395 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__2836832395 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__2836832395() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__2836832395.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__2836832395 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__2836832395(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__2836832395
}

var cache_Data_Semigroup_Foldable_foldMap1__3134302309 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__3134302309 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__3134302309() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__3134302309.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__3134302309 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__3134302309(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__3134302309
}

var cache_Data_Semigroup_Foldable_foldMap1__3810372805 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldMap1__3810372805 sync.Once
func Get_Data_Semigroup_Foldable_foldMap1__3810372805() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldMap1__3810372805.Do(func() {
		cache_Data_Semigroup_Foldable_foldMap1__3810372805 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldMap1__3810372805(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semigroup_Foldable_foldMap1__3810372805
}

var cache_Data_Semigroup_Foldable_foldRight1Semigroup__1201419834 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldRight1Semigroup__1201419834 sync.Once
func Get_Data_Semigroup_Foldable_foldRight1Semigroup__1201419834() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldRight1Semigroup__1201419834.Do(func() {
		cache_Data_Semigroup_Foldable_foldRight1Semigroup__1201419834 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := (*Constructor_Data_Semigroup_Foldable_FoldRight1)(v_0.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_FoldRight1{1, gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1)(v_0.UnsafePtr).V0, gopurs_runtime.Apply2(f_4, __local_var_2_0, gopurs_runtime.Apply2((*Constructor_Data_Semigroup_Foldable_FoldRight1)(v1_1.UnsafePtr).V0, a_3, f_4)), f_4)
})
}), (*Constructor_Data_Semigroup_Foldable_FoldRight1)(v1_1.UnsafePtr).V1})}
})
}))
	})
	return cache_Data_Semigroup_Foldable_foldRight1Semigroup__1201419834
}

var cache_Data_Semigroup_Foldable_foldableDual__189846079 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableDual__189846079 sync.Once
func Get_Data_Semigroup_Foldable_foldableDual__189846079() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableDual__189846079.Do(func() {
		cache_Data_Semigroup_Foldable_foldableDual__189846079 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableDual()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_Data_Semigroup_Foldable_foldableDual__189846079
}

var cache_Data_Semigroup_Foldable_foldableIdentity__189846079 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableIdentity__189846079 sync.Once
func Get_Data_Semigroup_Foldable_foldableIdentity__189846079() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableIdentity__189846079.Do(func() {
		cache_Data_Semigroup_Foldable_foldableIdentity__189846079 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableIdentity()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_Data_Semigroup_Foldable_foldableIdentity__189846079
}

var cache_Data_Semigroup_Foldable_foldableMultiplicative__189846079 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableMultiplicative__189846079 sync.Once
func Get_Data_Semigroup_Foldable_foldableMultiplicative__189846079() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableMultiplicative__189846079.Do(func() {
		cache_Data_Semigroup_Foldable_foldableMultiplicative__189846079 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableMultiplicative()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_Data_Semigroup_Foldable_foldableMultiplicative__189846079
}

var cache_Data_Semigroup_Foldable_foldableTuple__3696373503 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldableTuple__3696373503 sync.Once
func Get_Data_Semigroup_Foldable_foldableTuple__3696373503() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldableTuple__3696373503.Do(func() {
		cache_Data_Semigroup_Foldable_foldableTuple__3696373503 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableTuple()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V1
})
}))
	})
	return cache_Data_Semigroup_Foldable_foldableTuple__3696373503
}

var cache_Data_Semigroup_Foldable_foldl1__3059734942 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldl1__3059734942 sync.Once
func Get_Data_Semigroup_Foldable_foldl1__3059734942() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldl1__3059734942.Do(func() {
		cache_Data_Semigroup_Foldable_foldl1__3059734942 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldl1__3059734942(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldl1__3059734942
}

var cache_Data_Semigroup_Foldable_foldl1__2849185176 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldl1__2849185176 sync.Once
func Get_Data_Semigroup_Foldable_foldl1__2849185176() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldl1__2849185176.Do(func() {
		cache_Data_Semigroup_Foldable_foldl1__2849185176 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldl1__2849185176(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semigroup_Foldable_foldl1__2849185176
}

var cache_Data_Semigroup_Foldable_foldr1__3059734942 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldr1__3059734942 sync.Once
func Get_Data_Semigroup_Foldable_foldr1__3059734942() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldr1__3059734942.Do(func() {
		cache_Data_Semigroup_Foldable_foldr1__3059734942 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldr1__3059734942(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dict_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_foldr1__3059734942
}

var cache_Data_Semigroup_Foldable_foldr1__2849185176 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_foldr1__2849185176 sync.Once
func Get_Data_Semigroup_Foldable_foldr1__2849185176() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_foldr1__2849185176.Do(func() {
		cache_Data_Semigroup_Foldable_foldr1__2849185176 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_foldr1__2849185176(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semigroup_Foldable_foldr1__2849185176
}

var cache_Data_Semigroup_Foldable_getAct__3579306522 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_getAct__3579306522 sync.Once
func Get_Data_Semigroup_Foldable_getAct__3579306522() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_getAct__3579306522.Do(func() {
		cache_Data_Semigroup_Foldable_getAct__3579306522 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_getAct__3579306522(v_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_getAct__3579306522
}

var cache_Data_Semigroup_Foldable_getAct__3831010933 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_getAct__3831010933 sync.Once
func Get_Data_Semigroup_Foldable_getAct__3831010933() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_getAct__3831010933.Do(func() {
		cache_Data_Semigroup_Foldable_getAct__3831010933 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_getAct__3831010933(v_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_getAct__3831010933
}

var cache_Data_Semigroup_Foldable_joinee__1857126433 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_joinee__1857126433 sync.Once
func Get_Data_Semigroup_Foldable_joinee__1857126433() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_joinee__1857126433.Do(func() {
		cache_Data_Semigroup_Foldable_joinee__1857126433 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_joinee__1857126433(v_0_box)
})
	})
	return cache_Data_Semigroup_Foldable_joinee__1857126433
}

var cache_Data_Semigroup_Foldable_mkFoldRight1__364767315 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_mkFoldRight1__364767315 sync.Once
func Get_Data_Semigroup_Foldable_mkFoldRight1__364767315() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_mkFoldRight1__364767315.Do(func() {
		cache_Data_Semigroup_Foldable_mkFoldRight1__364767315 = gopurs_runtime.Apply(Get_Data_Semigroup_Foldable_FoldRight1(), Get_Data_Function_go__const())
	})
	return cache_Data_Semigroup_Foldable_mkFoldRight1__364767315
}

var cache_Data_Semigroup_Foldable_runFoldRight1__3719153019 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_runFoldRight1__3719153019 sync.Once
func Get_Data_Semigroup_Foldable_runFoldRight1__3719153019() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_runFoldRight1__3719153019.Do(func() {
		cache_Data_Semigroup_Foldable_runFoldRight1__3719153019 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_runFoldRight1__3719153019(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_FoldRight1](v_0_box))
})
	})
	return cache_Data_Semigroup_Foldable_runFoldRight1__3719153019
}

var cache_Data_Semigroup_Foldable_traverse1___3055398386 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_traverse1___3055398386 sync.Once
func Get_Data_Semigroup_Foldable_traverse1___3055398386() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_traverse1___3055398386.Do(func() {
		cache_Data_Semigroup_Foldable_traverse1___3055398386 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_traverse1___3055398386(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_traverse1___3055398386
}

var cache_Data_Semigroup_Foldable_traverse1___1509071858 gopurs_runtime.Value
var once_Data_Semigroup_Foldable_traverse1___1509071858 sync.Once
func Get_Data_Semigroup_Foldable_traverse1___1509071858() gopurs_runtime.Value {
	once_Data_Semigroup_Foldable_traverse1___1509071858.Do(func() {
		cache_Data_Semigroup_Foldable_traverse1___1509071858 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Foldable_traverse1___1509071858(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_1_box))
})
	})
	return cache_Data_Semigroup_Foldable_traverse1___1509071858
}

type Constructor_Data_Semigroup_Foldable_FoldRight1 struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


type Constructor_Data_Semigroup_Foldable_Foldable1 struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2465059545] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semigroup_Foldable_Foldable1)(ptr)
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


func Call_Data_Semigroup_Foldable_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Foldable_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Foldable_JoinWith(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Foldable_Foldable1_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Foldable_Act(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Foldable_semigroupJoinWith(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(v_1, j_3), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), j_3, gopurs_runtime.Apply(v1_2, j_3)))
})
})
}))
}

func Call_Data_Semigroup_Foldable_semigroupAct(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
})
}), v_1), v1_2)
})
}))
}

func Call_Data_Semigroup_Foldable_runFoldRight1(v_0_loop *Constructor_Data_Semigroup_Foldable_FoldRight1) gopurs_runtime.Value {
var v_0 *Constructor_Data_Semigroup_Foldable_FoldRight1 = v_0_loop
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

func Call_Data_Semigroup_Foldable_foldr1(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Semigroup_Foldable_foldl1(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Semigroup_Foldable_maximumBy(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V2), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
// TAST (Let): __local_var_4_0 -> uint32
__local_var_4_0 := uint32(gopurs_runtime.Apply2(cmp_1, x_2, y_3).IntVal)
_ = __local_var_4_0
var __t1 bool
{
if (__local_var_4_0 == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = x_2
goto end_branch_2
} else {

}
}
{
__t2 = y_3
}
end_branch_2:
return __t2
})
}))
}

func Call_Data_Semigroup_Foldable_minimumBy(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V2), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
// TAST (Let): __local_var_4_0 -> uint32
__local_var_4_0 := uint32(gopurs_runtime.Apply2(cmp_1, x_2, y_3).IntVal)
_ = __local_var_4_0
var __t1 bool
{
if (__local_var_4_0 == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = x_2
goto end_branch_2
} else {

}
}
{
__t2 = y_3
}
end_branch_2:
return __t2
})
}))
}

func Call_Data_Semigroup_Foldable_foldMap1DefaultR(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictFunctor_1_loop *Constructor_Data_Functor_Functor, dictSemigroup_2_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictFunctor_1 *Constructor_Data_Functor_Functor = dictFunctor_1_loop
_ = dictFunctor_1
var dictSemigroup_2 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_2_loop
_ = dictSemigroup_2
// TAST (Let): append_3_0 -> gopurs_runtime.Value
append_3_0 := gopurs_runtime.Box(dictSemigroup_2.V0)
_ = append_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_1.V0), f_4)
_ = __local_var_5_1
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V3), append_3_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
}

func Call_Data_Semigroup_Foldable_foldMap1DefaultL(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictFunctor_1_loop *Constructor_Data_Functor_Functor, dictSemigroup_2_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictFunctor_1 *Constructor_Data_Functor_Functor = dictFunctor_1_loop
_ = dictFunctor_1
var dictSemigroup_2 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_2_loop
_ = dictSemigroup_2
// TAST (Let): append_3_0 -> gopurs_runtime.Value
append_3_0 := gopurs_runtime.Box(dictSemigroup_2.V0)
_ = append_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFunctor_1.V0), f_4)
_ = __local_var_5_1
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V2), append_3_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
}

func Call_Data_Semigroup_Foldable_foldMap1(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semigroup_Foldable_foldl1Default(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), Get_Data_Semigroup_Foldable_semigroupDual(), Get_Data_Semigroup_Foldable_mkFoldRight1())
_ = __local_var_1_2
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(__local_var_1_2, x_2)
_ = __local_var_3_3
return gopurs_runtime.Apply((*Constructor_Data_Semigroup_Foldable_FoldRight1)(__local_var_3_3.UnsafePtr).V0, (*Constructor_Data_Semigroup_Foldable_FoldRight1)(__local_var_3_3.UnsafePtr).V1)
})
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_1, a_3, b_2)
})
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(x_2, a_4, b_3)
})
}))
})
}

func Call_Data_Semigroup_Foldable_foldr1Default(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Foldable_foldRight1Semigroup()))}, Get_Data_Semigroup_Foldable_mkFoldRight1())
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(__local_var_1_1, x_2)
_ = __local_var_3_2
return gopurs_runtime.Apply((*Constructor_Data_Semigroup_Foldable_FoldRight1)(__local_var_3_2.UnsafePtr).V0, (*Constructor_Data_Semigroup_Foldable_FoldRight1)(__local_var_3_2.UnsafePtr).V1)
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, a_3, b_2)
})
})
}

func Call_Data_Semigroup_Foldable_intercalateMap(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): semigroupJoinWith1_2_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupJoinWith1_2_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), gopurs_runtime.Apply(v_2, j_4), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), j_4, gopurs_runtime.Apply(v1_3, j_4)))
})
})
})}
_ = semigroupJoinWith1_2_0
return gopurs_runtime.Func(func(j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(foldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_2_0)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_1 -> gopurs_runtime.Value
__local_var_7_1 := gopurs_runtime.Apply(f_4, x_6)
_ = __local_var_7_1
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_1
})
}), foldable_5, j_3)
})
})
})
}

func Call_Data_Semigroup_Foldable_intercalate(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_1_loop
_ = dictSemigroup_1
// TAST (Let): semigroupJoinWith1_2_1 -> *Constructor_Data_Semigroup_Semigroup
semigroupJoinWith1_2_1 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), gopurs_runtime.Apply(v_2, j_4), gopurs_runtime.Apply2(gopurs_runtime.Box(dictSemigroup_1.V0), j_4, gopurs_runtime.Apply(v1_3, j_4)))
})
})
})}
_ = semigroupJoinWith1_2_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(foldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_2_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_2 -> gopurs_runtime.Value
__local_var_7_2 := gopurs_runtime.Apply(f_4, x_6)
_ = __local_var_7_2
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_2
})
}), foldable_5, j_3)
})
})
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_3, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}))
})
}

func Call_Data_Semigroup_Foldable_maximum(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): semigroupMax_1_0 -> gopurs_runtime.Value
semigroupMax_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_1 -> gopurs_runtime.Value
v_3_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), v_1, v1_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (uint32(v_3_1.IntVal) == 1527465420) {
__t2 = v1_2
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 902936544) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 380165415) {
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
})
}))
_ = semigroupMax_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), semigroupMax_1_0, Get_Unsafe_Coerce_unsafeCoerce())
})
}

func Call_Data_Semigroup_Foldable_minimum(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): semigroupMin_1_0 -> gopurs_runtime.Value
semigroupMin_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_3_1 -> gopurs_runtime.Value
v_3_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), v_1, v1_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (uint32(v_3_1.IntVal) == 1527465420) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 902936544) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 380165415) {
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
})
}))
_ = semigroupMin_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), semigroupMin_1_0, Get_Unsafe_Coerce_unsafeCoerce())
})
}

func Call_Data_Semigroup_Foldable_traverse1_(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictApply_1_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *Constructor_Control_Apply_Apply = dictApply_1_loop
_ = dictApply_1
// TAST (Let): Functor0_2_0 -> *Constructor_Data_Functor_Functor
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}))
_ = Functor0_2_0
// TAST (Let): semigroupAct1_3_1 -> *Constructor_Data_Semigroup_Semigroup
semigroupAct1_3_1 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), v_3), v1_4)
})
})}
_ = semigroupAct1_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
})
}

func Call_Data_Semigroup_Foldable_for1_(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictApply_1_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *Constructor_Control_Apply_Apply = dictApply_1_loop
_ = dictApply_1
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): semigroupAct1_3_2 -> *Constructor_Data_Semigroup_Semigroup
semigroupAct1_3_2 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), v_3), v1_4)
})
})}
_ = semigroupAct1_3_2
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_2)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_4, b_3)
})
})
}

func Call_Data_Semigroup_Foldable_sequence1_(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictApply_1_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *Constructor_Control_Apply_Apply = dictApply_1_loop
_ = dictApply_1
// TAST (Let): Functor0_2_0 -> *Constructor_Data_Functor_Functor
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}))
_ = Functor0_2_0
// TAST (Let): semigroupAct1_3_1 -> *Constructor_Data_Semigroup_Semigroup
semigroupAct1_3_1 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), v_3), v1_4)
})
})}
_ = semigroupAct1_3_1
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_1)}, gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
}), t_4))
})
}

func Call_Data_Semigroup_Foldable_fold1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictSemigroup_1_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *Constructor_Data_Semigroup_Semigroup = dictSemigroup_1_loop
_ = dictSemigroup_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_1)}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}

func Call_Data_Semigroup_Foldable_foldMap1__1646952192(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_String_NonEmpty_Internal_semigroupNonEmptyString()))})
}

func Call_Data_Semigroup_Foldable_foldMap1__3675913824(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_String_NonEmpty_Internal_semigroupNonEmptyString()))})
}

func Call_Data_Semigroup_Foldable_foldMap1__4118857699(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semigroup_Foldable_foldMap1__3342855683(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semigroup_Foldable_foldMap1__2273408387(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semigroup_Foldable_foldMap1__3962279491(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semigroup_Foldable_foldMap1__3881804011(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Foldable_foldRight1Semigroup()))})
}

func Call_Data_Semigroup_Foldable_foldMap1__3273114988(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))})
}

func Call_Data_Semigroup_Foldable_foldMap1__1749181674(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semigroup_Foldable_foldMap1__912539845(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semigroup_Foldable_foldMap1__4160988333(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Control_Monad_Gen_semigroupFreqSemigroup()))})
}

func Call_Data_Semigroup_Foldable_foldMap1__3028551755(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semigroup_Foldable_foldMap1__2836832395(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))})
}

func Call_Data_Semigroup_Foldable_foldMap1__3134302309(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(gopurs_runtime.Box(dict_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))})
}

func Call_Data_Semigroup_Foldable_foldMap1__3810372805(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldable1NonEmptyList(), "foldMap1"), __eta0_0, __eta1_1)
}

func Call_Data_Semigroup_Foldable_foldl1__3059734942(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Semigroup_Foldable_foldl1__2849185176(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldable1NonEmptyList(), "foldl1"), __eta0_0, __eta1_1)
}

func Call_Data_Semigroup_Foldable_foldr1__3059734942(dict_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Semigroup_Foldable_foldr1__2849185176(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_foldable1NonEmptyList(), "foldr1"), __eta0_0, __eta1_1)
}

func Call_Data_Semigroup_Foldable_getAct__3579306522(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Semigroup_Foldable_getAct__3831010933(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Semigroup_Foldable_joinee__1857126433(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Semigroup_Foldable_runFoldRight1__3719153019(v_0_loop *Constructor_Data_Semigroup_Foldable_FoldRight1) gopurs_runtime.Value {
var v_0 *Constructor_Data_Semigroup_Foldable_FoldRight1 = v_0_loop
_ = v_0
return gopurs_runtime.Apply((v_0).V0, (v_0).V1)
}

func Call_Data_Semigroup_Foldable_traverse1___3055398386(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictApply_1_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *Constructor_Control_Apply_Apply = dictApply_1_loop
_ = dictApply_1
// TAST (Let): Functor0_2_0 -> *Constructor_Data_Functor_Functor
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}))
_ = Functor0_2_0
// TAST (Let): semigroupAct1_3_1 -> *Constructor_Data_Semigroup_Semigroup
semigroupAct1_3_1 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), v_3), v1_4)
})
})}
_ = semigroupAct1_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
})
}

func Call_Data_Semigroup_Foldable_traverse1___1509071858(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1, dictApply_1_loop *Constructor_Control_Apply_Apply) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *Constructor_Control_Apply_Apply = dictApply_1_loop
_ = dictApply_1
// TAST (Let): Functor0_2_0 -> *Constructor_Data_Functor_Functor
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}))
_ = Functor0_2_0
// TAST (Let): semigroupAct1_3_1 -> *Constructor_Data_Semigroup_Semigroup
semigroupAct1_3_1 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictApply_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApply_1.V0), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), v_3), v1_4)
})
})}
_ = semigroupAct1_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
})
}


